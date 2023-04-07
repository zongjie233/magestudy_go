package sortt

import "common"

type Sorter interface {
	Sort([]*common.Product) []*common.Product
	Name() string
}
