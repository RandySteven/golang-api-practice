package configs

import (
	"fmt"
	"log"
	"test-api/entities/models"
	"test-api/interfaces"
	"test-api/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	interfaces.UserUsecase
	interfaces.ProductUsecase
	db *gorm.DB
}

func NewRepository(config *models.Config) (*Repositories, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
	log.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	return &Repositories{
		UserUsecase:    usecase.NewUserUsecase(db),
		ProductUsecase: usecase.NewProductUsecase(db),
		db:             db,
	}, nil
}

func (r *Repositories) Automigrate() error {
	return r.db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Product{},
		models.Transaction{},
		models.TransactionDetail{},
	)
}
