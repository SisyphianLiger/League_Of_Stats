document.addEventListener('DOMContentLoaded', () => {
	const button = document.querySelector('.HealthCheck');
	button.addEventListener('click', async () => {
		// Here we will send the on click 
		console.log('Button Clicked');
		await sendHealthCheck();
	});
});


async function sendHealthCheck() {

	try {
		const response = await fetch('http://127.0.0.1:8081/api/healthz');
		if (!response.ok) 
			throw new Error(`HTTP error! status ${response.status}`);

		const data = await response.text();
		console.log('Success:', data);
	}
	catch (error) {
		console.error('Error:', error);
	}
};
