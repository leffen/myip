package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	resp, err := http.Get("http://ifconfig.co")
	if err != nil {
		logrus.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(string(data))
}
