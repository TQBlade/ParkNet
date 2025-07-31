package tablas

import "gorm.io/gorm"

type Tarifa struct {
	gorm.Model
	TipoVehiculo  string  `json:"tipo_vehiculo" gorm:"unique"`
	TarifaPorHora float64 `json:"tarifa_por_hora"`
}
