package repository

import (
	"database/sql"
	_"log"
	_"payeasy/config"
	"payeasy/entity"
	"payeasy/shared/model"
	_"time"
)

type MerchantRepository interface {
	CreateMerchant(payload entity.Merchant) (entity.Merchant, error)
	ListMerchant(page, size int) ([]entity.Merchant, model.Paging, error)
	GetMerchantById(id string) (entity.Merchant, error)
	UpdateMerchant(payload entity.Merchant) (entity.Merchant, error)
	DeleteMerchant(id string) error
}

type merchantRepository struct {
	db *sql.DB
}

// CreateMerchant implements MerchantRepository.
func (*merchantRepository) CreateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	// var merchant entity.Merchant
	// payload.UpdatedAt = time.Now()

	// tx, err := m.db.Begin()
	// if err != nil {
	// 	log.Println("merchantRepository.BeginTrx: ", err.Error())
	// 	return entity.Merchant{}, err
	// }

	// err = tx.QueryRow(config.MerchantCreate,
	// 	payload.NameMerchant,
	// 	payload.IdUsers,
	// 	payload.Balance,
	// ).Scan(
	// 	&payload.Id,
	// 	&payload.CreatedAt,
	// 	&payload.UpdatedAt,
	// )
	// if err != nil{
	// 	log.Println("merchantRepository.QueryInsertData: ", err.Error())
	// 	return entity.Merchant{}, err
	// }
	panic("unimplemented")


}

// DeleteMerchant implements MerchantRepository.
func (*merchantRepository) DeleteMerchant(id string) error {
	panic("unimplemented")
}

// GetMerchantById implements MerchantRepository.
func (*merchantRepository) GetMerchantById(id string) (entity.Merchant, error) {
	panic("unimplemented")
}

// ListMerchant implements MerchantRepository.
func (*merchantRepository) ListMerchant(page int, size int) ([]entity.Merchant, model.Paging, error) {
	panic("unimplemented")
}

// UpdateMerchant implements MerchantRepository.
func (*merchantRepository) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	panic("unimplemented")
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
