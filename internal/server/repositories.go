package server

import (
	"fmt"
	"qlt/internal/model"

	uuid "github.com/satori/go.uuid"
)

//UUID ...
type UUID interface {
	GetNextUUID() string
}

//RandomUUID ...
type RandomUUID struct{}

//PaymentsRepository ...
type PaymentsRepository struct {
	payments []*model.PaymentPayload
	u        UUID
}

//GetNextUUID ...
func (randomUuid *RandomUUID) GetNextUUID() string {
	return uuid.NewV4().String()
}

//NewPaymentsRepository ...
func NewPaymentsRepository() *PaymentsRepository {
	return &PaymentsRepository{
		payments: make([]*model.PaymentPayload, 0),
		u:        &RandomUUID{},
	}
}

//Save ...
func (r *PaymentsRepository) Save(payment *model.PaymentPayload) *model.PaymentPayload {
	newPayment := model.PaymentPayload(*payment)
	newPayment.ID = r.u.GetNextUUID()

	r.payments = append(r.payments, &newPayment)

	return &newPayment
}

//Update ...
func (r *PaymentsRepository) Update(id string, payment *model.PaymentPayload) (*model.PaymentPayload, error) {
	for i := 0; i < len(r.payments); i++ {
		current := r.payments[i]
		if current.ID == id {
			updatedPayment := model.PaymentPayload(*payment)
			updatedPayment.ID = current.ID
			r.payments[i] = &updatedPayment
			return &updatedPayment, nil
		}
	}

	return &model.PaymentPayload{}, fmt.Errorf("Could not update payment with id %s, payment does not exist", id)
}

//Delete ...
func (r *PaymentsRepository) Delete(id string) error {
	for i := 0; i < len(r.payments); i++ {
		current := r.payments[i]
		if current.ID == id {
			r.payments = append(r.payments[:i], r.payments[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not delete payment with id %s, payment does not exist", id)
}

//FindByID ...
func (r *PaymentsRepository) FindByID(id string) (*model.PaymentPayload, error) {
	for i := 0; i < len(r.payments); i++ {
		current := r.payments[i]
		if current.ID == id {
			return current, nil
		}
	}

	return &model.PaymentPayload{}, fmt.Errorf("Could not find payment with id %s", id)
}

//FindAll ...
func (r *PaymentsRepository) FindAll() []*model.PaymentPayload {
	paymentsCopy := make([]*model.PaymentPayload, len(r.payments))
	copy(paymentsCopy, r.payments)

	return paymentsCopy
}
