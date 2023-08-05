package usecase

import (
	"github.com/Junkes887/queues/internal/entity"
)

type ProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewProductUseCase(repository entity.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepository: repository,
	}
}

func (u *ProductUseCase) Create(input ProductInputDto) (ProductOutputDto, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return ProductOutputDto{}, err
	}

	return ProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}

func (u *ProductUseCase) List() ([]*ProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ProductOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})

	}

	return productsOutput, nil
}
