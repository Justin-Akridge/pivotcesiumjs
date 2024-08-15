//async function fetchProtectedContent() {
//  const token = await JSON.parse(localStorage.getItem('authToken'));
//  console.log(token)
//  if (token) {
//    try {
//      const response = await fetch('/checkAuth', {
//        method: 'GET',
//        headers: {
//          'Authorization': `Bearer ${token}`
//        }
//      });
//
//      if (response.ok) {
//        window.location.href = '/';
//      } else {
//        window.location.href = '/login';
//      }
//    } catch (error) {
//      console.error('An error occurred:', error);
//    }
//  } else {
//    window.location.href = '/login';
//  }
//}
//
//window.onload = fetchProtectedContent;
//
document.getElementById('login-form').addEventListener('submit', async function(event) {
  event.preventDefault();

  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  console.log(email)
  console.log(password)

  const formData = { email, password };

  try {
    const response = await fetch(`/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (response.ok) {
      const tokenAndClaims = await response.json();
      localStorage.setItem('authToken', JSON.stringify(tokenAndClaims.token));
      localStorage.setItem('claims', JSON.stringify(tokenAndClaims.claims));
      const homeResponse = await fetch('/', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${tokenAndClaims.token}`
        }
      });

      if (homeResponse.ok) {
          window.location.href = '/';
      } else {
          document.getElementById('error-message').textContent = "Error accessing home page.";
      }
    } else {
      document.getElementById('error-message').textContent = "Error logging in. Invalid email or password";
    }
  } catch (error) {
    console.error("An error occurred trying to log user in", error);
  }
});
