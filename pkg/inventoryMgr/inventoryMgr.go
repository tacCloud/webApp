package inventoryMgr

import
(
	"github.com/tacCloud/webApp/pkg/dbMgr"
)

// InventoryItem is great
type InventoryItem struct {
	ItemName string
	Price    float32
}

func GetItems() []InventoryItem {
	//Get the json from the dB
	var items []InventoryItem

	m,_ := dbMgr.DumpDataBase()

	for k,v := range(m) {
		item := InventoryItem{k, v}
		items = append(items, item)
	}
	return items
}

func BuyItem(item InventoryItem) {
	dbMgr.BuyItem(item.ItemName)
}