package User

import (
	"errors"
	"fmt"
)

type UserStrct struct {
	Firstname string
	Lastname  string
	Birthdate string
}


func Somethingcleaner(User1 UserStrct) {
	fmt.Println(User1.Firstname, User1.Lastname, User1.Birthdate)
}

func SomethingDoneWithPointer(User1 *UserStrct) {
	fmt.Println(User1.Firstname, User1.Lastname, User1.Birthdate)
}

func Something(Firstname, Lastname, Birthdate string) UserStrct {
	fmt.Println(Firstname, Lastname, Birthdate)
	return UserName(Firstname, Lastname, Birthdate)
}

func GetUserData(txt string) string {
	fmt.Print(txt)
	var value string
	fmt.Scan(&value)
	return value
}

func (u *UserStrct) RemoveUserName() { 
	// this is a method with u User and the data here is copy of the original data, So to change the values pass the pointers
	u.Firstname = ""
	u.Lastname = ""
}

func UserName(Firstname, Lastname, Birthdate string) UserStrct {
	return UserStrct{
		Firstname: Firstname,
		Lastname:  Lastname,
		Birthdate: Birthdate,
	}
}

func UserNamePointer(Firstname, Lastname, Birthdate string) *UserStrct {
	return &UserStrct{
		Firstname: Firstname,
		Lastname:  Lastname,
		Birthdate: Birthdate,
	}
}

func ValidatingUserCreated(Firstname, Lastname, Birthdate string) (*UserStrct, error) {
	var Uservar *UserStrct
	if Firstname == "" || Lastname == "" || Birthdate == "" {
		return Uservar, errors.New("empty error")
	}
	return &UserStrct{
		Firstname: Firstname,
		Lastname:  Lastname,
		Birthdate: Birthdate,
	}, nil
}
