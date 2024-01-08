window.onload = () => {
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
		go.run(result.instance);
	});

	const addButton = document.getElementById('add-button');
	const resultSpan = document.getElementById('result');
	addButton.addEventListener('click', () => {
		const result = add(7,8);
		resultSpan.textContent = result;
	})
}