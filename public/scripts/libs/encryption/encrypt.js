import wasm from "../vendor/wasm/load.js";

function Encrypt(input) {
    return wasm.encrypt(input);
}

function Decrypt(input) {
    return wasm.decrypt(input);
}

export {Encrypt, Decrypt}
