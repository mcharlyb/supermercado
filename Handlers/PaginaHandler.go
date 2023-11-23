package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/utilitarios"

	"github.com/gorilla/mux"
)

func Pagina(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "pagina", nil)
}
func ListarPagina(w http.ResponseWriter, r *http.Request) {
	listapagina := CapaDatos.ListarPaginas()
	bytepagina, _ := json.Marshal(listapagina)
	fmt.Fprintf(w, string(bytepagina))
}
func ObtenerPaginaPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objpagina := CapaDatos.LeerPagina(numid)
	bytepagina, _ := json.Marshal(objpagina)
	fmt.Fprintf(w, string(bytepagina))
}
func GuardarPagina(w http.ResponseWriter, r *http.Request) {
	oPagina := Models.Pagina{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&oPagina)
	if err != nil {
		panic("Ocurrio un Error")
	}
	if oPagina.Idpagina == 0 {
		_, err := CapaDatos.InsertarPagina(oPagina)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	} else {
		_, err := CapaDatos.ActualizarPagina(oPagina)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	}

}
func EliminarPagina(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	_, errores := CapaDatos.EliminarPagina(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}

}
func LlenarMenu(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("idusuario")
	if err == http.ErrNoCookie {
		panic("no se encontro la cookie")
	}
	numid, _ := strconv.Atoi(cookie.Value)
	//fmt.Println(numid)
	//numid := 25
	objpagina := CapaDatos.BuscarPaginaPorUsuario(numid)
	bytepagina, _ := json.Marshal(objpagina)
	fmt.Fprintf(w, string(bytepagina))
}
