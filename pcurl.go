package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yosssi/gohtml"
	"io/ioutil"
	"net/http"
	"os"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func pp(s string, v string) {
	fmt.Printf("\x1b[30;1m%v \x1b[30;0m%s\n", s, v)
	return
}

func ppCode(title string, code int) {
	if code == 200 || code == 201 || code == 202 {
		fmt.Printf("\x1b[30;1m%s \x1b[22;32m%d\x1b[30;0m\n", title, code)
	} else {
		fmt.Printf("\x1b[30;1m%s \x1b[30;0m%d\x1b[30;0m\n", title, code)
	}
	return
}

func ppHeader(s string, v http.Header) {
	fmt.Printf("\x1b[30;1m%s \x1b[30;0m%s\x1b[30;0m\n", s, v)
	return
}

func main() {
	client := &http.Client{}
	method := flag.String("x", "GET", "a string")

	flag.Parse()
	url := flag.Args()[0]

	req, err := http.NewRequest(*method, url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\x1b[34m======================================")
	fmt.Println("                HEADER")
	fmt.Println("======================================\x1b[0m")
	pp("url:", url)
	pp("method:", *method)
	pp("HTTP Status:", res.Status)
	ppCode("Status Code:", res.StatusCode)
	pp("Proto:", res.Proto)
	ppHeader("Header:", res.Header)

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println("\x1b[34m======================================")
	fmt.Println("                BODY")
	fmt.Println("======================================\x1b[0m")
	if isJSON(string(resBody)) {
		var out bytes.Buffer
		json.Indent(&out, resBody, "", "    ")
		out.WriteTo(os.Stdout)
	} else {
		fmt.Println(gohtml.Format(string(resBody)))
	}
}
