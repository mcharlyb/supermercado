var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
    if (valcook==""){
      window.location.replace("/")
    }
     LlenarMenu("persona",true)
    $.datepicker.setDefaults($.datepicker.regional["es"]);
    $("#txtfechanacimiento").datepicker();
   // {dateFormat:"dd/mm/yy"}
    Listar(()=>{})
    PintarCombo("/listartipopersona","Nombre","Idtipopersona","cbotipopersona")
}
socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="guardarpersona" || data=="eliminarpersona"){
    var indiceActual=indiceActualFuncion("tabla")
    Listar(()=>{
        RecuperarPaginadoActual("tabla",indiceActual)

    })
   }   
}   

function Listar(callback){
     Pintar("/listarpersona",["Id Persona","Nombre Completo","Nombre Tipo","Fecha Nacimiento"],
     ["Idpersona","Nombrecompleto","Nombretipopersona","Fechanacimientocadena"],"divTabla",true,true,"Idpersona",
     undefined,true,false,true,false,undefined,undefined,function(){
        callback();
     })
}
function FiltrarPersona(){
    var nombre=document.getElementById("txtnombrebuscar").value;
    if(nombre==""){
        Listar();
    }else{
        Pintar("/filtrarpersona/"+nombre,["Id Persona","Nombre Completo","Nombre Tipo","Fecha Nacimiento"],
        ["Idpersona","Nombrecompleto","Nombretipopersona","Fechanacimientocadena"],"divTabla",true,true,"Idpersona",
         undefined,true,false,true,false,undefined,undefined,function(){
            callback();
         });
    } 
}
function Recuperar(id){
    fetchGet("/obtenerpersona/"+id,function(res){
        sets(["txtIdpersona","txtnombre","txtapellidopaterno","txtapellidomaterno","txtfechanacimiento",
        "cbotipopersona"],[res.Idpersona,res.Nombre,
        res.Appaterno,res.Apmaterno,res.Fechanacimientocadena,res.Idtipopersona]) 
    })
}
function GuardarDatos(){
    var bedit=document.getElementById("bledit").value;
    var objetopersona={
    Idpersona:get("txtIdpersona")==""?0:get("txtIdpersona") * 1,
    Nombre:get("txtnombre"),
    Appaterno:get("txtapellidopaterno"),
    Apmaterno:get("txtapellidomaterno"),
    Fechanacimiento:$("#txtfechanacimiento").datepicker("getDate"),
    Idtipopersona:get("cbotipopersona") * 1
	}
    if(bedit==0 & objetopersona.Idpersona!=0){
        var titulo="Error"
        var mensaje="Su perfil no le permite Editar"
        errorModal(titulo,mensaje)
        return
    }else{
      fetchPost("/guardarpersona",objetopersona,true,function(){
      alerta1()
      document.getElementById("btncerrar").click()
      socket.send("guardarpersona")
      })
    }
}
 function Eliminar(id){
    var bedelet=document.getElementById("bldelet").value;
    if(bedelet=="1"){
    fetchDelete("/eliminarpersona/"+id,true,function(){
        alerta1()
        socket.send("eliminarpersona")
    })
    }else{
      var titulo="Error"
      var mensaje="Su perfil no le permite Eliminar"
      errorModal(titulo,mensaje)
    }     
}