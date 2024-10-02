package service

import (
	"io"
	"net/http"
	"os"

	"github.com/jsusmachaca/tiksup/pkg/movie/validation"
)

func ApiService(body io.Reader) error {
	API_URL := os.Getenv("PROCESSOR_URL")

	client := &http.Client{}

	res, err := client.Post(API_URL+"/recommend", "application/json", body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != 200 {
		return validation.ErrRequest
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}
