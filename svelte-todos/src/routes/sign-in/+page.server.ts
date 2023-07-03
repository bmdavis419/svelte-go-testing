import { API_URL } from "$env/static/private";
import { redirect } from "@sveltejs/kit";

export const load = async (event) => {
  // get the sessionId from the cookie
  const sessionId = event.cookies.get("sessionId");

  // if there is a sessionId, redirect to the user page
  if (sessionId) {
    throw redirect(301, "/");
  }
};

export const actions = {
  default: async (event) => {
    const formData = await event.request.formData();
    const email = formData.get("email");
    const password = formData.get("password");
    const body = await JSON.stringify({ email, password });

    const res = await fetch(`${API_URL}/users/sign-in`, {
      body,
      method: "POST",
      headers: { "content-type": "application/json" },
    });

    // check the status
    if (res.ok) {
      // set the cookie
      const sessionId = res.headers.get("Authorization");
      event.cookies.set("sessionId", sessionId?.split("Bearer ")[1] ?? "", {
        path: "/",
      });

      // redirect to the user page
      throw redirect(301, "/me");
    }

    return {
      error: await res.text(),
    };
  },
};
