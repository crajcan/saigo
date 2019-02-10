package main

import (
	"database/sql"
	"fmt"
  "flag"
  "errors"
	_ "github.com/lib/pq"
)


func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

  //define flags
  insert := flag.Bool("new", false, "Insert a new user")
  destroy_id := flag.Int("delete", 0, "A person_id to delete from the db")
  update_id := flag.Int("update_id", 0, "A person_id to update")
  ssn := flag.Int("ssn", 0, "A new ssn for a person")
  name := flag.String("name", "", "A name to insert")

  //parse flags
  flag.Parse() 
 
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)
 

  //insert new record with name & ssn
  if *insert {
    if ((*ssn != 0) && (*name != "")){
      var id int
      err := db.QueryRow("INSERT INTO people(name, ssn) VALUES ($1,$2) RETURNING person_id", *name, *ssn).Scan(&id)
      PanicOn(err)
      fmt.Println("New record created with id: ", id)
    } else {
      err = errors.New("Please provide an ssn and a name for a new person")
      PanicOn(err)
    }
  }


  //update relevent columns when flag is set
  if *update_id != 0 { 
    if (*ssn != 0 && *name != "") {
      _, err = db.Exec("UPDATE people SET ssn = $1, name = $2 WHERE person_id = $3", *ssn, *name, *update_id)
    } else if (*ssn != 0) {
      _, err = db.Exec("UPDATE people SET ssn=$1 WHERE person_id=$2", *ssn, *update_id)  
    }  else if (*name != "") {
      _, err = db.Exec("UPDATE people SET name=$1 WHERE person_id=$2", *name, *update_id)  
    } 
  }
	PanicOn(err)


  //delete record if flag is set
  if (*destroy_id != 0) {
    _, err = db.Exec("DELETE FROM people WHERE (person_id=$1)", *destroy_id)
  }  
	PanicOn(err)


  //get rows 
	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)
  defer rows.Close()

  //print rows
	for rows.Next() {
		var id int
		var name string
		var ssn int
		err := rows.Scan(&id, &name, &ssn)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", id, name, ssn)
	}
}
