package db

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "chiruza1"
	dbname   = "bdsupermecado"
)

var cn *sql.DB

func Open() {
	conection := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic("ocurrio un error de coneccion")
	}
	cn = db
}
func Query(consulta string, args ...interface{}) (*sql.Rows, error) {
	rows, err := cn.Query(consulta, args...)
	if err != nil {
		return nil, errors.New("error al hacer la consulta")
	}
	return rows, err
}
func Begin() (*sql.Tx, error) {
	tx, err := cn.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func Close() {
	cn.Close()
}

// insert update delete
func Exec(consulta string, args ...interface{}) (sql.Result, error) {
	result, err := cn.Exec(consulta, args...)
	if err != nil {
		return nil, errors.New("error al insertar, actualizar o eliminar el registro")
	}
	return result, err
}
func QueryRow(consulta string, args ...interface{}) (error, int) {
	var cantidad int
	err := cn.QueryRow(consulta, args...).Scan(&cantidad)
	if err != nil {
		return err, -1
	}
	return nil, cantidad
}
