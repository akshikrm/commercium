package db

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"log"
	"time"
)

func (m *storage) GetAllUsers() ([]*types.User, error) {
	query := "select * from users where role_code != 'admin' AND deleted_at IS NULL;"

	rows, err := m.db.Query(query)
	if err != nil {
		log.Printf("failed to retrieve user %s", err)
		return nil, utils.ServerError
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := scanUserRows(rows)
		if err != nil {
			return nil, utils.ServerError
		}
		users = append(users, user)
	}

	return users, nil
}

func (m *storage) GetPasswordByEmail(email string) (*types.User, error) {
	query := "select user_id, password, role_code from users inner join profiles on users.id = profiles.user_id where email=$1 AND users.deleted_at IS NULL;"

	row := m.db.QueryRow(query, email)

	user := types.User{}
	if err := row.Scan(&user.ID, &user.Password, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("profile with email: %s not found", email)
			return nil, utils.NotFound
		}
		log.Printf("failed to retrieve for email: %s due to error:%s", email, err)
		return nil, utils.ServerError
	}
	return &user, nil
}

func (m *storage) GetUserById(id int) (*types.User, error) {
	query := "select id, role_code, created_at,updated_at from users where id=$1 AND deleted_at IS NULL"
	row := m.db.QueryRow(query, id)
	user := &types.User{}
	err := row.Scan(
		&user.ID,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("user with id %d not found due to %s", id, err)
		return nil, utils.NotFound
	}
	return user, nil
}

func (m *storage) GetUserByEmail(email string) (*types.User, error) {
	query := "select * from users where email=$1 AND deleted_at IS NULL"
	row := m.db.QueryRow(query, email)

	user, err := scanUserRow(row)
	if err != nil {
		log.Printf("user with email %s not found due to %s", email, err)
		return nil, utils.NotFound
	}
	return user, nil
}

func (m *storage) GetCustomerID(id uint) *string {
	query := "select customer_id from users where id=$1"
	row := m.db.QueryRow(query, id)
	var customer_id string
	if err := row.Scan(&customer_id); err != nil {
		log.Printf("failed to get customer_id due to %s", err)
		return nil
	}
	return &customer_id
}

func (m *storage) CreateUser(user types.CreateUserRequest) (*types.User, error) {
	query := `insert into 
	users (password, role_code, customer_id)
	values($1, $2, $3)
	returning id, role_code
	`
	role := "user"
	if user.Role != "" {
		role = user.Role
	}
	row := m.db.QueryRow(query,
		user.Password,
		role,
		user.CustomerID,
	)
	log.Printf("Created user %v", user)

	savedUser := types.User{}
	if err := row.Scan(&savedUser.ID, &savedUser.Role); err != nil {
		log.Printf("failed to scan user after writing %d %s", savedUser.ID, err)
		return nil, utils.ServerError
	}

	return &savedUser, nil
}

func (m *storage) UpdateUser(id int, user types.UpdateUserRequest) error {
	query := `update users set first_name=$1, last_name=$2, email=$3 where id=$4`
	result, err := m.db.Exec(query, user.FirstName, user.LastName, user.Email, id)

	if err != nil {
		log.Printf("failed to update user %v due to %s", user, err)
		return utils.ServerError
	}

	if count, _ := result.RowsAffected(); count == 0 {
		log.Printf("updated %d rows", count)
		return utils.NotFound
	}

	return nil
}

func (m *storage) DeleteUser(id int) error {
	query := "UPDATE users set deleted_at=$1 where id=$2"
	if _, err := m.db.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to delete %d due to %s", id, err)
	}
	return nil
}

func scanUserRows(rows *sql.Rows) (*types.User, error) {
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

func scanUserRow(row *sql.Row) (*types.User, error) {
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
