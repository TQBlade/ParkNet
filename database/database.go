package database

// Importamos paquetes necesarios
import (
	"control_horario/tablas" // Importa el modelo Empleado desde el paquete "tablas"
	"fmt"                    // Para construir cadenas de texto (como el DSN)
	"log"                    // Para imprimir mensajes de error o información
	"os"                     // Para leer variables de entorno

	"gorm.io/driver/postgres" // Driver de PostgreSQL para GORM
	"gorm.io/gorm"            // ORM GORM para trabajar con bases de datos
)

// DB es una variable global para acceder a la base de datos desde cualquier parte del proyecto
var DB *gorm.DB

// Conectar establece la conexión con PostgreSQL y ejecuta la migración
func Conectar() {
	var err error // Variable para capturar errores

	// Lee las variables de entorno para configuración de conexión
	dbHost := os.Getenv("DB_HOST")         // Dirección del host (ej. localhost)
	dbUser := os.Getenv("DB_USER")         // Usuario de la base de datos
	dbPassword := os.Getenv("DB_PASSWORD") // Contraseña de la base de datos
	dbName := os.Getenv("DB_NAME")         // Nombre de la base de datos
	dbPort := os.Getenv("DB_PORT")         // Puerto de PostgreSQL (por defecto 5432)

	// Si alguna variable no está definida, se asigna un valor por defecto para desarrollo local
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbName == "" {
		dbName = "parqueadero"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	// Validamos que se haya proporcionado la contraseña, si no, detenemos el programa
	if dbPassword == "" {
		log.Fatal("Error: La variable de entorno DB_PASSWORD no está definida.")
	}

	// Construimos el DSN (Data Source Name), que es la cadena de conexión a PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Intentamos abrir la conexión con GORM usando el DSN
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Si hay un error al conectar, se imprime y se detiene el programa
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	log.Println(" Conectado a la base de datos")

	// Ejecuta la migración automática para la estructura Empleado
	// Esto crea la tabla si no existe o la actualiza si hay cambios en el modelo
	log.Println("Migrando la base de datos...")
	err = DB.AutoMigrate(&tablas.Empleado{})
	if err != nil {
		log.Fatal("No se pudo migrar la base de datos:", err)
	}
	log.Println(" Base de datos migrada exitosamente.")
}
