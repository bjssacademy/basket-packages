package main

import (
	"fmt"
	"shopping/basket"
)

func main(){
	myBasket := basket.Basket{}

	item1 := basket.Item{Name: "Banana", Price: 999}
	item2 := basket.Item{Name: "Cat Food", Price: 99}

	myBasket.Add(item1)
	myBasket.Add(item2)

	fmt.Println("Total:", myBasket.Total)

	myBasket.Remove(item1)

	fmt.Println("Total:", myBasket.Total)

}