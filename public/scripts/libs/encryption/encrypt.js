import wasm from "../vendor/wasm/load.js";

function Encrypt(input) {
    return wasm.encrypt(input);
}

export {Encrypt}
