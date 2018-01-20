function init() {
	bootpage.show("home-page");

	var xhttp = new XMLHttpRequest();

	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			console.log(this.responseText);
		}
	};

	xhttp.open("GET", "library/v1/books", true);
	xhttp.send();
}
