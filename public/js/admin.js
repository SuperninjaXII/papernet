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
      image: imageLink.value,
      link: download.value,
      cartegory1: category1.value,
      cartegory2: category2.value,
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
      alert("book added");
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("error");
    });
};
addbtn.addEventListener("click", addBook);
