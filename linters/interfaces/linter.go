package interfaces

import (
	"fmt"
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var (
	osGetEnvFix = analysis.SuggestedFix{Message: "see config domain"}
)

const (
	fixMessage           = "use unexported interface only"
	exportedInterfaceErr = "exported interface %s"
	doc                  = "check interfaces is unexported"
	name                 = "interfaces"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: name,
		Doc:  doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// check unexported interfaces
			if strNode, ok := n.(*ast.TypeSpec); ok {
				firstSymbol := fmt.Sprintf("%s", strNode.Name)
				if unicode.IsUpper(rune(firstSymbol[0])) {
					pass.Report(analysis.Diagnostic{
						Pos:            strNode.Pos(),
						End:            strNode.End(),
						Category:       name,
						Message:        fmt.Sprintf(exportedInterfaceErr, strNode.Name),
						SuggestedFixes: []analysis.SuggestedFix{getSuggestedFix()},
					})
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
