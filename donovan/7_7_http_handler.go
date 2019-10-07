package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Dollars int64

func NewDollars(value string) Dollars {
	matched, ok := regexp.MatchString("^\\d+\\.\\d{0,2}?$", value)
	if matched == false || ok != nil {
		panic("Failed convert string to Dollars")
	}

	res, ok := strconv.Atoi(strings.Replace(value, ".", "", 1))
	if ok != nil {
		panic("Failed convert string to Dollars")
	}

	return Dollars(res)
}

func (d Dollars) String() string {
	return fmt.Sprintf("%d.%d", d/100, d%100)
}

type database map[string]Dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for key, value := range db {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Нет такого продукта %s", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{
		"shoes": NewDollars("12.20"),
		"socks": NewDollars("0.20"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/price", http.HandlerFunc(db.price))

	log.Fatal(http.ListenAndServe("localhost:7070", mux))
}
