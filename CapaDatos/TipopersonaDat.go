package CapaDatos

import (
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaTipoPersona []Models.TipoPersona

func ListarTipoPersonas() ListaTipoPersona {
	oLista := ListaTipoPersona{}
	sqlQuery := `select * from Listartipopersona()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		otipopersona := Models.TipoPersona{}
		rows.Scan(&otipopersona.Idtipopersona, &otipopersona.Nombre)
		oLista = append(oLista, otipopersona)
	}
	db.Close()
	return oLista
}
