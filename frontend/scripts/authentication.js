//// Function to fetch protected content
//async function fetchProtectedContent() {
//  const token = await JSON.parse(localStorage.getItem('authToken'));
//
//  console.log(token)
//  if (token) {
//    try {
//      const response = await fetch('/', {
//        method: 'GET',
//        headers: {
//          'Authorization': `Bearer ${token}`
//        }
//      });
//
//      if (response.ok) {
//        const data = await response.json();
//      } else {
//      }
//    } catch (error) {
//      console.error('An error occurred:', error);
//    }
//  } else {
//    window.location.href = '/login';
//  }
//}
//
//// Fetch protected content on page load
//window.onload = fetchProtectedContent;
