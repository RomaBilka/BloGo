package main

import (
	"net/http"

	"github.com/RomaBilka/BloGo/project-layout/project-layou_2/configs"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_2/internal/product"
	productHttp "github.com/RomaBilka/BloGo/project-layout/project-layou_2/internal/product/http"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_2/internal/user"
	userHttp "github.com/RomaBilka/BloGo/project-layout/project-layou_2/internal/user/http"
	"github.com/RomaBilka/BloGo/project-layout/project-layou_2/pkg/database/postgres"
	httpServer "github.com/RomaBilka/BloGo/project-layout/project-layou_2/pkg/http"
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
	userRepository := user.NewPostgreUserRepository(db)
	userService := user.NewUserService(userRepository)
	uHttp := userHttp.NewUserHttp(userService)

	productRepository := product.NewPostgreProductRepository(db)
	productService := product.NewProductService(productRepository)
	pHttp := productHttp.NewProductHttp(productService)

	r.HandleFunc("/user", uHttp.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", uHttp.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", uHttp.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", uHttp.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", uHttp.DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/product", pHttp.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products", pHttp.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", pHttp.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", pHttp.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id}", pHttp.DeleteProduct).Methods(http.MethodDelete)

	httpServer.Start(":8080", r)
}
