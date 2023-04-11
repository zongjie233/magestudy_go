package common

type Product struct {
	Id            int
	Name          string
	Size          int
	Sale          int //销量
	ShipAddress   string
	Price         float64
	PositiveRatio float64
	RatioCount    int // 评论量
}
