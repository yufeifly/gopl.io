package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// 创建

// 读取

// 更新
func (db database) update(w http.ResponseWriter, req *http.Request) {

	/*item := req.URL.Query().Get("item")
	if _,ok := db[item];!ok {
		fmt.Fprintf(w, "fuck, item dose not exist!\n")

	} else {
		price, _ := strconv.ParseFloat(req.URL.Query().Get("price"),32)
		if price < 0.0 {
			fmt.Fprintf(w, "fuck, price is illegal!\n")
		} else {
			fmt.Fprintf(w, "fuck:%s %f\n", item, dollars(price))
			mu.Lock()
			db[item] = dollars(price)
			mu.Unlock()
		}
	}

	 */

	query := req.URL.Query()
	for key, item := range query {
		fmt.Fprintln(w, key, item)
	}



}

// 删除