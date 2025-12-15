let searchResult = document.querySelector("#suggestions");

let searchInput = document.querySelector(".search-input");
let searchBtn = document.querySelector(".search-btn");
let wasBtnClicked = false;

const interaction1 = () => {
  searchInput.classList.add("interaction1");
  searchBtn.classList.add("interaction1");
};

const interaction2 = () => {
  searchBtn.classList.add("interaction2");
};

const interaction2Rev = () => {
  searchBtn.classList.add("interaction2rev");
};

const checkBtnState = () => {
  if (wasBtnClicked == false) {
    interaction2();
    wasBtnClicked = true;
  } else {
    interaction1();
    wasBtnClicked = false;
  }
};

searchBtn.addEventListener("click", checkBtnState);

const searchInputEvent = () => {
  interaction1();
  interaction2Rev();
};

searchInput.addEventListener("click", searchInputEvent);
let suggestionTimeout;

const getSuggestion = () => {
  fetch(`/result?q=${searchInput.value}`, {
    method: "POST",
  })
    .then((response) => {
      const contentType = response.headers.get("content-type");
      if (contentType.includes("application/json")) {
        return response.json().then((data) => {
          console.log("JSON data:", JSON.stringify(data));
        });
      } else if (contentType.includes("text/html")) {
        return response.text().then((text) => {
          console.log("HTML data:", text);
          searchResult.style.display = "flex";
          searchResult.innerHTML = "";
          searchResult.innerHTML = `${text}`;
          suggestionsPopAnimation();
          resetSuggestionTimeout();
        });
      } else {
        throw new Error("Unsupported content type: " + contentType);
      }
    })
    .catch((error) => console.error("Error:", error));
};

const suggestionsPopAnimation = () => {
  searchResult.style.display = "flex";
  let li = document.querySelectorAll("#suggestions li");
  gsap.fromTo(
    "#suggestions",
    {
      width: "10svw",
      height: "10svh",
      x: 4,
      borderRadius: "50svw",
      background: "pink",
    },
    {
      width: "70svw",
      height: "auto",
      duration: 0.31,
      borderRadius: ".5em",
      background: "#121212",
      x: 5,
    },
  );

  li.forEach((element) => {
    element.style.display = "block";
  });
};

const resetSuggestionTimeout = () => {
  clearTimeout(suggestionTimeout);
  suggestionTimeout = setTimeout(() => {
    searchResult.style.display = "none";
  }, 3000); // hide suggestions after 5 seconds
};

searchInput.addEventListener("input", getSuggestion);

const NextPage = () => {
  const searchQuery = searchInput.value;
  try {
    window.location.href = `/searchbooks?q=${searchQuery}`;
  } catch (error) {
    alert("somethings wrong");
  }
};

const limit = () => {
  const descriptions = document.querySelectorAll(".description");

  descriptions.forEach((description) => {
    const wordLimit =
      parseInt(description.getAttribute("data-word-limit"), 10) || 100;
    const fullText = description.textContent.trim();

    if (fullText.length > wordLimit) {
      const truncatedText = fullText.substring(0, wordLimit) + "...";
      description.textContent = truncatedText;
      console.log(truncatedText);
    }
  });
};

searchBtn.addEventListener("click", NextPage);
