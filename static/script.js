document.addEventListener("htmx:afterRequest", function () {
  const items = document.querySelectorAll(".headshot_txt div");
  items.forEach((div, i) => {
    div.style = "opacity: 1; transform: translateX(0);";
  });
});
document.addEventListener("DOMContentLoaded", function () {
  const items = document.querySelectorAll(".headshot_txt div");
  items.forEach((div, i) => {
    setTimeout(() => {
      div.style.animation =
        "slideInFromUnder 0.5s cubic-bezier(.5,1.5,.5,1) forwards";
    }, i * 100);
  });
});
