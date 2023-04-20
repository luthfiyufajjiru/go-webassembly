import {Encrypt} from "./libs/encryption/encrypt.js"

let inCharge = false

let updateTrigger = async (prom) => {
    inCharge = true
    prom.then(() => {
    var text = document.getElementById("inputText").value
    let res = Encrypt(text)
    document.getElementById("output").innerHTML = res
    inCharge = false
    })
}

function writeToDOM(event) {
    var text = document.getElementById("inputText").value
    let res = Encrypt(text)
    switch (event.type) {
        case "keyup":
            if (!inCharge){
                let prom = new Promise((resolve) => setTimeout(resolve, 1500))
                updateTrigger(prom)
            }
        break
        case "submit":
            event.preventDefault()
            document.getElementById("output").innerHTML = res
        break
    }
}

document.getElementById("inputText").addEventListener("keyup", writeToDOM)
document.getElementById("inputForm").addEventListener("submit", writeToDOM)