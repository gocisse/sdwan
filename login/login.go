package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// const (
// 	username = "j_username"
// 	password = "j_password"
// )
var err error

func Login(loginCredentials url.Values, link, user, passwd string, client *http.Client) {
	resp, err := client.PostForm(link, loginCredentials)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(string(data), "<html>") {
		fmt.Println("Login failed!")
		os.Exit(1)

	} else {
		fmt.Println("Access Granted!")

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

/*




 */
