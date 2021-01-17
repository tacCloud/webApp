package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var testInventory = flag.Bool("t", false, "test mode")

type InventoryItem struct {
	ItemName string
	Price    int
}

var fakeItems = [...]InventoryItem{
	{ItemName: "Foo", Price: 37},
	{ItemName: "Bar", Price: 42},
	{ItemName: "Bill", Price: 42},
	{ItemName: "Boo", Price: 42},
	{ItemName: "Pee", Price: 42},
}

//var homeTemplate = template.Must(template.New("").ParseFiles("tmp1.html"))

/*
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp1.html")
	//t.Execute(w, "Hello World!")
	t.Execute(w, fakeItems)
}
*/

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
		err = c.WriteMessage(mt, []byte("test1:hello:there")) //message)
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
	//homeTemplate.Execute(os.Stdout, "ws://"+r.Host+"/echo")
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	//fmt.Print(homeTemplate)
	flag.Parse()
	log.SetFlags(0)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)

	//	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
