package models

import "github.com/aysmdb/ojire-casetiga/pkg/database"

type Cart struct {
	User      User `gorm:"foreignKey:UserID"`
	UserID    uint
	Products  Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Quantity  uint
}

type CartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func AddToCart(cart Cart) error {
	db := database.DBConn
	if err := db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}

func GetCartByUserID(userID uint) ([]Cart, error) {
	db := database.DBConn
	var cartItems []Cart
	if err := db.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
