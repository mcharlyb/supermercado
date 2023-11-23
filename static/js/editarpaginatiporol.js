var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
  LlenarMenu("paginatiporol",false)  
  listar()
}
function listar(){
    var idtiporol=get("txtidtiporol")
    fetchGet("/obtenertiporol/"+idtiporol,function(data){
        set("txtidtiporolform",data.Idtiporol)
        set("txtnombretiporol",data.Nombre)
        set("txtdescripciontiporol",data.Descripcion)
    })
    Pintar("/listarpagina",["id Pagina","Mensaje","Ruta"],["Idpagina","Mensaje","Ruta"],"divTabla",false,false,"Idpagina",
    undefined,false,true,true,true,["Nueva","Editar","Eliminar"],["Crear","Modificar","Eliminar"],function(){
        fetchGet("/obtenerpaginatiporol/"+idtiporol,function(data){
            for(var i=0;i<data.length;i++){
                document.getElementById("chk"+data[i].Idpagina).checked=true
                if(data[i].Crear){
                   document.getElementById("chkCrear"+data[i].Idpagina).checked=true
                }
                if(data[i].Modificar){
                  document.getElementById("chkModificar"+data[i].Idpagina).checked=true
                }
                if(data[i].Eliminar){
                  document.getElementById("chkEliminar"+data[i].Idpagina).checked=true
                }
                
            }
        })
    });
}
function mostrarAlerta(){
    
    var id=get("txtidtiporolform")==""?0:get("txtidtiporolform")*1;
    var nombre=get("txtnombretiporol");
    var lnombre=nombre.length;  
    var descripcion=get("txtdescripciontiporol");
    var ldesc=descripcion.length; 
    var checkSeleccionados=ObtenerCheckSeleccionado()
    //alerta1(checkSeleccionados)
    var contenido="<ul>";
    var exito=true;
    if(lnombre==0 ){
      contenido+="<ol class='alert alert-danger mt-2'>No ha ingresado el Nombre (Campo Obligatorio)</ol>"
      exito=false;
    }
    if (lnombre > 150) {
      contenido+="<ol class='alert alert-danger mt-2'>El campo Nombre no puede tener mas de 150 caracteres</ol>"
      exito=false;
    }
    if(ldesc==0){
      contenido+="<ol class='alert alert-danger mt-2'>No ha ingresado la Descripcion (Campo Obligatorio)</ol>"
      exito=false;
    }
    if(ldesc>800){
      contenido+="<ol class='alert alert-danger mt-2'>El campo Descripción no puede tener mas de 800 caracteres</ol>"
      exito=false;
    }
    contenido+="</ul>";
    document.getElementById("divErrores").innerHTML=contenido;
    if(exito){
        confirmacion1().then((result) => {
           if (result.isConfirmed) {
             alerta1()
            var objetotiporol={
              Idtiporol:id,
              Nombre:nombre,
              Descripcion:descripcion,
              Idpaginascadena:checkSeleccionados
            }
            fetchPost("/guardartiporol",objetotiporol,true,function(){
              socket.send("editartiporol") 
              document.location.href="/paginatiporol"
             })
           }
        })
    } 
}