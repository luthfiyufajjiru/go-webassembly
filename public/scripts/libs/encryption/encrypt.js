import wasm from "../vendor/wasm/load.js";

async function Encrypt(input) {
    return await wasm().then((res) => {
        return res.encrypt(input)
    })
}

export {Encrypt}
