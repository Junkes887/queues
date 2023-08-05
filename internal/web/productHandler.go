package web

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/queues/internal/usecase"
	"github.com/Junkes887/queues/internal/utils"
)

type ProductHandler struct {
	ProductUsecase *usecase.ProductUseCase
}

func NewProductHandler(productUsecase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.ProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	utils.ErrorHttpStatusInternalServerError(err, w)

	output, err := p.ProductUsecase.Create(input)
	utils.ErrorHttpStatusInternalServerError(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	output, err := p.ProductUsecase.List()
	utils.ErrorHttpStatusInternalServerError(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
