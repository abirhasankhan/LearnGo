package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Bit more JSON :)")

	// Encodejson()
	DecodeJson()

}

func Encodejson() {

	leocourse := []course{
		{"Reat", 599, "TikTok", "dshauasdasd", []string{"code", "course"}},
		{"Java", 1999, "Youtuve", "asdasdas", []string{"code", "course"}},
		{"JS", 199, "TikTok", "asdasdasd", nil},
	}

	// fmt.Println(leocourse)

	finalJson, err := json.MarshalIndent(leocourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}


func DecodeJson() {

	jsonDataFromWeb := []byte(`
		{
			"coursename": "Reat",
			"price": 599,
			"website": "TikTok",
			"tags": ["code","course"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was vaild")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}


	// add data key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("key is %v, value %v\n", key, val)
	}

}
