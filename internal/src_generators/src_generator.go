package src_generators

import "deronyan.com/columbo/internal/common"

type SrcGenerator interface {
	Generate(schemaContext *common.SchemaContext) error
}
