package Handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/utilitarios"
)

type ListaProductoForm struct {
	ListaCategoria []Models.Categoria
	ListaProducto  []Models.Producto
	IdCategoria    string
}

func Producto(w http.ResponseWriter, r *http.Request) {
	var oproducto []Models.Producto
	ocategoria := CapaDatos.ListarCategorias()
	idcategoria := r.FormValue("idcategoria")
	if r.Method == "GET" {
		oproducto = CapaDatos.ListarProductos()
	} else {
		if idcategoria == "" {
			oproducto = CapaDatos.ListarProductos()
		} else {
			num, err := strconv.Atoi(idcategoria)
			if err != nil {
				panic("No se puede Convertir")
			} else {
				oproducto = CapaDatos.FiltrarProductosPorCategoria(num)
			}
		}
	}

	oListaProductoForm := ListaProductoForm{ListaCategoria: ocategoria, ListaProducto: oproducto, IdCategoria: idcategoria}
	utilitarios.Request(w, "producto", oListaProductoForm)
}
func CrearProducto(w http.ResponseWriter, r *http.Request) {

	listaCategoria := CapaDatos.ListarCategorias()
	oProducto := Models.Producto{Precio: 0, Stock: 0, ListaCategoria: listaCategoria}
	utilitarios.Request(w, "crearproducto", oProducto)
}
func EditarProducto(w http.ResponseWriter, r *http.Request) {
	listaCategoria := CapaDatos.ListarCategorias()
	mapa := mux.Vars(r)
	idproducto := mapa["id"]
	num, _ := strconv.Atoi(idproducto)
	oProductoObjeto := CapaDatos.LeerProducto(num)
	oProductoObjeto.ListaCategoria = listaCategoria
	//oProducto := Models.Producto{ListaCategoria: listaCategoria}
	utilitarios.Request(w, "editarproducto", oProductoObjeto)
}
func EliminarProducto(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	idproducto := mapa["id"]
	num, _ := strconv.Atoi(idproducto)
	_, err := CapaDatos.EliminarProducto(num)
	if err == nil {
		par := 301
		http.Redirect(w, r, "/productos", par)
	} else {
		fmt.Println("error al bajar")
	}
}
func GrabarProducto(w http.ResponseWriter, r *http.Request) {
	idproducto := r.FormValue("idproducto")
	nombre := r.FormValue("nombre")
	descripcion := r.FormValue("descripcion")
	precio := r.FormValue("precio")
	preciofloat, _ := strconv.ParseFloat(precio, 64)
	stock := r.FormValue("stock")
	numstock, _ := strconv.Atoi(stock)
	categoria := r.FormValue("categoria")
	numidcategoria, _ := strconv.Atoi(categoria)
	var errorMensaje []string
	errorNombre := utilitarios.MaxLength(nombre, "nombre", 100)
	if errorNombre != nil {
		errorMensaje = append(errorMensaje, errorNombre.Error())
	}
	errorNombremin := utilitarios.MinLength(nombre, "nombre", 3)
	if errorNombremin != nil {
		errorMensaje = append(errorMensaje, errorNombremin.Error())
	}
	errorDescripcion := utilitarios.MaxLength(descripcion, "descripcion", 200)
	if errorDescripcion != nil {
		errorMensaje = append(errorMensaje, errorDescripcion.Error())
	}
	errorDescripcionmin := utilitarios.MinLength(descripcion, "descripcion", 4)
	if errorDescripcionmin != nil {
		errorMensaje = append(errorMensaje, errorDescripcionmin.Error())
	}
	errorStock := utilitarios.ValidarEntero(stock, "stock")
	if errorStock != nil {
		errorMensaje = append(errorMensaje, errorStock.Error())
	}
	errorPrecio := utilitarios.ValidarDecimal(precio, "precio")
	if errorPrecio != nil {
		errorMensaje = append(errorMensaje, errorPrecio.Error())
	}
	listaCategoria := CapaDatos.ListarCategorias()
	oproducto := Models.Producto{Nombreproducto: nombre,
		Descripcionproducto: descripcion, Precio: preciofloat, Stock: numstock,
		Idcategoria: numidcategoria}
	oproducto.ListaCategoria = listaCategoria
	if idproducto == "" {
		//Insertar
		if len(errorMensaje) > 0 {
			oproducto.Errores = errorMensaje
			listaCategoria := CapaDatos.ListarCategorias()
			oproducto.ListaCategoria = listaCategoria

			utilitarios.Request(w, "crearproducto", oproducto)
			return
		}
		errorExist := utilitarios.ExisteInsertar("producto", "nombre", nombre)
		if errorExist != nil {
			errorMensaje = append(errorMensaje, errorExist.Error())
			oproducto.Errores = errorMensaje
			utilitarios.Request(w, "crearproducto", oproducto)
			return
		}

		_, err := CapaDatos.InsertarProducto(nombre, descripcion, preciofloat, numstock, numidcategoria)
		if err == nil {
			var par = 301
			http.Redirect(w, r, "/productos", par)
		} else {
			utilitarios.Request(w, "crearproducto", oproducto)
			return
		}

	} else {

		//Actualizar
		numidproducto, _ := strconv.Atoi(idproducto)
		oproducto.Idproducto = numidproducto
		if len(errorMensaje) > 0 {
			oproducto.Errores = errorMensaje
			listaCategoria := CapaDatos.ListarCategorias()
			oproducto.ListaCategoria = listaCategoria
			utilitarios.Request(w, "editarproducto", oproducto)
			return
		}
		errorExist := utilitarios.ExisteActualizar("producto", "nombre", nombre, "idproducto", numidproducto)
		if errorExist != nil {
			errorMensaje = append(errorMensaje, errorExist.Error())
			oproducto.Errores = errorMensaje
			utilitarios.Request(w, "editarproducto", oproducto)
			return
		}

		_, err := CapaDatos.ActualizarProducto(numidproducto, nombre, descripcion, preciofloat, numstock, numidcategoria)
		if err == nil {
			var par = 301
			http.Redirect(w, r, "/productos", par)
		} else {
			utilitarios.Request(w, "editarproducto", oproducto)
			return
		}
	}
}
