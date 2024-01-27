package usecase

import (
	"errors"
	"payeasy/entity"
	"payeasy/repository"
)

type HistoryUsecase interface {
	RegisterNewTransaction(history *entity.History) error
	GetById(id string) (*entity.History, error)
	GetByIdUsers(idUser string) ([]entity.History, error)
	GetByIdMerchant(idMerchant string) ([]entity.History, error)
	Update(history *entity.History) error
	ReadHistoryTransaction() ([]entity.History, error)
	GetBalanceByUserId(idUser string) int 
}

type historyUsecase struct {
	repo repository.HistoryRepository
}


func (u *historyUsecase) RegisterNewTransaction(history *entity.History) error {
	if !history.IsRequiredFields() {
		return errors.New("required fields are not filled")
	}

	if history.StatusPayment != "CREDIT" && history.StatusPayment != "DEBIT" {
		return errors.New("invalid transaction type")
	}

	if history.StatusPayment == "DEBIT" && history.TotalAmount > u.GetBalanceByUserId(history.IdUser) {
		return errors.New("insufficient balance")
	}

	return u.repo.CreateTrxHistory(history)
}

func (u *historyUsecase) GetById(id string) (*entity.History, error) {
	history, err := u.repo.ReadHistoryTransactionById(id)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (u *historyUsecase) GetByIdUsers(idUser string) ([]entity.History, error) {
	histories, err := u.repo.ReadHistoryTransactionByIdUser(idUser)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (u *historyUsecase) GetByIdMerchant(idMerchant string) ([]entity.History, error) {
	histories, err := u.repo.ReadHistoryTransactionByIdMerchant(idMerchant)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (u *historyUsecase) Update(history *entity.History) error {
	return u.repo.UpdateBalanceFromUsers(history.IdUser, -history.TotalAmount)
}

func (u *historyUsecase) GetBalanceByUserId(idUser string) int {
	user, err := u.repo.ReadBalanceHistory(idUser)
	if err != nil {
		return 0
	}

	return user.Balance
}

func (u *historyUsecase) ReadHistoryTransaction() ([]entity.History, error) {
    histories, err := u.repo.ReadHistoryTransaction()
    if err != nil {
        return nil, err
    }

    return histories, nil
}

func NewHistoryUsecase(repo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{repo}
}