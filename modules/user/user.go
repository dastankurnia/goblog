package user

import (
	"net/http"

	"github.com/dastankurnia/goblog/common"
)

// User : User struct
type User struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

// Address : User address
type Address struct {
	Street string `json:"street"`
	Region string `json:"region"`
}

// Login : Login user
func Login(w http.ResponseWriter, r *http.Request) {
	// TODO : login logic
	common.Response(w, r, 200, "Success", []string{})
}

// Register : Register User
func Register(w http.ResponseWriter, r *http.Request) {
	// TODO : Register Logic
	common.Response(w, r, 200, "Success", []string{})
}

// GetUsers : Get list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "Dastan",
		Email: "dastankurnia@gmail.com",
		Phone: "08245746232",
		Address: Address{
			Street: "Jl Timbul Jaya No.34 RT01 RW04",
			Region: "DKI Jakarta",
		},
	}
	common.Response(w, r, 200, "Success", user)
}
