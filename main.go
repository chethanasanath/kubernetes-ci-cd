package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/Pallinder/go-randomdata"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	var str1 string = randomdata.SillyName()
	var str2 string = randomdata.Title(randomdata.Male)
	var str3 string = randomdata.FullName(randomdata.Male)
	var minversion string
	var jsonOut string = "{" + "App Version : "+ minversion + "," + str1 + ":" + str2 + "," + str3 +"}"
	data, _ := json.Marshal(jsonOut)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
	_=str1
	_=str2
	_=str3
	_ =jsonOut
	_=minversion
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
