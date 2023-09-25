package parcel

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
