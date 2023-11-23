package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaProveedor []Models.Proveedor

func ListarProveedor() ListaProveedor {
	oLista := ListaProveedor{}
	sqlQuery := `select * from listarproveedor()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oproveedor := Models.Proveedor{}
		rows.Scan(&oproveedor.Idproveedor, &oproveedor.Nombre, &oproveedor.Direccion, &oproveedor.Telefono, &oproveedor.Nombrepais)
		oLista = append(oLista, oproveedor)
	}
	db.Close()
	return oLista
}
func FiltrarProveedor(idpais int) ListaProveedor {
	oLista := ListaProveedor{}
	sqlQuery := `select * from filtrarproveedor($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, idpais)
	for rows.Next() {
		oproveedor := Models.Proveedor{}
		rows.Scan(&oproveedor.Idproveedor, &oproveedor.Nombre, &oproveedor.Direccion, &oproveedor.Telefono, &oproveedor.Nombrepais)
		oLista = append(oLista, oproveedor)
	}
	db.Close()
	return oLista
}
func LeerProveedor(idproveedor int) Models.Proveedor {
	oproveedor := Models.Proveedor{}
	sqlQuery := "select * from recuperarProveedor($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idproveedor)
	for rows.Next() {
		rows.Scan(&oproveedor.Idproveedor, &oproveedor.Nombre, &oproveedor.Direccion, &oproveedor.Telefono, &oproveedor.Email,
			&oproveedor.Representantelegal, &oproveedor.Celular, &oproveedor.Idpais, &oproveedor.Ruc)
	}
	db.Close()
	return oproveedor
}
func InsertarProveedor(oproveedor Models.Proveedor) (sql.Result, error) {

	db.Open()
	sql := "select insertarproveedor($1,$2,$3,$4,$5,$6,$7,$8)"
	result, err := db.Exec(sql, oproveedor.Nombre, oproveedor.Direccion, oproveedor.Telefono, oproveedor.Email,
		oproveedor.Representantelegal, oproveedor.Celular, oproveedor.Idpais, oproveedor.Ruc)
	db.Close()
	return result, err
}
func ActualizarProveedor(oproveedor Models.Proveedor) (sql.Result, error) {

	db.Open()
	sql := "select actualizarproveedor($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	result, err := db.Exec(sql, oproveedor.Idproveedor, oproveedor.Nombre, oproveedor.Direccion, oproveedor.Telefono, oproveedor.Email,
		oproveedor.Representantelegal, oproveedor.Celular, oproveedor.Idpais, oproveedor.Ruc)
	db.Close()
	return result, err
}
func EliminarProveedor(id int) (sql.Result, error) {
	db.Open()
	sql := "select eliminarproveedor($1)"
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
