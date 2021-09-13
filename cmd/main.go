package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mercuryoio/go-config/internal/config"
)

func main() {
	singlechecker.Main(config.NewAnalyzer())
}
