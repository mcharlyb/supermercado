package Models

type Producto struct {
	Idproducto          int
	Nombreproducto      string
	Descripcionproducto string
	Precio              float64
	Stock               int
	Idcategoria         int
	Nombrecategoria     string
	ListaCategoria      []Categoria
	Errores             []string
}
