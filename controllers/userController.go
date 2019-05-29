package controllers

import (
	"net/http"
	"time"

	"github.com/ajangi/golang-rest-api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// RegisterUserBody : this struct type is to bind register user request body
type RegisterUserBody struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required,len=11"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}

var validate *validator.Validate

// RegisterUser : this method is to register users by phone number and password and name
func RegisterUser(c echo.Context) (err error) {
	body := new(RegisterUserBody)
	// this part is to check if the body is empty or a valid json
	if err = c.Bind(body); err != nil {
		emptyData := utils.ResponseData{}
		badRequestMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.InputErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: badRequestMessage, StatusCode: http.StatusBadRequest}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	validate = validator.New()
	err = validate.Struct(body)
	// this part is to check the request body validation
	if err != nil {
		emptyData := utils.ResponseData{}
		badRequestMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.InputErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: badRequestMessage, StatusCode: http.StatusBadRequest}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	// in this part we can get valid data to register user in database and return jwt token
	name := body.Name
	email := body.Email
	phone := body.Phone
	rawPassword := body.Password
	hashedPassword, _ := HashPassword(rawPassword)
	db := utils.DbConn()
	insUser, err := db.Prepare("INSERT INTO users(name, email,phone,password) VALUES(?,?,?,?)")
	if err != nil {
		// todo : check this error here
		panic(err.Error())
	}
	insUser.Exec(name, email, phone, hashedPassword)
	// Create token for the registered user
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 3000).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		// TODO : check this error here
		return err
	}
	emptyData := utils.ResponseData{
		map[string]string{"token": t},
	}
	registerMessage := utils.ResponseMessages{}
	registerResponse := utils.ResponseApi{Result: "SUCCESS", Data: emptyData, Messages: registerMessage, StatusCode: http.StatusOK}
	return c.JSON(http.StatusOK, registerResponse)
}

// HashPassword : this method is to make password hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash : this method is to check password and hashed one
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
