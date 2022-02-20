package server

import "qlt/internal/model"

//PaymentsService ...
type PaymentsService struct {
	PaymentsRepository *PaymentsRepository
}

//NewPaymentsService ...
func NewPaymentsService(paymentsRepository *PaymentsRepository) *PaymentsService {
	return &PaymentsService{paymentsRepository}
}

//CreatePayment ...
func (service *PaymentsService) CreatePayment(payment *model.PaymentPayload) *model.PaymentPayload {
	return service.PaymentsRepository.Save(payment)
}

//UpdatePayment ...
func (service *PaymentsService) UpdatePayment(id string, payment *model.PaymentPayload) (*model.PaymentPayload, error) {
	return service.PaymentsRepository.Update(id, payment)
}

//DeletePayment ...
func (service *PaymentsService) DeletePayment(id string) error {
	return service.PaymentsRepository.Delete(id)
}

//GetPaymentByID ...
func (service *PaymentsService) GetPaymentByID(id string) (*model.PaymentPayload, error) {
	return service.PaymentsRepository.FindByID(id)
}

//GetAllPayments ...
func (service *PaymentsService) GetAllPayments() []*model.PaymentPayload {
	return service.PaymentsRepository.FindAll()
}
