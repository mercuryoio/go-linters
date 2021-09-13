package config

import (
	"flag"
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	osGetEnvFix = analysis.SuggestedFix{Message: "see config domain"}
	flagSet     flag.FlagSet
)

const (
	emptyTagCfgErr   = "config attribute %s has no tag %s (%s)"
	osGetEnvErr      = "os.Getenv usage"
	cfgDoc           = "calculates cyclomatic complexity"
	cfgName          = "config"
	osGetEnvFuncName = "Getenv"
	osPackageName    = "os"

	confPackageName = "config"

	confTag    = "envconfig"
	defaultTag = "default"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  cfgName,
		Doc:   cfgDoc,
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
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
							SuggestedFixes: []analysis.SuggestedFix{osGetEnvFix},
						})
					}
				}
			}
			// check tag config values
			if fmt.Sprintf("%v", file.Name) == confPackageName {
				if strNode, ok := n.(*ast.StructType); ok {
					for _, configEl := range strNode.Fields.List {
						basicName := configEl.Names[0]
						filteredTag := strings.ReplaceAll(configEl.Tag.Value, "`", "")
						tag := reflect.StructTag(filteredTag)
						if _, ok := tag.Lookup(confTag); !ok {
							pass.Report(analysis.Diagnostic{
								Pos:            strNode.Pos(),
								End:            strNode.End(),
								Category:       cfgName,
								Message:        fmt.Sprintf(emptyTagCfgErr, basicName.Name, confTag, configEl.Tag.Value),
								SuggestedFixes: []analysis.SuggestedFix{osGetEnvFix},
							})
						}
						if _, ok := tag.Lookup(defaultTag); !ok {
							pass.Report(analysis.Diagnostic{
								Pos:            strNode.Pos(),
								End:            strNode.End(),
								Category:       cfgName,
								Message:        fmt.Sprintf(emptyTagCfgErr, basicName.Name, defaultTag, configEl.Tag.Value),
								SuggestedFixes: []analysis.SuggestedFix{osGetEnvFix},
							})
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
