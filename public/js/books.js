const limitDescriptions = () => {
  const descriptions = document.querySelectorAll(".description");

  descriptions.forEach((description) => {
    const fullText = description.textContent.trim();

    if (fullText.length > 100) {
      const truncatedText = fullText.substring(0, 100) + "...";
      description.textContent = truncatedText;
    }
  });
};

// Example call to limitDescriptions when needed
