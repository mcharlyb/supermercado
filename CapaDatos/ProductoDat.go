package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaProductos []Models.Producto

func ListarProductos() ListaProductos {
	oLista := ListaProductos{}
	sqlQuery := `select * from uspListarProductos()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oproducto := Models.Producto{}
		rows.Scan(&oproducto.Idproducto, &oproducto.Nombreproducto, &oproducto.Precio,
			&oproducto.Stock, &oproducto.Nombrecategoria)
		oLista = append(oLista, oproducto)
	}
	db.Close()
	return oLista
}
func FiltrarProductosPorCategoria(idcategoria int) ListaProductos {
	oLista := ListaProductos{}
	sqlQuery := `select * from uspfiltrarproductosporcategoria($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, idcategoria)
	for rows.Next() {
		oproducto := Models.Producto{}
		rows.Scan(&oproducto.Idproducto, &oproducto.Nombreproducto, &oproducto.Precio, &oproducto.Stock, &oproducto.Nombrecategoria)
		oLista = append(oLista, oproducto)
	}
	db.Close()
	return oLista
}
func LeerProducto(idproducto int) Models.Producto {
	oproducto := Models.Producto{}
	sqlQuery := "select * from uspobtenerproducto($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idproducto)
	for rows.Next() {
		rows.Scan(&oproducto.Idproducto, &oproducto.Nombreproducto, &oproducto.Descripcionproducto,
			&oproducto.Precio, &oproducto.Stock, &oproducto.Idcategoria)
	}
	db.Close()
	return oproducto
}
func EliminarProducto(id int) (sql.Result, error) {
	db.Open()
	//sql := "select uspeliminarproductol($1)"
	sql := `update producto set habilitado=false where idproducto=($1)`
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
func ActualizarProducto(id int, nombre, descripcion string, precio float64,
	stock int, idcategoria int) (sql.Result, error) {
	db.Open()
	sql := "select uspactualizarproducto($1,$2,$3,$4,$5,$6)"
	result, err := db.Exec(sql, id, nombre, descripcion, precio, stock, idcategoria)
	db.Close()
	return result, err
}
func InsertarProducto(nombre, descripcion string, precio float64,
	stock int, idcategoria int) (sql.Result, error) {
	db.Open()
	sql := "select uspinsertarproducto($1,$2,$3,$4,$5)"
	result, err := db.Exec(sql, nombre, descripcion, precio, stock, idcategoria)
	db.Close()
	return result, err
}
