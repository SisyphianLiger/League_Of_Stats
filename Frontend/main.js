// document.addEventListener('DOMContentLoaded', () => {
// 	const button = document.querySelector('.HealthCheck');
// 	button.addEventListener('click', async () => {
// 		// Here we will send the on click 
// 		console.log('Button Clicked');
// 		await sendHealthCheck();
// 	});
// });


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

export async function createCards(numberOfCards) {
	const cardHolder = document.querySelector('.CardHolder');
	cardHolder.innerHTML = '';

	for (let i = 0; i < numberOfCards; i++) {
		const card = document.createElement('div');
		const divider = document.createElement('div');
		const imageOf = document.createElement('div');
		const cardHeader = document.createElement('div');
		const cardStats = document.createElement('div');

		card.className = 'Card';
		divider.className = 'Divider';
		imageOf.className = 'ImageOf';
		cardHeader.className = 'CardHeader';
		cardStats.className = 'CardStats';
		cardHeader.innerHTML = '<p> Lee "Faker" Sang-hyeok</p>'; // Add this line
		const playerStats = {
			winPercentage: 100,
			mostPlayed: 'Sylas',
			mostLostAgainst: 'Chovy',
			worldTitles: 5
		};

		cardStats.innerHTML = createStatsTable(playerStats);	

		card.appendChild(cardStats);
		card.appendChild(divider);
		card.appendChild(imageOf);
		card.appendChild(cardHeader);

		cardHolder.appendChild(card);

		await new Promise(resolve => setTimeout(resolve, 300));
	}
}

function createStatsTable(stats) {
    return `
        <table class="StatsTable">
            <tr>
                <td>Win Ratio:</td>
                <td>${stats.winPercentage}%</td>
            </tr>
            <tr>
                <td>Most Played Champ:</td>
                <td>${stats.mostPlayed}</td>
            </tr>
            <tr>
                <td>Most Losses vs Player:</td>
                <td>${stats.mostLostAgainst}</td>
            </tr>
            <tr>
                <td>World Titles:</td>
                <td>${stats.worldTitles}</td>
            </tr>
        </table>
    `;
}


window.createCards = createCards;
