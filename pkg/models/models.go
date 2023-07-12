package models

type Order struct {
	Id        int64 `json:"id" gorm:"primarykey;auto_increment"`
	Price     int64 `json:"price"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}
