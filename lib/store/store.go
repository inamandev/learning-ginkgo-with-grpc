package store

import (
	"database/sql"
	"log"

	pb "github.com/inamandev/learning-ginkgo-with-grpc/students_proto"
	_ "github.com/mattn/go-sqlite3"
)

type Pool struct {
	*sql.DB
}

var (
	DB *Pool
)

func Connect() error {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		return err
	}
	DB = &Pool{db}
	return DB.Ping()
}

func (db Pool) GetStudent(payload *pb.StudentPayload) error {
	stmt, err := db.Prepare(`SELECT name, age, email FROM students WHERE id = ?`)
	if err != nil {
		log.Println("unable to prepare statement: ", err)
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(payload.Id).Scan(&payload.Name, &payload.Age, &payload.Email)
	if err != nil {
		log.Println("unable to execute statement: ", err)
		return err
	}
	return nil
}
