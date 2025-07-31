package tablas

import "gorm.io/gorm"

type Vehiculo struct {
	gorm.Model
	Placa string `json:"placa" gorm:"unique"`
	Tipo  string `json:"tipo"`
	Color string `json:"color"`
}
