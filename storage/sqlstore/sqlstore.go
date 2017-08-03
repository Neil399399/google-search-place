package sqlstore

import (
	"database/sql"
	"fmt"
	"google-search-place/datamodel"

	_ "github.com/go-sql-driver/mysql"
)

//"root:123456@tcp(localhost:3306)/hello"
var (
	cof datamodel.Coffee
)

type WriteToSQL struct {
	serverurl string
	database  string
	db        *sql.DB
}

func NewWriteToSQL(username, password, serverurl, database string) *WriteToSQL {
	// Create the database handle, confirm driver is present
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,utf8", username, password, serverurl, database))
	if err != nil {
		fmt.Println("Connect Error!!", err)
		panic(err)
	}

	return &WriteToSQL{db: DB}
}

func (w *WriteToSQL) Read(id string) error {

	return nil
}
func (w *WriteToSQL) Write(data datamodel.Coffee) error {

	_, err := w.db.Exec("INSERT INTO CoffeeInfo (PlaceID,Name,Rate) VALUES (?,?,?)", data.Id, data.Name, data.Rate)
	if err != nil {
		fmt.Println("Write Info Error!!")
		panic(err)
	}

	for i := 1; i < len(data.Reviews); i++ {

		_, err = w.db.Exec("INSERT INTO CoffeeComment (PlaceID,Comment) VALUES (?,?)", data.Reviews[i].StoreId, data.Reviews[i].Text)
		if err != nil {
			fmt.Println("Write Comment Error!!")
			fmt.Println("Index: ", i)
			fmt.Println(data.Reviews[i].Text)
			panic(err)
		}

	}

	return nil
}

/*
func (w *WriteToSQL) ReadReviewsByID(id string) []string {

}
*/
