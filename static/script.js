let form = document.getElementById("location_data");
let distance = document.getElementById("distance_box")

form.addEventListener('submit', handleSubmit);
function handleSubmit(event) {
	event.preventDefault();
	const data = new FormData(event.target);
	const formJSON = Object.fromEntries(data.entries());

	var fetchOptions = {
		method: "POST",
		header: new Headers({
			"Content-Type": "application/json",
		}),
		//cross origin mode is needed as we are not using the same domain
		mode: "cors",
		body: JSON.stringify(formJSON)
	}

	fetch("/getDistance", fetchOptions)
		.then(response => response.json())
		.then(data => {
			distance.innerHTML = `
				<p><b>Distance:</b> ${data["distance"]}</p>\n
				<p><b>Duration:</b> ${data["duration"]}</p>\n
			`;
		})
}

function initMap() {
	const directionsService = new google.maps.DirectionsService();
	const directionsRenderer = new google.maps.DirectionsRenderer();
	const map = new google.maps.Map(
		document.getElementById("map"),
		{
			zoom: 7,
			center: {lat: 41.85, lng: -87.65},
		}
	);

	directionsRenderer.setMap(map);
	const onChangeHandler = function () {
		calculateAndDisplayRoute(directionsService, directionsRenderer);
	};

	form.addEventListener("submit", onChangeHandler);
	form.addEventListener("submit", onChangeHandler);
}

function calculateAndDisplayRoute(
	directionsService,
	directionsRenderer,
) {
	directionsService
		.route({
			origin: {
				query: document.getElementById("origin").value,
			},
			destination: {
				query: document.getElementById("destination").value,
			},
			travelMode: google.maps.TravelMode.DRIVING,
		})
		.then((response) => {
			directionsRenderer.setDirections(response);
		})
		// if we catch an exception, reset map
		.catch((e) => {
			console.log(e);
			initMap();
		});
}

window.initMap = initMap;
