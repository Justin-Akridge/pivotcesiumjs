// Select all tabs and dropdowns
let currentDropDown = null;
let currentTab = null;

const tabs = document.querySelectorAll("#view, #file, #edit, #settings, #help, #environments");

function resetTabColor() {
  if (currentTab) {
    currentTab.style.backgroundColor = ""
  }
}

function setTabColor() {
  currentTab.style.backgroundColor = "#333"
}

function toggleDropdown(event) {
  console.log(event.target)
  if (currentTab) {
    resetTabColor();
  }
  currentTab = document.querySelector(`#${event.target.id}`);
  setTabColor();

  const targetId = "#" + event.target.id + "-dropdown";
  const targetDropdown = document.querySelector(targetId);

  if (currentDropDown && currentDropDown !== targetDropdown) {
    currentDropDown.style.display = "none";
  }

  currentDropDown = targetDropdown;

  // Toggle the target dropdown
  if (targetDropdown.style.display === "none" || targetDropdown.style.display === "") {
    targetDropdown.style.display = "flex";
  } else {
    targetDropdown.style.display = "none";
    resetTabColor();
  }
}

// Attach event listeners to all tabs
tabs.forEach(tab => {
  tab.addEventListener("click", toggleDropdown);
});

// Event listener for the Esc key
document.addEventListener("keydown", function(event) {
  if (currentDropDown !== null && event.key === "Escape") {
    closeCurrentTab();
  }
});

document.addEventListener("click", function(event) {
  if (currentDropDown && !currentDropDown.contains(event.target) && !Array.from(tabs).includes(event.target)) {
    closeCurrentTab()
  }
});

function closeCurrentTab() {
  currentDropDown.style.display = "none";
  currentDropDown = null;
  resetTabColor();
  currentTab = null;
}
