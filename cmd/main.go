package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/mercuryoio/go-config/linters/config"
)

func main() {
	multichecker.Main(config.NewAnalyzer())
}
