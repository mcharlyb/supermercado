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

func Usuario(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "usuario", nil)
}
func IngresarUsuario(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	usu := mapa["usu"]
	contra := mapa["contra"]
	cadenaCifrada := utilitarios.Cifrar(contra)
	encontro := CapaDatos.Validar(usu, cadenaCifrada)
	if encontro == "1" {
		//fmt.Fprintf(w, "1")
		idusuario := CapaDatos.ObtenerIdUsuario(usu, cadenaCifrada)
		utilitarios.CrearCookie(w, "idusuario", idusuario)
	}
	fmt.Fprintf(w, encontro)
}
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	listausuario := CapaDatos.ListarUsuarios()
	byteusuario, _ := json.Marshal(listausuario)
	fmt.Fprintf(w, string(byteusuario))
}
func ObtenerUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objusuario := CapaDatos.LeerUsuario(numid)
	byteusuario, _ := json.Marshal(objusuario)
	fmt.Fprintf(w, string(byteusuario))
}
func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "0")
	}
	objusuario := CapaDatos.LeerUsuario(numid)
	errores := CapaDatos.RegistrarBajaUsuarioTransaccion(objusuario)

	//_, errores := CapaDatos.EliminarPais(numid)
	if errores != nil {
		fmt.Fprintf(w, "0")
	}
	fmt.Fprintf(w, "1")
}

func GuardarUsuario(w http.ResponseWriter, r *http.Request) {
	oUsuario := Models.Usuario{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&oUsuario)

	if err != nil {
		panic("Ocurrio un Error")
	}

	if oUsuario.Idusuario == 0 {
		err := CapaDatos.RegistrarUsuarioTransaccion(oUsuario)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	} else {
		_, err := CapaDatos.ActualizarUsuario(oUsuario)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	}

}
