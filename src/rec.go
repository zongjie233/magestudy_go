package main

import (
	"common"
	"filter"
	"fmt"
	"log"
	"recall"
	"sortt"
	"time"
)

type Recommender struct {
	Recallers []recall.Recaller
	Sorter    sortt.Sorter
	Filters   []filter.Filter
}

func (rec *Recommender) Rec() []*common.Product {
	RecallMap := make(map[int]*common.Product, 100)
	// 顺序遍历每一路召回
	for _, recaller := range rec.Recallers {
		begin := time.Now()
		products := recaller.Recall(10)
		log.Printf("召回%s耗时%dns,找回了%d个商品\n", recaller.Name(), time.Since(begin).Nanoseconds(), len(products))
		for _, product := range products {
			RecallMap[product.Id] = product
		}
	}
	log.Printf("一共召回了%d个商品\n", len(RecallMap))
	// 把召回结果进行排序
	RecallSlice := make([]*common.Product, 0, len(RecallMap))
	for _, product := range RecallMap {
		RecallSlice = append(RecallSlice, product)
	}

	// 对召回结果进行排序
	begin := time.Now()
	SortedResult := rec.Sorter.Sort(RecallSlice)
	log.Printf("排序耗时%dns\n", time.Since(begin).Nanoseconds())
	println(time.Since(begin).Nanoseconds())

	// 顺序选择多个过滤规则
	FilteredResult := SortedResult
	prevCount := len(FilteredResult)
	for _, filter := range rec.Filters {
		begin := time.Now()
		FilteredResult = filter.Filter(FilteredResult)
		log.Printf("过滤规则%s耗时%dns,过滤掉了%d个商品", filter.Name(), time.Since(begin).Nanoseconds(), prevCount-len(FilteredResult))
		prevCount = len(FilteredResult)
	}

	return FilteredResult
}
func main() {
	recer := Recommender{
		Recallers: []recall.Recaller{
			recall.HotRecall{Tag: "hot"},
			recall.SizeRecall{Tag: "size"},
		},
		Sorter: sortt.RationSorter{Tag: "ratio"},
		Filters: []filter.Filter{
			filter.RatioFilter{Tag: "ratio"},
		},
	}
	result := recer.Rec()

	for i, p := range result {
		fmt.Printf("第%d名：%d %s\n", i, p.Id, p.Name)
	}

}
