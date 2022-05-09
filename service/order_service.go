package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/pandora-service/config"
	"github.com/tensuqiuwulu/pandora-service/entity"
	"github.com/tensuqiuwulu/pandora-service/repository/mysql"
	"gorm.io/gorm"
)

type OrderServiceInterface interface {
	ProsesPembayaranViaVa() error
	ProsesCompletedOrder() error
	ProsesPembatalanOrder() error
}

type OrderServiceImplementation struct {
	DB                       *gorm.DB
	Logger                   *logrus.Logger
	ConfigPayment            config.Payment
	OrderRepositoryInterface mysql.OrderRepositoryInterface
}

func NewOrderService(
	DB *gorm.DB,
	logger *logrus.Logger,
	configPayment config.Payment,
	orderRepositoryInterface mysql.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderServiceImplementation{
		DB:                       DB,
		Logger:                   logger,
		ConfigPayment:            configPayment,
		OrderRepositoryInterface: orderRepositoryInterface,
	}
}

func (service *OrderServiceImplementation) ProsesCompletedOrder() error {
	orders, err := service.OrderRepositoryInterface.FindOrderByOrderStatus(service.DB, "Sampai Di Tujuan")

	fmt.Println("Completed Url = ", service.ConfigPayment.ApiCompleted)

	for _, order := range orders {
		waktuSekarang := time.Now()
		waktu := order.CompleteDueDate.Time
		// fmt.Println("waktu sekarang = ", waktuSekarang)
		// fmt.Println("batas waktu penyelesaian = ", waktu)
		if waktuSekarang.After(waktu) {
			url, _ := url.Parse(service.ConfigPayment.ApiCompleted + order.Id)

			req := &http.Request{
				Method: "PUT",
				URL:    url,
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {"Bearer " + service.ConfigPayment.SecretToken},
				},
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()

			ss := fmt.Sprintf("%+v\n", resp)
			fmt.Println(ss)

			fmt.Println("penyelesaian orderan sukses = ", order.NumberOrder, " Nama = ", order.FullName)
		} else {
			fmt.Println("penyelesaian orderan belum bisa dilakukan = ", order.NumberOrder, " Nama = ", order.FullName)
		}
	}
	return err
}

func (service *OrderServiceImplementation) ProsesPembayaranViaVa() error {
	orders, err := service.OrderRepositoryInterface.FindOrderByOrderStatus(service.DB, "Menunggu Pembayaran")

	// fmt.Println("Pembayaran Url = ", service.ConfigPayment.ApiCompleted)

	for _, order := range orders {
		// cek status pembayaran ke ipaymu
		var ipaymu_va = string(service.ConfigPayment.IpaymuVa)
		var ipaymu_key = string(service.ConfigPayment.IpaymuKey)

		url, _ := url.Parse(string(service.ConfigPayment.IpaymuTranscationUrl))
		postBody, _ := json.Marshal(map[string]interface{}{
			"transactionId": order.TrxId,
		})

		bodyHash := sha256.Sum256([]byte(postBody))
		bodyHashToString := hex.EncodeToString(bodyHash[:])
		stringToSign := "POST:" + ipaymu_va + ":" + strings.ToLower(string(bodyHashToString)) + ":" + ipaymu_key

		h := hmac.New(sha256.New, []byte(ipaymu_key))
		h.Write([]byte(stringToSign))
		signature := hex.EncodeToString(h.Sum(nil))

		reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

		req := &http.Request{
			Method: "POST",
			URL:    url,
			Header: map[string][]string{
				"Content-Type": {"application/json"},
				"va":           {ipaymu_va},
				"signature":    {signature},
			},
			Body: reqBody,
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

		var dataPaymentStatus entity.PaymentStatusResponse

		if err := json.NewDecoder(resp.Body).Decode(&dataPaymentStatus); err != nil {
			fmt.Println(err)
		}

		if dataPaymentStatus.Data.Status == 1 || dataPaymentStatus.Data.Status == 6 {
			url, _ := url.Parse(service.ConfigPayment.ApiPembayaran)
			// fmt.Println("Proses Pembayaran Url = ", url)
			postBody, _ := json.Marshal(map[string]interface{}{
				"trx_id":       order.TrxId,
				"status":       "berhasil",
				"status_code":  dataPaymentStatus.Data.Status,
				"reference_id": order.NumberOrder,
			})

			reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

			req := &http.Request{
				Method: "POST",
				URL:    url,
				Header: map[string][]string{
					"Content-Type": {"application/json"},
					// "Authorization": {"Bearer " + service.ConfigPayment.SecretToken},
				},
				Body: reqBody,
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()

			ss := fmt.Sprintf("%+v\n", resp)
			fmt.Println(ss)

			fmt.Println("pembayaran sukses = ", order.NumberOrder, " Nama = ", order.FullName)

		} else {
			fmt.Println("gagal proses belum bayar order number = ", order.NumberOrder, " Nama = ", order.FullName)
		}

	}

	return err
}

func (service *OrderServiceImplementation) ProsesPembatalanOrder() error {
	orders, err := service.OrderRepositoryInterface.FindOrderByOrderStatus(service.DB, "Menunggu Pembayaran")

	// fmt.Println("Pembatalan Url = ", service.ConfigPayment.ApiCancel)

	for _, order := range orders {
		waktuSekarang := time.Now()
		waktu := order.PaymentDueDate.Time
		fmt.Println("waktu sekarang = ", waktuSekarang)
		fmt.Println("batas waktu penyelesauan = ", waktu)

		if waktuSekarang.After(waktu) {
			fmt.Println("token = ", service.ConfigPayment.SecretToken)
			url, _ := url.Parse(service.ConfigPayment.ApiCancel + order.Id)
			fmt.Println("Pembatalan Url = ", url)
			req := &http.Request{
				Method: "PUT",
				URL:    url,
				Header: map[string][]string{
					"Content-Type":  {"application/json"},
					"Authorization": {"Bearer " + service.ConfigPayment.SecretToken},
				},
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()

			ss := fmt.Sprintf("%+v\n", resp)
			fmt.Println(ss)

			fmt.Println("pembatalan orderan sukses = ", order.NumberOrder, " Nama = ", order.FullName)
		} else {
			fmt.Println("pembatalan orderan belum bisa dilakukan = ", order.NumberOrder, " Nama = ", order.FullName)
		}
	}

	return err
}
