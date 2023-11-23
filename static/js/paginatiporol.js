var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
    if (valcook==""){
      window.location.replace("/")
    }
    
    LlenarMenu("paginatiporol",true)
   
    Listar(()=>{})
}
function Listar(callback){
   Pintar("/listartiporol",["id Tipo Rol","Nombre","Descripcion"],["Idtiporol","Nombre","Descripcion"],
  "divTabla",true,true,"Idtiporol",undefined,false,false,true,false,undefined,undefined,function(){
        callback()
  });
}
function Recuperar(id){
    var bedit=document.getElementById("bledit").value;
    if(bedit=="1"){
        window.location.href="/editarpaginatiporol/"+id
    }else{
      var titulo="Error"
      var mensaje="Su perfil no le permite Editar"
      errorModal(titulo,mensaje)
    }
}

socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="editartiporol" || data=="agregartiporol" || data=="eliminartiporol"){
    var indiceActual=indiceActualFuncion("tabla")
    Listar(()=>{
        RecuperarPaginadoActual("tabla",indiceActual)

    })

   }   
} 
function Eliminar(id){
  var bedelet=document.getElementById("bldelet").value;
  if(bedelet=="1"){
    fetchDelete("/eliminartiporol/"+id,true,function(){
        alerta1()
        socket.send("eliminartiporol")
    })
  }else{
    var titulo="Error"
    var mensaje="Su perfil no le permite Eliminar"
    errorModal(titulo,mensaje)
  }
}

