package inventoryMgr

//Really dumb tests (as an example)
import (
	"testing"

	"github.com/tacCloud/webApp/pkg/dbMgr"
)

func init() {
	dbMgr.FakeDb = true
}

func TestGetItems(t *testing.T) {
	items := GetItems()
	if len(items) != 2 {
		t.Error("Bad length")
	}
	var total float32
	for _, e := range items {
		total += e.Price
	}
	if total >= 30 {
		t.Error("Bad price = ", total)
	}
	t.Log("items ", items)
}

func TestBuyItem(t *testing.T) {
	itemName := "The Idiot"
	BuyItem(InventoryItem{ItemName: itemName, Price: 1})
	//very its gone
	items := GetItems()
	for _, item := range items {
		t.Log("item = ", item)
		if item.ItemName == itemName {
			t.Errorf("%s was not bought!", itemName)
		}
	}
}
