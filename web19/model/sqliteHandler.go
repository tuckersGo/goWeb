package model

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) getTodos() []*Todo {
	return nil
}

func (s *sqliteHandler) addTodo(name string) *Todo {
	return nil
}

func (s *sqliteHandler) removeTodo(id int) bool {
	return false
}

func (s *sqliteHandler) completeTodo(id int, complete bool) bool {
	return false
}

func (s *sqliteHandler) close() {
	s.db.Close()
}

func newSqliteHandler() dbHandler {
	os.Remove("./test.db")
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	statement.Exec()
	return &sqliteHandler{db: database}
}
