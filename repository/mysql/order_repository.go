package mysql

import (
	"github.com/tensuqiuwulu/pandora-service/config"
	"github.com/tensuqiuwulu/pandora-service/entity"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	FindOrderByOrderStatus(DB *gorm.DB, orderStatus string) ([]entity.Order, error)
	UpdateOrderStatus(DB *gorm.DB, numberOrder string, order entity.Order) (entity.Order, error)
}

type OrderRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderRepository(configDatabase *config.Database) OrderRepositoryInterface {
	return &OrderRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderRepositoryImplementation) FindOrderByOrderStatus(DB *gorm.DB, orderStatus string) ([]entity.Order, error) {
	var order []entity.Order
	results := DB.Where("order_status = ?", orderStatus).Where("payment_method = ?", "va").Find(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) UpdateOrderStatus(DB *gorm.DB, NumberOrder string, order entity.Order) (entity.Order, error) {
	// var test entity.Order
	// fmt.Println("waktu 2", test.PaymentSuccessAt.Time)

	result := DB.
		Model(entity.Order{}).
		Where("number_order = ?", NumberOrder).
		Updates(entity.Order{
			PaymentStatus:    order.PaymentStatus,
			OrderSatus:       order.OrderSatus,
			PaymentSuccessAt: order.PaymentSuccessAt,
			PaymentMethod:    order.PaymentMethod,
			PaymentChannel:   order.PaymentChannel,
		})
	return order, result.Error
}
