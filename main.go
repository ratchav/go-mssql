package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ratchav/go-mssql/entities"
	"github.com/ratchav/go-mssql/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost" // or the Docker service name if running in another container
	port     = 5432        // default PostgreSQL port
	user     = "postgres"  // as defined in docker-compose.yml
	password = "P@ssw0rd"  // as defined in docker-compose.yml
	dbname   = "postgres"  // as defined in docker-compose.yml
)

func main() {
	// DESKTOP-334VGBG\MSSQLSERVER2014

	app := fiber.New()

	db := initialDB()
	repo := repositories.NewQuotationRepositoryDB(db)
	qutation, err := repo.GetQuotation(1)
	if err != nil {
		log.Fatalf("failed to get quotation %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(qutation)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		quotation := new(entities.Quotation)
		if err := c.BodyParser(quotation); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(200)
		}

		if err := repo.Create(quotation); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(400)
		}

		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":8000"))
}

func initialDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db

}
