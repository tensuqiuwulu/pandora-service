package service

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/pandora-service/entity"
	"github.com/tensuqiuwulu/pandora-service/repository/mysql"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	UpdateUserNotVerification()
}

type UserServiceImplementation struct {
	DB                      *gorm.DB
	Logger                  *logrus.Logger
	UserRepositoryInterface mysql.UserRepositoryInterface
}

func NewUserService(
	DB *gorm.DB,
	logger *logrus.Logger,
	userRepositoryInterface mysql.UserRepositoryInterface) UserServiceInterface {
	return &UserServiceImplementation{
		DB:                      DB,
		Logger:                  logger,
		UserRepositoryInterface: userRepositoryInterface,
	}
}

func (service *UserServiceImplementation) UpdateUserNotVerification() {
	// Get data user yng belum melakukan verifikasi
	users, _ := service.UserRepositoryInterface.FindNotActiveUser(service.DB)

	if len(users) > 0 {
		for _, user := range users {
			waktuSekarang := time.Now()
			waktu := user.VerificationDueDate

			if waktuSekarang.After(waktu) {
				userEntity := &entity.User{}
				userEntity.NotVerification = 1
				userEntity.NotVerificationDate = waktuSekarang
				_, err := service.UserRepositoryInterface.UpdateUserNotVerification(service.DB, user.Id, *userEntity)
				if err != nil {
					fmt.Println("Error saat update not verification id = ", user.Id)
				} else {
					fmt.Println("Sukses update not verification id = ", user.Id)
				}
			} else {
				fmt.Println("Belum bisa melakukan proses not verification = ", user.Id)
			}

		}
	} else {
		fmt.Println("Belum ada user dalam status tidak aktif")
	}
}
