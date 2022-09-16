package main

import (
	"crud/dal"
	"crud/model"
	"fmt"
)

func main() {
	item := model.Item{0, "Item 100", 15.5, 12, 1, ""}
	rowsAffected, lastInsertedId, err := dal.InsertItem(item)

	if err != nil {
		fmt.Println("Insert error.")
	} else {
		if rowsAffected > 0 {
			fmt.Println("Insert completed.")
			item.ItemId = lastInsertedId
			fmt.Printf("new item id is %d\n", item.ItemId)
		}
	}
}
