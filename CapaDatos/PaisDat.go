package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaPais []Models.Pais

func ListarPaises() ListaPais {
	oLista := ListaPais{}
	sqlQuery := `select * from ListarPais()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		opais := Models.Pais{}
		rows.Scan(&opais.Idpais, &opais.Nombre, &opais.Capital)
		oLista = append(oLista, opais)
	}
	db.Close()
	return oLista
}
func FiltrarPaises(nombre string) ListaPais {
	oLista := ListaPais{}
	sqlQuery := `select * from filtrarpais($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombre)
	for rows.Next() {
		opais := Models.Pais{}
		rows.Scan(&opais.Idpais, &opais.Nombre, &opais.Capital)
		oLista = append(oLista, opais)
	}
	db.Close()
	return oLista
}
func LeerPais(idpais int) Models.Pais {
	opais := Models.Pais{}
	sqlQuery := "select * from uspobtenerpais($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idpais)
	for rows.Next() {
		rows.Scan(&opais.Idpais, &opais.Nombre, &opais.Capital)
	}
	db.Close()
	return opais
}
func EliminarPais(id int) (sql.Result, error) {
	db.Open()
	sql := "select eliminarpais($1)"
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
func InsertarPais(opais Models.Pais) (sql.Result, error) {

	db.Open()
	sql := "select insertarpais($1,$2)"
	result, err := db.Exec(sql, opais.Nombre, opais.Capital)
	db.Close()
	return result, err
}
func ActualizarPais(opais Models.Pais) (sql.Result, error) {

	db.Open()
	sql := "select actualizarpais($1,$2,$3)"
	result, err := db.Exec(sql, opais.Idpais, opais.Nombre, opais.Capital)
	db.Close()
	return result, err
}
