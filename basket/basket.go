package basket

import (
	"shopping/helpers"
)

type Basket struct {
	Items []Item
	Total int
}

func (b *Basket) calculateTotal() {
	total := 0
	for _, item := range b.Items {
		total += item.Price
	}
	b.Total = total
}

func (b *Basket) Add(item Item) {
	b.Items = append(b.Items, item)
	b.calculateTotal()
}

func (b *Basket) Remove(itemToRemove Item) {
	b.Items = helpers.Delete[Item](b.Items, itemToRemove)
	b.calculateTotal()
}

