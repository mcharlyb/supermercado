package Models

import "errors"

type Categoria struct {
	Idcategoria  int
	Nombre       string
	Descripcion  string
	ExisteError  bool
	MensajeError string
}

func MaximoNombreCategoria(nombre string) error {
	if len(nombre) > 150 {
		return errors.New("el nombre de la categoria no puede tener mas de 150 caracteres")
	}
	return nil
}
func MaximoDescripcionCategoria(nombre string) error {
	if len(nombre) > 800 {
		return errors.New("la descripcion de la categoria puede tener un mÃ¡ximo de 800 caracteres")
	}
	return nil
}
