<script>
    import { onMount } from 'svelte';

    let role ='';

    function parseJwt(token) {
    try {
      const base64Url = token.split('.')[1];
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
      const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
          return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
      }).join(''));
      return JSON.parse(jsonPayload);
    } catch (error) {
      console.error('Error decoding token:', error);
      return null;
    }
  }

  onMount(async () => {
    try {
      const token = sessionStorage.getItem('token');
      const decoded = parseJwt(token);
      console.log(decoded)
      if (decoded && decoded.role) {
        role = decoded.role;
      }
    } catch (error) {
      console.error('Error decoding token:', error);
    }
  });
  </script>

<h2>Hello {role}!</h2>