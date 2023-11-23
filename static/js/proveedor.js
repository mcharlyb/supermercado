var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
    if (valcook==""){
      window.location.replace("/")
    }
    LlenarMenu("proveedor",true)
    Listar(()=>{})
    llenarComboPais()
}
socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="guardarproveedor" || data=="eliminarproveedor" ){
    var indiceActual=indiceActualFuncion("tabla")
    Listar(()=>{
        RecuperarPaginadoActual("tabla",indiceActual)
    })
   }else if(data=="guardarpais"){
         llenarComboPais()
   }
}

function llenarComboPais(){
    PintarCombo("/listarpais","Nombre","Idpais","cbopais")
    PintarCombo("/listarpais","Nombre","Idpais","cboPaisBuscar")
}

function Listar(callback){
    Pintar("/listarproveedor",["Id Proveedor","Nombre","Direccion","Telefono","Nombre Pais"],
    ["Idproveedor","Nombre","Direccion","Telefono","Nombrepais"],"divTabla",true,true,"Idproveedor",
    undefined,true,false,true,false,undefined,undefined,function(){
        callback()
    });
    llenarComboPais()
}


function FiltrarProveedor(){
    var idpais=document.getElementById("cboPaisBuscar").value
    if(idpais==""){
        Listar()
    }else{
         Pintar("/filtrarproveedor/"+idpais,["Id Proveedor","Nombre","Direccion","Telefono","Nombre Pais"],
    ["Idproveedor","Nombre","Direccion","Telefono","Nombrepais"],"divTabla",true,true,"Idproveedor",
    undefined,true,false,true,false,undefined,undefined,function(){
        callback()
     });
    }
   
}
function Recuperar(id){
    fetchGet("/obtenerproveedor/"+id,function(res){
        sets(["txtIdproveedor","txtnombreproveedor","txtdireccionproveedor","txttelefonoproveedor","txtemailproveedor",
        "txtrepresentanteproveedor","txtcelularproveedor","txtrucproveedor","cbopais"],[res.Idproveedor,res.Nombre,
        res.Direccion,res.Telefono,res.Email,res.Representantelegal,res.Celular,res.Ruc,res.Idpais]) 
    })
}
function GuardarDatos(){
   // var frm=document.getElementById("frm")
   var bedit=document.getElementById("bledit").value;
   
    var objetoproveedor={
    Idproveedor:get("txtIdproveedor")==""?0:get("txtIdproveedor") * 1,
    Nombre:get("txtnombreproveedor"),
    Direccion:get("txtdireccionproveedor"),
    Telefono:get("txttelefonoproveedor"),
    Email:get("txtemailproveedor"),
    Representantelegal:get("txtrepresentanteproveedor"),
	Celular:get("txtcelularproveedor"),
	Idpais:get("cbopais") * 1,
	Ruc:get("txtrucproveedor")
   }
   if(bedit==0 & objetoproveedor.Idproveedor!=0){
      var titulo="Error"
      var mensaje="Su perfil no le permite Editar"
      errorModal(titulo,mensaje)
      return
   }else{

   fetchPost("/guardarproveedor",objetoproveedor,true,function(){
        //alert("ejecuta el callback")
        Limpiar();
        document.getElementById("btncerrar").click()
        socket.send("guardarproveedor")
   })
   
    }
 }
 function Eliminar(id){
    var bedelet=document.getElementById("bldelet").value;
    if(bedelet=="1"){
    fetchDelete("/eliminarproveedor/"+id,true,function(){
        alerta1()
        socket.send("eliminarproveedor")
    })
    }else{
      var titulo="Error"
      var mensaje="Su perfil no le permite Eliminar"
      errorModal(titulo,mensaje)
    }
    
}