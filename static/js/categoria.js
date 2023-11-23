var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
 
  var valcook=leercookie("idusuario")
  if (valcook==""){
    window.location.replace("/")
  }
  LlenarMenu("categoria",true)

  Paginar("table")
}
socket.onopen=function(){
  document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
  document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
 var data=e.data 
 if(data=="editarcategoria" || data=="agregarcategoria"){
  setTimeout(()=>{
    document.location.reload() 
  
  },1000)
 }else if(data.includes("bajarcategoria")){
  var array=data.split("*")
  var id=array[1]
  var frm=document.getElementById("frm")
      frm.action="/categoria/eliminar/"+id
      frm.submit()
 }
} 


function mostrarModal(id){
  var bedelet=document.getElementById("bldelet").value;
  if(bedelet=="1"){
  document.getElementById("txtid").value=id;
  texconf="Desea eliminar el Id: "+id
  subtitulo="Tambien se Eliminaran los Productos Asociados"
  confirmacion(texconf,subtitulo).then((result) => {
     if (result.value) {
      socket.send("bajarcategoria*"+id)
      alerta()
      }
   })
  }else{
    var titulo="Error"
    var mensaje="Su perfil no le permite Eliminar"
    errorModal(titulo,mensaje)
  }
}
function mostrarEditar(id){
  var frm=document.getElementById("frm")
   var bedit=document.getElementById("bledit").value;
  if(bedit=="1"){
  frm.action="/categoria/editar/"+id
  frm.submit()
  }else{
    var titulo="Error"
    var mensaje="Su perfil no le permite Editar"
    errorModal(titulo,mensaje)
  }
}

function alerta(titulo="Exito",mensaje="Se Actualizó correctamente"){
    Swal.fire(
        titulo,
        mensaje,
        'success'
      )
}
function confirmacion(titulo="Desea Guardar los Cambios?",
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



