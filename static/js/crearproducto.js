var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
  var idcategoria=document.getElementById("cbocategoria").getAttribute("data-idcategoria")
  if(idcategoria!="0"){
    document.getElementById("cbocategoria").value=idcategoria
  } 

}


function mostrarAlertaa(){
  obj=ValidarObligatorios()
  
  if(obj.error==true){
    document.getElementById("divErrores").innerHTML=obj.contenido;
  }else{
    confirmaciona().then((result) => {
      if (result.isConfirmed) {
        alertaa()
        document.getElementById("frmEnviar").submit();
        socket.send("agregarproducto")
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