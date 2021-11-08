package bfd

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

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	DstIP           string `json:"dst-ip"`
	SrcIP           string `json:"src-ip"`
	Color           string `json:"color"`
	TxPkts          int    `json:"tx-pkts"`
	VdeviceName     string `json:"vdevice-name"`
	SrcPort         int    `json:"src-port"`
	TimeDate        int64  `json:"time-date"`
	SystemIP        string `json:"system-ip"`
	Index           int    `json:"index"`
	DstPort         int    `json:"dst-port"`
	Del             int    `json:"del"`
	SiteID          int    `json:"site-id"`
	VdeviceHostName string `json:"vdevice-host-name"`
	VdeviceDataKey  string `json:"vdevice-dataKey"`
	Proto           string `json:"proto"`
	Lastupdated     int64  `json:"lastupdated"`
	Time            string `json:"time"`
	State           string `json:"state"`
	RxPkts          int    `json:"rx-pkts"`
}
type Responses struct {
	Data []Datas `json:"data"`
}

type Datas struct {
	SessionsFlap    int    `json:"sessions-flap"`
	IfName          string `json:"if-name"`
	SessionsTotal   int    `json:"sessions-total"`
	VdeviceDataKey  string `json:"vdevice-dataKey"`
	VdeviceName     string `json:"vdevice-name"`
	Lastupdated     int64  `json:"lastupdated"`
	Encap           string `json:"encap"`
	SessionsUp      int    `json:"sessions-up"`
	VdeviceHostName string `json:"vdevice-host-name"`
}

func BfdSessionHistory(baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	cl := client
	url := baseURL + "dataservice/device/bfd/history?deviceId=" + edge + "&&"
	device, err := cl.Get(url)
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

func BfdSessionEdges(baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	fmt.Printf("Destination Edge System-ip >")
	sysIP, _ := reader.ReadString('\n')
	sysIP = strings.Replace(sysIP, "\n", "", -1)

	cl := client
	url := baseURL + "dataservice/device/bfd/history?deviceId=" + edge + "&system-ip=" + sysIP + "&"
	device, err := cl.Get(url)
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
func BfdSessionCoulor(baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	fmt.Printf("Destination Edge System-ip >")
	sysIP, _ := reader.ReadString('\n')
	sysIP = strings.Replace(sysIP, "\n", "", -1)
	fmt.Printf("Transport Colour>")
	colour, _ := reader.ReadString('\n')
	colour = strings.Replace(colour, "\n", "", -1)

	cl := client
	url := baseURL + "dataservice/device/bfd/history?deviceId=" + edge + "&" + "system-ip=" + sysIP + "&" + "color=" + colour
	device, err := cl.Get(url)
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

func BfdSessionTlocs(baseURL string, client *http.Client) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Edge System-ip >")
	edge, _ := reader.ReadString('\n')
	edge = strings.Replace(edge, "\n", "", -1)
	cl := client
	url := baseURL + "dataservice/device/bfd/tloc?deviceId=" + edge + "&&"
	device, err := cl.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer device.Body.Close()

	data, err := ioutil.ReadAll(device.Body)
	if err != nil {
		log.Fatal(err)
	}
	parseData := &Responses{}

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
