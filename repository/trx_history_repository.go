package repository

import (
	"database/sql"
	"payeasy/config"
	"payeasy/entity"
)

type HistoryRepository interface {
	CreateTrxHistory(history *entity.History) error 
	ReadHistoryTransaction() ([]entity.History, error) 
	ReadHistoryTransactionById(id string) (*entity.History, error) 
	ReadHistoryTransactionByIdUser(idUser string) ([]entity.History, error) 
	ReadBalanceHistory(idUser string) (*entity.Users, error)
	ReadHistoryTransactionByIdMerchant(idMerchant string) ([]entity.History, error)
	UpdateBalanceFromUsers(idUser string, amount int) error
	UpdateBalanceFromMerchants(idMerchant string, amount int) error
	
}

type historyRepository struct {
	db *sql.DB
}

func (r *historyRepository) CreateTrxHistory(history *entity.History) error {
	stmt, err := r.db.Prepare(config.HistoryCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		history.IdUser,
		history.IdMerchant,
		history.StatusPayment,
		history.TotalAmount,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *historyRepository) ReadHistoryTransaction() ([]entity.History, error) {
	
	rows, err := r.db.Query(config.SelectAllHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History
	for rows.Next() {
		var history entity.History
		err := rows.Scan(
			&history.Id,
			&history.IdUser,
			&history.IdMerchant,
			&history.StatusPayment,
			&history.TotalAmount,
			&history.CreatedAt,
			&history.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}

	return histories, nil
}

func (r *historyRepository) ReadHistoryTransactionById(id string) (*entity.History, error) {
	row := r.db.QueryRow(config.SelectHistoryById, id)

	var history entity.History
	err := row.Scan(
		&history.Id,
		&history.IdUser,
		&history.IdMerchant,
		&history.StatusPayment,
		&history.TotalAmount,
		&history.CreatedAt,
		&history.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &history, nil
}

func (r *historyRepository) ReadHistoryTransactionByIdUser(idUser string) ([]entity.History, error) {
	rows, err := r.db.Query(config.SelectHistoryByIdUsers, idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

var histories []entity.History
for rows.Next() {
    var history entity.History
    err := rows.Scan(
        &history.Id,
        &history.IdUser,
        &history.IdMerchant,
        &history.StatusPayment,
        &history.TotalAmount,
        &history.CreatedAt,
        &history.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    histories = append(histories, history)
}

return histories, nil
}

func (r *historyRepository) ReadBalanceHistory(idUser string) (*entity.Users, error) {
    row := r.db.QueryRow(config.SelectBalanceUser, idUser)

    var user entity.Users
    err := row.Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.Number,
        &user.Address,
        &user.Role,
        &user.Balance,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &user, nil
}

func (r *historyRepository) ReadHistoryTransactionByIdMerchant(idMerchant string) ([]entity.History, error) {
    rows, err := r.db.Query(config.SelectHistoryByIdMerchant, idMerchant)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var histories []entity.History
    for rows.Next() {
        var history entity.History
        err := rows.Scan(
            &history.Id,
            &history.IdUser,
            &history.IdMerchant,
            &history.StatusPayment,
            &history.TotalAmount,
            &history.CreatedAt,
            &history.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        histories = append(histories, history)
    }

    return histories, nil
}

func (r *historyRepository) UpdateBalanceFromUsers(idUser string, amount int) error {
    
    stmt, err := r.db.Prepare(config.UpdateBalanceFromUsers)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(amount, idUser)
    if err != nil {
        return err
    }

    return nil
}

func (r *historyRepository) UpdateBalanceFromMerchants(idMerchant string, amount int) error {
    stmt, err := r.db.Prepare(config.UpdateHistory)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(amount, idMerchant)
    if err != nil {
        return err
    }

    return nil
}



func NewHistoryRepository(db *sql.DB) HistoryRepository {
	return &historyRepository{db}
}
