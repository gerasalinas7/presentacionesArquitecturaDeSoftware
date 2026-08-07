const iframe = document.querySelector(".embed-frame iframe");
const fallback = document.querySelector("[data-embed-fallback]");

if (iframe && fallback) {
  const timeout = window.setTimeout(() => {
    fallback.hidden = false;
  }, 2200);

  iframe.addEventListener("load", () => {
    window.clearTimeout(timeout);
  });
}
