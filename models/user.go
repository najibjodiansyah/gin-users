package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	PinCode string `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Email   string  `json:"email" bson:"email"`
	Phone   string  `json:"phone" bson:"phone"`
	Address Address `json:"address" bson:"address"`
}
