package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Driver struct {
	Uuid  string `json: "uuid"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

type Drivers struct {
	Drivers []Driver
}

func getDrivers() []byte {
	jsonFile, err := os.Open("data.json")
	fmt.Println(jsonFile)
	if err != nil {
		panic(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	print(data)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func ShowDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := getDrivers()
	w.Write([]byte(drivers))
}

func GetDriversByUuid(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	data := getDrivers()

	var drivers Drivers
	json.Unmarshal(data, &drivers)

	for _, d := range drivers.Drivers {
		if d.Uuid == query["uuid"] {
			driver, _ := json.Marshal(d)
			w.Write([]byte(driver))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", ShowDrivers)

	http.ListenAndServe(":8001", r)

}
