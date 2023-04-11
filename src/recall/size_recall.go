package recall

import "common"

type SizeRecall struct {
	Tag string
}

func (s SizeRecall) Name() string {
	return s.Tag
}

func (s SizeRecall) Recall(n int) []*common.Product {
	rect := make([]*common.Product, 0, n)
	for _, ele := range allProducts {
		if ele.Size < 200 {
			rect = append(rect, ele)
			if len(rect) >= n {
				break
			}
		}
	}
	return rect
}
