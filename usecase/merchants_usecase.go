package usecase

import (
	"errors"
	"fmt"
	"payeasy/entity"
	"payeasy/repository"
	"payeasy/shared/model"
)

type MerchantUseCase interface {
	FindMerchantByID(id string) (entity.Merchant, error)
	RegisterNewMerchant(payload entity.Merchant) (entity.Merchant, error)
	UpdateMerchant(payload entity.Merchant) (entity.Merchant, error)
	ListAll(page, size int) ([]entity.Merchant, model.Paging, error)
	DeleteMerchant(id string) error
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

// DeleteMerchant implements MerchantUseCase.
func (m *merchantUseCase) DeleteMerchant(id string) error {
	if _, err := m.repo.GetMerchantById(id); err != nil {
		return err
	}
	return m.repo.DeleteMerchant(id)
}

// FindMerchantByID implements MerchantUseCase.
func (m *merchantUseCase) FindMerchantByID(id string) (entity.Merchant, error) {
	if id == "" {
		return entity.Merchant{}, errors.New("id harus diisi")
	}
	return m.repo.GetMerchantById(id)
}

// ListAll implements MerchantUseCase.
func (m *merchantUseCase) ListAll(page int, size int) ([]entity.Merchant, model.Paging, error) {
	return m.repo.ListMerchant(page, size)
}

// RegisterNewMerchant implements MerchantUseCase.
func (m *merchantUseCase) RegisterNewMerchant(payload entity.Merchant) (entity.Merchant, error) {
	if payload.NameMerchant == "" ||  payload.Balance == 0 {
		return entity.Merchant{}, fmt.Errorf("oops, field required")
	}

	merchant, err := m.repo.CreateMerchant(payload)
	if err != nil {
		return entity.Merchant{}, fmt.Errorf("oppps, failed to save data users :%v", err.Error())
	}
	return merchant, nil
}

// UpdateMerchant implements MerchantUseCase.
func (m *merchantUseCase) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	if payload.NameMerchant == "" || payload.Balance == 0 || payload.Id == "" {
		return entity.Merchant{}, fmt.Errorf("oops, field required")
	}

	merchant, err := m.repo.UpdateMerchant(payload)
	if err != nil {
		return entity.Merchant{}, fmt.Errorf("oppps, failed to save data merchant :%v", err.Error())
	}
	return merchant, nil
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
