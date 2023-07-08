<script lang="ts">
  import { enhance } from "$app/forms";
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<svelte:head>
  <title>Todos Demo</title>
  <meta name="description" content="A sample using todos..." />
</svelte:head>

<div class="w-full h-screen p-16 text-neutral-950 bg-slate-50">
  <h1 class="text-3xl font-bold font-sans">Basic Todo App Example</h1>
  <h4 class="font-light text-lg italic">by Ben Davis</h4>

  {#each data.todos as todo (todo.id)}
    <div class="flex mt-4 bg-white rounded-lg w-[500px] p-4 shadow-lg flex-col">
      <h3
        class="flex flex-row items-center gap-x-2 {todo.completed &&
          'line-through'}"
      >
        <span
          class="w-[12px] h-[12px] rounded-full inline-block {todo.completed
            ? 'bg-green-500'
            : 'bg-red-500'}"
        />
        <span class="font-bold text-lg">{todo.title}</span>
      </h3>
      <p
        class="font-light text-sm text-neutral-600 {todo.completed &&
          'line-through'}"
      >
        {todo.description}
      </p>
      <div class="flex flex-row justify-end gap-x-4 mt-4">
        <form
          method="POST"
          action="?/complete"
          use:enhance={() => {
            return async ({ result }) => {
              if (result.status === 200) {
                // update the todo to be completed
                const todoIdx = data.todos.findIndex((t) => t.id === todo.id);
                data.todos[todoIdx].completed = true;
              }
            };
          }}
        >
          <input class="hidden" name="id" value={todo.id} />
          <button
            class="border-2 p-1 border-blue-600 rounded-lg text-blue-600 disabled:border-gray-600 disabled:text-gray-600"
            disabled={todo.completed}
            type="submit">Complete</button
          >
        </form>
        <form
          method="POST"
          action="?/delete"
          use:enhance={() => {
            return async ({ result }) => {
              if (result.status === 200) {
                data.todos = data.todos.filter((t) => t.id !== todo.id);
              }
            };
          }}
        >
          <input class="hidden" name="id" value={todo.id} />
          <button
            class="border-2 p-1 border-red-600 rounded-lg text-red-600"
            type="submit"
          >
            Delete</button
          >
        </form>
      </div>
    </div>
  {/each}
  <form method="POST" action="?/create">
    <input
      type="text"
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      placeholder="Title"
      name="title"
    />
    <input
      type="text"
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      placeholder="Description"
      name="description"
    />
    <button
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      type="submit"
    >
      Add Todo
    </button>
  </form>
  <a href="/sign-out" class="" data-sveltekit-preload-data="off">Sign Out</a>
</div>
