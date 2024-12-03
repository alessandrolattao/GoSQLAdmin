function showErrorToast(message) {
  // Create toast element
  const toast = document.createElement("div");
  toast.className =
    "alert alert-error shadow-lg fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-fit max-w-xl";

  toast.innerHTML = `
        <div class="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 mr-2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" />
            </svg>
            <span>${message}</span>
        </div>
    `;

  // Append toast to body
  document.body.appendChild(toast);

  // Remove toast after 3 seconds
  setTimeout(() => {
    toast.remove();
  }, 3000);
}

document.addEventListener("htmx:responseError", (event) => {
  console.error(
    "HTMX error:",
    event.detail.xhr.status,
    event.detail.xhr.responseText,
  );
  showErrorToast(event.detail.xhr.responseText);
});
