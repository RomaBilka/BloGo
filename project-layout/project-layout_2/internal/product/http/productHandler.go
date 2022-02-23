package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RomaBiliak/BloGo/project-layout/project-layou_2/internal/product"
	"github.com/RomaBiliak/BloGo/project-layout/project-layou_2/pkg/response"
	"github.com/gorilla/mux"
)

type ProductService interface {
	CreateProduct(product product.Product) (product.Product, error)
	GetProduct(id product.ProductId) (product.Product, error)
	DeleteProduct(id product.ProductId) error
	UpdateProduct(id product.ProductId, Product product.Product) (product.Product, error)
	GetProducts() ([]product.Product, error)
}

type ProductHTTP struct {
	ProductService ProductService
}

func NewProductHttp(ProductService ProductService) *ProductHTTP {
	return &ProductHTTP{ProductService: ProductService}
}

type createProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"email"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
}

type createProductResponse struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"email"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
}

func (h *ProductHTTP) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productRequest := &createProductRequest{}

	err := json.NewDecoder(r.Body).Decode(productRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	productModel := product.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Image:       productRequest.Image,
		Price:       productRequest.Price,
		Count:       productRequest.Count,
	}

	newProduct, err := h.ProductService.CreateProduct(productModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createProductResponse{Id: uint64(newProduct.Id), Name: newProduct.Name, Description: newProduct.Description, Image: newProduct.Image, Price: newProduct.Price, Count: newProduct.Count})
}

func (h *ProductHTTP) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := getProductId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	productRequest := &createProductRequest{}

	err = json.NewDecoder(r.Body).Decode(productRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	productModel := product.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Image:       productRequest.Image,
		Price:       productRequest.Price,
		Count:       productRequest.Count,
	}

	Product, err := h.ProductService.UpdateProduct(productId, productModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createProductResponse{Id: uint64(Product.Id), Name: Product.Name, Description: Product.Description, Image: Product.Image, Price: Product.Price, Count: Product.Count})
}

func (h *ProductHTTP) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := getProductId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	productInDb, err := h.ProductService.GetProduct(productId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createProductResponse{Id: uint64(productInDb.Id), Name: productInDb.Name, Description: productInDb.Description, Image: productInDb.Image, Price: productInDb.Price, Count: productInDb.Count})
}

func (h *ProductHTTP) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetProducts()

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	var productsResponse []createProductResponse

	for _, p := range products {
		productsResponse = append(productsResponse, createProductResponse{Id: uint64(p.Id), Name: p.Name, Description: p.Description, Image: p.Image, Price: p.Price, Count: p.Count})
	}

	response.WriteJSON(w, http.StatusCreated, productsResponse)
}

func (h *ProductHTTP) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := getProductId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.ProductService.DeleteProduct(productId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusNoContent, createProductResponse{})
}

func getProductId(r *http.Request) (product.ProductId, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		return 0, err
	}

	return product.ProductId(id), nil
}
