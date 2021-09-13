package development

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func runReport(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if funcNode, ok := n.(*ast.StructType); ok {
				pass.Report(analysis.Diagnostic{
					Pos:            funcNode.Pos(),
					End:            funcNode.End(),
					Category:       lintCategory,
					Message:        lintMessage,
					SuggestedFixes: []analysis.SuggestedFix{getSuggestedFix()},
				})
			}
			return true
		})
	}
	return nil, nil
}

func getSuggestedFix() analysis.SuggestedFix {
	return analysis.SuggestedFix{Message: fixMessage}
}
