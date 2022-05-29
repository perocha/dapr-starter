package order

type Order struct {
	OrderId    string `json:"orderId"`
	OrderName  string `json:"orderName"`
	OrderPrice string `json:"orderPrice"`
}
