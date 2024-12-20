package sales

type ProductSelected struct {
	ProductId string
	Quantity  int
}

type Sale struct {
	Products   []ProductSelected
	TotalValue float32
}
