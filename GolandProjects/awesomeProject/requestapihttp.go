package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Id    string
	Title string
}

func main() {

	var item Item
	url := "https://api.mercadolibre.com/items/MLA1149305242"
	resp, _ := http.Get(url) //agarra el url

	data, _ := ioutil.ReadAll(resp.Body) //lee el json del body 
	json.Unmarshal(data, &item)
	fmt.Println(item)
	fmt.Println(item.Id)
	fmt.Println(item.Title)

}

