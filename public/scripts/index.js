import {Encrypt} from "./libs/encryption/encrypt.js"

function writeToDOM() {
    var text = document.getElementById("inputText").value;
    let res = Encrypt(text)
    document.getElementById("output").innerHTML = res
    
}

document.getElementById("inputText").addEventListener("keyup", writeToDOM);