package controller

import (
	"fmt"
	"time"

	"github.com/tensuqiuwulu/pandora-service/service"
)

type UserControllerInterface interface {
	ProsesUpdateUserNotVerification()
}

type UserControllerImplementation struct {
	UserServiceInterface service.UserServiceInterface
}

func NewUserController(userServiceInterface service.UserServiceInterface) UserControllerInterface {
	return &UserControllerImplementation{
		UserServiceInterface: userServiceInterface,
	}
}

func (controller *UserControllerImplementation) ProsesUpdateUserNotVerification() {
	fmt.Println("Start Proses Update Not Verification User = ", time.Now())
	controller.UserServiceInterface.UpdateUserNotVerification()
}
