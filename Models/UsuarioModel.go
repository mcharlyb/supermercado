package Models

type Usuario struct {
	Idusuario     int
	Nombreusuario string
	Contra        string
	Idpersona     int
	Idtiporol     int
	Nombrepersona string
	Nombretiporol string
	Errores       []string
}
