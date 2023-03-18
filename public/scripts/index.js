import {Encrypt} from "./libs/encryption/encrypt.js"

function writeToDOM() {
    var text = document.getElementById("inputText").value;
    Encrypt(text).then((res) =>{
        document.getElementById("output").innerHTML = res
    })
    
}

document.getElementById("inputText").addEventListener("keyup", writeToDOM);