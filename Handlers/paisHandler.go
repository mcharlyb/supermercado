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

func Pais(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "pais", nil)
}
func ListarPais(w http.ResponseWriter, r *http.Request) {
	listapais := CapaDatos.ListarPaises()
	bytepais, _ := json.Marshal(listapais)
	fmt.Fprintf(w, string(bytepais))
}
func FiltrarPais(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	nombre := mapa["nombre"]
	listapais := CapaDatos.FiltrarPaises(nombre)
	bytepais, _ := json.Marshal(listapais)
	fmt.Fprintf(w, string(bytepais))
}
func ObtenerPaisPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objpais := CapaDatos.LeerPais(numid)
	bytepais, _ := json.Marshal(objpais)
	fmt.Fprintf(w, string(bytepais))
}
func GuardarPais(w http.ResponseWriter, r *http.Request) {
	oPais := Models.Pais{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&oPais)
	if err != nil {
		panic("Ocurrio un Error")
	}
	if oPais.Idpais == 0 {
		_, err := CapaDatos.InsertarPais(oPais)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := CapaDatos.ActualizarPais(oPais)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}
	fmt.Fprintf(w, "1")

}
func EliminarPais(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	_, errores := CapaDatos.EliminarPais(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}

}
