package Models

import "time"

type Persona struct {
	Idpersona             int
	Nombre                string
	Appaterno             string
	Apmaterno             string
	Nombrecompleto        string
	Fechanacimiento       time.Time
	Fechanacimientocadena string
	Idtipopersona         int
	Nombretipopersona     string
}
