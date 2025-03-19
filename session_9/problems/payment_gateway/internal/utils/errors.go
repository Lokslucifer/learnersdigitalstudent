package utils

type PaypalPaymentError struct {
	Msg string
}

func NewPayalPaymentError(msg string) PaypalPaymentError {
	return PaypalPaymentError{Msg: msg}
}

func (err PaypalPaymentError) Error() string {
	return err.Msg

}

type PaypalTransactionError struct {
	Msg string
}

func NewPayalTransactionError(msg string) PaypalTransactionError {
	return PaypalTransactionError{Msg: msg}
}

func (err PaypalTransactionError) Error() string {
	return err.Msg

}

type StripePaymentError struct {
	Msg string
}

func NewStripePaymentError(msg string) StripePaymentError {
	return StripePaymentError{Msg: msg}
}

func (err StripePaymentError) Error() string {
	return err.Msg

}

type StripeTransactionError struct {
	Msg string
}

func NewStripeTransactionError(msg string) StripeTransactionError {
	return StripeTransactionError{Msg: msg}
}
func (err StripeTransactionError) Error() string {
	return err.Msg

}

type RazorpayPaymentError struct {
	Msg string
}

func NewRazorpayPaymentError(msg string) RazorpayPaymentError {
	return RazorpayPaymentError{Msg: msg}
}

func (err RazorpayPaymentError) Error() string {
	return err.Msg

}

type RazorpayTransactionError struct {
	Msg string
}

func NewRazorpayTransactionError(msg string) RazorpayTransactionError {
	return RazorpayTransactionError{Msg: msg}
}
func (err RazorpayTransactionError) Error() string {
	return err.Msg

}
