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

func Persona(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "persona", nil)
}
func ListarPersonas(w http.ResponseWriter, r *http.Request) {
	listapersona := CapaDatos.ListarPersonas()
	bytepersona, _ := json.Marshal(listapersona)
	fmt.Fprintf(w, string(bytepersona))
}
func ListarPersonaSinUsuario(w http.ResponseWriter, r *http.Request) {
	listapersonasu := CapaDatos.ListarPersonaSinUsuario()
	bytepersonasu, _ := json.Marshal(listapersonasu)
	fmt.Fprintf(w, string(bytepersonasu))
}
func FiltrarPersona(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	nombre := mapa["nombre"]
	listapersona := CapaDatos.FiltrarPersona(nombre)
	bytepersona, _ := json.Marshal(listapersona)
	fmt.Fprintf(w, string(bytepersona))
}
func ObtenerPersonaPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objpersona := CapaDatos.LeerPersona(numid)
	bytepersona, _ := json.Marshal(objpersona)
	fmt.Fprintf(w, string(bytepersona))
}
func GuardarPersona(w http.ResponseWriter, r *http.Request) {
	oPersona := Models.Persona{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&oPersona)
	if err != nil {
		panic("Ocurrio un Error")
	}
	if oPersona.Idpersona == 0 {
		_, err := CapaDatos.InsertarPersona(oPersona)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	} else {
		_, err := CapaDatos.ActualizarPersona(oPersona)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}
	fmt.Fprintf(w, "1")

}
func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	_, errores := CapaDatos.EliminarPersona(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	}
	fmt.Fprintf(w, "1")
}
