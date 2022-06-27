package mysql

import (
	"github.com/tensuqiuwulu/pandora-service/config"
	"github.com/tensuqiuwulu/pandora-service/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	UpdateUserNotVerification(DB *gorm.DB, idUser string, user entity.User) (entity.User, error)
	FindNotActiveUser(DB *gorm.DB) ([]entity.User, error)
}

type UserRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewUserRepository(configDatabase *config.Database) UserRepositoryInterface {
	return &UserRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *UserRepositoryImplementation) UpdateUserNotVerification(DB *gorm.DB, idUser string, user entity.User) (entity.User, error) {
	result := DB.
		Model(entity.User{}).
		Where("id = ?", idUser).
		Updates(entity.User{
			NotVerification:     user.NotVerification,
			NotVerificationDate: user.NotVerificationDate,
		})
	return user, result.Error
}

func (repository *UserRepositoryImplementation) FindNotActiveUser(DB *gorm.DB) ([]entity.User, error) {
	var user []entity.User
	results := DB.Where("is_active = ?", 0).Where("not_verification = ?", 0).Find(&user)
	return user, results.Error
}
