package main

import (
	"calmoon/application/usecases"
	"calmoon/domain/repository"
	"calmoon/domain/service"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=calmoon port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	signRepo := repository.NewSignRepository(db)
	signService := service.NewSignService(signRepo)
	signUseCase := usecases.NewSignUseCase(signService)

	signs, err := signUseCase.GetAllSigns()

	if err != nil {
		panic(err)
	}

	for _, sign := range signs {
		fmt.Printf("Sign: %s, Element: %s\n", sign.Name, sign.Element)
	}
}
