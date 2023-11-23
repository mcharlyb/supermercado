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

func Proveedor(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "proveedor", nil)
}
func ListarProveedor(w http.ResponseWriter, r *http.Request) {
	listaproveedor := CapaDatos.ListarProveedor()
	byteproveedor, _ := json.Marshal(listaproveedor)
	fmt.Fprintf(w, string(byteproveedor))
}
func FiltrarProveedor(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	idpais := mapa["idpais"]
	numidpais, _ := strconv.Atoi(idpais)
	listaproveedor := CapaDatos.FiltrarProveedor(numidpais)
	byteproveedor, _ := json.Marshal(listaproveedor)
	fmt.Fprintf(w, string(byteproveedor))
}
func ObtenerProveedorPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objproveedor := CapaDatos.LeerProveedor(numid)
	byteproveedor, _ := json.Marshal(objproveedor)
	fmt.Fprintf(w, string(byteproveedor))
}
func GuardarProveedor(w http.ResponseWriter, r *http.Request) {
	oProveedor := Models.Proveedor{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&oProveedor)
	if err != nil {
		panic("Ocurrio un Error")
	}
	if oProveedor.Idproveedor == 0 {
		_, err := CapaDatos.InsertarProveedor(oProveedor)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := CapaDatos.ActualizarProveedor(oProveedor)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}
	fmt.Fprintf(w, "1")

}
func EliminarProveedor(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	_, errores := CapaDatos.EliminarProveedor(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	}
	fmt.Fprintf(w, "1")
}
