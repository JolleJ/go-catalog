package models

type Item struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

var items []Item

func init() {
	items = append(items, Item{Id: "1", Description: "Test1", Category: "Category 1"})
	items = append(items, Item{Id: "2", Description: "Test2", Category: "Category 2"})
}

func GetItems() []Item {
	return items
}

func GetItem(id string) Item {
	for _, item := range items {
		if item.Id == id {
			return item
		}
	}
	return Item{}
}

func AddItem(item Item) Item {
	items = append(items, item)
	return item
}

func DeleteItem(id string) {
	for index, item := range items {
		if item.Id == id {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
}
