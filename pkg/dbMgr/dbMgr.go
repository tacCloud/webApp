package dbMgr

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var FakeDb bool = false
var rdb *redis.Client
var ctx = context.Background()

var fakeItems = map[string]float32{
	"Crime and Punishment": 12.50,
	"The Idiot":            10.00,
}

// DumpDataBase is great
func DumpDataBase() (map[string]float32, error) {

	if FakeDb {
		return fakeItems, nil
	} else if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
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
	if FakeDb {
		delete(fakeItems, item)
	} else {
		rdb.Del(ctx, item)
	}
}

func AddItem(item string, price float32) {
	if FakeDb {
		fakeItems[item] = price
	} else {
		s := strconv.FormatFloat(float64(price), 'f', -1, 32)
		rdb.Set(ctx, item, s, 0)
	}
}
