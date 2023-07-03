import { API_URL } from "$env/static/private";
import { redirect } from "@sveltejs/kit";

export const load = async ({ cookies, fetch }) => {
  // fetch the current user's todos from the server
  const sessionId = cookies.get("sessionId");

  if (!sessionId) {
    throw redirect(301, "/sign-in");
  }

  const res = await fetch(`${API_URL}/todos`, {
    headers: {
      Authorization: `Bearer ${sessionId}`,
    },
  });

  if (!res.ok) {
    throw new Error(res.statusText);
  }

  const respBody = (await res.json()) as {
    todos: {
      completed: boolean;
      description: string;
      title: string;
      id: number;
    }[];
  };

  return respBody;
};

export const actions = {
  delete: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const id = formData.get("id") || "";

    // ensure the user is logged in
    const sessionId = cookies.get("sessionId");
    if (!sessionId) {
      throw redirect(301, "/sign-in");
    }

    const res = await fetch(`${API_URL}/todos/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
    });

    if (!res.ok) {
      throw new Error(res.statusText);
    }

    return { success: true };
  },
  complete: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const id = formData.get("id") || "";

    // ensure the user is logged in
    const sessionId = cookies.get("sessionId");
    if (!sessionId) {
      throw redirect(301, "/sign-in");
    }

    const res = await fetch(`${API_URL}/todos/${id}/complete`, {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${sessionId}`,
      },
    });

    if (!res.ok) {
      throw new Error(res.statusText);
    }

    return { success: true };
  },
  create: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const title = formData.get("title") || "";
    const description = formData.get("description") || "";
    const body = await JSON.stringify({ title, description });

    // ensure the user is logged in
    const sessionId = cookies.get("sessionId");
    if (!sessionId) {
      throw redirect(301, "/sign-in");
    }

    // create the todo
    const res = await fetch(`${API_URL}/todos`, {
      method: "POST",
      body,
      headers: {
        "content-type": "application/json",
        Authorization: `Bearer ${sessionId}`,
      },
    });

    if (!res.ok) {
      throw new Error(res.statusText);
    }

    return { success: true };
  },
};
