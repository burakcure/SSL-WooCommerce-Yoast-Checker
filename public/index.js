
function altButton(event){
    event.preventDefault();
    document.getElementById('file').click();
}


function altButtonDrag(event){
    event.preventDefault();
    document.getElementById('file').ondragend();

}
function ssl(uri){
    document.getElementById('drop').action=uri;

}
