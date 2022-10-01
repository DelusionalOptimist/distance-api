const form = document.getElementById("location-data");
const distance = document.getElementById("result-box")
const placesInput = document.getElementsByClassName("pac-input");

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

	// for displaying route
	directionsRenderer.setMap(map);
	const onChangeHandler = function () {
		calculateAndDisplayRoute(directionsService, directionsRenderer);
	};

	const options = {
		// TODO:look into fields
		fields: ["formatted_address", "geometry", "name"],
		strictBounds: false,
		types: ["establishment"],
	};

	let placesAutocomplete = {};
	for (let i = 0; i < placesInput.length; i++) {
		placesAutocomplete[i] = new google.maps.places.Autocomplete(placesInput[i], options);

		placesAutocomplete[i].addListener("place_changed", () => {
			//infowindow.close();
			//marker.setVisible(false);

			const place = placesAutocomplete[i].getPlace();

			if (!place.geometry || !place.geometry.location) {
				// User entered the name of a Place that was not suggested and
				// pressed the Enter key, or the Place Details request failed.
				window.alert("No details available for input: '" + place.name + "'");
				return;
			}

			// If the place has a geometry, then present it on a map.
			//if (place.geometry.viewport) {
			//	map.fitBounds(place.geometry.viewport);
			//} else {
			//	map.setCenter(place.geometry.location);
			//	map.setZoom(17);
			//}

			//marker.setPosition(place.geometry.location);
			//marker.setVisible(true);

			//infowindowContent.children["place-name"].textContent = place.name;
			//infowindowContent.children["place-address"].textContent =
			//	place.formatted_address;
			//infowindow.open(map, marker);
		});

	}

	form.addEventListener("submit", onChangeHandler);

	//const infowindow = new google.maps.InfoWindow();
	//const infowindowContent = document.getElementById(
	//	"infowindow-content"
	//);
	//infowindow.setContent(infowindowContent);

	//const marker = new google.maps.Marker({
	//	map,
	//	anchorPoint: new google.maps.Point(0, -29),
	//});

	//const origin = document.getElementById("pac-origin");
	//const destination = document.getElementById("pac-destination");
	//const originAutocomplete = new google.maps.places.Autocomplete(origin, options);

	//const destinationAutocomplete = new google.maps.places.Autocomplete(destination, options);


	//var originPlace,destinationPlace;

	//originAutocomplete.addListener("place_changed", () => {
	//	infowindow.close();
	//	marker.setVisible(false);

	//	originPlace = originAutocomplete.getPlace();

	//	if (!originPlace.geometry || !originPlace.geometry.location) {
	//		// User entered the name of a Place that was not suggested and
	//		// pressed the Enter key, or the Place Details request failed.
	//		window.alert("No details available for input: '" + place.name + "'");
	//		return;
	//	}

	//	// If the place has a geometry, then present it on a map.
	//	if (place.geometry.viewport) {
	//		map.fitBounds(place.geometry.viewport);
	//	} else {
	//		map.setCenter(place.geometry.location);
	//		map.setZoom(17);
	//	}

	//	marker.setPosition(place.geometry.location);
	//	marker.setVisible(true);

	//	infowindowContent.children["place-name"].textContent = place.name;
	//	infowindowContent.children["place-address"].textContent =
	//		place.formatted_address;
	//	infowindow.open(map, marker);
	//});

	//destinationAutocomplete.addListener("place_changed", () => {
	//	infowindow.close();
	//	marker.setVisible(false);

	//	const place = destinationAutocomplete.getPlace();

	//	if (!place.geometry || !place.geometry.location) {
	//		// User entered the name of a Place that was not suggested and
	//		// pressed the Enter key, or the Place Details request failed.
	//		window.alert("No details available for input: '" + place.name + "'");
	//		return;
	//	}

	//	// If the place has a geometry, then present it on a map.
	//	if (place.geometry.viewport) {
	//		map.fitBounds(place.geometry.viewport);
	//	} else {
	//		map.setCenter(place.geometry.location);
	//		map.setZoom(17);
	//	}

	//	marker.setPosition(place.geometry.location);
	//	marker.setVisible(true);

	//	infowindowContent.children["place-name"].textContent = place.name;
	//	infowindowContent.children["place-address"].textContent =
	//		place.formatted_address;
	//	infowindow.open(map, marker);
	//});

	//form.addEventListener("change", onChangeHandler);
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
