package main

import (
	"flavioamaral-dev/go-experts-lab-temperatura/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/temperatura", controllers.GetWeatherHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
