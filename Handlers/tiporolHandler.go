package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/utilitarios"
)

func TipoRol(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "tiporol", nil)
}
func ListarTipoRol(w http.ResponseWriter, r *http.Request) {
	listatiporol := CapaDatos.ListarTipoRol()
	byttiporol, _ := json.Marshal(listatiporol)
	fmt.Fprintf(w, string(byttiporol))
}
func GuardarTipoRol(w http.ResponseWriter, r *http.Request) {
	otiporol := Models.TipoRol{}
	listarol := []Models.PaginaTipoRol{}
	datatiporol := json.NewDecoder(r.Body)
	err := datatiporol.Decode(&otiporol)
	if err != nil {
		panic("ocurrio un error en el objeto")
	}
	if otiporol.Idpaginascadena != "" {
		cadenaIds := strings.Split(otiporol.Idpaginascadena, "*")
		for _, valor := range cadenaIds {
			valid := valor[3:]
			entero, _ := strconv.Atoi(valid)
			var varcrea, varmodi, vardele bool
			if string(valor[0]) == "t" {
				varcrea = true
			} else {
				varcrea = false
			}
			if string(valor[1]) == "t" {
				varmodi = true
			} else {
				varmodi = false
			}
			if string(valor[2]) == "t" {
				vardele = true
			} else {
				vardele = false
			}

			listarol = append(listarol, Models.PaginaTipoRol{Idpagina: entero, Idtiporol: otiporol.Idtiporol,
				Crear: varcrea, Modificar: varmodi, Eliminar: vardele})
		}
	}
	if otiporol.Idtiporol == 0 {
		err := CapaDatos.RegistrarAltaTipoRol(otiporol, listarol)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	} else {
		err := CapaDatos.ActualizarTipoRolItems(otiporol, listarol)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}

	}
}
func ObtenerTipoRolPorId(w http.ResponseWriter, r *http.Request) {
	mapa := mux.Vars(r)
	id := mapa["id"]
	numid, _ := strconv.Atoi(id)
	objtiporol := CapaDatos.LeerTipoRolCabecera(numid)
	bytetiporol, _ := json.Marshal(objtiporol)
	fmt.Fprintf(w, string(bytetiporol))
}

/*func GuardarTipoRol(w http.ResponseWriter, r *http.Request) {
	otiporol := Models.TipoRol{
		Idtiporol:   0,
		Nombre:      "Supervisor",
		Descripcion: "Personal que supervisa tareas Generales",
	}
	listarol := []Models.PaginaTipoRol{
		Models.PaginaTipoRol{Idpagina: 1},
		Models.PaginaTipoRol{Idpagina: 2},
	}
	if otiporol.Idtiporol == 0 {
		err := CapaDatos.RegistrarAltaTipoRol(otiporol, listarol)
		if err != nil {
			fmt.Fprintf(w, "0")
		} else {
			fmt.Fprintf(w, "1")
		}
	}
}*/
