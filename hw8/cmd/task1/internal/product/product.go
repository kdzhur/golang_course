package product

type Product struct {
	Name         string  `json:"name"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}
