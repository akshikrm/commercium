package main

import (
	"akshidas/e-com/pkg/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"slices"
)

type database struct {
	store *sql.DB
}

func (s *database) INIT() {
	s.createTable()
	s.createFunction()
	s.createTrigger()
}

func (s *database) DROP() {
	s.dropTable()
	s.dropFunction()
	s.dropTrigger()
}

func (s *database) createTable() {
	for _, key := range KEYS {
		schema := SCHEMA[key]
		fmt.Printf("CREATING table %s...", key)
		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", key, schema)
		_, err := s.store.Exec(query)
		if err != nil {
			fmt.Println("FAILED")
			fmt.Println("ERR:", err.Error())
		} else {
			fmt.Println("SUCCESS")
		}
	}
}

func (s *database) dropTable() {
	slices.Reverse(KEYS)
	for _, key := range KEYS {
		fmt.Printf("DROPPING table %s...", key)
		query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", key)
		_, err := s.store.Exec(query)
		if err != nil {
			fmt.Println("FAILED")
			fmt.Println("ERR:", err.Error())
		} else {
			fmt.Println("SUCCESS")
		}
	}
	slices.Reverse(KEYS)
}

func (s *database) createFunction() {
	fmt.Print("CREATING function...")
	query := `CREATE FUNCTION update_updated_on_user_task() RETURNS TRIGGER AS
	$$ BEGIN NEW.updated_at = now(); RETURN NEW; END; $$ language 'plpgsql';`
	_, err := s.store.Exec(query)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Println("ERR:", err.Error())
	} else {
		fmt.Println("SUCCESS")
	}
}

func (s *database) dropFunction() {
	fmt.Print("DROPPING function...")
	query := "DROP FUNCTION IF EXISTS update_updated_on_user_task"
	_, err := s.store.Exec(query)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Println("ERR:", err.Error())
	} else {
		fmt.Println("SUCCESS")
	}
}

func (s *database) createTrigger() {
	for key := range SCHEMA {
		fmt.Printf("CREATING trigger...")
		query := fmt.Sprintf(`CREATE TRIGGER update_user_task_updated_on BEFORE UPDATE ON %s FOR EACH ROW EXECUTE PROCEDURE update_updated_on_user_task();`, key)
		_, err := s.store.Exec(query)
		if err != nil {
			fmt.Println("Failed")
			fmt.Printf("ERR: %s\n\n", err)
		} else {
			fmt.Println("SUCCESS")
		}
	}
}

func (s *database) dropTrigger() {
	for key := range SCHEMA {
		fmt.Printf("DROPPING trigger...")
		query := fmt.Sprintf("DROP TRIGGER IF EXISTS update_user_task_updated_on on %s", key)
		_, err := s.store.Exec(query)
		if err != nil {
			fmt.Println("Failed")
			fmt.Printf("ERR: %s\n\n", err)
		} else {
			fmt.Println("SUCCESS")
		}
	}
}

func NewDatabase(store *repository.Storage) *database {
	d := new(database)
	d.store = store.DB
	return d
}
