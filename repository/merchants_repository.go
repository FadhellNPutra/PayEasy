package repository

import (
	"database/sql"
	"log"
	"math"
	"payeasy/config"
	"payeasy/entity"
	"payeasy/shared/model"
	"time"
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
func (m *merchantRepository) CreateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	var merchant entity.Merchant

	payload.UpdatedAt = time.Now()

	err := m.db.QueryRow(config.InsertMerchants,
		payload.NameMerchant,
		payload.Balance,
	).Scan(&merchant.Id, &merchant.CreatedAt, &merchant.UpdatedAt)

	if err != nil {
		log.Println("merchantRepository.QueryRow: ", err.Error())
		return entity.Merchant{}, err
	}

	merchant.NameMerchant = payload.NameMerchant
	merchant.Balance = payload.Balance

	return merchant, nil
}

// DeleteMerchant implements MerchantRepository.
func (m *merchantRepository) DeleteMerchant(id string) error {
	err := m.db.QueryRow(config.DeleteMerchants, id).Err()
	if err != nil {
		log.Println("merchantRepository.Delete.QueryRow: ", err.Error())
		return err
	}

	log.Println(err)
	return nil
}

// GetMerchantById implements MerchantRepository.
func (m *merchantRepository) GetMerchantById(id string) (entity.Merchant, error) {
	var merchant entity.Merchant
	err := m.db.QueryRow(config.SelectMerchantsById, id).Scan(
		&merchant.Id,
		&merchant.NameMerchant,
		&merchant.Balance,
		&merchant.CreatedAt,
		&merchant.UpdatedAt)
	if err != nil {
		log.Println("merchantRepository.GetMerchantByID.QueryRow: ", err.Error())
		return entity.Merchant{}, err
	}
	return merchant, nil
}

// ListMerchant implements MerchantRepository.
func (m *merchantRepository) ListMerchant(page int, size int) ([]entity.Merchant, model.Paging, error) {
	var merchants []entity.Merchant
	offset := (page - 1) * size
	rows, err := m.db.Query(config.SelectAllMerchants, size, offset)
	if err != nil {
		log.Println("merchantRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}
	for rows.Next() {
		var merchant entity.Merchant
		err := rows.Scan(
			&merchant.Id,
			&merchant.NameMerchant,
			&merchant.Balance,
			&merchant.CreatedAt,
			&merchant.UpdatedAt,
		)
		if err != nil {
			log.Println("merchantRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}

		merchants = append(merchants, merchant)
	}

	totalRows := 0
	if err := m.db.QueryRow("SELECT COUNT(*) FROM merchants").Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return merchants, paging, nil
}

// UpdateMerchant implements MerchantRepository.
func (m *merchantRepository) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	var merchant entity.Merchant

	payload.UpdatedAt = time.Now()

	err := m.db.QueryRow(config.UpdateMerchants,
		payload.NameMerchant,
		payload.Balance,
		payload.Id,
	).Scan(&merchant.CreatedAt, &merchant.UpdatedAt)

	if err != nil {
		log.Println("merchantRepository.QueryRow: ", err.Error())
		return entity.Merchant{}, err
	}

	merchant.Id = payload.Id
	merchant.NameMerchant = payload.NameMerchant
	merchant.Balance = payload.Balance

	return merchant, nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
