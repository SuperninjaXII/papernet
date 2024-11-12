const title = document.querySelector("#title");
const description = document.querySelector("#description");
const imageLink = document.querySelector("#img-link");
const download = document.querySelector("#download");
const category1 = document.querySelector("#cartegory1");
const category2 = document.querySelector("#cartegory2");
const addbtn = document.querySelector(".add-btn");
const addBook = () => {
  const data = JSON.stringify([
    {
      title: title.value,
      description: description.value,
      imageLink: imageLink.value,
      download: download.value,
      category1: category1.value,
      category2: category2.value,
    },
  ]);

  fetch("/admin/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: data,
  })
    .then((response) => response.json())
    .then((data) => {
      console.log("Success:", data);
      title.innerHTML = "";
      description.innerHTML = "";
      imageLink.innerHTML = "";
      download.innerHTML = "";
      category1.innerHTML = "";
      category2.innerHTML = "";
      alert("book added");
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("error");
    });
};
addbtn.addEventListener("click", addBook);
