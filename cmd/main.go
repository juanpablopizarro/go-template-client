package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("APP")
	viper.BindEnv("PORT")
	viper.BindEnv("ENDPOINT")

	port := viper.GetInt("PORT")
	endpoint := viper.Get("ENDPOINT")

	url := fmt.Sprintf("http://%s:%v/hash", endpoint, port)
	fmt.Println(url)

	resp, err := http.Post(url, "application/json; charset=utf-8", strings.NewReader(`{"string": "cjsndkj"}`))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
