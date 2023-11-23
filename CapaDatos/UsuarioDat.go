package CapaDatos

import (
	"database/sql"
	"strconv"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
	"github.com/mcharlyb/cursogo/utilitarios"
)

type ListaUsuario []Models.Usuario

func ListarUsuarios() ListaUsuario {
	oLista := ListaUsuario{}
	//sqlQuery := `select * from Listarusuarios()`
	sqlQuery := `SELECT u.idusuario, u.nombreusuario,
	 p.nombre||' '||p.appaterno||' '||p.apmaterno , t.nombre
	FROM public.usuario u inner join persona p
	on u.idpersona = p.idpersona
	inner join tiporol t
	on u.idtiporol = t.idtiporol
	where u.habilitado=true`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		ousuario := Models.Usuario{}
		rows.Scan(&ousuario.Idusuario, &ousuario.Nombreusuario, &ousuario.Nombrepersona, &ousuario.Nombretiporol)
		oLista = append(oLista, ousuario)
	}
	db.Close()
	return oLista
}
func LeerUsuario(idusuario int) Models.Usuario {
	ousuario := Models.Usuario{}
	sqlQuery := "select * from recuperarusuario($1)"
	db.Open()
	rows, _ := db.Query(sqlQuery, idusuario)
	for rows.Next() {
		rows.Scan(&ousuario.Idusuario, &ousuario.Nombreusuario, &ousuario.Idpersona, &ousuario.Idtiporol, &ousuario.Nombretiporol,
			&ousuario.Nombrepersona)
	}
	db.Close()
	return ousuario
}

func ActualizarPersonaTransaccion(idpersona int, opcion bool, tx *sql.Tx) error {
	sqlQuery := "update persona set btieneusuario=$1 where idpersona=$2"
	_, err := tx.Exec(sqlQuery, opcion, idpersona)
	if err != nil {
		return err
	}
	return nil
}
func InsertarUsuarioTransaccion(ousuario Models.Usuario, tx *sql.Tx) error {

	clavecifrada := utilitarios.Cifrar(ousuario.Contra)
	sqlQuery := `insert into usuario(nombreusuario,contra,idpersona,idtiporol,habilitado)
	values($1,$2,$3,$4,true)`
	_, err := tx.Exec(sqlQuery, ousuario.Nombreusuario, clavecifrada,
		ousuario.Idpersona, ousuario.Idtiporol)
	if err != nil {
		return err
	}
	return nil
}
func RegistrarUsuarioTransaccion(ousuario Models.Usuario) error {
	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = ActualizarPersonaTransaccion(ousuario.Idpersona, true, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = InsertarUsuarioTransaccion(ousuario, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Close()
	return nil
}
func ActualizarUsuario(ousuario Models.Usuario) (sql.Result, error) {
	db.Open()
	sql := "select uspactualizarusuario($1,$2,$3)"
	result, err := db.Exec(sql, ousuario.Idusuario, ousuario.Nombreusuario, ousuario.Idtiporol)
	db.Close()
	return result, err
}

func EliminarUsuarioTransaccion(id int, tx *sql.Tx) error {
	sql := "select uspeliminarusuario($1)"
	_, err := tx.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
func RegistrarBajaUsuarioTransaccion(ousuario Models.Usuario) error {
	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = ActualizarPersonaTransaccion(ousuario.Idpersona, false, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = EliminarUsuarioTransaccion(ousuario.Idusuario, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Close()
	return nil
}
func Validar(usuario, contra string) string {
	var cantidad int
	sqlQuery := `select count(*) from usuario where nombreusuario=$1 and contra=$2`
	db.Open()
	err, cantidad := db.QueryRow(sqlQuery, usuario, contra)
	if err != nil {
		//panic(err)
		return "0"
	}
	return strconv.Itoa(cantidad)
}
func ObtenerIdUsuario(usuario, contra string) string {
	var idusuario int
	sqlQuery := `select idusuario from usuario where nombreusuario=$1 and contra=$2`
	db.Open()
	err, idusuario := db.QueryRow(sqlQuery, usuario, contra)
	if err != nil {
		//panic(err)
		return "0"
	}
	return strconv.Itoa(idusuario)
}
