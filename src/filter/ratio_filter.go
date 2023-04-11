package filter

import "common"

type RatioFilter struct {
	Tag string
}

func (r RatioFilter) Name() string {
	return r.Tag
}

func (r RatioFilter) Filter(products []*common.Product) []*common.Product {
	rect := make([]*common.Product, 0, len(products))
	for _, product := range products {
		if product.RatioCount > 10 && product.PositiveRatio > 0.8 {
			rect = append(rect, product)
		}
	}
	return rect
}
