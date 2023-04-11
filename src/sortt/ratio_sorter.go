package sortt

import (
	"common"
	"sort"
)

type RationSorter struct {
	Tag string
}

func (r RationSorter) Name() string {
	return r.Tag
}

func (r RationSorter) Sort(products []*common.Product) []*common.Product { // 原地排序
	sort.Slice(products, func(i, j int) bool {
		return products[i].PositiveRatio > products[j].PositiveRatio
	})
	return products
}
