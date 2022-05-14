package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/tensuqiuwulu/pandora-service/config"
	"github.com/tensuqiuwulu/pandora-service/controller"
	"github.com/tensuqiuwulu/pandora-service/repository/mysql"
	"github.com/tensuqiuwulu/pandora-service/service"
	"github.com/tensuqiuwulu/pandora-service/utilities"
)

func main() {
	appConfig := config.GetConfig()

	mysqlDBConnection := mysql.NewDatabaseConnection(&appConfig.Database)

	// Timezone
	location, err := time.LoadLocation(appConfig.Timezone.Timezone)
	time.Local = location
	fmt.Println("location:", location, err)

	// Logger
	logrusLogger := utilities.NewLogger(appConfig.Log)

	scheduler := cron.New(cron.WithLocation(location), cron.WithLogger(cron.DefaultLogger))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// Repository
	orderRepository := mysql.NewOrderRepository(&appConfig.Database)

	// Service
	orderService := service.NewOrderService(mysqlDBConnection, logrusLogger, appConfig.Payment, orderRepository)

	// Controler
	orderController := controller.NewOrderController(orderService)
	// orderController.FindOrderByOrderStatus("Menunggu Pembayaran")

	scheduler.AddFunc("*/1 * * * *", func() { orderController.ProsesPembayaranViaVa() })

	scheduler.AddFunc("*/5 * * * *", func() { orderController.ProsesCompletedOrder() })

	scheduler.AddFunc("*/3 * * * *", func() { orderController.ProsesPembatalanOrder() })

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
