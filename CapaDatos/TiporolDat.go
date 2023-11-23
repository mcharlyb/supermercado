package CapaDatos

import (
	"database/sql"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

type ListaTipoRol []Models.TipoRol

func ListarTipoRol() ListaTipoRol {
	oLista := ListaTipoRol{}
	sqlQuery := `select idtiporol,nombre,descripcion from tiporol where habilitado=true`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		otiporol := Models.TipoRol{}
		rows.Scan(&otiporol.Idtiporol, &otiporol.Nombre, &otiporol.Descripcion)
		oLista = append(oLista, otiporol)
	}
	db.Close()
	return oLista
}
func ActualizarTipoRol(otiporol Models.TipoRol, tx *sql.Tx) error {

	sql := `update tiporol set nombre=$1,descripcion=$2 where idtiporol=$3`
	_, err := tx.Exec(sql, otiporol.Nombre, otiporol.Descripcion, otiporol.Idtiporol)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
func EliminarTipoRol(idtiporol int, tx *sql.Tx) error {

	sql := `update tiporol set habilitado=false where idtiporol=$1`
	_, err := tx.Exec(sql, idtiporol)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

func InsertarTipoRol(otiporol Models.TipoRol, tx *sql.Tx) (int, error) {
	var idtiporol int
	sql := `Insert into tiporol (nombre,descripcion,habilitado) values($1,$2,true) returning idtiporol`
	err := tx.QueryRow(sql, otiporol.Nombre, otiporol.Descripcion).Scan(&idtiporol)
	if err != nil {
		//panic(err)
		return 0, err
	}
	return idtiporol, nil
}
func BuscarPaginaTipoRolDB(opaginatiporol Models.PaginaTipoRol, tx *sql.Tx) (int, error) {
	var cantidad int
	sql := `select count(*) from paginatiporol where idtiporol=$1 and idpagina=$2`
	err := tx.QueryRow(sql, opaginatiporol.Idtiporol, opaginatiporol.Idpagina).Scan(&cantidad)
	if err != nil {
		//panic(err)
		return 0, err
	}
	return cantidad, nil
}

func InsertarPaginaTipoRol(opaginatiporol Models.PaginaTipoRol, tx *sql.Tx) error {
	sql := `Insert into paginatiporol (Idtiporol,Idpagina,Crear,Modificar,Eliminar,habilitado) values($1,$2,$3,$4,$5,true)`
	_, err := tx.Exec(sql, opaginatiporol.Idtiporol, opaginatiporol.Idpagina, opaginatiporol.Crear,
		opaginatiporol.Modificar, opaginatiporol.Eliminar)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
func LimpiarPaginaTiporol(otiporol Models.TipoRol, tx *sql.Tx) error {
	sql := `update paginatiporol set habilitado=false where idtiporol=$1`
	_, err := tx.Exec(sql, otiporol.Idtiporol)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
func ActualizarPaginaTipoDB(opaginatiporol Models.PaginaTipoRol, tx *sql.Tx) error {
	sql := `update paginatiporol set habilitado=true,Crear=$1,Modificar=$2,Eliminar=$3 where idtiporol=$4 and idpagina=$5`
	_, err := tx.Exec(sql, opaginatiporol.Crear, opaginatiporol.Modificar, opaginatiporol.Eliminar, opaginatiporol.Idtiporol, opaginatiporol.Idpagina)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
func ActualizarPaginaTipoRol(opaginatiporol Models.PaginaTipoRol, tx *sql.Tx) error {
	sql := `update paginatiporol set habilitado=false where idtiporol=$1`
	_, err := tx.Exec(sql, opaginatiporol.Idtiporol)
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

func RegistrarAltaTipoRol(otiporol Models.TipoRol, listarol []Models.PaginaTipoRol) error {
	var idtiporol int
	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	idtiporol, err = InsertarTipoRol(otiporol, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, valor := range listarol {
		valor.Idtiporol = idtiporol
		err = InsertarPaginaTipoRol(valor, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	db.Close()
	return nil
}
func ActualizarTipoRolItems(otiporol Models.TipoRol, listarol []Models.PaginaTipoRol) error {
	//var idtiporol int
	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = ActualizarTipoRol(otiporol, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = LimpiarPaginaTiporol(otiporol, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, valor := range listarol {
		//valor.Idtiporol = otiporol.Idtiporol
		cantidad, err := BuscarPaginaTipoRolDB(valor, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		if cantidad == 0 {
			err = InsertarPaginaTipoRol(valor, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			err = ActualizarPaginaTipoDB(valor, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	db.Close()
	return nil
}
func RegistrarBajaTipoRol(idtiporol int) error {
	otiporol := Models.TipoRol{Idtiporol: idtiporol}
	db.Open()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = EliminarTipoRol(idtiporol, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = LimpiarPaginaTiporol(otiporol, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	db.Close()
	return nil
}
