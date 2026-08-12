package orders

var catalog = map[string]int{
	"keyboard": 7500,
	"monitor":  24000,
	"mouse":    3500,
}

func Price(productID string) (int, bool) {
	price, ok := catalog[productID]
	return price, ok
}
