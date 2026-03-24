package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//USER Collection (like tables)

type User struct{

	ID 
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_token
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
}

type Product struct{


}

type ProductUser struct{

}

type Adress struct{


}

type Order struct{


}

type Payment struct{

	
}