package golang

import (
	"github.com/deronyan-llc/dwim/internal/common"
	"github.com/deronyan-llc/rdf/rdf"
)

// combine all the imports and types from all the classes together
type GoStructs struct {
	Package string
	Imports []string
	Classes []*common.SchemaClass
}

var (
	SourceMap = make(map[rdf.Term]*GoStructs)
)
