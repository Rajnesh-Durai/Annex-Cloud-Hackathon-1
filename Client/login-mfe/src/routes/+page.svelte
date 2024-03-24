<script>
    import { Router, Route } from 'svelte-routing';
    import Login from '../pages/login.svelte';
</script>

<Router >
    <Route path="/"><Login/></Route>
</Router>