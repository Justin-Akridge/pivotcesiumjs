const fileModal = document.querySelector("#myModal");
const newJobTab = document.querySelector("#new-job-tab");
const closeButton = document.querySelector(".close-modal");
closeButton.addEventListener("click", function() {
  fileModal.style.display = "none";
})
// Add event listener to show the modal when the tab is clicked
newJobTab.addEventListener("click", function() {
  fileModal.style.display = "flex";
});


const jobForm = document.querySelector(".job-form");
const jobName = document.querySelector(".job-name");
const company = document.querySelector(".company-name");

jobForm.addEventListener("submit", async function(e) {
  e.preventDefault(); 
  fileModal.style.display = "none";
  const token = await JSON.parse(localStorage.getItem('authToken'));

  const formData = new FormData();
  formData.append('jobName', jobName.value);
  formData.append('clientName', company.value);


  try {
    const response = await fetch("/jobs", {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: formData,
    })

    if (!response.ok) {
      throw new Error('Network response was not ok ' + response.statusText);
    }

  } catch (error) {
    console.error('There was a problem with your fetch operation:', error);
  }
  jobName.value = ""
  company.value = ""
})
