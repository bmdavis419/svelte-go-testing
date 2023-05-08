<script lang="ts">
  import { fly } from "svelte/transition";
  import type { PageData } from "./$types";

  export let data: PageData;

  // new todo form state
  let title = "";
  let description = "";
  let color = "blue";
</script>

<svelte:head>
  <title>Todos Demo</title>
  <meta name="description" content="A sample using todos..." />
</svelte:head>

<div class="w-full h-screen p-16 text-neutral-950 bg-slate-50">
  <h1 class="text-3xl font-bold font-sans">Basic Todo App Example</h1>
  <h4 class="font-light text-lg italic">by Ben Davis</h4>

  {#each data.todos as todo (todo.id)}
    <div
      class="flex mt-4 bg-white rounded-lg w-[500px] p-4 shadow-lg flex-col"
      in:fly={{ x: 50, duration: 5000 }}
      out:fly={{ x: 50, duration: 1000 }}
    >
      <h3
        class="flex flex-row items-center gap-x-2 {todo.completed &&
          'line-through'}"
      >
        <span
          class="w-[12px] h-[12px] rounded-full inline-block {todo.color ===
            'blue' && 'bg-blue-600'} {todo.color === 'red' &&
            'bg-red-600'} {todo.color === 'green' && 'bg-green-600'}"
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
        <button
          class="border-2 p-1 border-blue-600 rounded-lg text-blue-600"
          on:click={() => {
            todo.completed = !todo.completed;
          }}>Complete</button
        >
        <button
          class="border-2 p-1 border-red-600 rounded-lg text-red-600"
          on:click={() => {
            data.todos = data.todos.filter((t) => t.id !== todo.id);
          }}
        >
          Delete</button
        >
      </div>
    </div>
  {/each}
  <form
    on:submit={() => {
      data.todos = [
        ...data.todos,
        {
          id: Math.random(),
          title,
          completed: false,
          description,
          color,
        },
      ];
      title = "";
      description = "";
      color = "blue";
    }}
  >
    <input
      type="text"
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      placeholder="Title"
      bind:value={title}
    />
    <input
      type="text"
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      placeholder="Description"
      bind:value={description}
    />
    <select
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      placeholder="Color"
      bind:value={color}
    >
      <option value="blue">Blue</option>
      <option value="red">Red</option>
      <option value="green">Green</option>
    </select>
    <button
      class="border-2 border-neutral-800 rounded-lg p-2 mt-4"
      type="submit"
    >
      Add Todo
    </button>
  </form>
</div>
