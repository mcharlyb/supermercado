function Ingresar(){
    var usuario=get("txtusuario")
    var contra=get("txtcontra")
    if (usuario==""){
        errorModal("Debe Ingresar el Usuario")
        return
    }
    if (contra==""){
        errorModal("Debe Ingresar la Contraseña")
        return
    }
    fetchGet("/login/"+usuario+"/"+contra,function(data){
        if(data=="1"){
            window.location.href="/index"
        }else{
            errorModal("Usuario o ontraseña incorrecta") 
        }
    })
}
