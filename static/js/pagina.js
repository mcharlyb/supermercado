var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
    if (valcook==""){
      window.location.replace("/")
    }
    LlenarMenu("pagina",false)
    Listar(()=>{})
}
function Listar(callback){
     
   Pintar("/listarpagina",["id Pagina","Mensaje","Ruta"],["Idpagina","Mensaje","Ruta"],"divTabla",
   true,true,"Idpagina",undefined,false,false,true,false,undefined,undefined,function(){
    callback();
})
}
socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="guardarpagina" || data=="eliminarpagina"){
    var indiceActual=indiceActualFuncion("tabla")
    Listar(()=>{
        RecuperarPaginadoActual("tabla",indiceActual)

    })
   }
}
function Recuperar(id){
    fetchGet("/obtenerpagina/"+id,function(res){sets(["txtIdpagina","txtmensaje","txtruta"],
    [res.Idpagina,res.Mensaje,res.Ruta])})
}
function guardarDatosPagina(){
    var bedit=document.getElementById("bledit").value;
  
    obj= ValidarObligatorios()
    if(obj.error==true){
      setI("divErrores",obj.contenido)
     return
    }
    var objetopagina={
        Idpagina:get("txtIdpagina")==""?0:get("txtIdpagina")*1,
        Mensaje:get("txtmensaje"),
        Ruta:get("txtruta")
    }
    if(bedit==0 & objetopagina.Idpagina!=0){
        var titulo="Error"
        var mensaje="Su perfil no le permite Editar"
        errorModal(titulo,mensaje)
        return
    }else{ 
       fetchPost("/guardarpagina",objetopagina,true,function(){
        Listar();
        Limpiar();
        socket.send("guardarpagina")
       })
    }
}
function Eliminar(id){
    var bedelet=document.getElementById("bldelet").value;
  if(bedelet=="1"){
    fetchDelete("/eliminarpagina/"+id,true,function(){
        alerta1()
        socket.send("eliminarpagina")
    })
    }else{
        var titulo="Error"
        var mensaje="Su perfil no le permite Eliminar"
        errorModal(titulo,mensaje)
    }     
}