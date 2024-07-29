let searchInput = document.querySelector(".search-input")
let searchBtn = document.querySelector('.search-btn')
let searchResult =document.querySelector("#suggestions")

let wasBtnClicked = false

  
interaction1=() => {
    searchBtn.style.display='flex'
    let tl1 =gsap.timeline()
    tl1.to(searchInput,{
      width:"70svw"
    })
    tl1.to(searchBtn,{
      y:2,
      rotation:"3deg",
      x:5,
    })
    
}
interaction2=()=>{
  let tl2 = gsap.timeline()
  tl2.to(searchBtn,{
    scale:0.9
  })
  tl2.to(searchBtn,{
  borderTopRightRadius: "2rem",
  borderBottomRightRadius: "2rem",
  })

}
interaction2Rev=()=>{
  let tl2rev =gsap.timeline()
  tl2rev.to(searchBtn,{
    scale:1
  })
  tl2rev.to(searchBtn,{
    borderRadius:0
  })
}
checkBtnState=()=>{
  if (wasBtnClicked==false) {
    interaction2()
    wasBtnClicked=true
  } else if (wasBtnClicked=true) {
    interaction1()
    wasBtnClicked=false
  } 
}
searchInputEvent=()=>{
  interaction1()
  interaction2Rev()
  
}
  searchInput.addEventListener("click",searchInputEvent)
searchBtn.addEventListener("click",checkBtnState)
let suggestionTimeout;

const getSuggestion = () => {
  fetch(`/result?q=${searchInput.value}`, {
    method: 'POST',
  })
    .then(response => {
      const contentType = response.headers.get('content-type');
      if (contentType.includes('application/json')) {
        return response.json().then(data => {
          console.log('JSON data:', JSON.stringify(data));
        });
      } else if (contentType.includes('text/html')) {
        return response.text().then(text => {
          console.log('HTML data:', text);
          searchResult.style.display = "flex";
          searchResult.innerHTML = "";
          searchResult.innerHTML = `${text}`;
          suggestionsPopAnimation();
          resetSuggestionTimeout();
        });
      } else {
        throw new Error('Unsupported content type: ' + contentType);
      }
    })
    .catch(error => console.error('Error:', error));
};

const suggestionsPopAnimation = () => {
  let li = document.querySelectorAll('#suggestions li');
  gsap.fromTo("#suggestions", {
    width: "10svw",
    height: "10svh",
    x: 4,
    borderRadius: "50svw",
    background: "pink"
  }, {
    width: "70svw",
    height: "auto",
    duration: 0.31,
    borderRadius: ".5em",
    background: "#121212",
    x: 5
  });
  
  li.forEach(element => {
    element.style.display = 'block';
  });
};

const resetSuggestionTimeout = () => {
  clearTimeout(suggestionTimeout);
  suggestionTimeout = setTimeout(() => {
    searchResult.style.display = "none";
  }, 3000); // hide suggestions after 5 seconds
};

searchInput.addEventListener("input", getSuggestion);
