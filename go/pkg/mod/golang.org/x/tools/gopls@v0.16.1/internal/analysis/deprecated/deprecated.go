// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package deprecated

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	_ "embed"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/internal/analysisinternal"
)

//go:embed doc.go
var doc string

var Analyzer = &analysis.Analyzer{
	Name:             "deprecated",
	Doc:              analysisinternal.MustExtractDoc(doc, "deprecated"),
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	Run:              checkDeprecated,
	FactTypes:        []analysis.Fact{(*deprecationFact)(nil)},
	RunDespiteErrors: true,
	URL:              "https://pkg.go.dev/golang.org/x/tools/gopls/internal/analysis/deprecated",
}

// checkDeprecated is a simplified copy of staticcheck.CheckDeprecated.
func checkDeprecated(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	deprs, err := collectDeprecatedNames(pass, inspector)
	if err != nil || (len(deprs.packages) == 0 && len(deprs.objects) == 0) {
		return nil, err
	}

	reportDeprecation := func(depr *deprecationFact, node ast.Node) {
		// TODO(hyangah): staticcheck.CheckDeprecated has more complex logic. Do we need it here?
		// TODO(hyangah): Scrub depr.Msg. depr.Msg may contain Go comments
		// markdown syntaxes but LSP diagnostics do not support markdown syntax.

		buf := new(bytes.Buffer)
		if err := format.Node(buf, pass.Fset, node); err != nil {
			// This shouldn't happen but let's be conservative.
			buf.Reset()
			buf.WriteString("declaration")
		}
		pass.ReportRangef(node, "%s is deprecated: %s", buf, depr.Msg)
	}

	nodeFilter := []ast.Node{(*ast.SelectorExpr)(nil)}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		// Caveat: this misses dot-imported objects
		sel, ok := node.(*ast.SelectorExpr)
		if !ok {
			return
		}

		obj := pass.TypesInfo.ObjectOf(sel.Sel)
		if fn, ok := obj.(*types.Func); ok {
			obj = fn.Origin()
		}
		if obj == nil || obj.Pkg() == nil {
			// skip invalid sel.Sel.
			return
		}

		if obj.Pkg() == pass.Pkg {
			// A package is allowed to use its own deprecated objects
			return
		}

		// A package "foo" has two related packages "foo_test" and "foo.test", for external tests and the package main
		// generated by 'go test' respectively. "foo_test" can import and use "foo", "foo.test" imports and uses "foo"
		// and "foo_test".

		if strings.TrimSuffix(pass.Pkg.Path(), "_test") == obj.Pkg().Path() {
			// foo_test (the external tests of foo) can use objects from foo.
			return
		}
		if strings.TrimSuffix(pass.Pkg.Path(), ".test") == obj.Pkg().Path() {
			// foo.test (the main package of foo's tests) can use objects from foo.
			return
		}
		if strings.TrimSuffix(pass.Pkg.Path(), ".test") == strings.TrimSuffix(obj.Pkg().Path(), "_test") {
			// foo.test (the main package of foo's tests) can use objects from foo's external tests.
			return
		}

		if depr, ok := deprs.objects[obj]; ok {
			reportDeprecation(depr, sel)
		}
	})

	for _, f := range pass.Files {
		for _, spec := range f.Imports {
			var imp *types.Package
			var obj types.Object
			if spec.Name != nil {
				obj = pass.TypesInfo.ObjectOf(spec.Name)
			} else {
				obj = pass.TypesInfo.Implicits[spec]
			}
			pkgName, ok := obj.(*types.PkgName)
			if !ok {
				continue
			}
			imp = pkgName.Imported()

			path, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				continue
			}
			pkgPath := pass.Pkg.Path()
			if strings.TrimSuffix(pkgPath, "_test") == path {
				// foo_test can import foo
				continue
			}
			if strings.TrimSuffix(pkgPath, ".test") == path {
				// foo.test can import foo
				continue
			}
			if strings.TrimSuffix(pkgPath, ".test") == strings.TrimSuffix(path, "_test") {
				// foo.test can import foo_test
				continue
			}
			if depr, ok := deprs.packages[imp]; ok {
				reportDeprecation(depr, spec.Path)
			}
		}
	}
	return nil, nil
}

