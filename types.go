package main

type FixerResponse struct {
	Success bool
	Timestamp int
	Base string
	Rates map[string]float64
}

type ConverterRequest struct{
	inputCurrency string
	outputCurrency string
	amount float64
}

type NotImplemented struct{}