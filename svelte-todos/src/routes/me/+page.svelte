<script>
  export let data;
</script>

{#if data.props}
  <main>
    <h1>Welcome {data.props.user.FirstName}!</h1>
    <p>Here is your full payload:</p>
    <pre>{JSON.stringify(data.props, null, 2)}</pre>
    <a href="/" class="block">View Todos</a>
    <a href="/sign-out" data-sveltekit-preload-data="off">Log Out</a>
  </main>
{/if}
