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

func GetAllRecords() ([]*VegetableRecord, error) {
	db := connectDB()
	defer db.Close()

	var vegetables []*VegetableRecord
	rows, err := db.Query("SELECT * FROM Vegetables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		vr := new(VegetableRecord)
		if err := rows.Scan(vr.ID, vr.Name, vr.Calories); err != nil {
			log.Fatal(err)
		}
	}
	return vegetables, nil
}

func GetRecord(myid int) []*VegetableRecord {
	db := connectDB()
	defer db.Close()
	v := new(VegetableRecord)
	db.QueryRow("SELECT * FROM Vegetables where id = myid").Scan(v.ID, v.Name, v.Calories)
	return []*VegetableRecord{v}
}

func AddRecord(myid int, name string, calories int, db *sql.DB) {
	defer db.Close()
	db.QueryRow("INSERT INTO Vegetables (id, name, calories) values (myid, name, calories)")
}
