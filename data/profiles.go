package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type User struct {
	UserId       int    `json:"id"`
	Nickname     string `json:"nick"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func GetProfiles() Users {
	return userList
}

func AddUser(u *User) {
	u.UserId = getId()
	userList = append(userList, u)
}

func getId() int {
	return userList[len(userList)-1].UserId + 1
}

var ErrorUserNotFound = fmt.Errorf("User not found")

func UpdateUser(id int, u *User) error {
	pos, err := findUser(id)
	if err != nil {
		return err
	}

	userList[pos] = u
	userList[pos].UserId = id

	return nil
}

func findUser(id int) (int, error) {
	for i, v := range userList {
		if v.UserId == id {
			return i, nil
		}
	}
	return -1, ErrorUserNotFound
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

var userList = []*User{
	{
		UserId:       1,
		Nickname:     "Vasya",
		PasswordHash: "vdskaj3rbwerjbwi3",
		Email:        "vasya2007@gmail.com",
	},
	{
		UserId:       2,
		Nickname:     "Petya",
		PasswordHash: "fwe5349f34fm9f0aa",
		Email:        "petya2004@gmail.com",
	},
	{
		UserId:       3,
		Nickname:     "Sanya",
		PasswordHash: "fwefdskgfwerhdfio",
		Email:        "sanya2002@gmail.com",
	},
}
