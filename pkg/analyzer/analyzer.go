package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "typenaming",
	Doc:      "Checks that user declared types have not the `Type` suffix.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		typeDecl, ok := node.(*ast.TypeSpec)
		if !ok {
			return
		}

		if !strings.HasSuffix(typeDecl.Name.Name, "Type") {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     typeDecl.Name.Pos(),
			Message: "trim suffix `Type` from type name",
		})
	})

	return nil, nil
}
