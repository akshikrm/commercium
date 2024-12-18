package main

import (
	"database/sql"
	"fmt"
	"slices"

	_ "github.com/lib/pq"
)

type Database struct {
	store *sql.DB
}

func (s *Database) INIT() {
	s.createTable()
	s.createFunction()
	s.createTrigger()
}

func (s *Database) DROP() {
	s.dropTable()
	s.dropFunction()
	s.dropTrigger()
}

func (s *Database) createTable() {
	for _, key := range KEYS {
		schema := SCHEMA[key]
		fmt.Printf("Creating table %s...", key)
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

func (s *Database) dropTable() {
	slices.Reverse(KEYS)
	for _, key := range KEYS {
		fmt.Printf("Dropping table %s...", key)
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

func (s *Database) createFunction() {
	fmt.Print("Creating function...")
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

func (s *Database) dropFunction() {
	fmt.Print("dropping function...")
	query := "DROP FUNCTION IF EXISTS update_updated_on_user_task"
	_, err := s.store.Exec(query)
	if err != nil {
		fmt.Println("FAILED")
		fmt.Println("ERR:", err.Error())
	} else {
		fmt.Println("SUCCESS")
	}
}

func (s *Database) createTrigger() {
	for key := range SCHEMA {
		fmt.Printf("Creating trigger...")
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

func (s *Database) dropTrigger() {
	for key := range SCHEMA {
		fmt.Printf("Dropping trigger...")
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
