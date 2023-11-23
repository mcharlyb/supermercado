package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaCategoria []Models.Categoria

func ListarCategorias() ListaCategoria {
	oLista := ListaCategoria{}
	sqlQuery := `select * from uspListarCategorias()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		ocategoria := Models.Categoria{}
		rows.Scan(&ocategoria.Idcategoria, &ocategoria.Nombre, &ocategoria.Descripcion)
		oLista = append(oLista, ocategoria)
	}
	db.Close()
	return oLista
}
func LeerCategoria(idcategoria int) Models.Categoria {
	ocategoria := Models.Categoria{}
	sqlQuery := "select * from uspbuscarcategoriaporid($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idcategoria)
	for rows.Next() {
		rows.Scan(&ocategoria.Idcategoria, &ocategoria.Nombre, &ocategoria.Descripcion)
	}
	db.Close()
	return ocategoria
}

func FiltrarCategorias(Nombrecategoria string) ListaCategoria {
	oLista := ListaCategoria{}
	sqlQuery := `select * from uspFiltrarCategorias($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, Nombrecategoria)
	for rows.Next() {
		ocategoria := Models.Categoria{}
		rows.Scan(&ocategoria.Idcategoria, &ocategoria.Nombre, &ocategoria.Descripcion)
		oLista = append(oLista, ocategoria)
	}
	db.Close()
	return oLista
}

func InsertarCategoria(nombre, descripcion string) (sql.Result, error) {

	//validar
	var errorCategoria error
	errorCategoria = Models.MaximoNombreCategoria(nombre)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	errorCategoria = Models.MaximoDescripcionCategoria(descripcion)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	db.Open()
	sql := "select insertarcategoria($1,$2)"
	result, err := db.Exec(sql, nombre, descripcion)
	db.Close()
	return result, err
}
func ActualizarCategoria(id int, nombre, descripcion string) (sql.Result, error) {

	var errorCategoria error
	errorCategoria = Models.MaximoNombreCategoria(nombre)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	errorCategoria = Models.MaximoDescripcionCategoria(descripcion)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	db.Open()
	sql := "select actualizarcategoria($1,$2,$3)"
	result, err := db.Exec(sql, id, nombre, descripcion)
	db.Close()
	return result, err
}

/*func EliminarlogicaCategoria(id int) (sql.Result, error) {
	db.Open()
	//sql := "select uspeliminarlogicacategoria($1)"
	sql := "select eliminarcategoriayproductos($1)"
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}*/

func RegistrarBajaCategoria(id int) error {

	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = EliminarCategoriaLogica(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = EliminarProductosCategoriaLogica(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Close()
	return nil
}
func EliminarCategoriaLogica(id int, tx *sql.Tx) error {
	sql := `update categoria set bhabilitado=false where idcategoria=$1`
	_, err := tx.Exec(sql, id)
	if err != nil {

		return err
	}
	return nil
}
func EliminarProductosCategoriaLogica(id int, tx *sql.Tx) error {
	sql := `update producto set habilitado=false where idcategoria=$1`
	_, err := tx.Exec(sql, id)
	if err != nil {

		return err
	}
	return nil
}
