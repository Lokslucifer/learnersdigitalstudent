package service

type Payment_Interface interface {
	Pay(clientId string,merchantId string,amount float64) (string, error)
	Refund(transactionID string) (string, error)
	GetProviderName() string
	AddTransaction(t Transaction) string
	FindTransaction(transactionId string) (Transaction, error)
}

type Transaction struct {
	TransactionId string
	Amount        float64
	ClientId      string
	MerchantID    string
	Provider      string
}
