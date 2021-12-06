package coinbase

type WebsocketChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}
