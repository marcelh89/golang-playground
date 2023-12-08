package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// add headers
	req.Header.Add("X-Pattern", "DWAWEBFE")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseBody

}

func searchCars() []byte {
	return get("https://vtpapi.dasweltauto.com/restapi/v1/dwafrkons/search/car")
}

func configurationCarCriteria() []byte {
	return get("https://vtpapi.dasweltauto.com/restapi/v1/dwafrkons/configuration/car/criteria")
}
