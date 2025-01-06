import { wasmModule } from './wasm';

async function main() {
    try {
        await wasmModule.initialize();

        const fibForm = document.getElementById('fib-form') as HTMLFormElement;
        const calcForm = document.getElementById('calc-form') as HTMLFormElement;

        fibForm.addEventListener('submit', (e) => {
            e.preventDefault();
            const input = document.getElementById('fib-input') as HTMLInputElement;
            const result = document.getElementById('fib-result') as HTMLDivElement;
            const n = parseInt(input.value);
            result.textContent = `Fibonacci(${n}) = ${wasmModule.fibonacci(n)}`;
            alert(`Fibonacci(${n}) = ${wasmModule.fibonacci(n)}`);
        });

        calcForm.addEventListener('submit', (e) => {
            e.preventDefault();
            const xInput = document.getElementById('x-input') as HTMLInputElement;
            const yInput = document.getElementById('y-input') as HTMLInputElement;
            const result = document.getElementById('calc-result') as HTMLDivElement;
            const x = parseFloat(xInput.value);
            const y = parseFloat(yInput.value);
            result.textContent = `Complex Calculation Result: ${wasmModule.complexCalculation(x, y)}`;
            alert(`Complex Calculation Result: ${wasmModule.complexCalculation(x, y)}`);
        });
    } catch (error) {
        console.error('Failed to initialize application:', error);
    }
}

main();