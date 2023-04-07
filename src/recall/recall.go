package recall

import "common"

type Recaller interface {
	Recall(n int) []*common.Product
	Name() string
}
