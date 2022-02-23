package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RomaBiliak/BloGo/project-layout/project-layou_1/internal/models"
	"github.com/RomaBiliak/BloGo/project-layout/project-layou_1/pkg/response"
	"github.com/gorilla/mux"
)

type productService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProduct(id models.ProductId) (models.Product, error)
	DeleteProduct(id models.ProductId) error
	UpdateProduct(id models.ProductId, Product models.Product) (models.Product, error)
	GetProducts() ([]models.Product, error)
}

type ProductHTTP struct {
	productService productService
}

func NewProductHttp(productService productService) *ProductHTTP {
	return &ProductHTTP{productService: productService}
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
	ProductRequest := &createProductRequest{}

	err := json.NewDecoder(r.Body).Decode(ProductRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ProductModel := models.Product{
		Name:        ProductRequest.Name,
		Description: ProductRequest.Description,
		Image:       ProductRequest.Image,
		Price:       ProductRequest.Price,
		Count:       ProductRequest.Count,
	}

	newProduct, err := h.productService.CreateProduct(ProductModel)

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

	ProductModel := models.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Image:       productRequest.Image,
		Price:       productRequest.Price,
		Count:       productRequest.Count,
	}

	product, err := h.productService.UpdateProduct(productId, ProductModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createProductResponse{Id: uint64(product.Id), Name: product.Name, Description: product.Description, Image: product.Image, Price: product.Price, Count: product.Count})
}

func (h *ProductHTTP) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := getProductId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.productService.GetProduct(productId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createProductResponse{Id: uint64(product.Id), Name: product.Name, Description: product.Description, Image: product.Image, Price: product.Price, Count: product.Count})
}

func (h *ProductHTTP) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetProducts()

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	var productsResponse []createProductResponse

	for _, product := range products {
		productsResponse = append(productsResponse, createProductResponse{Id: uint64(product.Id), Name: product.Name, Description: product.Description, Image: product.Image, Price: product.Price, Count: product.Count})
	}

	response.WriteJSON(w, http.StatusCreated, productsResponse)
}

func (h *ProductHTTP) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := getProductId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.productService.DeleteProduct(productId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusNoContent, createProductResponse{})
}

func getProductId(r *http.Request) (models.ProductId, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		return 0, err
	}

	return models.ProductId(id), nil
}
