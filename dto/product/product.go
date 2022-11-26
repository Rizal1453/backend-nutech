package productdto

type ProductRequest struct {
	Name       string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Image      string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Buy        int    `json:"buy" form:"buy" gorm:"type: int"`
	Sale       int    `json:"sale" form:"sale" gorm:"type: int"`
	Qty        int    `json:"qty" form:"qty" gorm:"type: int"`
	
}

type UpdateProductRequest struct {
	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Buy   int    `json:"buy" form:"buy" gorm:"type: int"`
	Sale  int    `json:"sale" form:"sale" gorm:"type: int"`
	Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
}