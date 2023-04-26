package p2

import pkg1 "Goproject/p1"

func main() {
	// 不可以你依赖我，我依赖你
	pkg1.Sub(3, 2)
	//easyjson.Marshal()
}
