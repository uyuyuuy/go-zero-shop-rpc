package model

type Order struct {
	OrderId   string `db:"order_id" json:"order_id"`
	UserId    uint64 `db:"user_id" json:"user_id"`
	ProductId uint64 `db:"product_id" json:"product_id"`
	Number    uint64 `db:"number" json:"number"`
	Status    uint8  `db:"status" json:"status"`
}
