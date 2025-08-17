package models

type Cart struct {
	User      User `gorm:"foreignKey:UserID"`
	UserID    uint
	Products  Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Quantity  uint
}
