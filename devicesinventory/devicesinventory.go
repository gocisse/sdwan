package devicesinventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	DeviceType    string `json:"deviceType"`
	Validity      string `json:"validity"`
	ChasisNumber  string `json:"chasisNumber"`
	SerialNumber  string `json:"serialNumber"`
	HostName      string `json:"host-name,omitempty"`
	SiteID        string `json:"site-id,omitempty"`
	SystemIP      string `json:"system-ip,omitempty"`
	LocalSystemIP string `json:"local-system-ip,omitempty"`
}

func DevicesInventory(baseURL string, client *http.Client) {
	// client = client
	url := baseURL + "dataservice/device/vedgeinventory/detail?api_key=vedges"
	device, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer device.Body.Close()

	data, err := ioutil.ReadAll(device.Body)
	if err != nil {
		log.Fatal(err)
	}

	parseData := &Response{}

	err = json.Unmarshal(data, parseData)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(parseData.Data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	devices := string(jsonData)
	if len(devices) > 0 {

		fmt.Printf("Data : %s\n", devices)
	}
	
}