type deprecationFact struct{ Msg string }

func (*deprecationFact) AFact()           {}
func (d *deprecationFact) String() string { return "Deprecated: " + d.Msg }

type deprecatedNames struct {
	objects  map[types.Object]*deprecationFact
	packages map[*types.Package]*deprecationFact
}

// collectDeprecatedNames collects deprecated identifiers and publishes
// them both as Facts and the return value. This is a simplified copy
// of staticcheck's fact_deprecated analyzer.
func collectDeprecatedNames(pass *analysis.Pass, ins *inspector.Inspector) (deprecatedNames, error) {
	extractDeprecatedMessage := func(docs []*ast.CommentGroup) string {
		for _, doc := range docs {
			if doc == nil {
				continue
			}
			parts := strings.Split(doc.Text(), "\n\n")
			for _, part := range parts {
				if !strings.HasPrefix(part, "Deprecated: ") {
					continue
				}
				alt := part[len("Deprecated: "):]
				alt = strings.Replace(alt, "\n", " ", -1)
				return strings.TrimSpace(alt)
			}
		}
		return ""
	}

	doDocs := func(names []*ast.Ident, docs *ast.CommentGroup) {
		alt := extractDeprecatedMessage([]*ast.CommentGroup{docs})
		if alt == "" {
			return
		}

		for _, name := range names {
			obj := pass.TypesInfo.ObjectOf(name)
			pass.ExportObjectFact(obj, &deprecationFact{alt})
		}
	}

	var docs []*ast.CommentGroup
	for _, f := range pass.Files {
		docs = append(docs, f.Doc)
	}
	if alt := extractDeprecatedMessage(docs); alt != "" {
		// Don't mark package syscall as deprecated, even though
		// it is. A lot of people still use it for simple
		// constants like SIGKILL, and I am not comfortable
		// telling them to use x/sys for that.
		if pass.Pkg.Path() != "syscall" {
			pass.ExportPackageFact(&deprecationFact{alt})
		}
	}
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.TypeSpec)(nil),
		(*ast.ValueSpec)(nil),
		(*ast.File)(nil),
		(*ast.StructType)(nil),
		(*ast.InterfaceType)(nil),
	}
	ins.Preorder(nodeFilter, func(node ast.Node) {
		var names []*ast.Ident
		var docs *ast.CommentGroup
		switch node := node.(type) {
		case *ast.GenDecl:
			switch node.Tok {
			case token.TYPE, token.CONST, token.VAR:
				docs = node.Doc
				for i := range node.Specs {
					switch n := node.Specs[i].(type) {
					case *ast.ValueSpec:
						names = append(names, n.Names...)
					case *ast.TypeSpec:
						names = append(names, n.Name)
					}
				}
			default:
				return
			}
		case *ast.FuncDecl:
			docs = node.Doc
			names = []*ast.Ident{node.Name}
		case *ast.TypeSpec:
			docs = node.Doc
			names = []*ast.Ident{node.Name}
		case *ast.ValueSpec:
			docs = node.Doc
			names = node.Names
		case *ast.StructType:
			for _, field := range node.Fields.List {
				doDocs(field.Names, field.Doc)
			}
		case *ast.InterfaceType:
			for _, field := range node.Methods.List {
				doDocs(field.Names, field.Doc)
			}
		}
		if docs != nil && len(names) > 0 {
			doDocs(names, docs)
		}
	})

	// Every identifier is potentially deprecated, so we will need
	// to look up facts a lot. Construct maps of all facts propagated
	// to this pass for fast lookup.
	out := deprecatedNames{
		objects:  map[types.Object]*deprecationFact{},
		packages: map[*types.Package]*deprecationFact{},
	}
	for _, fact := range pass.AllObjectFacts() {
		out.objects[fact.Object] = fact.Fact.(*deprecationFact)
	}
	for _, fact := range pass.AllPackageFacts() {
		out.packages[fact.Package] = fact.Fact.(*deprecationFact)
	}

	return out, nil
}
