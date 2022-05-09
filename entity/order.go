package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Order struct {
	Id                      string    `gorm:"primaryKey;column:id;"`
	NumberOrder             string    `gorm:"column:number_order;"`
	TrxId                   int       `gorm:"column:trx_id;"`
	IdUser                  string    `gorm:"column:id_user;"`
	FullName                string    `gorm:"column:full_name;"`
	Email                   string    `gorm:"column:email;"`
	Address                 string    `gorm:"column:address;"`
	Phone                   string    `gorm:"column:phone;"`
	CourierNote             string    `gorm:"column:courier_note;"`
	TotalBillBeforeDiscount float64   `gorm:"column:total_bill_before_discount;"`
	TotalBillAfterDiscount  float64   `gorm:"column:total_bill_after_discount;"`
	TotalBill               float64   `gorm:"column:total_bill;"`
	OrderSatus              string    `gorm:"column:order_status;"`
	OrderedAt               time.Time `gorm:"column:ordered_at;"`
	PaymentMethod           string    `gorm:"column:payment_method;"`
	PaymentChannel          string    `gorm:"column:payment_channel;"`
	PaymentStatus           string    `gorm:"column:payment_status;"`
	PaymentNo               string    `gorm:"column:payment_no;"`
	PaymentName             string    `gorm:"column:payment_name;"`
	PaymentByPoint          float64   `gorm:"column:payment_by_point;"`
	PaymentByCash           float64   `gorm:"column:payment_by_cash;"`
	ShippingMethod          string    `gorm:"column:shipping_method;"`
	ShippingCost            float64   `gorm:"column:shipping_cost;"`
	ShippingStatus          string    `gorm:"column:shipping_status;"`
	PaymentDueDate          null.Time `gorm:"column:payment_due_date;"`
	PaymentSuccessAt        null.Time `gorm:"column:payment_success_at;"`
	ProcessingDueDate       null.Time `gorm:"column:processing_due_date;"`
	ProcessedAt             null.Time `gorm:"column:processed_at;"`
	DeliveryDueDate         null.Time `gorm:"column:delivery_due_date;"`
	DeliveredAt             null.Time `gorm:"column:delivered_at;"`
	CompleteDueDate         null.Time `gorm:"column:complete_due_date;"`
	CanceledAt              null.Time `gorm:"column:delivered_at;"`
	CompletedAt             null.Time `gorm:"column:completed_at;"`
}

func (Order) TableName() string {
	return "orders_transaction"
}
