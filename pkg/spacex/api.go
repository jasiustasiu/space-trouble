package spacex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"space-trouble/internal/httpError"
	"time"
)

const (
	launchesUri = "/v3/launches"
)

func NewAPI(baseURL string) *API {
	return &API{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: "https://api.spacexdata.com",
	}
}

type API struct {
	client  *http.Client
	baseURL string
}

func (a *API) ListLaunches() ([]Launch, error) {
	response, err := a.client.Get(fmt.Sprintf("%v%v", a.baseURL, launchesUri))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		log.Printf(fmt.Sprintf("Could not list launches. Reason: %v", string(output)))
		return nil, httpError.NewHTTPError(response.StatusCode, "Could not list launches")
	}
	var launches []Launch
	err = json.Unmarshal(output, &launches)
	return launches, err

}