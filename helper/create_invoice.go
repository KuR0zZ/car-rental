package helper

import (
	"bytes"
	"car-rental/models"
	"encoding/json"
	"net/http"
	"os"
)

func CreateInvoice(user models.User, rental models.Rental, car models.Car) (*models.Invoice, error) {
	bodyRequest := map[string]interface{}{
		"external_id":      "1",
		"amount":           rental.TotalCosts,
		"description":      "Car Rental Invoice",
		"invoice_duration": 86400,
		"customer": map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
		"currency": "IDR",
		"items": []interface{}{
			map[string]interface{}{
				"name":     car.Name,
				"quantity": rental.Duration,
				"price":    car.RentalCosts,
				"category": car.Category,
			},
		},
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", os.Getenv("XENDIT_URL"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(os.Getenv("XENDIT_APIKEY"), "")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var resInvoice models.Invoice
	if err := json.NewDecoder(response.Body).Decode(&resInvoice); err != nil {
		return nil, err
	}

	return &resInvoice, nil
}
