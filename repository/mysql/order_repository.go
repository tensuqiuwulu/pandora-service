package mysql

import (
	"github.com/tensuqiuwulu/pandora-service/config"
	"github.com/tensuqiuwulu/pandora-service/entity"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	FindOrderByOrderStatusVa(DB *gorm.DB, orderStatus string) ([]entity.Order, error)
	FindOrderByOrderStatus(DB *gorm.DB, orderStatus string) ([]entity.Order, error)
}

type OrderRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderRepository(configDatabase *config.Database) OrderRepositoryInterface {
	return &OrderRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderRepositoryImplementation) FindOrderByOrderStatusVa(DB *gorm.DB, orderStatus string) ([]entity.Order, error) {
	var order []entity.Order
	results := DB.Where("order_status = ?", orderStatus).Where("payment_method = ?", "va").Or("payment_method = ?", "qris").Or("payment_method = ?", "cc").Find(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) FindOrderByOrderStatus(DB *gorm.DB, orderStatus string) ([]entity.Order, error) {
	var order []entity.Order
	results := DB.Where("order_status = ?", orderStatus).Find(&order)
	return order, results.Error
}
