package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/kayceeDev/CRM-Project/utils"
)

func GetData (filename string) Customers{
    err := utils.CheckFile(filename)
    if err != nil {
        log.Fatal(err)
    }

	content, err := ioutil.ReadFile("./customers.json")
	if err != nil {
		log.Fatal(err)
	}
	data := Customers{}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}