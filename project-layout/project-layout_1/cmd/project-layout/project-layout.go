package main

import (
	"net/http"

	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/configs"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/internal/handlers"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/internal/repositories"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/internal/services"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/pkg/database/postgres"
	httpServer "github.com/RomaBilka/BloGo/project-layout/project-layou_1/pkg/http"
	"github.com/gorilla/mux"
)

func main() {

	dbConfig := postgres.Config{
		configs.PG_USER,
		configs.PG_PASSWORD,
		configs.PG_DATABASE,
		configs.PG_HOST,
	}
	db := postgres.Run(dbConfig)
	defer db.Close()

	r := mux.NewRouter()
	userRepository := repositories.NewPostgreUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHttp := handlers.NewUserHttp(userService)

	productRepository := repositories.NewPostgreProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHttp := handlers.NewProductHttp(productService)

	r.HandleFunc("/user", userHttp.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHttp.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHttp.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", userHttp.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", userHttp.DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/product", productHttp.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products", productHttp.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", productHttp.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", productHttp.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", productHttp.DeleteProduct).Methods(http.MethodDelete)

	httpServer.Start(":8080", r)
}
