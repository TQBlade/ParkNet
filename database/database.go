package database

import (
	"control_horario/tablas"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variable global para acceder a la base de datos

func Conectar() {
	var err error

	// Lee las credenciales de las variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Proporciona valores por defecto para facilitar la ejecución local
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

	// Valida que la contraseña fue proporcionada
	if dbPassword == "" {
		log.Fatal("Error: La variable de entorno DB_PASSWORD no está definida.")
	}

	// Construye la cadena de conexión (DSN) usando los valores leídos
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	log.Println("✅ Conectado a la base de datos")

	// 2. Ejecuta la migración automática para la tabla de Empleado
	log.Println("Migrando la base de datos...")
	err = DB.AutoMigrate(&tablas.Empleado{})
	if err != nil {
		log.Fatal("No se pudo migrar la base de datos:", err)
	}
	log.Println("✅ Base de datos migrada exitosamente.")
}
