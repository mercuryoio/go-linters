package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/mercuryoio/go-config/linters/config"
	"github.com/mercuryoio/go-config/linters/interfaces"
)

func main() {
	multichecker.Main(config.NewAnalyzer(), interfaces.NewAnalyzer())
}
