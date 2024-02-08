package generators

import "github.com/deronyan-llc/dwim/internal/common"

type SrcGenerator interface {
	Generate(schemaContext *common.SchemaContext) error
}
