<script>
  export let data;
</script>

<main>
  <h1>Welcome {data.firstName}!</h1>
  <p>Here is your full payload:</p>
  <pre>{JSON.stringify(data, null, 2)}</pre>
  <a href="/" class="block">View Todos</a>
  <a href="/sign-out">Log Out</a>
</main>
