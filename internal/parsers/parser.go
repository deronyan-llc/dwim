package parsers

import "github.com/deronyan-llc/dwim/internal/common"

type Parser interface {
	Parse(file string) (*common.SchemaContext, error)
}
