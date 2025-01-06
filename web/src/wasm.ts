interface WasmExports {
    fibonacci: (n: number) => number;
    complex_calculation: (x: number, y: number) => number;
}

class WasmModule {
    private static instance: WasmModule;
    private wasmExports: WasmExports | null = null;

    private constructor() { }

    static getInstance(): WasmModule {
        if (!WasmModule.instance) {
            WasmModule.instance = new WasmModule();
        }
        return WasmModule.instance;
    }

    async initialize(): Promise<void> {
        try {
            const response = await fetch('/wasm/calc.wasm');
            const buffer = await response.arrayBuffer();
            const obj = await WebAssembly.instantiate(buffer);
            this.wasmExports = obj.instance.exports as unknown as WasmExports;
        } catch (error) {
            console.error('Failed to load WASM module:', error);
            throw error;
        }
    }

    fibonacci(n: number): number {
        if (!this.wasmExports) throw new Error('WASM module not initialized');
        return this.wasmExports.fibonacci(n);
    }

    complexCalculation(x: number, y: number): number {
        if (!this.wasmExports) throw new Error('WASM module not initialized');
        return this.wasmExports.complex_calculation(x, y);
    }
}

export const wasmModule = WasmModule.getInstance();