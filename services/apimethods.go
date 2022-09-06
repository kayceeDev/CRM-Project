package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
	
)

type Customer struct {
	ID        string  `json:"id"`
	Name      *string `json:"name" validate:"required"`
	Role      *string `json:"role" validate:"required"`
	Email     *string `json:"email" validate:"required"`
	Phone     *int    `json:"phone"`
	Contacted *bool   `json:"contacted"`
}

type Customers []Customer

// func HomeLink(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	http.StripPrefix("/",http.FileServer(http.Dir("static")))
// }

func AddCustomer(filename string) http.HandlerFunc {
	data := GetData(filename)
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer Customer
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.
				Fprintf(w, "Kindly enter data with the customer id, name, role and email only in order to update")
		}
		_ = json.Unmarshal(reqBody, &newCustomer)
		validate := validator.New()
		err = validate.Struct(newCustomer)
		if err != nil {
			// log out this error
			logrus.Error(err)
			// return a bad request and a helpful error message
			// if you wished, you could concat the validation error into this
			// message to help point your consumer in the right direction.
			http.Error(w, "failed to validate struct"+err.Error(), 400)
			return
		}
		id := (uuid.New()).String()
		newCustomer.ID = id
		data = append(data, newCustomer)
		dataBytes, err := json.MarshalIndent(data, "", "")
		if err != nil {
			logrus.Error(err)
		}

		err = ioutil.WriteFile(filename, dataBytes, 0644)
		if err != nil {
			logrus.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCustomer)
	}
}
func GetAllCustomers(filename string) http.HandlerFunc {
	data := GetData(filename)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func GetOneCustomer(filename string) http.HandlerFunc {
	data := GetData(filename)
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := mux.Vars(r)["id"]

		for _, singleCustomer := range data {
			if singleCustomer.ID == customerID {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(singleCustomer)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateCustomer(filename string) http.HandlerFunc {
	data := GetData(filename)
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := mux.Vars(r)["id"]
		var updatedCustomer Customer

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the valid customer deatails in order to update")
		}
		json.Unmarshal(reqBody, &updatedCustomer)

		for i, singleCustomer := range data {
			if singleCustomer.ID == customerID {
				if updatedCustomer.Name != nil {
					singleCustomer.Name = updatedCustomer.Name
				}
				if updatedCustomer.Role != nil {
					singleCustomer.Role = updatedCustomer.Role
				}
				if updatedCustomer.Email != nil {
					singleCustomer.Email = updatedCustomer.Email
				}
				if updatedCustomer.Phone != nil {
					singleCustomer.Phone = updatedCustomer.Phone
				}
				if updatedCustomer.Contacted != nil {
					singleCustomer.Contacted = updatedCustomer.Contacted
				}
				if len(data) == i { // nil or empty slice or after last element
					data = append(data, singleCustomer)
				}
				data[i] = singleCustomer
				dataBytes, err := json.MarshalIndent(data, " ", "")
				if err != nil {
					logrus.Error(err)
				}

				err = ioutil.WriteFile(filename, dataBytes, 0644)
				if err != nil {
					logrus.Error(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(singleCustomer)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "The event with ID %v is not found", customerID)
	}
}

func DeleteCustomer(filename string) http.HandlerFunc {
	data := GetData(filename)
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := mux.Vars(r)["id"]

		for i, singleCustomer := range data {
			if singleCustomer.ID == customerID {
				data = append(data[:i], data[i+1:]...)
				dataBytes, err := json.MarshalIndent(data, "", "")
				if err != nil {
					logrus.Error(err)
				}

				err = ioutil.WriteFile(filename, dataBytes, 0644)
				if err != nil {
					logrus.Error(err)
				}
				fmt.Fprintf(w, "The event with ID %v has been deleted successfully", customerID)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "The event with ID %v is not found", customerID)
	}
}
