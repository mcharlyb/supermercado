var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    document.getElementById("cbocategoria").value=
    document.getElementById("cbocategoria").getAttribute("data-idcategoria")
}
function mostrarAlertae(){
    confirmacione().then((result) => {
        if (result.isConfirmed) {
          alertae()
          document.getElementById("frmEnviar").submit();
          socket.send("editarproducto")
        }
    })
} 
function alertae(titulo="Exito",mensaje="Se guardo correctamente"){
    Swal.fire(
        titulo,
        mensaje,
        'success'
      )
 }
 function confirmacione(titulo="Desea Guardar los Cambios?",
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