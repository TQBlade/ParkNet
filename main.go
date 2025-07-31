package main

import (
	"control_horario/database"
	"control_horario/tablas"

	"github.com/gin-gonic/gin"
)

// CRUD HANDLERS PARA EMPLEADOS

func CrearEmpleado(c *gin.Context) {
	var empleado tablas.Empleado

	if err := c.ShouldBindJSON(&empleado); err != nil { //esta sirve para analizar el cuerpo de la solicitud y da error si asi lo es por ello no se usa el .Error
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

	result := database.DB.Find(&empleados)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los empleados"})
		return
	}
	c.JSON(200, empleados)
}

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

func EliminarEmpleado(c *gin.Context) {

	id := c.Param("id")
	var empleado tablas.Empleado

	if err := database.DB.First(&empleado, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "No se encontro al empleado"})
		return
	}

	database.DB.Delete(&empleado)
	c.JSON(200, gin.H{"mensaje": "se elimino al empleado"})
}

func CrearVehiculo(c *gin.Context) {

	var vehiculo tablas.Vehiculo

	if err := c.ShouldBindJSON(&vehiculo); err != nil {
		c.JSON(400, gin.H{"error": "Datos invalido:" + err.Error()})
		return
	}
	result := database.DB.Create(&vehiculo)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear vehiculo" + result.Error.Error()})
	}

	c.JSON(201, vehiculo)
}

func ObtenerVehiculos(c *gin.Context) {

	var vehiculo []tablas.Vehiculo

	result := database.DB.Find(&vehiculo)
	if result.Error != nil {
		c.JSON(500, gin.H{"eror": "No se encontraron los empleados" + result.Error.Error()})
		return
	}
	c.JSON(200, vehiculo)
}

func ObtenerVehiculo(c *gin.Context) {

	ID := c.Param("id")
	var vehiculo tablas.Vehiculo

	result := database.DB.First(&vehiculo, ID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "No se encontro al empleado" + result.Error.Error()})
		return
	}
	c.JSON(200, vehiculo)
}

func ActualizarVehiculo(c *gin.Context) {

	ID := c.Param("id")
	var vehiculo tablas.Vehiculo

	if err := database.DB.First(&vehiculo, ID).Error; err != nil {
		c.JSON(400, gin.H{"error": "No se encontro al empleado" + err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&vehiculo); err != nil {
		c.JSON(400, gin.H{"error": "Estos datos son invalidos" + err.Error()})
		return
	}
	database.DB.Save(&vehiculo)
	c.JSON(200, vehiculo)
}

func EliminarVehiculo(c *gin.Context) {

	ID := c.Param("id")
	var vehiculo tablas.Vehiculo

	if err := database.DB.First(&vehiculo, ID).Error; err != nil {
		c.JSON(400, gin.H{"error": "No se encontro el vehiculo" + err.Error()})
		return
	}
	database.DB.Delete(&vehiculo)
	c.JSON(200, gin.H{"mensaje": "Se elimino el vehiculo"})
}
func main() {
	database.Conectar()
	router := gin.Default()

	router.POST("/empleados", CrearEmpleado)

	router.GET("/empleados", ObtenerEmpleados)

	router.GET("/empleados/:id", ObtenerEmpleado)

	router.PUT("/empleados/:id", ActualizarEmpleado)

	router.DELETE("/empleados/:id", EliminarEmpleado)

	router.POST("/vehiculo", CrearVehiculo)

	router.GET("/vehiculo", ObtenerVehiculos)

	router.GET("/vehiculo/:id", ObtenerVehiculo)

	router.PUT("/vehiculo/:id", ActualizarVehiculo)

	router.DELETE("/vehiculo/:id", EliminarVehiculo)

	// Iniciar el servidor
	router.Run(":8081")
}
