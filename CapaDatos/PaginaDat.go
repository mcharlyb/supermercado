package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaPagina []Models.Pagina

func ListarPaginas() ListaPagina {
	oLista := ListaPagina{}
	sqlQuery := `select * from usplistarpagina()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		opagina := Models.Pagina{}
		rows.Scan(&opagina.Idpagina, &opagina.Mensaje, &opagina.Ruta)
		oLista = append(oLista, opagina)
	}
	db.Close()
	return oLista
}
func LeerPagina(idpagina int) Models.Pagina {
	opagina := Models.Pagina{}
	sqlQuery := "select * from recuperarpagina($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idpagina)
	for rows.Next() {
		rows.Scan(&opagina.Idpagina, &opagina.Mensaje, &opagina.Ruta)
	}
	db.Close()
	return opagina
}
func EliminarPagina(id int) (sql.Result, error) {
	db.Open()
	sql := "select eliminarpagina($1)"
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
func InsertarPagina(oPagina Models.Pagina) (sql.Result, error) {

	//validar
	/*var errorPais error
	errorCategoria = Models.MaximoNombrePais(opais.Nombre)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	errorCategoria = Models.MaximoCapitalpais(opais.Capital)
	if errorCategoria != nil {
		return nil, errorCategoria
	}*/
	db.Open()
	sql := "select insertarpagina($1,$2)"
	result, err := db.Exec(sql, oPagina.Mensaje, oPagina.Ruta)
	db.Close()
	return result, err
}
func ActualizarPagina(oPagina Models.Pagina) (sql.Result, error) {

	/*var errorPagina error
	errorCategoria = Models.MaximoNombreCategoria(oPagina.Nombre)
	if errorCategoria != nil {
		return nil, errorCategoria
	}
	errorCategoria = Models.MaximoDescripcionCategoria(oPagina.Capital)
	if errorCategoria != nil {
		return nil, errorCategoria
	}*/
	db.Open()
	sql := "select actualizarpagina($1,$2,$3)"
	result, err := db.Exec(sql, oPagina.Idpagina, oPagina.Mensaje, oPagina.Ruta)
	db.Close()
	return result, err
}
func BuscarPaginaPorUsuario(id int) []Models.Pagina {
	oLista := ListaPagina{}
	sqlQuery := `select pa.idpagina,pa.mensaje,pa.ruta,p.crear,p.modificar,p.eliminar
	from usuario u
	inner join tiporol r on u.idtiporol=r.idtiporol
	inner join paginatiporol p on p.idtiporol=r.idtiporol
	inner join pagina pa on pa.idpagina=p.idpagina
	where u.idusuario = $1 and pa.habilitado=true
	order by pa.mensaje
	`
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	for rows.Next() {
		opagina := Models.Pagina{}
		rows.Scan(&opagina.Idpagina, &opagina.Mensaje, &opagina.Ruta, &opagina.Crear, &opagina.Modificar, &opagina.Eliminar)
		opagina.Acceder = false
		oLista = append(oLista, opagina)
	}
	db.Close()
	return oLista
}
