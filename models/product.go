package models

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name" gorm:"unique" `
	Image string `json:"image"`
	Buy   int    `json:"buy"`
	Sale  int    `json:"sale"`
	Qty   int    `json:"qty"`
}