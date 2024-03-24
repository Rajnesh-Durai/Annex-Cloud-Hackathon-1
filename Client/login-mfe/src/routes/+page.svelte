<script>
    import { Router, Route } from 'svelte-routing';
    import Login from '../pages/login.svelte';
    import Dashboard from '../pages/dashboard.svelte';
</script>

<Router >
    <Route path="/"><Login/></Route>
    <Route path="/dashboard"><Dashboard/></Route>
</Router>