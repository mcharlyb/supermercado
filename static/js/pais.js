var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
    if (valcook==""){
      window.location.replace("/")
    }
    LlenarMenu("pais",true)
    Listar(()=>{})
}
socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="guardarpais" || data=="eliminarpais"){
    Listar()
   }   
}   

function Listar(callback){
     Pintar("/listarpais",["Id Pais","Nombre","Capital"],["Idpais","Nombre","Capital"],"divTabla",
     true,true,"Idpais",undefined,true,false,true,false,undefined,undefined,function(){
        callback();
    });
}
function FiltrarPais(){
    var nombre=document.getElementById("txtnombrebuscar").value;
    if(nombre==""){
        Listar();
    }else{
        Pintar("/filtrarpais/"+nombre,["Id Pais","Nombre","Capital"],["Idpais","Nombre","Capital"],"divTabla",true,true,"Idpais");
    } 
}
function Recuperar(id){
   
    fetchGet("/obtenerpais/"+id,function(res){sets(["txtIdpais","txtnombrepais","txtcapitalpais"],
    [res.Idpais,res.Nombre,res.Capital])})
    
}
   
function GuardarDatos(){
    var bedit=document.getElementById("bledit").value;
    var objetopais={
    Idpais:get("txtIdpais")==""?0:get("txtIdpais")*1,
    Nombre:get("txtnombrepais"),
    Capital:get("txtcapitalpais")
   }
   if(bedit==0 & objetopais.Idpais!=0){
      var titulo="Error"
      var mensaje="Su perfil no le permite Editar"
      errorModal(titulo,mensaje)
      return
    }else{
     fetchPost("/guardarpais",objetopais,true,function(){
     document.getElementById("btncerrar").click()  
     socket.send("guardarpais")
    })
   }
}
function Eliminar(id){
    var bedelet=document.getElementById("bldelet").value;
    if(bedelet=="1"){
      fetchDelete("/eliminarpais/"+id,true,function(){
        alerta1()
        socket.send("eliminarpais")
      })
    }else{
      var titulo="Error"
      var mensaje="Su perfil no le permite Eliminar"
      errorModal(titulo,mensaje)
    }     
}