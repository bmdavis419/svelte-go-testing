import type { PageLoad } from "./$types";

export const load = (({ params }) => {
  return {
    todos: [
      {
        id: 1,
        title: "Todo 1",
        completed: false,
        description: "This is a todo",
        color: "red",
      },
      {
        id: 2,
        title: "Todo 2",
        completed: true,
        description: "This is another todo",
        color: "green",
      },
    ],
  };
}) satisfies PageLoad;
