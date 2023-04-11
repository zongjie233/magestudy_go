package filter

import "common"

type Filter interface {
	Filter([]*common.Product) []*common.Product
	Name() string
}
