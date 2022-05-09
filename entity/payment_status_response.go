package entity

type PaymentResponse struct {
	Status  int
	Message string
	Data    Data
}

type PaymentStatusResponse struct {
	Status  int
	Data    PaymentStatus
	Message string
}

type PaymentStatus struct {
	TransactionId  int
	SessionId      string
	ReferenceId    string
	RelatedId      int
	Sender         string
	Recevier       string
	Amount         string
	Fee            string
	Status         int
	StatusDesc     string
	Type           int
	TypeDesc       string
	Notes          string
	CreatedDate    string
	ExpiredDate    string
	SuccessDate    string
	SettlementDate string
}

type Data struct {
	SessionId     string
	TransactionId int
	ReferenceId   string
	Via           string
	Channel       string
	PaymentNo     string
	PaymentName   string
	Total         float64
	Fee           float64
	Expired       string
}

type ListPaymentChannelPayment struct {
	PaymentMethod string
	BankCode      string
	BankName      string
	BankLogo      string
}
