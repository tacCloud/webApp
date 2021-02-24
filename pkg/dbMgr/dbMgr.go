package dbMgr

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

var fakeItems = map[string]float32{
	"Crime and Punishment": 12.50,
	"The Idiot":            10.00,
}

func InitializeDatabase(fakeDb bool, dbEndpoint string) {
	if fakeDb {
		return
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     dbEndpoint,
		Password: "",
		DB:       0,
	})
	statusCmd := rdb.Ping(ctx)
	if statusCmd.Err() != nil {
		panic(statusCmd.Err())
	}
}

// DumpDataBase is great
func DumpDataBase() (map[string]float32, error) {

	if rdb == nil {
		return fakeItems, nil
	}

	snapshot := make(map[string]float32)

	var cursor uint64
	for {
		var keys []string
		var err error

		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			panic(err)
		}
		for _, k := range keys {
			val, _ := rdb.Get(ctx, k).Result()
			valFloat, _ := strconv.ParseFloat(val, 32)
			snapshot[k] = float32(valFloat)
		}
		if cursor == 0 {
			break
		}
	}
	return snapshot, nil
}

func BuyItem(item string) {
	fmt.Println("Buying ", item)
	if rdb == nil {
		delete(fakeItems, item)
	} else {
		rdb.Del(ctx, item)
	}
}

func AddItem(item string, price float32) {
	if rdb == nil {
		fakeItems[item] = price
	} else {
		s := strconv.FormatFloat(float64(price), 'f', -1, 32)
		rdb.Set(ctx, item, s, 0)
	}
}

/*
Testing
-> % docker run --rm --name redis-test-instance -p 6379:6379 -d redis
-> % docker run -it --rm redis redis-cli -h 172.17.0.2
172.17.0.2:6379>
172.17.0.2:6379>
172.17.0.2:6379> SET Bahamas Nassau
OK
172.17.0.2:6379> GET Bahamas
"Nassau"
172.17.0.2:6379> GET poop
"13.2"
172.17.0.2:6379> SET "A book" 13.5
OK
172.17.0.2:6379> SET "A Better Book" 1.0
OK

*/
