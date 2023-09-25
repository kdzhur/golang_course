package parcel

import "fmt"

type Parcel struct {
	Name      string
	Sender    *Sender
	Recipient *Recipient
	sorting   *sorting
}

type sorting struct {
	name string
}

type ParcelSorter interface {
	SortParcel() string
}

type ParcelSender interface {
	SendParcel()
}

type Parceler interface {
	ParcelSender
	ParcelSorter
}

func SendParcelTo(p Parceler) {
	fmt.Println("The parcel was sorted to", p.SortParcel())
	p.SendParcel()
}
