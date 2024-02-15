function createTabs(containerSelector) {
    const container = document.querySelector(containerSelector);
    const tabs = container.querySelectorAll(".tabs .items");
    const contents = container.querySelectorAll(".contents .content-items");
  
    function showContent(index) {
      contents.forEach((content, i) => {
        if (i === index) {
          content.style.display = "block";
          content.style.Color = "#fff" 

          tabs[i].classList.add("active");
        } else {
          content.style.display = "none";
          tabs[i].classList.remove("active");
        }
      });
    }
  
    showContent(0);
  
    tabs.forEach((tab, i) => {
      tab.addEventListener("click", () => {
        showContent(i);
      });
    });
  }
  
  createTabs("#containers");