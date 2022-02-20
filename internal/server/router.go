package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"qlt/internal/model"
	"time"

	"github.com/gorilla/mux"

)

const baseURL = "/v1/payments"
//Router ...
type Router struct {
	Router *mux.Router
}

//SinglePaymentPayload ...
type SinglePaymentPayload struct {
	Data  *model.PaymentPayload `json:"data"`
	Links LinksPayload          `json:"links"`
}

//PaymentsListPayload ...
type PaymentsListPayload struct {
	Data  []*model.PaymentPayload `json:"data"`
	Links LinksPayload            `json:"links"`
}

//LinksPayload ...
type LinksPayload struct {
	Self string `json:"self"`
}

//NewRouter ...
func NewRouter(paymentsRepository *PaymentsRepository) *Router {
	paymentsService := NewPaymentsService(paymentsRepository)

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc(baseURL, GetAllPaymentsHandler(paymentsService)).Methods("GET")
	muxRouter.HandleFunc(baseURL+"/{id}", GetPaymentHandler(paymentsService)).Methods("GET")
	muxRouter.HandleFunc(baseURL, CreatePaymentHandler(paymentsService)).Methods("POST")
	muxRouter.HandleFunc(baseURL+"/{id}", UpdatePaymentHandler(paymentsService)).Methods("PUT")
	muxRouter.HandleFunc(baseURL+"/{id}", DeletePaymentHandler(paymentsService)).Methods("DELETE")

	return &Router{
		Router: muxRouter,
	}
}

//GetAllPaymentsHandler ...
func GetAllPaymentsHandler(paymentsService *PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newPayment := paymentsService.GetAllPayments()

		url := fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
		payload, _ := marshallPaymentsListPayload(newPayment, url)

		writeJSONResponse(w, http.StatusOK, payload)
	}
}

//GetPaymentHandler ...
func GetPaymentHandler(paymentsService *PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		newPayment, err := paymentsService.GetPaymentByID(vars["id"])
		if err != nil {
			log.Println("Error getting payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		url := fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
		payload, _ := marshalSinglePaymentPayload(newPayment, url)

		writeJSONResponse(w, http.StatusOK, payload)
	}
}

//CreatePaymentHandler ...
func CreatePaymentHandler(paymentsService *PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		var payment model.PaymentPayload
		err = json.Unmarshal(b, &payment)
		if err != nil {
			log.Println("Error parsing body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newPayment := paymentsService.CreatePayment(&payment)

		url := fmt.Sprintf("http://%s%s/%s", r.Host, r.URL.Path, newPayment.ID)
		payload, _ := marshalSinglePaymentPayload(newPayment, url)

		writeJSONResponse(w, http.StatusCreated, payload)
	}
}

//UpdatePaymentHandler ...
func UpdatePaymentHandler(paymentsService *PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		var payment model.PaymentPayload
		err = json.Unmarshal(b, &payment)
		if err != nil {
			log.Println("Error parsing body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		updatedPayment, err := paymentsService.UpdatePayment(vars["id"], &payment)
		if err != nil {
			log.Println("Error updating payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		url := fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
		payload, _ := marshalSinglePaymentPayload(updatedPayment, url)

		writeJSONResponse(w, http.StatusOK, payload)
	}
}

//DeletePaymentHandler ...
func DeletePaymentHandler(paymentsService *PaymentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := paymentsService.DeletePayment(vars["id"])
		if err != nil {
			log.Println("Error updating payment:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}


func writeJSONResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(payload)
}

func createPaymentPayload(payment *model.PaymentPayload) *model.PaymentPayload {
	return &model.PaymentPayload{
		ID:       "0",
		Name:     "Mustafa",
		Price:    220.00,
		Date:     time.Now().Format("January 2 15:04"),
		Type:     "income",
		Comment:  "user",
		Category: "student",
	}
}

func marshalPayment(payment *model.PaymentPayload) ([]byte, error) {
	return json.Marshal(createPaymentPayload(payment))
}

func marshalSinglePaymentPayload(payment *model.PaymentPayload, url string) ([]byte, error) {
	return json.Marshal(SinglePaymentPayload{
		Data:  createPaymentPayload(payment),
		Links: LinksPayload{Self: url},
	})
}

func marshallPaymentsListPayload(payments []*model.PaymentPayload, url string) ([]byte, error) {
	paymentsPayload := make([]*model.PaymentPayload, len(payments))
	for i := 0; i < len(payments); i++ {
		payload := createPaymentPayload(payments[i])
		paymentsPayload[i] = payload
	}

	return json.Marshal(PaymentsListPayload{
		Data:  paymentsPayload,
		Links: LinksPayload{Self: url},
	})
}
