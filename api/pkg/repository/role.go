package repository

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"log"
	"time"
)

type role struct {
	store *sql.DB
}

func (r *role) GetAll() ([]*types.Role, bool) {
	query := `select * from roles`
	rows, err := r.store.Query(query)

	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("Failed to query db for roles due to %s", err)
		return nil, false
	}

	roles, err := scanRows(rows)
	return roles, true
}

func (r *role) GetOne(id int) (*types.Role, bool) {
	query := `select * from roles where id=$1`
	row := r.store.QueryRow(query, id)
	role, err := scanRow(row)
	if err != nil {
		log.Printf("Failed to query db for role %d due to %s", id, err)
		return nil, false
	}
	return role, true
}

func (r *role) Create(newRole *types.CreateRoleRequest) bool {
	query := `INSERT INTO roles(name, code, description)
	VALUES($1,$2, $3)
	`
	if _, err := r.store.Exec(query,
		newRole.Name,
		newRole.Code,
		newRole.Description,
	); err != nil {
		return false
	}
	return true
}

func (r *role) Update(id int, newRole *types.CreateRoleRequest) (*types.Role, bool) {
	query := `UPDATE roles SET name=$1, code=$2, description=$3 returning *`
	row := r.store.QueryRow(query,
		newRole.Name,
		newRole.Code,
		newRole.Description,
	)
	role, err := scanRow(row)
	if err != nil {
		log.Printf("failed to update role %d due to %s", id, err)
		return nil, false
	}
	return role, true
}

func (r *role) Delete(id int) bool {
	query := "UPDATE roles set deleted_at=$1 where id=$2"
	if _, err := r.store.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to delete role %d due to %s", id, err)
		return false
	}
	return true
}

func scanRows(rows *sql.Rows) ([]*types.Role, error) {
	roles := []*types.Role{}
	for rows.Next() {
		role := &types.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Code,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func scanRow(row *sql.Row) (*types.Role, error) {
	role := types.Role{}
	err := row.Scan(
		&role.ID,
		&role.Code,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, utils.NotFound
	}

	if err != nil {
		return nil, utils.ServerError
	}

	return &role, nil

}

func newRole(store *sql.DB) types.RoleRepository {
	return &role{
		store: store,
	}
}
