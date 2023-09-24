package parcel

import "fmt"

func (e *Envelope) SortParcel() string {
	e.sorting = &sorting{
		name: "Mailbox",
	}
	return "Mailbox"
}

func (e *Envelope) SendParcel() {
	fmt.Printf("The envelope %v was sent by %v to %v at the address %v\n", e.Name, e.Sender.Name, e.Recipient.Name, e.Recipient.Address)
}
