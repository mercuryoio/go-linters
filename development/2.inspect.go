package development

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func runInspect(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			return true
		})
	}
	return nil, nil
}
