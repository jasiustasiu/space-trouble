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
	upcomingLaunchesUri = "/v3/launches/upcoming"
	launchpadsUri = "/v3/launchpads"
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

func (a *API) ListUpcomingLaunches(siteID string) ([]Launch, error) {
	response, err := a.client.Get(fmt.Sprintf("%v%v?site_id=%v", a.baseURL, upcomingLaunchesUri, siteID))
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

func (a *API) GetLaunchPad(siteID string) (*LaunchPad, error) {
	response, err := a.client.Get(fmt.Sprintf("%v%v/%v", a.baseURL, launchpadsUri, siteID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		log.Printf(fmt.Sprintf("Could not get launch pad. Reason: %v", string(output)))
		return nil, httpError.NewHTTPError(response.StatusCode, "Could not get launch pad")
	}
	var launches LaunchPad
	err = json.Unmarshal(output, &launches)
	return &launches, err
}