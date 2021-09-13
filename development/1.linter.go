package development

import "golang.org/x/tools/go/analysis"

// pass analysis.Pass see golang.org/x/tools/go/analysis/analysis.go
// method should implement run(pass *analysis.Pass) (interface{}, error)
// default we should return reflect.Type and error or can change it
func basic(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
