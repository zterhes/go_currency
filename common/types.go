package common

type FixerResponse struct {
	Success bool
	Timestamp int
	Base string
	Rates map[string]float64
}

type ConverterRequest struct{
	InputCurrency string
	OutputCurrency string
	Amount float64
}

type NotImplemented struct{}