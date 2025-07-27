package main

// Importamos los paquetes necesarios
import (
	"control_horario/database" // Paquete que gestiona la conexión a la base de datos
	"control_horario/tablas"   // Paquete que contiene el modelo Empleado
	"time"                     // Para manejar fechas y horas

	"github.com/gin-gonic/gin" // Framework web Gin para crear la API
)

// 📦 Estructura de datos para los registros (asistencia)
type Registro struct {
	ID        int       `json:"id"`         // ID único del registro
	Empleado  string    `json:"empleado"`   // Nombre del empleado
	FechaHora time.Time `json:"fecha_hora"` // Fecha y hora del registro
	Tipo      string    `json:"tipo"`       // Tipo: entrada o salida
}

// 📥 Estructura para recibir datos del cliente (sin ID ni fecha)
type RegistroEntrada struct {
	Empleado string `json:"empleado"` // Campo requerido en el JSON
	Tipo     string `json:"tipo"`     // Entrada o salida
}

// Almacenamiento temporal (slice en memoria, no guarda en base de datos)
var registros []Registro

func main() {

	// 🔌 Conexión a la base de datos
	database.Conectar()

	// 🚀 Crea una instancia del router con middleware por defecto (logger, recovery, etc.)
	router := gin.Default()

	// 🧪 Ruta de prueba para saber si el servidor está activo
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 📥 Ruta POST para registrar una asistencia (NO se guarda en la base de datos todavía)
	router.POST("/registro", func(c *gin.Context) {
		var datos RegistroEntrada // Creamos una variable para guardar el JSON recibido

		// 🔍 Intentamos leer el cuerpo del request como JSON
		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(400, gin.H{"error": err.Error()}) // Si falla, respondemos con error
			return
		}

		// 🏗 Construimos el nuevo registro completo con fecha y ID
		nuevo := Registro{
			ID:        len(registros) + 1, // ID simple basado en la cantidad de registros
			Empleado:  datos.Empleado,
			FechaHora: time.Now(), // Fecha y hora actuales
			Tipo:      datos.Tipo,
		}

		registros = append(registros, nuevo) // Guardamos en el slice en memoria
		c.JSON(200, nuevo)                   // Respondemos con el registro creado
	})

	// 📤 Ruta GET para obtener todos los registros guardados (en memoria)
	router.GET("/registros", func(c *gin.Context) {
		c.JSON(200, registros)
	})

	// 🧑 Ruta POST para guardar un empleado en la base de datos real (PostgreSQL)
	router.POST("/empleado", func(c *gin.Context) {
		var nuevo tablas.Empleado // Creamos una variable del modelo

		// Intentamos leer el JSON recibido
		if err := c.ShouldBindJSON(&nuevo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()}) // Si hay error, devolvemos 400
			return
		}

		// Guardamos el nuevo empleado en la base de datos usando GORM
		result := database.DB.Create(&nuevo)

		// Si hay un error al guardar, devolvemos un error 500
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		// Si todo va bien, devolvemos el empleado guardado con su ID
		c.JSON(200, nuevo)
	})

	// 🔊 Iniciamos el servidor en el puerto 8081 (http://localhost:8081)
	router.Run(":8081")
}
