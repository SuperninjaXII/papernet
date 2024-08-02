blob=()=>{
let tl = gsap.timeline();

tl.to(".blob", {
  rotation: "6deg",
  x: 2,
  y: 70,
  background: "linear-gradient(hsla(265, 15%, 31%, 1),hsla(200,20%,31%))", // Bright gradient colors
  borderBottomLeftRadius: "60%",
  borderTopLeftRadius: "20%",
  borderBottomRightRadius: "60%",
  borderTopRightRadius: "20%",
  duration: 7,
  filter: "blur(30px)",
  ease: "power2.out"
})
.to(".blob", {
  rotation: "70deg",
  x: 60,
  y: 50,
  scale: 1.1,
  borderBottomLeftRadius: "50%",
  borderTopLeftRadius: "10%",
  borderBottomRightRadius: "70%",
  borderTopRightRadius: "40%",
  duration: 7,
  ease: "power2.out"
})
.to(".blob", {
  rotateY: "30deg",
  x: 10,
  y: 190,
  scale: 1.7,
  background: "linear-gradient(hsla(205, 15%, 31%, 1),hsla(190,20%,31%))", // Bright gradient colors
  borderBottomLeftRadius: "60%",
  borderTopLeftRadius: "60%",
  borderBottomRightRadius: "4%",
  borderTopRightRadius: "10%",
  duration: 7,
  ease: "power2.out",
  opacity: 0.8,
  filter: "blur(22px)"
})
.to(".blob", {
  rotateZ: "12deg",
  x: -20,
  y: 150,
  scale: 1.5,
  background: "linear-gradient(hsla(235, 15%, 31%, 1),hsla(280,20%,31%))", // Bright gradient colors
  borderBottomLeftRadius: "20%",
  borderTopLeftRadius: "60%",
  borderBottomRightRadius: "20%",
  borderTopRightRadius: "70%",
  duration: 7,
  opacity: 0.6,
  ease: "power2.out",
  filter: "blur(19px)"
})
.to(".blob", {
  rotation: "1turn",
  x: 0,
  y: 50,
  scale: 1.3,
  background: "linear-gradient(hsla(305, 15%, 31%, 1),hsla(100,20%,31%))", // Bright gradient colors
  borderBottomLeftRadius: "59%",
  borderTopLeftRadius: "38%",
  borderBottomRightRadius: "50%",
  borderTopRightRadius: "7%",
  duration: 7,
  ease: "power2.out",
  opacity: 0.2,
  filter: "blur(10px)"
});
}
window.addEventListener("load",blob)
