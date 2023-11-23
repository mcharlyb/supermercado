function alerta1(titulo="Exito",mensaje="Se guardo correctamente"){
    Swal.fire(
        titulo,
        mensaje,
        'success'
      )
}
function confirmacion1(titulo="Desea Guardar los Cambios?",
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
function Paginar(idtabla){
    $('#'+idtabla).DataTable()
}
function Guardar(){
    obj= ValidarObligatorios()
    if(obj.error==true){
      setI("divErrores",obj.contenido)
     return
    }
    confirmacion1("Desea guardar los cambios","Presione Aceptar o Cancelar").then(res=>{
        //aceptar
        if(res.value){
             GuardarDatos()
        }
    })
   
}
function ValidarObligatorios(){
    var error=false;
    var contenido="<ul>";
    var obligatorios=document.getElementsByClassName("obligatorio")
    var ncontroles=obligatorios.length;
    for (var i=0;i<ncontroles;i++){
        if(obligatorios[i].value==""){
            error=true;
            if(obligatorios[i].name!=""){
                var campo=obligatorios[i].name
            }else{
                var campo=obligatorios[i].id.replace("txt","").replace("cbo","")
            }
            contenido+=`<ol class='alert alert-danger mt-2'>Debe ingresar el campo ${campo} </ol>`
        }
    }
    contenido+="</ul>";
    return {error,contenido}
}
function get(id){
    return document.getElementById(id).value;
}
function getI(id){
    return document.getElementById(id).innerHTML;
}
function set(idcontrol,valor){
    document.getElementById(idcontrol).value=valor;
}
function sets(arrayId, arrayValores){
    for(var i=0;i<arrayId.length;i++){
        set(arrayId[i],arrayValores[i])
    }
}
function setI(idcontrol,valor){
    document.getElementById(idcontrol).innerHTML=valor;
}
function fetchGet(url,callback){
    fetch(url).then(res=>res.json()).then(res=>{callback(res)})
}
function fetchDelete(url,Iscallback=false,Callback){
    fetch(url).then(res=>res.text())
    .then(res=>{
        if(res=="1"){
           if(Iscallback==false){
                alerta1()
                Listar()
            }else{
              Callback()  
            }
        }else{
            errorModal()
        }    
    })
}

function fetchPost(url,objeto,IsCallback=false,Callback){
    fetch(url,{headers:{"content-type":"application/json"},
    method:"POST",body:JSON.stringify(objeto)
    }).then(res=>res.text()).then(res=>{
        
        if(res=="1"){
             alerta1()
            
             if(IsCallback==false){
                document.getElementById("btncerrar").click()
                //alert("no entro en callback")
                Listar()
                 
             }else{
                //alert("entro en callback")
                Callback()
            }    
        }else{
            errorModal()
        }
    })   
}
function errorModal(title="Error",texto="Ocurrio un Error al guardar"){
    Swal.fire({
        icon:'error',
        title: title,
        text: texto

    })
}
function mostrarAlerta(id){
    confirmacion1("Desea Eliminar el Registro "+id,"De click en Aceptar o Cancelar").then(res=>{
        if(res.value){
            Eliminar(id)
        }
    })
}
function ObtenerCheckSeleccionado(){
   //var checkboxs=document.getElementsByClassName("checkbox")
    var checkboxs=document.getElementsByName("checkid")
    var ncheckbox=checkboxs.length;
    var cadena="";
    var id;
    var esp="";
    for(var i=0;i<ncheckbox;i++){
        if(checkboxs[i].checked==true){
            var blok=""
            id=checkboxs[i].id.replace("chk","")
            var chekper=document.getElementById("chkCrear"+id)
            if(chekper.checked==true){
                blok+="t"
            }else{
                blok+="f"
            }
            var chekper=document.getElementById("chkModificar"+id)
            if(chekper.checked==true){
                blok+="t"
            }else{
                blok+="f"
            }
            var chekper=document.getElementById("chkEliminar"+id)
            if(chekper.checked==true){
                blok+="t"
            }else{
                blok+="f"
            }
            blok+=id
            cadena+=esp+blok;
            esp="*"
        }
    }
    if (cadena.length>0){
        //alert(cadena)
        return cadena
    }
}


function Pintar(url,cabeceras,camposMostrar,divImprimir="divTabla",
mostrarEditar=false,mostrarEliminar=false,propiedadId="",divTabla="tabla",tienepopup=true,tienecheck=false,
IsCallback=false,tieneperm=false,cabeperm,camposperm,Callback,){
    var contenido=""
    contenido+="<table data-t='"+propiedadId+"' id='"+divTabla+"' class='table' >"
    fetch(url).then(
        res=>res.json()
    ).then(res=>{
       // alert(res)
       contenido+="<thead class='table-primary' >"
       contenido+="<tr>"
       if(tienecheck){
        contenido+="<th>"

        contenido+="</th>"
       }
       for(var i=0;i<cabeceras.length;i++){
        contenido+="<th>"+cabeceras[i]+"</th>"
       }
       if(tieneperm){
        for(var i=0;i<cabeperm.length;i++){
            contenido+="<th>"+cabeperm[i]+"</th>" 
        }
       }
       if(mostrarEditar==true || mostrarEliminar==true){
        contenido+="<th>Operaciones</th>"
       }
       contenido+="</tr>" 
       contenido+="</thead>"
       contenido+="<tbody>"
       for(var i=0;i<res.length;i++){
        contenido+="<tr>";
        //objeto
        fila=res[i]
        if(tienecheck){
            contenido+="<td>"
            contenido+= `<input type='checkbox' 
            class='checkbox' name='checkid' id='chk${fila[propiedadId]}'/>`
            contenido+="</td>"
           }
        
        for(var j=0;j<camposMostrar.length;j++){
            nombrePropiedad=camposMostrar[j]
            contenido+="<td>"
            contenido+=fila[nombrePropiedad]
            contenido+="</td>"
        }
        if(tieneperm){
            for(var j=0;j<camposperm.length;j++){
                contenido+="<td>"
                contenido+= `<input type='checkbox' 
                class='checkbox' id='chk${camposperm[j]}${fila[propiedadId]}'/>`
                contenido+="</td>" 
            }
        }
        if(mostrarEditar==true || mostrarEliminar==true){
            contenido+="<td>"
            if(mostrarEditar==true){
                contenido+=`<a class="btn btn-primary" 
                ${
                    tienepopup ?
                    `
                     data-bs-toggle="modal" data-bs-target="#staticBackdrop"
                     onclick='AbrirModal(${fila[propiedadId]},"${divTabla}");Recuperar(${fila[propiedadId]})'
                    `
                    :
                    `
                     onclick='Recuperar(${fila[propiedadId]})'
                    `
                }
               >
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pencil-square" viewBox="0 0 16 16">
                  <path d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z"/>
                  <path fill-rule="evenodd" d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z"/>
                </svg>
              </a>`
            }
            if(mostrarEliminar==true){
                contenido+=`
                <a class="btn btn-danger" onclick="mostrarAlerta(${fila[propiedadId]})">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
                  <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                  <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                </svg>
              </a>
              `
            }

            contenido+="</td>"
           }
       }
       contenido+="</tbody>"
        contenido+="</table>" 
        document.getElementById(divImprimir).innerHTML=contenido; 
        Paginar(divTabla);
        if(IsCallback==true){
            Callback()
        }
    }) 
    
}
function Limpiar(){
    var controles=document.getElementsByClassName("form-control")
    var ncontroles=controles.length;
    for(var i=0;i<ncontroles;i++){
        controles[i].value="";
    }
    var controles=document.getElementsByClassName("form-select")
    var ncontroles=controles.length;
    for(var i=0;i<ncontroles;i++){
        controles[i].value="";
    }
    var contenido=""
    setI("divErrores",contenido)
}
function PintarCombo(url,propiedadMostrar,propiedadValue,idcombo){
var contenido="";
fetch(url).then(res=>res.json())
    .then(res=>{
        contenido+="<option value=''>--Seleccione--</option>"
        var fila;
        for(var i=0;i<res.length;i++){
            fila=res[i]
            contenido+="<option value='"+fila[propiedadValue]+"'>"+fila[propiedadMostrar]+"</option>"
           }
    document.getElementById(idcombo).innerHTML=contenido;
    })
}
function AbrirModal(id,divTabla="tabla"){
    Limpiar()
    setI("divErrores","")
    var propiedad=document.getElementById(divTabla).getAttribute("data-t").replace("Id","").toLowerCase();
    if(id==0){
        document.getElementById("lbltitulo").innerHTML="Agregar "+propiedad
    }else{
        document.getElementById("lbltitulo").innerHTML="Editar "+propiedad
    }
}
function Pintarsubpopup(data,cabeceras,propiedades,idpro,nombrepro,IsCallback=false,Callback){
    var contenido="";
    contenido+="<table id='tablaSubPopup' class='table'>"
    
    contenido+="<thead>"
    contenido+="<tr>"
    for(var i=0;i<cabeceras.length;i++){
        contenido+="<th>"+cabeceras[i]+"</th>"
    }
    contenido+="<th>Operaciones</th>"
    contenido+="</tr>"; 
    contenido+="</thead>";
    contenido+="<tbody>";
    var nombrePropiedad;
    var objetoActual;
    for(var i=0;i<data.length;i++){
        contenido+="<tr>";
        //objeto
        objetoActual=data[i]
        for(var j=0;j<propiedades.length;j++){
            nombrePropiedad=propiedades[j]
            contenido+="<td>"
            contenido+=objetoActual[nombrePropiedad]
            contenido+="</td>"
        }
    contenido+="<td>";  
    contenido+=`
    <button onclick=
    "AsignarValores(${objetoActual[idpro]},'${objetoActual[nombrepro]}')"
        class="btn btn-success">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check-circle" viewBox="0 0 16 16">
        <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
        <path d="M10.97 4.97a.235.235 0 0 0-.02.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-1.071-1.05z"/>
        </svg>
        </button>
        `
    contenido+="</td>"    
    contenido+="</tr>"
           
    }
    contenido+="</tbody>"
    contenido+="</table>" 
    document.getElementById("divTablaSubPopup").innerHTML=contenido; 
    Paginar("tablaSubPopup");
    if(IsCallback==true){
        Callback();
    }
}
function LlenarMenu(activo,btnue){
           
    if(document.getElementById("ulmenu")){
        fetchGet("/llenarmenu",function(data){
            var contenido="";
            var redirec=true;

            
            for(var i=0;i<data.length;i++){
                if(activo==data[i].Ruta.replace("/","")){
                    if(data[i].Crear==false){
                        if(btnue){
                        const myButton = document.getElementById('btnuevo');
                        myButton.style.display = "none";
                        }
                    }
                    if(data[i].Modificar==false){
                        document.getElementById("bledit").value="0";
                    }
                    if(data[i].Eliminar==false){
                        document.getElementById("bldelet").value="0";
                    }
                    redirec=false;
                    contenido+=`<li class="nav-item"> 
                    <a class="nav-link active" href="${data[i].Ruta}">${data[i].Mensaje}</a>
                    </li>
                    `
                   // alert(opagina.Ruta)
                }else{
                  
                    
                    contenido+=`<li class="nav-item"> 
                    <a class="nav-link" href="${data[i].Ruta}">${data[i].Mensaje}</a>
                    </li>
                    `
                }
               
            }
            if(redirec==true & activo!="index"){
                window.location.replace("/cerrarsesion")
            }
        document.getElementById("ulmenu").innerHTML=contenido;
        })
     
    }
    
    
    
}
function RecuperarPaginadoActual(idtabla,indiceActual){
    
    while(1==1){
        var obj=document.querySelectorAll
        ("#"+idtabla+"_paginate :not(#"+idtabla+"_previous) a")
        var link
        var indiceBucle
        //No encontro la pagina seleccionada
        var encontro=false
        for(var i=0;i<obj.length;i++){
            link= obj[i]
            indiceBucle=document.querySelector
            ("#"+idtabla+"_paginate .current").innerHTML
            if(indiceBucle==indiceActual){
                encontro=true;
                break;
            }else{
                document.getElementById(idtabla+"_next").click()
                indiceSiguientePagina=document.querySelector
                ("#"+idtabla+"_paginate .current").innerHTML
                if(indiceBucle==indiceSiguientePagina){
                    encontro=true;
                    break;
                }
            }

        }
        if(encontro==true) break;
    }
}
function indiceActualFuncion(idtabla){
    var objPaginaActual=document.
    querySelector("#"+idtabla+"_paginate .current")
    return objPaginaActual.innerHTML
}
function leercookie(cname){
    var name= cname+"=";
    var ca = document.cookie.split(';');
    for(var i = 0;i <ca.length; i++){
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1);
        if(c.indexOf(name)== 0) return (c.substring(name.length, c.length));

    }
    return "";
}
