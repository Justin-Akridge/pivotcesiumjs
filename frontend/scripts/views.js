const perspectiveViewCheckbox = document.querySelector("#perspective-view-checkbox");
const mapViewCheckbox = document.querySelector("#map-view-checkbox");
const profileViewCheckbox = document.querySelector("#profile-view-checkbox");
const imageryViewCheckbox = document.querySelector("#imagery-view-checkbox");

const cesiumContainer = document.querySelector("#cesium-container");
const mapsContainer = document.querySelector("#maps-container");
const map = document.querySelector("#map");
const cesium = document.querySelector("#cesium");
const profile = document.querySelector("#profile-view");
const assets = document.querySelector("#assets-container");

function updateViewVisibility(event) {
  switch (event.target.name) {
    case "perspective-view":
      if (perspectiveViewCheckbox.checked) {
        cesium.style.display = "flex"
        if (mapViewCheckbox.checked) {
          mapsContainer.style.display = "flex";
        } else {
          mapsContainer.style.display = "none";
        }
      } else {
        cesium.style.display = "none"
      }
    case "map-view":
      if (mapViewCheckbox.checked) {
        mapsContainer.style.display = "flex";
        map.style.display = "flex"
        map.style.width = "100%";
      } else {
        map.style.display = "none"
      }
    case "imagery-view":
      if (imageryViewCheckbox.checked) {
        assets.style.display = "flex"
      } else {
        assets.style.display = "none"
      }
    case "profile-view":
      if (profileViewCheckbox.checked) {
        profile.style.display = "flex"
        if (!perspectiveViewCheckbox.checked) {
          mapsContainer.style.display = "none";
        }
        profile.style.height = "100%";
      } else {
        profile.style.display = "none"
      }

  }
}

function resetCheckboxes() {
  perspectiveViewCheckbox.checked = true; 
  mapViewCheckbox.checked = false;
  profileViewCheckbox.checked = false;
  imageryViewCheckbox.checked = false;
}
// Attach event listeners to checkboxes
perspectiveViewCheckbox.addEventListener('change', updateViewVisibility);
mapViewCheckbox.addEventListener('change', updateViewVisibility);
profileViewCheckbox.addEventListener('change', updateViewVisibility);
imageryViewCheckbox.addEventListener('change', updateViewVisibility);

resetCheckboxes()
