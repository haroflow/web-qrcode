function generate() {
	error.innerText = "";
	error.style.display = "none";

	fetch("/generate", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded"
		},
		body: "size=" + sizeInput.value + "&data=" + dataInput.value,
	})
		.then(res => {
			if (res.status != 200)
				throw new Error("Failed to generate QR code");

			return res.blob();
		})
		.then(res => {
			url = URL.createObjectURL(res);
			result.src = url;
		})
		.catch(err => {
			error.innerText = err.message;
			error.style.display = "block";
		})
}
