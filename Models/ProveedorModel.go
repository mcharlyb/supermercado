package Models

type Proveedor struct {
	Idproveedor        int
	Nombre             string
	Direccion          string
	Telefono           string
	Email              string
	Representantelegal string
	Celular            string
	Idpais             int
	Nombrepais         string
	Ruc                string
	Habilitado         bool
	Errores            []string
}
