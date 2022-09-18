package services

import (
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/controllers/views"
	"sesi6-gin/httpserver/repositories"
	"sesi6-gin/httpserver/repositories/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *params.UserCreateRequest) *views.Response {
	// step : (4) buat model
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = 1
	user.Password = string(hash)
	user.Username = req.Username

	// step : (5) kirim ke repositories
	// err = repositories.CreateUser(&model)
	repositories.Users = append(repositories.Users, user)

	// step : (7) buat sebuah views
	v := views.SuccessCreateResponse(user, "created success!")

	// step : (8) kembalikan views ke controller
	return v
}

func GetUsers() *views.Response{
	user := repositories.Users
	return views.SuccessRequest(user, "ok")
}
