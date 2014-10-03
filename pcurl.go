package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/yosssi/gohtml"
)

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

	Header("HEADER")
	PPKeyVal("URL:", url)
	PPKeyVal("Method:", *method)
	PPKeyVal("HTTP Status:", res.Status)
	PPrintCode("Status Code:", res.StatusCode)
	PPKeyVal("Protocol:", res.Proto)
	Header("BODY")

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if IsJSON(string(resBody)) {
		var out bytes.Buffer
		json.Indent(&out, resBody, "", "    ")
		out.WriteTo(os.Stdout)
	} else {
		fmt.Println(gohtml.Format(string(resBody)))
	}
}
