package recall

import (
	"common"
	"sort"
)

var allProducts = []*common.Product{
	{Id: 1, Sale: 33, Name: "p1"},
	{Id: 2, Sale: 234, Name: "p2"},
	{Id: 3, Sale: 323, Name: "p3"},
	{Id: 4, Sale: 313, Name: "p4"},
	{Id: 5, Sale: 3533, Name: "p5"},
	{Id: 6, Sale: 333, Name: "p6"},
}

type HotRecall struct {
	Tag string
}

func (h HotRecall) Name() string {
	return h.Tag
}

func (h HotRecall) Recall(n int) []*common.Product {
	sort.Slice(allProducts, func(i, j int) bool { // 定义第一个与第二个元素排序的方法
		return allProducts[i].Sale > allProducts[j].Sale
	})
	rect := make([]*common.Product, 0, n)

	for _, product := range allProducts {
		rect = append(rect, product)
		if len(rect) >= n {
			break
		}
	}
	return rect
}
