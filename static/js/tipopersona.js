window.onload=function(){
    LlenarMenu("tipopersona")
    Listar()
}
function Listar(){
     Pintar("/listartipopersona",["Id Tipo Persona","Nombre"],["Idtipopersona","Nombre"],"divTabla",true,true,"Idtipopersona");
}