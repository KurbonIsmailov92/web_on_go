package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	AvgGrates, Happiness float64
}

func (u *User) GetAllInfo() string {
	return fmt.Sprintf("Username is: %s. He is %d and he has %d mooney", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newName string) {
	u.Name = newName
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	Bob := User{Name: "Bob", Age: 25, Money: -50, AvgGrates: 4.2, Happiness: 0.8}
	Bob.SetNewName("Alex")
	fmt.Fprintf(w, Bob.GetAllInfo())
	//fmt.Fprintf(w, "Go is super easy")
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Cotact`s page")
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", Contacts)
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleRequest()
}
