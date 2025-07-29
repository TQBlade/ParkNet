package main

import (
	"control_horario/database"
	"control_horario/tablas"

	"github.com/gin-gonic/gin"
)

// CRUD HANDLERS PARA EMPLEADOS

func CrearEmpleado(c *gin.Context) {
	var empleado tablas.Empleado

	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(400, gin.H{"error": "Datos de entrada inválidos: " + err.Error()})
		return
	}

	result := database.DB.Create(&empleado)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el empleado: " + result.Error.Error()})
		return
	}

	c.JSON(201, empleado) // 201 Created es más apropiado para creaciones exitosas.
}

// ObtenerEmpleados responde con la lista de todos los empleados.
func ObtenerEmpleados(c *gin.Context) {
	var empleados []tablas.Empleado

	// Usamos 'result' como nombre estándar para el resultado de la operación de GORM.
	result := database.DB.Find(&empleados)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los empleados"})
		return
	}
	c.JSON(200, empleados)
}

// ObtenerEmpleado responde con un solo empleado según su ID.
func ObtenerEmpleado(c *gin.Context) {
	// Usamos 'id' como nombre estándar para el parámetro de la URL.
	id := c.Param("id")

	var empleado tablas.Empleado
	result := database.DB.First(&empleado, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Empleado no encontrado"})
		return
	}
	c.JSON(200, empleado)
}

// ActualizarEmpleado modifica un empleado existente según su ID.
func ActualizarEmpleado(c *gin.Context) {
	id := c.Param("id")

	// La primera búsqueda es para asegurarnos de que el registro existe.
	var empleado tablas.Empleado
	if err := database.DB.First(&empleado, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Empleado no encontrado"})
		return
	}

	// Vinculamos el JSON recibido al registro que ya encontramos.
	// Esto actualiza los campos de la variable 'empleado' en memoria.
	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(400, gin.H{"error": "Datos de entrada inválidos: " + err.Error()})
		return
	}

	// Guardamos el registro actualizado en la base de datos.
	database.DB.Save(&empleado)

	c.JSON(200, empleado)
}

func main() {
	database.Conectar()
	router := gin.Default()

	router.POST("/empleados", CrearEmpleado)

	router.GET("/empleados", ObtenerEmpleados)

	router.GET("/empleados/:id", ObtenerEmpleado)

	router.PUT("/empleados/:id", ActualizarEmpleado)

	// Iniciar el servidor
	router.Run(":8081")
}
