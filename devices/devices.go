package devices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Data []Data `json:"data"`
}

type Data struct {
	DeviceID            string   `json:"deviceId"`
	SystemIP            string   `json:"system-ip"`
	HostName            string   `json:"host-name"`
	Reachability        string   `json:"reachability"`
	Status              string   `json:"status"`
	Personality         string   `json:"personality"`
	DeviceType          string   `json:"device-type"`
	Timezone            string   `json:"timezone"`
	DeviceGroups        []string `json:"device-groups"`
	Lastupdated         int64    `json:"lastupdated"`
	DomainID            string   `json:"domain-id,omitempty"`
	BoardSerial         string   `json:"board-serial"`
	CertificateValidity string   `json:"certificate-validity"`
	MaxControllers      string   `json:"max-controllers,omitempty"`
	UUID                string   `json:"uuid"`
	ControlConnections  string   `json:"controlConnections,omitempty"`
	DeviceModel         string   `json:"device-model"`
	Version             string   `json:"version"`
	ConnectedVManages   []string `json:"connectedVManages"`
	SiteID              string   `json:"site-id"`
	Latitude            string   `json:"latitude"`
	Longitude           string   `json:"longitude"`
	IsDeviceGeoData     bool     `json:"isDeviceGeoData"`
	Platform            string   `json:"platform"`
	UptimeDate          int64    `json:"uptime-date"`
	StatusOrder         int      `json:"statusOrder"`
	DeviceOs            string   `json:"device-os"`
	Validity            string   `json:"validity"`
	State               string   `json:"state"`
	StateDescription    string   `json:"state_description"`
	ModelSku            string   `json:"model_sku"`
	LocalSystemIP       string   `json:"local-system-ip"`
	TotalCPUCount       string   `json:"total_cpu_count"`
	TestbedMode         bool     `json:"testbed_mode"`
	LayoutLevel         int      `json:"layoutLevel"`
	OmpPeers            string   `json:"ompPeers,omitempty"`
	LinuxCPUCount       string   `json:"linux_cpu_count,omitempty"`
	BfdSessionsUp       int      `json:"bfdSessionsUp,omitempty"`
	BfdSessions         string   `json:"bfdSessions,omitempty"`
}

func Devices(baseURL string, client *http.Client) {
	cl := client
	url := baseURL + "dataservice/device"
	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatalln("unable to unmarchal json", jsonErr)
	}

	device := data["data"].([]interface{})

	fmt.Println("####################################################")
	for value := range device {

		currentdevice := device[value].(map[string]interface{})

		fmt.Println("| Hostname |", currentdevice["host-name"], "| System-IP |", currentdevice["system-ip"], "| Reachability |", currentdevice["reachability"], "| Version |", currentdevice["version"])
	}

}

/*


	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	parseData := &response{}

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





*/
