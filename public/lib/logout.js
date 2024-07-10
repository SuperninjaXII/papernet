window.c = {
  print:function(msg) {
    console.info(`%c${msg}`, "color:blue;font-size:20px;font-style:italic;");
  },
  log: function(msg) {
    console.log(msg);
  },
  panic: function(msg) {
    console.error(`your error😂: ${msg}`);
  },
  warn: function(msg) {
    console.warn(`Warning: ${msg}`);
  },
  clear: function() {
    console.clear();
  },
  getTime: function() {
    return new Date().toLocaleTimeString();
  },
  measure: function(callback) {
    const start = performance.now();
    callback();
    const end = performance.now();
    const duration = end - start;
    console.log(`Operation took ${duration.toFixed(2)} milliseconds.`);
  },
  alert: function(message) {
    const modal = document.createElement("div");
    const modalContent = document.createElement("div");
    const text=document.createElement("p")
    let btn=document.createElement("button")
    modal.style.position = "fixed";
    modal.style.top = "50%";
    modal.style.left = "50%";
    modal.style.top="0";
    modal.style.left="0";
    modal.style.width="100%";
   modal.style.height="100%";
  modal.style.backgroundColor="rgba(0, 0, 0, 0.5)";
    modal.style.justifyContent="center";
   modal.style.alignItems="center";
   modalContent.style.backgroundColor="#fff";
    modalContent.style.padding="20px";
   modalContent.style.borderRadius="8px";
    modalContent.style.boxShadow="0 4px 8px rgba(0, 0, 0, 0.2)";
    modalContent.style.textAlign="center";
    modalContent.style.maxWidth="400px";
    modalContent.style.width="100%";
    text.innerHTML = `<p>${message}`;
    btn.style.color=" #0078D4"
  btn.style.backgroundcolor=" transparent"
  btn.style.padding=" 8px 12px"
  btn.style.border=" 1px solid #0078D4"
    btn.style.borderRadius=" 4px"
   btn.style.cursor=" pointer"
   btn.style.fontSize=" 14px"
  btn.style.marginTop=" 16px"
  btn.innerHTML =`close`
  btn.addEventListener("click",function(){
    modalContent.parentElement.remove()
  })
    document.body.appendChild(modal);
    modal.appendChild(modalContent)
    modalContent.appendChild(text)
    modalContent.appendChild(btn)
  }
};
