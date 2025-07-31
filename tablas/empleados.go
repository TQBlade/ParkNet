package tablas

import "gorm.io/gorm"

type Empleado struct {
	// gorm:"primaryKey" es correcto, pero no es necesario si usas el nombre 'ID'
	// GORM automáticamente trata un campo llamado 'ID' de tipo uint como llave primaria.
	gorm.Model        // Se añade la etiqueta JSON
	Nombre     string `json:"nombre"`               // Se añade la etiqueta JSON
	Cedula     string `json:"cedula" gorm:"unique"` // Se añade la etiqueta JSON
}
