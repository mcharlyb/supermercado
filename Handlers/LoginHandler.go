package Handlers

import (
	"net/http"

	"github.com/mcharlyb/cursogo/utilitarios"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "login", nil)
}
func CerrarSesion(w http.ResponseWriter, r *http.Request) {
	utilitarios.EliminarCookie(w, "idusuario", "")
	utilitarios.Request(w, "login", nil)
}
func Pinicio(w http.ResponseWriter, r *http.Request) {
	utilitarios.Request(w, "index", nil)
}
