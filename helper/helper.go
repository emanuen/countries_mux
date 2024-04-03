package helper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emanuen/countries-mux/config"
	"github.com/emanuen/countries-mux/models"
)

func Api() {

	var country []models.Country
	url := config.Config()["api"]

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "countries-mux")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &country)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
