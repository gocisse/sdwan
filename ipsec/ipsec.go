package ipsec

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

type response struct {
	Data []Data `json:"data"`
}

type Data struct {
	DestIP             string `json:"dest-ip"`
	SourcePort         int    `json:"source-port"`
	VdeviceName        string `json:"vdevice-name"`
	Spi                int    `json:"spi"`
	AuthKeyHash        string `json:"auth-key-hash"`
	VdeviceHostName    string `json:"vdevice-host-name"`
	RemoteTlocAddress  string `json:"remote-tloc-address"`
	DestPort           int    `json:"dest-port"`
	VdeviceDataKey     string `json:"vdevice-dataKey"`
	TunnelMtu          int    `json:"tunnel-mtu"`
	AuthenticationUsed string `json:"authentication-used"`
	Lastupdated        int64  `json:"lastupdated"`
	SourceIP           string `json:"source-ip"`
	RemoteTlocColor    string `json:"remote-tloc-color"`
	EncryptKeyHash     string `json:"encrypt-key-hash"`
}


func Ipsec (baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	cl := client
	url := baseURL + "dataservice/device/ipsec/outbound?deviceId=" + edge + "&&"
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
