package orders_entity

type BindCreateOrder struct {
	ProductId    []int `json:"product_id"`
	AddressId    int   `json:"address_id"`
	SummaryPrice int   `json:"summary_price"`
	UserId       int   `json:"user_id"`
}
