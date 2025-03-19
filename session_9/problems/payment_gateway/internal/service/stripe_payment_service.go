package service

import (
	"fmt"
	"payment_gateway/payment_gateway/internal/utils"
	"strconv"
)

type Stripe_payment_service struct {
	APIKey           string
	Transactionlst   []Transaction
	PayError         utils.StripePaymentError
	TransactionError utils.StripeTransactionError
}

func New_Stripe_service() Payment_Interface {
	return &Stripe_payment_service{APIKey: "STRIPE123", Transactionlst: make([]Transaction, 0), PayError: utils.NewStripePaymentError("Payment failed(Stripe)"), TransactionError: utils.NewStripeTransactionError("Transaction id not found(Stripe)")}

}

func (p *Stripe_payment_service) AddTransaction(t Transaction) string {
	t.TransactionId = "STRIPE" + strconv.Itoa(len(p.Transactionlst))
	p.Transactionlst = append(p.Transactionlst, t)
	return t.TransactionId
}

func (p *Stripe_payment_service) FindTransaction(transactionId string) (Transaction, error) {
	for _, trans := range p.Transactionlst {
		if trans.TransactionId == transactionId {
			return trans, nil
		}
	}
	return Transaction{}, p.TransactionError
}

func (p *Stripe_payment_service) Pay(clientId string, merchantId string, amount float64) (string, error) {

	fmt.Println("Payment of Rs.", amount, " using Stripe from ", clientId, "to ", merchantId, " is success")
	newtrans := Transaction{ClientId: clientId, MerchantID: merchantId, Amount: amount}
	transid := p.AddTransaction(newtrans)
	fmt.Println("Transaction Id:", transid)
	return transid, nil
}

func (p *Stripe_payment_service) Refund(transactionID string) (string, error) {
	trans, err := p.FindTransaction(transactionID)
	if err != nil {
		return "", err
	}
	fmt.Println("Refund initiated")
	transid, err := p.Pay(trans.MerchantID, trans.ClientId, trans.Amount)
	if err != nil {
		fmt.Println("Refund failed with error ", err)
		return "", err
	}
	fmt.Println("Refund succeed. Transaction Id:", transid)

	return transid, nil
}

func (p Stripe_payment_service) GetProviderName() string {
	fmt.Println("using Stripe payment provider")
	return "stripe"
}
