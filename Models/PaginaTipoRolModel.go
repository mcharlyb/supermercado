package Models

type PaginaTipoRol struct {
	Idpaginatiporol int
	Idtiporol       int
	Idpagina        int
	Crear           bool
	Modificar       bool
	Eliminar        bool
	Errores         []string
}
