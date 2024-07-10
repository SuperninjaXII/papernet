let sendBtn = document.getElementById('add');

function addBook() {
  console.log("clicked");

  // Get input values
  let Title = document.querySelector("#Title").value;
  let Description = document.querySelector("#Description").value;
  let Link = document.querySelector("#Link").value;
  let ImageUrl = document.querySelector("#imgUrl").value;

  // Create payload
  let payload = [{
    "title": Title,
    "description": Description,
    "link": Link,
    "image": `${ImageUrl}`
  }];

  // Send POST request to server
  fetch('/upload', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json' // Adjust based on your API's requirements
    },
    body: JSON.stringify(payload)
  })
  .then(response => response.json())
  .then(data => {
    // Handle successful response
    c.alert(JSON.stringify(data));
    console.log(data);
  })
  .catch(error => {
    // Handle errors
    c.alert("An error occurred. Please try again later.");
    console.error(error);
  });
}

// Add click event listener to the button
sendBtn.addEventListener("click", addBook);
