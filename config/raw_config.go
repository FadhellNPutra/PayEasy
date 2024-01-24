package config

const(
	// User
	InsertUser = "INSERT INTO users(name, email, pasword, number, address, role, balance, updated_at) VALUES($1, $2, crypt($3, gen_salt('md5')), $4, $5, $6, $7, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at"
	SelectAllUser = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users"
	SelectUserById = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users WHERE id = $1"
	SelectUserByEmail = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users WHERE email = $1"
	SelectUserForLogin = "SELECT id, name, email, password, number, address, role, balance, created_at, updated_at FROM users WHERE email = $1 AND password = $2"
	UpdateUser = "UPDATE FROM users SET name = $1, email = $2, password = crypt($3, password), number = $4, address = $5, role = $6, balance = $7, updated_at = CURRENT_TIMESTAMP  WHERE id = $8 RETURNING created_at, updated_at"
	DeleteUser = "DELETE FROM users WHERE id = $1"
	
	// Merchant
	InsertMerchants = "INSERT INTO merchants(name_merchants, id_user, balance, updated_at) VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	SelectAllMerchants = "SELECT id, name_merchants, id_users, balance, created_at, updated_at FROM users"
	SelectMerchantsById = "SELECT id, name_merchants, id_users, balance, created_at, updated_at FROM users WHERE id = $1"
	UpdateMerchants = "UPDATE FROM merchants SET name_merchants = $1, id_users = $2, balance = $3, updated_at = CURRENT_TIMESTAMP  WHERE id = $4 RETURNING id, created_at"
	DeleteMerchants = "DELETE FROM merchants WHERE id = $1"
	
	// Transaction
)