package main

/*
References:
https://github.com/gorilla/websocket/tree/master/examples/echo
*/

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/tacCloud/webApp/pkg/dbMgr"
	"github.com/tacCloud/webApp/pkg/inventoryMgr"
)

var testInventory = flag.Bool("t", false, "test mode")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	//Callback from javascript websocket
	//In non-test mode we would query the database here
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		if string(message)[0] == '#' {
			log.Println("update")
		} else {
			log.Println("Customer bought!")
			itemPrice := strings.Split(string(message), " ")
			price, _ := strconv.ParseFloat(itemPrice[len(itemPrice)-1][1:], 32)
			inventoryMgr.BuyItem(inventoryMgr.InventoryItem{
				strings.Join(itemPrice[0:len(itemPrice)-1], " "),
				float32(price)})
		}
		items := inventoryMgr.GetItems()
		s := ""
		for _, e := range items {
			s += fmt.Sprintf("%s $%.2f:", e.ItemName, e.Price)
		}
		err = c.WriteMessage(mt, []byte(s))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("HERE" + r.Host)
	homeTemplate, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	ver := os.Getenv("XYZ_MARKETPLACE_VERSION")
	if ver == "" {
		ver = "1.0.0"
	}
	data := struct {
		Host    string
		Version string
	}{
		Host:    "ws://" + r.Host + "/echo",
		Version: ver,
	}
	//homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
	fmt.Println(data)
	homeTemplate.Execute(w, data)
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Println("Starting server!")

	if *testInventory {
		dbMgr.FakeDb = true
	}

	server := http.Server{
		//Addr: "127.0.0.1:8080",
		Addr: ":8080",
	}
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	server.ListenAndServe()
}
