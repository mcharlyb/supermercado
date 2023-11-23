var socket=new WebSocket("ws://localhost:8000/socket")
window.onload=function(){
    var valcook=leercookie("idusuario")
     if (valcook==""){
     window.location.replace("/")
  }
     
   LlenarMenu("productos",true)
   
   Paginar("tabla")
    var idbuscado=document.getElementById("idcategoriabuscado").value;
    document.getElementById("idcategoriabusqueda").value=idbuscado;
}
socket.onopen=function(){
    document.getElementById("lblestado").innerHTML="Conectado";
}
socket.onclose=function(){
    document.getElementById("lblestado").innerHTML="Desonectado";
}
socket.onmessage=function(e){
   var data=e.data 
   if(data=="editarproducto" || data=="agregarproducto" ){
    setTimeout(()=>{
        //document.getElementById("frmProductoForm").submit()
        document.location.reload()
    },1000)
   }else if(data.includes("bajarproducto")){
    var array=data.split("*")
    var id=array[1]
    var frm=document.getElementById("frm")
        frm.action="/productos/eliminar/"+id
        frm.submit()
   } 
}

  function mostrarAlertad(id){
    var bedelet=document.getElementById("bldelet").value;
    if(bedelet=="1"){
    document.getElementById("txtid").value=id;
    texconf="Desea eliminar el Id: "+id
    confirmacion(texconf).then((result) => {
       if (result.isConfirmed) {
        socket.send("bajarproducto*"+id)
        alertad()
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
    frm.action="/productos/editar/"+id
    frm.submit()
    }else{
      var titulo="Error"
      var mensaje="Su perfil no le permite Editar"
      errorModal(titulo,mensaje)
    }
  }


  function alertad(titulo="Exito",mensaje="Se Actualizó correctamente"){
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