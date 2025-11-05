package models

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Qty  int32  `json:"qty"`
}
