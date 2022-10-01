const form = document.getElementById("location-data");
const distance = document.getElementById("result-box")
const placesInput = document.getElementsByClassName("pac-input");

// initializes the map for displaying route
function initMap() {
	// initially the map is centered on Delhi, India
	const map = new google.maps.Map(
		document.getElementById("map"),
		{
			zoom: 5,
			center: {lat: 28.70, lng: 77.10},
		}
	);

	// for rendering the route
	const directionsService = new google.maps.DirectionsService();
	const directionsRenderer = new google.maps.DirectionsRenderer();

	directionsRenderer.setMap(map);
	const onChangeHandler = function () {
		calculateAndDisplayRoute(directionsService, directionsRenderer);
	};

	// autocomplete options
	const options = {
		fields: ["formatted_address", "geometry", "name"],
		strictBounds: false,
		types: ["establishment"],
	};

	// validate entered object
	for (let i = 0; i < placesInput.length; i++) {
		const autocomplete = new google.maps.places.Autocomplete(placesInput[i], options);

		autocomplete.addListener("place_changed", () => {

			const place = autocomplete.getPlace();
			if (!place.geometry || !place.geometry.location) {
				// User entered the name of a Place that was not suggested and
				// pressed the Enter key, or the Place Details request failed.
				window.alert("No details available for input: '" + place.name + "'");
				return;
			}
		});
	}

	// display route on submit
	form.addEventListener("submit", onChangeHandler);
}

// caluclates and displays route
function calculateAndDisplayRoute(directionsService, directionsRenderer) {
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

// when the form is submitted, fetch the distance from backend
form.addEventListener("submit", getDistance);
function getDistance(event) {
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

	// get distance in json form and updates the page
	fetch("/getDistance", fetchOptions)
		.then(response => response.json())
		.then(data => {
			if (data["distance"] === "undefined" || data["duration"] == "undefined") {
				window.alert("No route between '" + formJSON["origin"] + "' and '" + formJSON["destination"] + "'");
				return;
			}
			distance.innerHTML = `
				<p><b>Distance:</b> ${data["distance"]}</p>\n
				<p><b>Duration:</b> ${data["duration"]}</p>\n
			`;
		})
}

window.initMap = initMap;
