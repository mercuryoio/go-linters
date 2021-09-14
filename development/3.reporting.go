package development

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func runReport(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// check os.Getenv calls
			if funcNode, ok := n.(*ast.CallExpr); ok {
				if selExpr, ok := funcNode.Fun.(*ast.SelectorExpr); ok {
					if selExpr.Sel.Name == osGetEnvFuncName &&
						fmt.Sprintf("%v", selExpr.X) == osPackageName {
						pass.Report(analysis.Diagnostic{
							Pos:            selExpr.Sel.Pos(),
							End:            selExpr.Sel.End(),
							Category:       cfgName,
							Message:        osGetEnvErr,
							SuggestedFixes: []analysis.SuggestedFix{getSuggestedFix()},
						})
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func getSuggestedFix() analysis.SuggestedFix {
	return analysis.SuggestedFix{Message: fixMessage}
}
