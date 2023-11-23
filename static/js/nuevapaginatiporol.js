var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
  LlenarMenu("paginatiporol")   
  Listar()
}
function Listar(){
  Pintar("/listarpagina",["id Pagina","Mensaje","Ruta"],["Idpagina","Mensaje","Ruta"],"divTabla",false,false,"Idpagina",
  undefined,false,true,true,true,["Nueva","Editar","Eliminar"],["Crear","Modificar","Eliminar"],function(){})  
 
}
function mostrarAlerta(){
    
    var id=get("txtidtiporol")==""?0:get("txtidtiporol")
    var nombre=get("txnombretiporol");
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
              socket.send("agregartiporol")
              document.location.href="paginatiporol"
            })
           }
        })
    } 
}