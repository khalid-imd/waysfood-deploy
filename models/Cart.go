package models

type Cart struct {
	ID        int     `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int     `json:"user_id" gorm:"type:varchar(25)"`
	Status    string  `json:"status"  gorm:"type:varchar(25)"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
	Qty       int     `json:"qty" `
}
