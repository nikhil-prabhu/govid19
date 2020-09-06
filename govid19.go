package main;

import (
	"os"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	apiUrl := "https://api.coronatracker.com/v3/stats/worldometer/topCountry"; // API url

	// Create client
	client := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	// Create new request
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil);

	// Log error
	if err != nil {
		log.Fatalln(err);
	}

	// Perform GET request
	resp, getErr := client.Do(req);

	// Log GET error
	if getErr != nil {
		log.Fatalln(getErr);
	}

	// Close response body
	if resp.Body != nil {
		defer resp.Body.Close();
	}

	// Read response body
	body, readErr := ioutil.ReadAll(resp.Body);

	// Log read response body error
	if readErr != nil {
		log.Fatalln(readErr);
	}
}
