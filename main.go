package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"gocisse/sdwan/bfd"
	"gocisse/sdwan/features"
	"gocisse/sdwan/ipsec"
	"log"
	"gocisse/sdwan/control"
	"gocisse/sdwan/devices"
	"gocisse/sdwan/devicesinventory"
	"gocisse/sdwan/login"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const (
	username = "j_username"
	password = "j_password"
)

var (
	baseURL = "https://192.168.110.31:8443/"
)
var err error

func main() {
	// Here is to bypass ssl https security check
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//Same as above
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// handling coockie jar
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	// baseUrl  := bufio.NewReader(os.Stdin)
	// fmt.Printf("vManage URL >")
	// baseURL, _ := baseUrl.ReadString('\n')
	// baseURL = strings.Replace(baseURL, "\n", "", -1)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("vManage Username>")
	uname, _ := reader.ReadString('\n')
	uname = strings.Replace(uname, "\n", "", -1)

	fmt.Printf("vManage Password>")
	passwd, _ := reader.ReadString('\n')
	passwd = strings.Replace(passwd, "\n", "", -1)

	loginLink := baseURL + "j_security_check"
	loginCredentials := url.Values{
		username: {uname},
		password: {passwd},
	}

	login.Login(loginCredentials, loginLink, uname, passwd, client)
	var (
		menu = `
		#  ███╗   ███╗███████╗███╗   ██╗██╗   ██╗    
		#  ████╗ ████║██╔════╝████╗  ██║██║   ██║    
		#  ██╔████╔██║█████╗  ██╔██╗ ██║██║   ██║    
		#  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║██║   ██║    
		#  ██║ ╚═╝ ██║███████╗██║ ╚████║╚██████╔╝    
		#  ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝  
		==============================================
		
		1 - For all devices in the sdwan Fabrc
		2 - For all Devices Inventory 
		3 - For Control connections per system-ip
		4 - For Feature Templates
		5 - For ipsec outboud connections per system-ip
		6 - For BFD sessuion history per sys-ip 
		7 - For BFD session history Edge to Edges sys-ip required
		8 - For BFD session history between Edges sys-ip and tloc colour required
		9 - For BFD session tlocs sysIP required

		0 - To quit the program 
		`
	)

	count := Menus(menu)
	for {

		switch {
		case count == 0:
			fmt.Println("Exiting program .....")
			os.Exit(0)
		case count == 1:
			fmt.Println("Devices in sdwan Fabric")
			devices.Devices(baseURL, client)
		case count == 2:
			fmt.Println("Devices inventory")
			devicesinventory.DevicesInventory(baseURL, client)

		case count == 3:
			control.Controls(baseURL, client)
		case count == 4:
			features.Features(baseURL, client)
		case count == 5:
			ipsec.Ipsec(baseURL, client)
		case count == 6:
			bfd.BfdSessionHistory(baseURL, client)
		case count == 7:
			bfd.BfdSessionEdges(baseURL, client)
		case count == 8:
			bfd.BfdSessionCoulor(baseURL, client)
		case count == 9:
			bfd.BfdSessionTlocs(baseURL, client)
		default:
			fmt.Println("You have to choose a number ")
			return
		}
		count = Menus(menu)
	}

}

func Menus(menu string) int {

	fmt.Printf("%v\n", menu)
	var count int
	fmt.Printf("Choose a number >")
	_, err = fmt.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
