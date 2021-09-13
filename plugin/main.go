package main

import (
	"github.com/mercuryoio/go-config/internal/config"

	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		config.NewAnalyzer(),
	}
}

var AnalyzerPlugin analyzerPlugin
