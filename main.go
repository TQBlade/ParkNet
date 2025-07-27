package main

import (
	"control_horario/database"
	"control_horario/tablas"
	"time"

	"github.com/gin-gonic/gin"
)

// Estructura completa del registro (incluye ID y fecha asignados por el servidor)
type Registro struct { //estructura de registro de empleado
	ID        int       `json:"id"`
	Empleado  string    `json:"empleado"`
	FechaHora time.Time `json:"fecha_hora"`
	Tipo      string    `json:"tipo"`
}

// Estructura solo para recibir los datos del cliente
type RegistroEntrada struct {
	Empleado string `json:"empleado"`
	Tipo     string `json:"tipo"`
}

var registros []Registro

func main() {

	database.Conectar() // Aquí se conecta la base de datos

	router := gin.Default() // Se crea un router (servidor) con configuraciones por defecto

	router.GET("/ping", func(c *gin.Context) { // Ruta GET simple para probar que el servidor funciona
		c.JSON(200, gin.H{"message": "pong"}) // Devuelve un JSON {"message":"pong"}i
	})

	router.POST("/registro", func(c *gin.Context) { // Ruta POST que recibe un nuevo registro desde el cliente
		var datos RegistroEntrada // Variable para recibir los datos enviados por el cliente

		// Recibimos solo lo que el cliente debe enviar
		if err := c.ShouldBindJSON(&datos); err != nil { // Intenta leer y convertir el JSON recibido a la estructura RegistroEntrada
			c.JSON(400, gin.H{"error": err.Error()}) // Si falla el parseo del JSON, responde con error 400
			return
		}

		// Ahora construimos el registro completo
		nuevo := Registro{
			ID:        len(registros) + 1, // ID incremental simple
			Empleado:  datos.Empleado,     // Asigna el nombre del empleado recibido
			FechaHora: time.Now(),         // Fecha y hora actual
			Tipo:      datos.Tipo,         // "entrada" o "salida"
		}

		registros = append(registros, nuevo) // Agrega el nuevo registro a la lista
		c.JSON(200, nuevo)                   // Devuelve el registro completo como respuesta
	})

	router.GET("/registros", func(c *gin.Context) { // Devuelve la lista completa como JSON
		c.JSON(200, registros) // Devuelve la lista completa como JSON
	})

	// Endpoint POST para guardar empleados
	router.POST("/empleado", func(c *gin.Context) {
		var nuevo tablas.Empleado

		if err := c.ShouldBindJSON(&nuevo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result := database.DB.Create(&nuevo)

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, nuevo)
	})

	router.Run(":8081") // El servidor comienza a escuchar en el puerto 8081
}
