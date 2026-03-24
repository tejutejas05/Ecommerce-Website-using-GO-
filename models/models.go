package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//USER Collection (like tables)

type User struct{

	ID							primitive.ObjectID
	First_Name					*string
	Last_Name					*string
	Password					*string
	Email						*string
	Phone						*string
	Token						*string
	Refresh_token				*string
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
}

type Product struct{

	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type ProductUser struct{

	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type Adress struct{
	Address_id
	House
	Street
	City
	Pincode

}

type Order struct{
	Order_ID
	Order_Cart
	Ordered_At
	Price
	Discount
	Payment_Method


}

type Payment struct{

	Digital bool
	COD bool


}