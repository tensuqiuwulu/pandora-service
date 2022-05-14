package controller

import (
	"fmt"
	"time"

	"github.com/tensuqiuwulu/pandora-service/service"
)

type OrderControllerInterface interface {
	ProsesPembayaranViaVa()
	ProsesCompletedOrder()
	ProsesPembatalanOrder()
}

type OrderControllerImplementation struct {
	OrderServiceInterface service.OrderServiceInterface
}

func NewOrderController(orderServiceInterface service.OrderServiceInterface) OrderControllerInterface {
	return &OrderControllerImplementation{
		OrderServiceInterface: orderServiceInterface,
	}
}

func (controller *OrderControllerImplementation) ProsesPembayaranViaVa() {
	fmt.Println("Start Proses Pembayaran = ", time.Now())
	controller.OrderServiceInterface.ProsesPembayaranViaVa()
}

func (controller *OrderControllerImplementation) ProsesCompletedOrder() {
	fmt.Println("Start Proses Penyelesaian = ", time.Now())
	controller.OrderServiceInterface.ProsesCompletedOrder()
}

func (controller *OrderControllerImplementation) ProsesPembatalanOrder() {
	fmt.Println("Start Proses Pembatalan = ", time.Now())
	controller.OrderServiceInterface.ProsesPembatalanOrder()
}
