package CapaDatos

import (
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaPaginaTipoRol []Models.PaginaTipoRol

func FiltrarPaginaTipoRolPorId(id int) ListaPaginaTipoRol {
	oLista := ListaPaginaTipoRol{}
	sqlQuery := `select idpagina,crear,modificar,eliminar from paginatiporol where idtiporol=$1 and habilitado=true`
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	for rows.Next() {
		opaginatiporol := Models.PaginaTipoRol{}
		rows.Scan(&opaginatiporol.Idpagina, &opaginatiporol.Crear, &opaginatiporol.Modificar, &opaginatiporol.Eliminar)
		oLista = append(oLista, opaginatiporol)
	}
	db.Close()
	return oLista
}

/*
	func ListarPaginaTipoRol() ListaPaginaTipoRol {
		oLista := ListaPaginaTipoRol{}
		sqlQuery := `select listarpaginatiporol()`
		db.Open()
		rows, _ := db.Query(sqlQuery)
		for rows.Next() {
			opaginatiporol := Models.PaginaTipoRol{}
			rows.Scan(&opaginatiporol.Idpaginatiporol, &opaginatiporol., &opaginatiporol.)
			oLista = append(oLista, opaginatiporol)
		}
		db.Close()
		return oLista
	}
*/
func LeerTipoRolCabecera(id int) Models.TipoRol {
	otiporol := Models.TipoRol{}
	sqlQuery := "select idtiporol,nombre,descripcion from tiporol where idtiporol=$1"
	db.Open()
	rows, _ := db.Query(sqlQuery, id)
	for rows.Next() {
		rows.Scan(&otiporol.Idtiporol, &otiporol.Nombre, &otiporol.Descripcion)
	}
	db.Close()
	return otiporol
}
