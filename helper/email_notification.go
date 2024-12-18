package helper

import (
	"bytes"
	"car-rental/dtos"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func EmailNotification(message dtos.Message) ([]byte, error) {
	reqBody, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, os.Getenv("MAILTRAP_URL"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("MAILTRAP_APITOKEN"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
