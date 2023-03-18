const go = new Go();

const wasm = async () => {
    const data = await WebAssembly.instantiateStreaming(fetch('./scripts/libs/vendor/wasm/main.wasm'), go.importObject).then((result) => {
        let mod = result.module;
        let inst = result.instance;
        go.run(inst);
        return {mod, encrypt}
    })
    return data
}

export default wasm