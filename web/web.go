package main

/*
References:
https://github.com/gorilla/websocket/tree/master/examples/echo
*/

import (
	"flag"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var testInventory = flag.Bool("t", false, "test mode")

type InventoryItem struct {
	ItemName string
	Price    int
}

var fakeItems = [...]InventoryItem{
	{ItemName: "Fooaaaaaaaaaaaaaa", Price: 37},
	{ItemName: "Bar", Price: 42},
	{ItemName: "Bill", Price: 42},
	{ItemName: "Boo", Price: 42},
	{ItemName: "Pee", Price: 42},
}

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
		}

		numItems := rand.Intn(len(fakeItems))
		startItem := rand.Intn(len(fakeItems))

		s := ""
		for i := 0; i < numItems; i++ {
			s += fakeItems[(i+startItem)%len(fakeItems)].ItemName + ":"
		}
		err = c.WriteMessage(mt, []byte(s))
		//err = c.WriteMessage(mt, []byte("test1:hello:there")) //message)
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
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Println("Starting server");

	server := http.Server{
		//Addr: "127.0.0.1:8080",
		Addr: ":8080",
	}
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	server.ListenAndServe()
}
