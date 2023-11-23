package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaPersona []Models.Persona

func ListarPersonas() ListaPersona {
	oLista := ListaPersona{}
	sqlQuery := `select * from listarpersona()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		opersona := Models.Persona{}
		rows.Scan(&opersona.Idpersona, &opersona.Nombrecompleto, &opersona.Nombretipopersona, &opersona.Fechanacimiento)
		opersona.Fechanacimientocadena = opersona.Fechanacimiento.Format("02/01/2006")
		oLista = append(oLista, opersona)
	}
	db.Close()
	return oLista
}
func FiltrarPersona(nombre string) ListaPersona {
	oLista := ListaPersona{}
	sqlQuery := `select * from filtrarpersona($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombre)
	for rows.Next() {
		opersona := Models.Persona{}
		rows.Scan(&opersona.Idpersona, &opersona.Nombrecompleto, &opersona.Nombretipopersona, &opersona.Fechanacimiento)
		opersona.Fechanacimientocadena = opersona.Fechanacimiento.Format("02/01/2006")
		oLista = append(oLista, opersona)
	}
	db.Close()
	return oLista
}
func ListarPersonaSinUsuario() ListaPersona {
	oLista := ListaPersona{}
	sqlQuery := `select * from listarpersonassinusuario()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		opersona := Models.Persona{}
		rows.Scan(&opersona.Idpersona, &opersona.Nombrecompleto)
		oLista = append(oLista, opersona)
	}
	db.Close()
	return oLista
}
func LeerPersona(idpersona int) Models.Persona {
	opersona := Models.Persona{}
	sqlQuery := "select * from recuperarpersona($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idpersona)
	for rows.Next() {
		rows.Scan(&opersona.Idpersona, &opersona.Nombre, &opersona.Appaterno, &opersona.Apmaterno,
			&opersona.Idtipopersona, &opersona.Fechanacimiento)
		opersona.Fechanacimientocadena = opersona.Fechanacimiento.Format("02/01/2006")
	}
	db.Close()
	return opersona
}
func InsertarPersona(opersona Models.Persona) (sql.Result, error) {

	db.Open()
	sql := "select insertarpersona($1,$2,$3,$4,$5)"
	result, err := db.Exec(sql, opersona.Nombre, opersona.Appaterno, opersona.Apmaterno,
		opersona.Idtipopersona, opersona.Fechanacimiento)
	db.Close()
	return result, err
}
func ActualizarPersona(opersona Models.Persona) (sql.Result, error) {

	db.Open()
	sql := "select actualizarpersona($1,$2,$3,$4,$5,$6)"
	result, err := db.Exec(sql, opersona.Idpersona, opersona.Nombre, opersona.Appaterno, opersona.Apmaterno,
		opersona.Idtipopersona, opersona.Fechanacimiento)
	db.Close()
	return result, err
}
func EliminarPersona(id int) (sql.Result, error) {
	db.Open()
	sql := "select eliminarpersona($1)"
	result, err := db.Exec(sql, id)
	db.Close()
	return result, err
}
