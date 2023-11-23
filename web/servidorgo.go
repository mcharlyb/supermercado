package main

import (
	"net/http"

	"github.com/mcharlyb/cursogo/Handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	r := mux.NewRouter()
	archivos := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", archivos))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	//producto
	r.HandleFunc("/productos", Handlers.Producto)
	r.HandleFunc("/productos/nuevo", Handlers.CrearProducto)
	r.HandleFunc("/productos/guardar", Handlers.GrabarProducto)
	r.HandleFunc("/productos/editar/{id}", Handlers.EditarProducto)
	r.HandleFunc("/productos/eliminar/{id}", Handlers.EliminarProducto)
	//categoria
	r.HandleFunc("/categoria", Handlers.Categoria)
	r.HandleFunc("/categoria/nueva", Handlers.CrearCategoria)
	r.HandleFunc("/categoria/editar/{id}", Handlers.EditarCategoria)
	r.HandleFunc("/categoria/eliminar/{id}", Handlers.EliminarCategoria)
	//persona
	r.HandleFunc("/persona", Handlers.Persona)
	r.HandleFunc("/listarpersona", Handlers.ListarPersonas)
	r.HandleFunc("/filtrarpersona/{nombre}", Handlers.FiltrarPersona)
	r.HandleFunc("/obtenerpersona/{id}", Handlers.ObtenerPersonaPorId)
	r.HandleFunc("/guardarpersona", Handlers.GuardarPersona)
	r.HandleFunc("/eliminarpersona/{id}", Handlers.EliminarPersona)
	r.HandleFunc("/listarpersonasinusuario", Handlers.ListarPersonaSinUsuario)
	//tipopersona
	r.HandleFunc("/tipopersona", Handlers.TipoPersona)
	r.HandleFunc("/listartipopersona", Handlers.ListarTipoPersonas)

	//pais
	r.HandleFunc("/pais", Handlers.Pais)
	r.HandleFunc("/listarpais", Handlers.ListarPais)
	r.HandleFunc("/filtrarpais/{nombre}", Handlers.FiltrarPais)
	r.HandleFunc("/obtenerpais/{id}", Handlers.ObtenerPaisPorId)
	r.HandleFunc("/guardarpais", Handlers.GuardarPais)
	r.HandleFunc("/eliminarpais/{id}", Handlers.EliminarPais)
	//proveedor
	r.HandleFunc("/proveedor", Handlers.Proveedor)
	r.HandleFunc("/listarproveedor", Handlers.ListarProveedor)
	r.HandleFunc("/filtrarproveedor/{idpais}", Handlers.FiltrarProveedor)
	r.HandleFunc("/obtenerproveedor/{id}", Handlers.ObtenerProveedorPorId)
	r.HandleFunc("/guardarproveedor", Handlers.GuardarProveedor)
	r.HandleFunc("/eliminarproveedor/{id}", Handlers.EliminarProveedor)
	//usuario
	r.HandleFunc("/usuario", Handlers.Usuario)
	r.HandleFunc("/obtenerusuario/{id}", Handlers.ObtenerUsuarioPorId)
	r.HandleFunc("/guardarusuario", Handlers.GuardarUsuario)
	r.HandleFunc("/listarusuario", Handlers.ListarUsuarios)
	r.HandleFunc("/eliminarusuario/{id}", Handlers.EliminarUsuario)
	//tiporol
	r.HandleFunc("/listartiporol", Handlers.ListarTipoRol)
	r.HandleFunc("/eliminartiporol/{id}", Handlers.EliminarTipoRol)
	//pagina
	r.HandleFunc("/pagina", Handlers.Pagina)
	r.HandleFunc("/listarpagina", Handlers.ListarPagina)
	r.HandleFunc("/obtenerpagina/{id}", Handlers.ObtenerPaginaPorId)
	r.HandleFunc("/guardarpagina", Handlers.GuardarPagina)
	r.HandleFunc("/eliminarpagina/{id}", Handlers.EliminarPagina)
	//Paginatiporol
	r.HandleFunc("/paginatiporol", Handlers.PaginaTipoRol)
	r.HandleFunc("/guardartiporol", Handlers.GuardarTipoRol)
	r.HandleFunc("/nuevapaginatiporol", Handlers.NuevaPaginaTipoRol)
	r.HandleFunc("/editarpaginatiporol/{id}", Handlers.EditarPaginaTipoRol)
	r.HandleFunc("/obtenertiporol/{id}", Handlers.ObtenerTipoRolPorId)
	r.HandleFunc("/obtenerpaginatiporol/{id}", Handlers.FiltrarPaginaTipoRolPorId)

	// Login
	r.HandleFunc("/", Handlers.Login)
	r.HandleFunc("/login/{usu}/{contra}", Handlers.IngresarUsuario)
	r.HandleFunc("/cerrarsesion", Handlers.CerrarSesion)
	r.HandleFunc("/llenarmenu", Handlers.LlenarMenu)
	r.HandleFunc("/index", Handlers.Pinicio)
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	var clients = make(map[*websocket.Conn]bool)
	r.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		ws, _ := upgrader.Upgrade(w, r, nil)
		clients[ws] = true
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				delete(clients, ws)
			}
			for client := range clients {
				err := client.WriteMessage(msgType, msg)
				if err != nil {
					return
				}
			}
		}
	})
	http.ListenAndServe(":8000", r)
}
