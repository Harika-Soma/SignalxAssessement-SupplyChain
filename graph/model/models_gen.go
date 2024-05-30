// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type InventoryItem struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Sku       string `json:"sku" bson:"sku"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	Warehouse string `json:"warehouse" bson:"warehouse"`
}

type Mutation struct {
}

type Query struct {
}

type Shipment struct {
	ID                string `json:"id" bson:"_id"`
	Origin            string `json:"origin" bson:"origin"`
	Destination       string `json:"destination" bson:"destination"`
	Status            string `json:"status" bson:"status"`
	EstimatedDelivery string `json:"estimatedDelivery" bson:"estimatedDelivery"`
}

type Supplier struct {
	ID            string `json:"id" bson:"_id"`
	Name          string `json:"name" bson:"name"`
	ContactPerson string `json:"contactPerson" bson:"contactPerson"`
	Phone         string `json:"phone" bson:"phone"`
	Email         string `json:"email" bson:"email"`
}

type UserLoginData struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type LoginData struct {
	ID        string    `json:"id" bson:"_id"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	LoginTime time.Time `json:"loginTime" bson:"loginTime"`
}
