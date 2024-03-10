package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "Ulyasar10389#"
	dbname   = "user"
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic((err))
	}

	fmt.Println("succesfully connected to database")
	CreateEmployee()
	GetEmployees()
	UpdateEmployee()
	DeleteEmployee()

}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (email) DO NOTHING
	Returning * `

	err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No duplicate email found, data inserted successfully")
		} else {
			panic(err)
		}
	} else {
		fmt.Println("Email already exists in the database, no data inserted")

	}
}

func GetEmployees() {
	var results = []Employee{}
	sqlStatement := `SELECT * from employees`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}
		results = append(results, employee)
	}
	fmt.Println("Employee datas:", results)

}

func UpdateEmployee() {
	sqlStatement := `
        UPDATE employees
        SET full_name = $2, email = $3, division = $4, age = $5
        WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airellhidayat@gmail.com", "CurDevs", 24)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Jumlah data yang diupdate:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}

/*CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL
); */

/*CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); */
/*
CREATE TABLE books (
    ID          INT PRIMARY KEY AUTO_INCREMENT,
    Title       VARCHAR(255) NOT NULL,
    Description TEXT,
    ImageURL    VARCHAR(255),
    ReleaseYear INT,
    Price       DECIMAL(10, 2),
    TotalPage   INT,
    Thickness   VARCHAR(50),
    CreatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CategoryID  INT,
    FOREIGN KEY (CategoryID) REFERENCES categories(id)
); */

/*
CREATE TABLE books (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    imageURL    VARCHAR(255),
    releaseYear INT,
    price       DECIMAL(10, 2),
    totalPage   INT,
    thickness   VARCHAR(50),
    created_At   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_At   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    categoryID  INT,
    FOREIGN KEY (CategoryID) REFERENCES categories(id)
); */

/*
CREATE TABLE books (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    imageURL    VARCHAR(255),
    releaseYear INT,
    price       DECIMAL(10, 2),
    totalPage   INT,
    thickness   VARCHAR(50),
    create_dAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_At   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    categoryID  INT,
    FOREIGN KEY (CategoryID) REFERENCES categories(id)
);
*/
