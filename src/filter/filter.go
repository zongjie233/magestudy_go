package filter

import "common"

type Filter interface {
	Filter([]*common.Product)
	Name() string
}
