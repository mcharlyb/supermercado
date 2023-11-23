var socket=new WebSocket("ws://localhost:8000/socket")
function mostrarAlertaa(){
  var nombre=document.getElementById("txtnombre").value;
  var lnombre=nombre.length;  
  var descripcion=document.getElementById("txtdescripcion").value;
  var ldesc=descripcion.length; 
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
 
  confirmaciona().then((result) => {
       if (result.isConfirmed) {
         alertaa()
         document.getElementById("frmEnviar").submit();
         socket.send("agregarcategoria")
       }
     })
  } 
}
function alertaa(titulo="Exito",mensaje="Se guardo correctamente"){
   Swal.fire(
       titulo,
       mensaje,
       'success'
     )
}
function confirmaciona(titulo="Desea Guardar los Cambios?",
subtitulo="Por favor de Click en Aceptar o Cancelar"){
       return Swal.fire({
       title: titulo,
       text: subtitulo,
       icon: 'warning',
       showCancelButton: true,
       confirmButtonColor: '#3085d6',
       cancelButtonColor: '#d33',
       confirmButtonText: 'Aceptar !'
     })
}