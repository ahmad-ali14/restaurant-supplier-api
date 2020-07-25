package supplier

import (
	"restaurant-supplier-api/utils/dbHandler"
)

type Supplier struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type SupplierList struct {
	Suppliers []Supplier
}

func (s *SupplierList) AddNew(sup Supplier) string {
	s.Suppliers = append(s.Suppliers, sup)
	return dbHandler.AddItemToDb()

}

func (s *SupplierList) GetAll() []Supplier {
	return s.Suppliers
}
