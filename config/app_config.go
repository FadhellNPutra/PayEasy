package config

const (
	ApiGroup = "/api/v1"

	// user
	UserCreate     = "/user"
	UserUpdate     = "/user"
	UserList       = "/user"
	UserGetById    = "/user/:id"
	UserDelete     = "/user/:id"
	UserGetByEmail = "/user/email/:email"

	// auth
	AuthLogin  = "/login"
	AuthLogout = "/logout"

	// merchant
	MerchantCreate  = "/merchant"
	MerchantUpdate  = "/merchant"
	MerchantList    = "/merchant"
	MerchantGetById = "/merchant/:id"
	MerchantDelete  = "/merchant/:id"

	// transaction
	HistoryCreate          = "/history"
	HistoryList            = "/history"
	HistoryGetByIdUsers    = "/history/:user"
	HistoryGetByIdMerchant = "/history/:merchant"
	HistoryGetBalance      = "/history/balance/:balance"
)
