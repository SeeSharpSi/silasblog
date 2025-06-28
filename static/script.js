document.addEventListener("htmx:afterRequest", function () {
    const items = document.querySelectorAll(".headshot_txt li");
    items.forEach((li, i) => {
        li.style = "opacity: 1; transform: translateX(0);";
    });
});
document.addEventListener("DOMContentLoaded", function () {
    const items = document.querySelectorAll(".headshot_txt li");
    items.forEach((li, i) => {
        setTimeout(() => {
            li.style.animation = "slideInFromUnder 0.5s cubic-bezier(.5,1.5,.5,1) forwards";
        }, i * 100); 
    });
});
