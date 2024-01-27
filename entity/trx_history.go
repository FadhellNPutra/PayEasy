package entity

import "time"

type History struct {
	Id            string    `json:"id"`
	IdUser        string    `json:"idUser"`
	IdMerchant    string    `json:"idMerchant"`
	StatusPayment string    `json:"statusPayment"`
	TotalAmount   int       `json:"totalAmount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (h History) IsTransactionTypeValid() bool {
	return h.StatusPayment == "CREDIT" || h.StatusPayment == "DEBIT"
}

func (h History) IsRequiredFields() bool {
	return h.TotalAmount > 0 || h.StatusPayment != "" 
}
