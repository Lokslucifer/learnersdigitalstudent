package main

import (
	"payment_gateway/payment_gateway/internal/service"
)

func main() {
	paypal_provider := service.New_paypal_service()
	razor_provider := service.New_Razorpay_service()
	stripe_provider := service.New_Stripe_service()
	payment_manager := service.New_Payment_Gateway()
	payment_manager.Register_Provider(paypal_provider)
	payment_manager.Register_Provider(razor_provider)
	payment_manager.Register_Provider(stripe_provider)
	transid, err := payment_manager.ProcessPayment("paypal", "lokesh02", "lucy01", 100)

	if err == nil {

		payment_manager.IssueRefund("paypal", transid)

	}

}
