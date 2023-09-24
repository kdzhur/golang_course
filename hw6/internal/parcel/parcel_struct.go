package parcel

//
type Customer struct {
	Name    string
	Address string
}

type Sender struct {
	Customer
}

type Recipient struct {
	Customer
}

//
type Parcel struct {
	Name      string
	Sender    *Sender
	Recipient *Recipient
	sorting   *sorting
}

type Envelope struct {
	Parcel
}

type Box struct {
	Parcel
}

//
type sorting struct {
	name string
}
