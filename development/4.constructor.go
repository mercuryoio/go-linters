package development

import "golang.org/x/tools/go/analysis"

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: lintCategory,
		Doc:  lintDoc,
		Run:  runReport,
	}
}
