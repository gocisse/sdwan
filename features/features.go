package features

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
// 	"net/url"
 )

// // const (
// // 	baseURL = "https://192.168.110.31:8443/"
// // )

// const (
// 	Username = "j_username"
// 	Password = "j_password"
// )

type Response struct {
	//	Header  Header    `json:"header"`
	Data []Data `json:"data"`
}

type Data struct {
	//DeviceType               []string `json:"deviceType"`
	TemplateMinVersion string `json:"templateMinVersion"`
	//TemplateType             string   `json:"templateType"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	//EditedTemplateDefinition string   `json:"editedTemplateDefinition"`
	//TemplateDefinition       string   `json:"templateDefinition"`
	AttachedMastersCount int    `json:"attachedMastersCount"`
	TemplateID           string `json:"templateId"`
	ConfigType           string `json:"configType"`
	CreatedOn            int64  `json:"createdOn"`
	Rid                  int    `json:"@rid"`
	FactoryDefault       bool   `json:"factoryDefault"`
	Feature              string `json:"feature"`
	CreatedBy            string `json:"createdBy"`
	TemplateName         string `json:"templateName"`
	DevicesAttached      int    `json:"devicesAttached"`
	TemplateDescription  string `json:"templateDescription"`
	LastUpdatedOn        int64  `json:"lastUpdatedOn"`
}



func Features(baseURL string, client *http.Client) {
	
	feature, err := client.Get(baseURL + "dataservice/template/feature")
	if err != nil {
		log.Fatal(err)
	}

	defer feature.Body.Close()

	data, err := ioutil.ReadAll(feature.Body)
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
