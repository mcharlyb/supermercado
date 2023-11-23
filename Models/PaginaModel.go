package Models

type Pagina struct {
	Idpagina  int
	Mensaje   string
	Ruta      string
	Crear     bool
	Modificar bool
	Eliminar  bool
	Acceder   bool
	Errores   []string
}
