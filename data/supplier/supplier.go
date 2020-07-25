package supplier

type Supplier struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type SupplierList struct {
	Suppliers []Supplier
}
