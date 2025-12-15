// Show notification function
function showNotification(message, duration = 3000) {
  const notification = document.createElement("div");
  notification.id = "notification";
  notification.style.cssText = `
    display: none;
    position: fixed;
    bottom: 20px;
    right: 20px;
    background: #333;
    color: #fff;
    padding: 15px;
    border-radius: 5px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
    font-size: 14px;
    z-index: 1000;
  `;
  document.body.appendChild(notification);

  notification.textContent = message;
  notification.style.display = "block";
  setTimeout(() => (notification.style.display = "none"), duration);
}

// Wait for the page to load
window.addEventListener("DOMContentLoaded", async () => {
  // Get form elements
  const loginForm = document.getElementById("login-form");
  const emailInput = document.getElementById("email");
  const passwordInput = document.getElementById("password");

  // Create notification element
  // Redirect to login and clear storage
  function redirectToLogin() {
    localStorage.clear();
    window.location.href = "/login";
  }

  // Login form submission
  loginForm.addEventListener("submit", async (event) => {
    event.preventDefault();

    const email = emailInput.value.trim();
    const password = passwordInput.value.trim();

    if (!email || !password) {
      showNotification("Email and Password are required!");
      return;
    }

    try {
      const response = await fetch("/authentication/login", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: new URLSearchParams({ email, password }),
      });

      if (!response.ok) {
        throw new Error("Failed to login");
      }

      const data = await response.json();
      const { token, userID } = data;

      localStorage.setItem("token", token);
      localStorage.setItem("userID", userID);

      showNotification("Login successful!");

      try {
        const adminResponse = await fetch("/admin", {
          method: "GET",
          headers: { Authorization: `Bearer ${token}`, userID: userID },
        });

        if (!adminResponse.ok) {
          throw new Error(
            (await adminResponse.text()) || "Authentication failed",
          );
        }

        document.body.innerHTML = await adminResponse.text();
      } catch (error) {
        console.error("Admin access error:", error);
        showNotification("Authentication failed. Please log in again.");
        redirectToLogin();
      }
    } catch (error) {
      console.error("Error:", error);
      showNotification("An error occurred. Please try again.");
    }
  });

  // Auto-login logic
  const token = localStorage.getItem("token");
  const userID = localStorage.getItem("userID");

  if (token && userID) {
    showNotification("Attempting auto-login...");
    try {
      const response = await fetch("/admin", {
        method: "GET",
        headers: { Authorization: `Bearer ${token}`, "User-ID": userID },
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || "Authentication failed");
      }

      // Attempt to parse response as JSON
      try {
        const jsonData = await response.json();
        // If parsing succeeds, handle JSON data (e.g., display a message)
        console.log("Admin response (JSON):", jsonData);
        // Do NOT replace the body content if it's JSON. Update elements as needed.
        // Example:
        // if (jsonData.message) {
        //    showNotification(jsonData.message);
        // }
      } catch (jsonError) {
        // If parsing fails, assume it's HTML and replace the body
        if (response.headers.get("content-type").includes("text/html")) {
          document.body.innerHTML = await response.text();
        } else {
          console.error("Response in html", response);
          showNotification("Unexpected response from the server.");
        }
      }
    } catch (error) {
      console.error("Auto-login error:", error);
      showNotification("Session expired. Please log in again.");
      redirectToLogin();
    }
  }
});

// Similar logic can be applied to the addBook functio
// n
async function addBook(addBtn) {
  const token = localStorage.getItem("token");
  const userID = localStorage.getItem("userID");

  if (!token || !userID) {
    showNotification("You are not authenticated. Please log in again.");
    redirectToLogin();
    return;
  }

  const title = document.getElementById("title").value.trim();
  const description = document.getElementById("description").value.trim();
  const author = document.getElementById("author").value.trim();
  const imageFile = document.getElementById("image").files[0];
  const downloadFile = document.getElementById("file").files[0];
  const category1 = document.getElementById("cartegory1").value.trim();
  const category2 = document.getElementById("cartegory2").value.trim();

  if (!title || !description || !imageFile || !downloadFile || !category1) {
    showNotification("Please fill in all required fields."); // Use showNotification
    return;
  }

  addBtn.disabled = true; // Disable the button while processing
  addBtn.style.background = "grey";
  addBtn.innerHTML = "adding book";

  const formData = new FormData();
  formData.append("title", title);
  formData.append("description", description);
  formData.append("author", author);
  formData.append("image", imageFile);
  formData.append("file", downloadFile);
  formData.append("cartegory1", category1);
  formData.append("cartegory2", category2);

  try {
    const response = await fetch("/admin/addbook", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
        userID: userID,
      },
      body: formData,
    });
    showNotification("uploading your book please wait");
    if (!response.ok) {
      const errorText = await response.text(); // Get error message from server
      throw new Error(errorText || "Failed to add the book");
    }

    showNotification("Book added successfully!");

    // Clear the form (more efficiently)
    title.value = "";
    description.value = "";
    author.value = "";
    imageFile.value = "";
    downloadFile.value = "";
  } catch (error) {
    console.error("Error adding book:", error);
    if (
      error.message.toLowerCase().includes("authentication") ||
      error.message.toLowerCase().includes("unauthorized")
    ) {
      showNotification("Session expired. Please log in again.");
      redirectToLogin();
    }
  } finally {
    addBtn.disabled = false; // Re-enable the button
    addBtn.style.background = "var(--accent)";
    addBtn.innerHTML = "add book";
  }
}
