package parsers

import "github.com/deronyan-llc/columbo/internal/common"

type Parser interface {
	Parse(file string) (*common.SchemaContext, error)
}
