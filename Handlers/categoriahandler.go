package Handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/utilitarios"
)

type ListaCategoriaForm struct {
	ListaCategoria []Models.Categoria
	Nombre         string
}

func Categoria(w http.ResponseWriter, r *http.Request) {
	var nombreCategoria string
	var oCategoria []Models.Categoria
	if r.Method == "GET" {
		oCategoria = CapaDatos.ListarCategorias()
	} else {
		nombreCategoria = r.FormValue("nombre")
		oCategoria = CapaDatos.FiltrarCategorias(nombreCategoria)
	}
	obj := ListaCategoriaForm{ListaCategoria: oCategoria, Nombre: nombreCategoria}
	utilitarios.Request(w, "categoria", obj)
}

func CrearCategoria(w http.ResponseWriter, r *http.Request) {
	//idcategoria := 0;
	nombre := ""
	descripcion := ""
	categoria := Models.Categoria{Nombre: nombre, Descripcion: descripcion,
		ExisteError: false, MensajeError: ""}
	if r.Method == "GET" {
		utilitarios.Request(w, "nuevacategoria", categoria)
	} else {
		idcategoria := r.FormValue("idcategoria")
		nombre = r.FormValue("nombre")
		descripcion = r.FormValue("descripcion")
		categoria.Nombre = nombre
		categoria.Descripcion = descripcion
		//agregar categoria
		if idcategoria == "" {
			errorExist := utilitarios.ExisteInsertar("categoria", "nombre", nombre)
			if errorExist != nil {
				categoria.ExisteError = true
				categoria.MensajeError = errorExist.Error()
				utilitarios.Request(w, "nuevacategoria", categoria)
				return
			}
			_, err := CapaDatos.InsertarCategoria(nombre, descripcion)
			if err == nil {
				par := 301
				http.Redirect(w, r, "/categoria", par)
			} else {
				categoria.ExisteError = true
				categoria.MensajeError = err.Error()
				utilitarios.Request(w, "nuevacategoria", categoria)
				return
			}
		} else {
			num, _ := strconv.Atoi(idcategoria)
			categoria.Idcategoria = num
			errorExist := utilitarios.ExisteActualizar("categoria", "nombre", nombre, "idcategoria", num)
			//fmt.Println(errorExist)
			if errorExist != nil {
				categoria.ExisteError = true
				categoria.MensajeError = errorExist.Error()
				utilitarios.Request(w, "editarcategoria", categoria)
				return
			}

			_, err := CapaDatos.ActualizarCategoria(num, nombre, descripcion)
			if err == nil {
				par := 301
				http.Redirect(w, r, "/categoria", par)
			} else {
				categoria.ExisteError = true
				categoria.MensajeError = err.Error()
				utilitarios.Request(w, "editarcategoria", categoria)
				return
			}
		}
	}

}
func EditarCategoria(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	idcat, err := strconv.Atoi(id)
	if err != nil {
		panic("Ocurrio un error")
	}
	oCategoria := CapaDatos.LeerCategoria(idcat)

	utilitarios.Request(w, "editarcategoria", oCategoria)
}
func EliminarCategoria(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	idcat, err := strconv.Atoi(id)
	if err != nil {
		panic("Ocurrio un error")
	}
	//_, errorCategoria := CapaDatos.EliminarlogicaCategoria(idcat)
	errorCategoria := CapaDatos.RegistrarBajaCategoria(idcat)
	if errorCategoria == nil {
		par := 301
		http.Redirect(w, r, "/categoria", par)

	} else {
		panic("Ocurrio un error")
	}
}
