package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	u "sexsquare-api/utils"
	"strings"
)

/*
JWT claims struct
*/
type Token struct {
	ID uint
	jwt.StandardClaims
}

type AccountSettings struct {
	DefaultSex string `json:"sex,omitempty"`
}

type Achievements struct {
	*Achievement
	Date string `json:"date,omitempty"`
}

//a struct to rep user account
type Account struct {
	gorm.Model
	Nickname        string           `json:"Nickname,omitempty"`
	Email           string           `json:"Email,omitempty"`
	Password        string           `json:"Password,omitempty"`
	Gender          *Gender          `json:"Gender,omitempty"`
	CurrentPartners []*Partner       `json:"CurrentPartners,omitempty"`
	PartnersHistory []*Partner       `json:"PartnersHistory,omitempty"`
	Achievements    []*Achievements  `json:"Achievements,omitempty"`
	Settings        *AccountSettings `json:"Settings,omitempty"`
	Friends         []*Account		 `json:"Friends,omitempty"`
	Token           string           `json:"Token",omitempty;sql:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if account.Nickname == "" {
		return u.Message(false, "Nickname is required"), false
	}

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	errEmail := GetDB().Table("accounts").Where("Email = ?", account.Email).First(temp).Error
	errNickname := GetDB().Table("accounts").Where("Nickname = ?", account.Nickname).First(temp).Error
	if errEmail != nil && errNickname != nil && errEmail != gorm.ErrRecordNotFound && errNickname != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Nickname != "" {
		return u.Message(false, "Nickname already in use by another user."), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{ID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{ID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
