package generators

import "github.com/deronyan-llc/columbo/internal/common"

type SrcGenerator interface {
	Generate(schemaContext *common.SchemaContext) error
}
