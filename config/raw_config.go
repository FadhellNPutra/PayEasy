package config

const(
	// User
	InsertUser = "INSERT INTO users(name, email, password, number, address, role, balance, updated_at) VALUES($1, $2, crypt($3, gen_salt('md5')), $4, $5, $6, $7, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at"
	SelectAllUser = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users LIMIT $1 OFFSET $2"
	SelectUserById = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users WHERE id = $1"
	SelectUserOnlyId = "SELECT id FROM users WHERE id = $1"
	SelectUserByEmail = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users WHERE email = $1"
	SelectUserForLogin = "SELECT id, name, email, password, role FROM users WHERE email = $1 AND password = crypt($2, password)"
	UpdateUser = "UPDATE users SET name = $1, email = $2, password = crypt($3, password), number = $4, address = $5, role = $6, balance = $7, updated_at = CURRENT_TIMESTAMP  WHERE id = $8 RETURNING created_at, updated_at"
	DeleteUser = "DELETE FROM users WHERE id = $1"
	
	// Merchant
	InsertMerchants = "INSERT INTO merchants(name_merchants, balance, updated_at) VALUES($1, $2, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at"
	SelectAllMerchants = "SELECT id, name_merchants, balance, created_at, updated_at FROM merchants LIMIT $1 OFFSET $2"
	SelectMerchantsById = "SELECT id, name_merchants, balance, created_at, updated_at FROM merchants WHERE id = $1"
	UpdateMerchants = "UPDATE merchants SET name_merchants = $1, balance = $2, updated_at = CURRENT_TIMESTAMP  WHERE id = $3 RETURNING created_at, updated_at"
	DeleteMerchants = "DELETE FROM merchants WHERE id = $1"
	
	// Transaction
	// InsertHistory = "INSERT INTO trx_history (id_users, id_merchants, status_payment, total_amount, updated_at) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at"
	// SelectAllHistory = "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history LIMIT $1 OFFSET $2"
	// SelectHistoryById = "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history WHERE id = $1"
	SelectBalanceUser = "SELECT u.id, u.name, u.email, u.password, u.number, u.address, u.role, u.balance FROM users u WHERE u.id = $1"
	InsertHistory = "INSERT INTO trx_history (id_users, id_merchants, status_payment, total_amount) VALUES ($1, $2, $3, $4)"
	SelectAllHistory = "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history"
	SelectHistoryById= "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history WHERE id = $1"
	SelectHistoryByIdUsers = "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history WHERE id_users = $1"
	SelectHistoryByIdMerchant = "SELECT id, id_users, id_merchants, status_payment, total_amount, created_at, updated_at FROM trx_history WHERE id_merchants = $1"
	UpdateHistory = "UPDATE merchants SET balance = balance + $1 WHERE id = $2"
	UpdateBalanceFromUsers = "UPDATE users SET balance = balance - $1 WHERE id = $2"
)