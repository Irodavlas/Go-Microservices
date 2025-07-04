package handler

import (
	"fmt"
	"net/http"

	response "github.com/irodavlas/common-response"
	"github.com/irodavlas/user-service/database"
	"github.com/irodavlas/user-service/model"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	Database *database.Database
	Sender   *response.Sender
}
type IUserService interface {
	CreateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUser(c echo.Context) error
	//...
}

func NewUserService(db *database.Database) *UserService {
	return &UserService{
		Database: db,
		Sender:   response.NewSender(),
	}
}

func (service *UserService) CreateUser(c echo.Context) error {
	User := &model.User{}

	if err := c.Bind(User); err != nil {
		return service.Sender.Error(c, http.StatusBadRequest, "Error binding body of request to User struct", err.Error())
	}
	fmt.Println("Username:", User.Username)
	get_user, err := service.Database.Get(*User)
	if err != nil {
		return service.Sender.Error(c, http.StatusBadRequest, "User doesn't exists", err.Error())
	}
	if get_user != nil {
		return service.Sender.Error(c, http.StatusBadRequest, "Username not available", "")
	}
	err = service.Database.Create(User)
	if err != nil {
		return service.Sender.Error(c, http.StatusInternalServerError, "Error Creating user in the database", err.Error())
	}

	return service.Sender.Success(c, http.StatusOK, "Success", User)
}
func (service *UserService) DeleteUser(c echo.Context) error {
	usermame := c.QueryParam("username")
	if usermame == "" {
		return service.Sender.Error(c, http.StatusBadRequest, "No username provided", "")
	}
	err := service.Database.Delete(usermame)
	if err != nil {
		return service.Sender.Error(c, http.StatusInternalServerError, "Error deleting user from database", err.Error())
	}
	return service.Sender.Success(c, http.StatusOK, "Success", usermame)
}
func (service *UserService) UpdateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return service.Sender.Error(c, http.StatusBadRequest, "Error binding body of request to User struct", err.Error())
	}

	err := service.Database.Update(*user)
	if err != nil {
		return service.Sender.Error(c, http.StatusInternalServerError, "Error updating user in the Database", err.Error())
	}
	return service.Sender.Created(c, http.StatusCreated, "User Updated", user)
}
func (service *UserService) GetUser(c echo.Context) error {
	user := &model.User{}

	user.Username = c.QueryParam("username")
	if user.Username == "" {
		return service.Sender.Error(c, http.StatusBadRequest, "Username cannot be empty", "")
	}

	databaseUser, err := service.Database.Get(*user)
	if err != nil {
		return service.Sender.Error(c, http.StatusInternalServerError, "Error retrieving the user", "")
	}
	if databaseUser == nil {
		//ig should redirect to sign in page or tell the user
		return model.Failed(c, http.StatusNotFound, "User not found", fmt.Sprintf("User not found:%s", user.Username))
	}

	return model.Success(c, http.StatusOK, "Success", *databaseUser)
}

//define all the other functions
