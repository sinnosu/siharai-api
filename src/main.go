package main

import (
	"siharai-api/src/controller"
	"siharai-api/src/db"
	"siharai-api/src/repository"
	"siharai-api/src/router"
	"siharai-api/src/usecase"
	"siharai-api/src/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	invoiceValidator := validator.NewInvoiceValidator()
	userRepository := repository.NewUserRepository(db)
	invoiceRepository := repository.NewInvoiceRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepository, invoiceValidator)
	userController := controller.NewUserController(userUsecase)
	invoiceController := controller.NewInvoiceController(invoiceUsecase)
	e := router.NewRouter(userController, invoiceController)
	e.Logger.Fatal(e.Start(":8080"))
}
