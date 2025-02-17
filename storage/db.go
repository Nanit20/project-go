package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/Nanit20/project-go/router"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewStorage() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var errConn error
	DB, errConn = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errConn != nil {
		log.Fatal("Error al conectar con la base de datos:", errConn)
	}
	fmt.Println("Base de datos conectada en Supabase")

	if err := DB.Migrator().DropTable(&router.Task{}); err != nil {
		log.Fatalf("Error al eliminar la tabla 'tasks': %v", err)
	}

	fmt.Println("Se ha limpiado la tabla")

	if err := DB.AutoMigrate(&router.Task{}); err != nil {
		log.Fatalf("Error al realizar migración: %v", err)
	}

	fmt.Println("Base de datos migrada correctamente")
	return DB
}
