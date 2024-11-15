package movie

import (
	"io"
	"net/http"
	"os"
)

func ApiService(client *http.Client, body io.Reader) error {
	PROCESSOR_URL := os.Getenv("PROCESSOR_URL")

	res, err := client.Post(PROCESSOR_URL+"/recommend", "application/json", body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != 200 {
		return ErrRequest
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}
