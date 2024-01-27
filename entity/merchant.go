package entity

import "time"

type Merchant struct {
	Id           string    `json:"id"`
	NameMerchant string    `json:"name_merchant"`
	Balance      int       `json:"balance"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
