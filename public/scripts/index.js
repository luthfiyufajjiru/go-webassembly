import {Decrypt, Encrypt} from "./libs/encryption/encrypt.js"

let inChargeEncrypt = false

let updateTriggerEncrypt = async (prom) => {
    inChargeEncrypt = true
    prom.then(() => {
    var text = document.getElementById("inputText").value
    if (text == "") {
        document.getElementById("output").innerHTML = "Encrypted:"
        inChargeEncrypt = false
        return
    }
    let res = Encrypt(text)
    res = `Encrypted: ${res}`
    document.getElementById("output").innerHTML = res
    inChargeEncrypt = false
    })
}

function writeToDOM(event) {
    let encryptMode = event.target.id == "inputText"
    let decryptMode = event.target.id == "inputDecrypt"
    switch (event.type) {
        case "keyup":
            if (!inChargeEncrypt && encryptMode){
                let prom = new Promise((resolve) => setTimeout(resolve, 1500))
                updateTriggerEncrypt(prom)
                break
            }
            if (decryptMode && event.keyCode == 13){
                let text = document.getElementById("inputDecrypt").value
                if (text == "") {
                    document.getElementById("decrypted").innerHTML = "Decrypted:"
                    break
                }
                let res = Decrypt(text)
                res = `Decrypted: ${res}`
                document.getElementById("decrypted").innerHTML = res
            }
        break
        case "submit":
            event.preventDefault()
        break
    }
}

document.getElementById("inputDecrypt").addEventListener("keyup", writeToDOM)
document.getElementById("inputText").addEventListener("keyup", writeToDOM)
document.getElementById("inputForm").addEventListener("submit", writeToDOM)