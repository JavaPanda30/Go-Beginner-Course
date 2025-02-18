package main

import (
	"fmt"

	"example.com/structs/User"
)

type aman string //custom data structure named aman

func main() {
	//custom data type
	var nameVal aman = "aman"
	nameVal.log()

	//sample data input from user
	firstName := User.GetUserData("Please enter first name: ")
	lastName := User.GetUserData("Please enter last name: ")
	birthdate := User.GetUserData("Please enter birthdate (MM/DD/YYYY): ")
	//Making Struct
	appUser := User.Something(firstName, lastName, birthdate)
	User.Somethingcleaner(appUser)
	User.SomethingDoneWithPointer(&appUser)
	(appUser).RemoveUserName()
	// direct calling struct
	var appUsercreated2 User.UserStrct = User.UserName(firstName, lastName, birthdate)
	User.Somethingcleaner(appUsercreated2)
	var appUsercreated3 *User.UserStrct = User.UserNamePointer(firstName, lastName, birthdate)
	User.Somethingcleaner(*appUsercreated3)
	//validation function
	var UserChkd *User.UserStrct
	UserChkd, err := User.ValidatingUserCreated(firstName, lastName, birthdate)
	if err == nil {
		User.Somethingcleaner(*UserChkd)
	}
}

func (name aman) log() {
	fmt.Println(name)
}
