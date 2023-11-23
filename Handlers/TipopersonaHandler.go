package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcharlyb/cursogo/CapaDatos"
	"github.com/mcharlyb/cursogo/utilitarios"
)

func TipoPersona(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "tipopersona", nil)
}
func ListarTipoPersonas(w http.ResponseWriter, r *http.Request) {
	listatipopersona := CapaDatos.ListarTipoPersonas()
	byttipopersona, _ := json.Marshal(listatipopersona)
	fmt.Fprintf(w, string(byttipopersona))
}
