// +build ignore
package main

import (
	"fmt"

	"github.com/tacCloud/webApp/pkg/dbMgr"
)

func main() {
	//dbMgr.RedisEndpoint = "localhost:6379"
	dbMgr.InitializeDatabase(false, "localhost:6379")
	dbMgr.AddItem("Item1", 14.20)
	dbMgr.AddItem("Item2", 14.20)

	snapshot, _ := dbMgr.DumpDataBase()

	fmt.Println(snapshot)
}
