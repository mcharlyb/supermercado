package utilitarios

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"

	"github.com/mcharlyb/cursogo/Models"
	"github.com/mcharlyb/cursogo/db"
)

func GenerarURL(uri, host, protocolo string, mapa map[string]string) string {
	u, _ := url.Parse(uri)
	u.Host = host
	u.Scheme = protocolo
	mapaFuncion := u.Query()
	for key, value := range mapa {
		mapaFuncion.Add(key, value)
	}
	u.RawQuery = mapaFuncion.Encode()
	return u.String()
}
func RequestURL(metodo, url string) string {

	r, e := http.NewRequest(metodo, url, nil)
	if e != nil {
		panic("Hubo un error en el request")
	}
	cliente := &http.Client{}
	resp, e := cliente.Do(r)
	if e != nil {
		panic("Hubo un error en el cliente")
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Hubo un error al leer")
	}
	return string(bytes)

}

var TemplateTodos = template.Must(template.New("T").Funcs(mapa).
	ParseGlob("../html/**/*.html"))

var templateError = template.Must(template.ParseFiles("../html/error/error.html"))

func Request(w http.ResponseWriter, nombrepagina string, estructura interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := TemplateTodos.ExecuteTemplate(w, nombrepagina, estructura)
	if err != nil {
		w.WriteHeader(500)
		templateError.Execute(w, nil)
	}

}

func Saludar(nombre string) {
	fmt.Printf("Hola amigo %s como estas", nombre)
}

func Bienvenido() string {
	return "Bienvenido a la Página"
}

var mapa = template.FuncMap{
	"Bienvenido": Bienvenido,
}

func Cifrar(claveAcifrar string) string {
	bytecifrado := sha256.Sum256([]byte(claveAcifrar))
	claveCifrada := hex.EncodeToString(bytecifrado[:])
	return claveCifrada
}

func Required(valor, nombrecampo string) error {
	if valor == "" {
		return errors.New("el campo " + nombrecampo + " debe ingresarse obligatoiramente")
	}
	return nil
}
func MaxLength(valor, nombrecampo string, longitudMaxima int) error {
	if len(valor) > longitudMaxima {
		return errors.New("la longitud maxima del campo " + nombrecampo + " es " + strconv.Itoa(longitudMaxima))
	}
	return nil
}
func MinLength(valor, nombrecampo string, longitudMinima int) error {
	if len(valor) < longitudMinima {
		return errors.New("la longitud minima del campo " + nombrecampo + " es " + strconv.Itoa(longitudMinima))
	}
	return nil
}
func ValidarEntero(valor, nombrecampo string) error {
	_, err := strconv.Atoi(valor)
	if err != nil {
		return errors.New("El valor del campo " + nombrecampo + " debe ser entero ")
	}
	return nil
}
func ValidarDecimal(valor, nombrecampo string) error {
	_, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		return errors.New("El valor del campo " + nombrecampo + " debe ser numerico decimal ")
	}
	return nil
}
func ExisteInsertar(tabla, campo, valor string) error {

	sql := "select count(*) from " + tabla + " where upper(" + campo + ") ='" + strings.ToUpper(valor) + "'"

	var cantidad int
	db.Open()
	rows, _ := db.Query(sql)
	for rows.Next() {
		rows.Scan(&cantidad)

	}
	db.Close()
	if cantidad > 0 {

		return errors.New("Ya existe el campo " + campo + " llamado " + valor)
	}
	return nil

}
func ExisteActualizar(tabla, campo, valor, campoid string, id int) error {

	sql := "select count(*) from " + tabla + " where upper(" +
		campo + ") ='" + strings.ToUpper(valor) + "' and " + campoid +
		"!=" + strconv.Itoa(id)

	var cantidad int
	db.Open()
	rows, _ := db.Query(sql)
	for rows.Next() {
		rows.Scan(&cantidad)

	}
	db.Close()
	if cantidad > 0 {

		return errors.New("Ya existe el campo " + campo + " llamado " + valor)
	}
	return nil

}
func CrearCookie(w http.ResponseWriter, nombre string, valor string) {
	cookie := &http.Cookie{
		Name:  nombre,
		Value: valor,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}
func EliminarCookie(w http.ResponseWriter, nombre string, valor string) {
	cookie := &http.Cookie{
		Name:   nombre,
		Value:  valor,
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
func Acceso(w http.ResponseWriter, r *http.Request, opagina Models.Pagina) {
	if !opagina.Acceder {
		par := 301
		http.Redirect(w, r, "/cerrarsesion", par)
	}

}
