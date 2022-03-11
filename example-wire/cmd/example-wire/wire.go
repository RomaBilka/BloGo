package main

/*
import (
	"database/sql"

	"github.com/RomaBilka/BloGo/example-wire/internal/handlers"
	"github.com/RomaBilka/BloGo/example-wire/internal/repositories"
	"github.com/RomaBilka/BloGo/example-wire/internal/services"
	"github.com/google/wire"
)

func InitializeUserHttp(db *sql.DB) *handlers.UserHTTP {
	wire.Build(
		wire.Bind(new(handlers.UserService), new(*services.UserService)),
		wire.Bind(new(services.UserRepository), new(*repositories.UserRepository)),
		repositories.NewPostgreUserRepository,
		services.NewUserService,
		handlers.NewUserHttp)

	return &handlers.UserHTTP{}
}

func InitializeProductHttp(db *sql.DB) *handlers.ProductHTTP {
	wire.Build(
		wire.Bind(new(handlers.ProductService), new(*services.ProductService)),
		wire.Bind(new(services.ProductRepository), new(*repositories.ProductRepository)),
		repositories.NewPostgreProductRepository,
		services.NewProductService,
		handlers.NewProductHttp)

	return &handlers.ProductHTTP{}
}
*/
