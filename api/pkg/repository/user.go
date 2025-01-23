package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"log"
	"time"
)

type user struct {
	store *sql.DB
}

func (m *user) Get() ([]*types.User, bool) {
	query := "select * from users where role_code != 'admin' AND deleted_at IS NULL;"

	rows, err := m.store.Query(query)
	if err != nil {
		log.Printf("failed to retrieve user %s", err)
		return nil, false
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := ScanRows(rows)
		if err != nil {
			return nil, false
		}
		users = append(users, user)
	}

	return users, true
}

func (m *user) GetPasswordByEmail(email string) (*types.User, bool) {
	query := "select user_id, password, role_code from users inner join profiles on users.id = profiles.user_id where email=$1 AND users.deleted_at IS NULL;"

	row := m.store.QueryRow(query, email)

	user := types.User{}
	if err := row.Scan(&user.ID, &user.Password, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("profile with email: %s not found", email)
			return nil, true
		}
		log.Printf("failed to retrieve for email: %s due to error:%s", email, err)
		return nil, false
	}
	return &user, true
}

func (m *user) GetOne(id uint32) (*types.User, bool) {
	query := "select id, role_code, created_at,updated_at from users where id=$1 AND deleted_at IS NULL"
	row := m.store.QueryRow(query, id)
	user := &types.User{}
	err := row.Scan(
		&user.ID,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("user with id %d not found due to %s", id, err)
		return nil, true
	}
	return user, true
}

func (m *user) GetUserByEmail(email string) (*types.User, bool) {
	query := "select * from users where email=$1 AND deleted_at IS NULL"
	row := m.store.QueryRow(query, email)

	user, err := ScanRow(row)
	if err != nil {
		log.Printf("user with email %s not found due to %s", email, err)
		return nil, true
	}
	return user, true
}

func (m *user) GetCustomerID(id uint32) *string {
	query := "select customer_id from users where id=$1"
	row := m.store.QueryRow(query, id)
	var customer_id string
	if err := row.Scan(&customer_id); err != nil {
		log.Printf("failed to get customer_id due to %s", err)
		return nil
	}
	return &customer_id
}

func (m *user) Create(user types.CreateUserRequest) (*types.User, bool) {
	query := `insert into 
	users (password, role_code, customer_id)
	values($1, $2, $3)
	returning id, role_code
	`
	role := "user"
	if user.Role != "" {
		role = user.Role
	}

	row := m.store.QueryRow(query,
		user.Password,
		role,
		user.CustomerID,
	)
	savedUser := types.User{}
	if err := row.Scan(&savedUser.ID, &savedUser.Role); err != nil {
		log.Printf("failed to scan user after writing %d %s", savedUser.ID, err)
		return nil, false
	}

	return &savedUser, true
}

func (m *user) Update(id uint32, user types.UpdateUserRequest) bool {
	query := `update users set first_name=$1, last_name=$2, email=$3 where id=$4`
	result, err := m.store.Exec(query, user.FirstName, user.LastName, user.Email, id)

	if err != nil {
		log.Printf("failed to update user %v due to %s", user, err)
		return false
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return false
	}
	return true
}

func (m *user) Delete(id uint32) bool {
	query := "UPDATE users set deleted_at=$1 where id=$2"
	if _, err := m.store.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to delete %d due to %s", id, err)
		return false
	}
	return true
}

func ScanRows(rows *sql.Rows) (*types.User, error) {
	user := &types.User{}
	err := rows.Scan(
		&user.ID,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		log.Printf("scan into user failed due to %s", err)
	}
	return user, err
}

func ScanRow(row *sql.Row) (*types.User, error) {
	user := &types.User{}
	err := row.Scan(
		&user.ID,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	return user, err
}

func newUser(store *sql.DB) types.UserRepository {
	return &user{
		store: store,
	}
}
