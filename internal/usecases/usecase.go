package usecases

import "ecomm/internal/domain"

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return u.repo.CreateProduct(product)
}

func (u *UseCase) GetProductByID(id string) (*domain.Product, error) {
	return u.repo.GetProductByID(id)
}

func (u *UseCase) GetAllProducts() ([]domain.Product, error) {
	return u.repo.ListProducts()
}

func (u *UseCase) UpdateProduct(product *domain.Product) error {
	return u.repo.UpdateProduct(product)
}

func (u *UseCase) DeleteProduct(id string) error {
	return u.repo.DeleteProduct(id)
}
