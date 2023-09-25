package parcel

import "fmt"

type Box struct {
	Parcel
}

func (b *Box) SortParcel() string {
	b.sorting = &sorting{
		name: "Stock",
	}
	return "Stock"
}

func (b *Box) SendParcel() {
	fmt.Printf("The box %v was sent by %v to %v at the address %v\n", b.Name, b.Sender.Name, b.Recipient.Name, b.Recipient.Address)
}
