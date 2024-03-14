<script>
    import Login from './routes/login/+login.svelte'
    import { Router, Route } from 'svelte-routing';
    import Dashboard from './routes/dashboard/+dashboard.svelte';
</script>

<Router >
    <Route path="/"><Login /></Route>
    <Route path="dashboard"><Dashboard/></Route>
</Router>