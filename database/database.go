package database // Define el nombre del paquete, "database" porque este archivo se encarga de la conexión a la base de datos

import (
	"control_horario/tablas"
	"fmt" // Para construir cadenas de texto (usamos fmt.Sprintf)
	"log" // Para imprimir mensajes de error o información
	"os"  // Para acceder a variables del sistema (entorno)

	"github.com/joho/godotenv" // Librería para cargar el archivo .env
	"gorm.io/driver/postgres"  // Driver de PostgreSQL para GORM
	"gorm.io/gorm"             // ORM principal que usaremos para conectarnos y trabajar con la base de datos
)

var DB *gorm.DB // Variable global que contendrá la conexión a la base de datos

// Conectar es una función que se encarga de establecer la conexión con PostgreSQL
func Conectar() {
	// Intenta cargar las variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" Error al cargar .env") // Si falla, detiene la ejecución y muestra mensaje
	}

	// Lee las variables del entorno usando os.Getenv
	host := os.Getenv("DB_HOST")         // Dirección del servidor (por lo general, localhost)
	user := os.Getenv("DB_USER")         // Usuario de PostgreSQL
	password := os.Getenv("DB_PASSWORD") // Contraseña del usuario
	dbname := os.Getenv("DB_NAME")       // Nombre de la base de datos
	port := os.Getenv("DB_PORT")         // Puerto (por defecto 5432)

	// Construye el DSN (Data Source Name) con los datos obtenidos
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Abre la conexión a la base de datos usando GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Error al conectar con la base de datos:", err) // Si falla la conexión, muestra el error y termina
	}

	log.Println("Migrando la base de datos")
	err = DB.AutoMigrate(
		&tablas.Empleado{},
		&tablas.Vehiculo{},
	)
	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}
	// Si todo sale bien, imprime mensaje de éxito
	log.Println(" Conectado a la base de datos")
}
