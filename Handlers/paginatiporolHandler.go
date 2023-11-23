package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/utilitarios"
)

func FiltrarPaginaTipoRolPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]

	num, _ := strconv.Atoi(id)
	listapaginatiporol := CapaDatos.FiltrarPaginaTipoRolPorId(num)
	bytepaginatiporol, _ := json.Marshal(listapaginatiporol)
	fmt.Fprintf(w, string(bytepaginatiporol))
}
func PaginaTipoRol(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "paginatiporol", nil)
}
func NuevaPaginaTipoRol(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "nuevapaginatiporol", nil)
}
func EditarPaginaTipoRol(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	num, _ := strconv.Atoi(id)

	otiporol := Models.TipoRol{Idtiporol: num}

	utilitarios.Request(w, "editarpaginatiporol", otiporol)
}

func alert(otiporol Models.TipoRol) {
	panic("unimplemented")
}
func EliminarTipoRol(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	errores := CapaDatos.RegistrarBajaTipoRol(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}

}
