package Handlers

import (
	"net/http"

	"github.com/mcharlyb/cursogo/utilitarios"
)

func Principal(w http.ResponseWriter, r *http.Request) {

	utilitarios.Request(w, "index", nil)
}
