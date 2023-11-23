var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
  if (valcook==""){
    window.location.replace("/")
  } 
        
    LlenarMenu("usuario",true)
     Listar(()=>{})
     //PintarCombo("/listarpersonasinusuario","Nombrecompleto","Idpersona","cbopersona")
     PintarCombo("/listartiporol","Nombre","Idtiporol","cbotiporol")
     //PintarCombo("/listartiporol","Nombre","Idtiporol","cbotiporole")
}
socket.onopen=function(){
     document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
     document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
    var data=e.data 
    if(data=="guardarusuario" || data=="eliminarusuario"){
     var indiceActual=indiceActualFuncion("tabla")
     Listar(()=>{
         RecuperarPaginadoActual("tabla",indiceActual)
 
        })
    } else if(data=="agregartiporol" || data=="editartiporol"){
        PintarCombo("/listartiporol","Nombre","Idtiporol","cbotiporol")
    } else if(data=="guardarpersona" || data=="eliminarpersona"){
        var indiceActual=indiceActualFuncion("tablaSubPopup")
        fetchGet("/listarpersonasinusuario",function(rpta){
            Pintarsubpopup(rpta,["Id persona","Nombre Completo"],["Idpersona","Nombrecompleto"],
            "Idpersona","Nombrecompleto",true,function(){
                RecuperarPaginadoActual("tablaSubPopup",indiceActual)
            })
        });
    }   
}


function Listar(callback){
     Pintar("/listarusuario",["Id Usuario","Nombre Usuario","Nombre Completo","Tipo Rol"],
     ["Idusuario","Nombreusuario","Nombrepersona","Nombretiporol"],"divTabla",true,true,"Idusuario",
     undefined,true,false,true,false,undefined,undefined,function(){
          callback();
     })
}
function MostrarCampos(){
     document.getElementById("divContra").style.display="block"
     document.getElementById("divselpersona").style.display="block"
     //document.getElementById("cbopersona").setAttribute("class","form-select form-select-sm obligatorio")
     document.getElementById("divNomPersona").style.display="none"
}

function Recuperar(id){
    document.getElementById("divContra").style.display="none"
    document.getElementById("divselpersona").style.display="none"
    //document.getElementById("cbopersona").setAttribute("class","form-select form-select-sm")
    document.getElementById("divNomPersona").style.display="block"
    set("txtcontra","1234")    
    
     fetchGet("/obtenerusuario/"+id,function(res){sets(["txtIdusuario","txtnombreusuario","txtnombrepersonaro","txtidpersona","cbotiporol"],
     [res.Idusuario,res.Nombreusuario,res.Nombrepersona,res.Idpersona,res.Idtiporol])})
 }
 function AbrirSubPopup(nombre){
     if(nombre=="Persona"){
         document.getElementById("lbltituloSubPopup").innerHTML="Personas sin Usuario"
        fetchGet("/listarpersonasinusuario",function(rpta){
          Pintarsubpopup(rpta,["Id persona","Nombre Completo"],["Idpersona","Nombrecompleto"],"Idpersona","Nombrecompleto")
         })
    }

 }
 function AsignarValores(id,nombre){
    document.getElementById("btncerrarSubPopup").click()
    set("txtnombrepersona",nombre)
    set("txtidpersona",id)
 }
 function GuardarDatos(){
    var bedit=document.getElementById("bledit").value;
    var objetousuario={
    Idusuario:get("txtIdusuario")==""?0:get("txtIdusuario") * 1,
    Nombreusuario:get("txtnombreusuario"),
    Contra:get("txtcontra"),
    Idpersona:get("txtidpersona") * 1,
    Idtiporol:get("cbotiporol") * 1,
   }
   if(bedit==0 & objetousuario.Idusuario!=0){
    var titulo="Error"
    var mensaje="Su perfil no le permite Editar"
    errorModal(titulo,mensaje)
    return
   }else{

   fetchPost("/guardarusuario",objetousuario,true,function(){
     alerta1()
     document.getElementById("btncerrar").click();
     socket.send("guardarusuario")
    })
  }
}
function Eliminar(id){
    var bedelet=document.getElementById("bldelet").value;
    if(bedelet=="1"){ 
    fetchDelete("/eliminarusuario/"+id,true,function(){
        alerta1()
        socket.send("eliminarusuario")
     })
    }else{
        var titulo="Error"
        var mensaje="Su perfil no le permite Eliminar"
        errorModal(titulo,mensaje)
    } 
 }