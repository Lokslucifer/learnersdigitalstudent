package service

import (
	"errors"
	"fmt"
)

type Payment_Gateway struct {
	Payment_Providers []Payment_Interface
}

func New_Payment_Gateway() Payment_Gateway {
	return Payment_Gateway{
		Payment_Providers: []Payment_Interface{},
	}
}

func (p *Payment_Gateway) Register_Provider(provider Payment_Interface) {
	p.Payment_Providers = append(p.Payment_Providers, provider)

}

func (p *Payment_Gateway) ProcessPayment(providerName string, clientId string, merchantId string, amount float64) (string, error) {
	for _, provider := range p.Payment_Providers {
		if provider.GetProviderName() == providerName {
			transid, err := provider.Pay(clientId, merchantId, amount)
			return transid, err

		}

	}
	return "", errors.New("cant find provider")

}

func (p *Payment_Gateway) IssueRefund(providerName string, transactionId string) (string, error) {
	fmt.Println("Refund request recieved-", providerName, "-", transactionId)
	for _, provider := range p.Payment_Providers {
		if provider.GetProviderName() == providerName {
			transid, err := provider.Refund(transactionId)
			return transid, err

		}

	}
	return "", errors.New("cant find provider")

}

func Extract_details(prov Payment_Interface) error {
	if prov.GetProviderName() == "paypal" {
		provdata, ok := prov.(*Paypal_payment_service)
		if ok {
			fmt.Println("API key:", provdata.APIKey)
			fmt.Println("Transaction list:", provdata.Transactionlst)
			return nil
		}
		return errors.New("error during type assertion")

	}
	if prov.GetProviderName() == "stripe" {
		provdata, ok := prov.(*Stripe_payment_service)
		if ok {
			fmt.Println("API key:", provdata.APIKey)
			fmt.Println("Transaction list:", provdata.Transactionlst)

			return nil
		}
		return errors.New("error during type assertion")
	}
	if prov.GetProviderName() == "razorpay" {
		provdata, ok := prov.(*Razorpay_payment_service)
		if ok {
			fmt.Println("API key:", provdata.APIKey)
			fmt.Println("Transaction list:", provdata.Transactionlst)

			return nil
		}
		return errors.New("error during type assertion")
	}
	return errors.New("cant find provider")

}

func (p *Payment_Gateway) Display_Provider_Details() {

	for _, provider := range p.Payment_Providers {
		err := Extract_details(provider)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
