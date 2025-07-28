package main

import (
	"control_horario/database" // Conexión a base de datos
	"control_horario/tablas"   // Modelo de Empleado

	"github.com/gin-gonic/gin" // Framework web
)

func ObtenerEmpleados(c *gin.Context) {

	var empleados []tablas.Empleado
	result_empleados := database.DB.Find(&empleados)

	if result_empleados.Error != nil {
		c.JSON(500, gin.H{"erro": "error al obtener empleados"})
		return
	}
	c.JSON(200, empleados)
}

func ObtenerEmpleado(c *gin.Context) {

	ID_Empleado := c.Param("id")
	var empleado tablas.Empleado
	result_empleado := database.DB.First(&empleado, ID_Empleado)

	if result_empleado.Error != nil {
		c.JSON(404, gin.H{"error": "Empleado no encontrado"})
		return
	}
	c.JSON(200, empleado)
}

func main() {
	//  Conectar a PostgreSQL
	database.Conectar()

	//  Crear el router
	router := gin.Default()

	// Ruta para guardar empleados en la base de datos
	router.POST("/empleado", func(c *gin.Context) {
		var nuevoEmpleado tablas.Empleado

		// Leer JSON recibido
		if err := c.ShouldBindJSON(&nuevoEmpleado); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Guardar en la base de datos
		result := database.DB.Create(&nuevoEmpleado)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, nuevoEmpleado)
	})

	router.GET("/obtenerempleados", ObtenerEmpleados)

	//  Iniciar el servidor en el puerto 8081
	router.Run(":8081")
}
