package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type VegetableRecord struct {
	ID       int
	Name     string
	Calories int
}

func connectDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=diy_nutrition sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetAllRecords() ([]VegetableRecord, error) {
	db := connectDB()
	defer db.Close()

	var vegetables []VegetableRecord
	rows, err := db.Query("SELECT id, name, calories FROM Vegetables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		vr := new(VegetableRecord)
		if err := rows.Scan(&vr.ID, &vr.Name, &vr.Calories); err != nil {
			log.Fatal(err)
		}
		vegetables = append(vegetables, *vr)
	}
	return vegetables, nil
}

func GetRecord(myid int) ([]VegetableRecord, error) {
	var (
		id       int
		name     string
		calories int
	)

	db := connectDB()
	defer db.Close()
	queryStr := `SELECT id, name, calories FROM vegetables where id=$1`
	err := db.QueryRow(queryStr, myid).Scan(&id, &name, &calories)
	return []VegetableRecord{{ID: id, Name: name, Calories: calories}}, err
}

func AddRecord(name string, calories int) error {
	db := connectDB()
	defer db.Close()
	queryStr := `INSERT INTO vegetables (name, calories) values ($1, $2)`
	_, err := db.Exec(queryStr, name, calories)
	return err
}

func UpdateRecord(id int, name string, calories int) error {
	db := connectDB()
	defer db.Close()
	queryStr := `UPDATE vegetables SET name=$1, calories=$2 WHERE  id=$3`
	_, err := db.Exec(queryStr, name, calories, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteRecord(id int, name string, calories int) error {
	db := connectDB()
	defer db.Close()
	queryStr := `DELETE FROM vegetables WHERE id = $1`
	_, err := db.Exec(queryStr, id)
	return err
}
