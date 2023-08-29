package model

type Product struct {
	Id     uint64 `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Number uint64 `db:"number" json:"number"`
}
