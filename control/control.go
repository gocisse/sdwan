package control

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	
)

var edge int

type response struct {
	Data []Data `json:"data"`
}

type Data struct {
	DomainID          int    `json:"domain-id"`
	Instance          int    `json:"instance"`
	VdeviceName       string `json:"vdevice-name"`
	BehindProxy       string `json:"behind-proxy"`
	SystemIP          string `json:"system-ip"`
	RemoteColor       string `json:"remote-color"`
	SiteID            int    `json:"site-id"`
	PrivatePort       int    `json:"private-port"`
	ControllerGroupID int    `json:"controller-group-id"`
	VdeviceHostName   string `json:"vdevice-host-name"`
	LocalColor        string `json:"local-color"`
	Uptime            string `json:"uptime"`
	PeerType          string `json:"peer-type"`
	Protocol          string `json:"protocol"`
	VdeviceDataKey    string `json:"vdevice-dataKey"`
	PublicIP          string `json:"public-ip"`
	PublicPort        int    `json:"public-port"`
	Lastupdated       int64  `json:"lastupdated"`
	State             string `json:"state"`
	PrivateIP         string `json:"private-ip"`
	UptimeDate        int64  `json:"uptime-date"`
}

func Controls(baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	cl := client
	url := baseURL + "dataservice/device/control/connections?deviceId=" + edge + "&&"
	device, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer device.Body.Close()

	data, err := ioutil.ReadAll(device.Body)
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
}
