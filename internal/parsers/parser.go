package parsers

import "deronyan.com/columbo/internal/common"

type Parser interface {
	Parse(file string) (*common.SchemaContext, error)
}
