import { redirect } from "@sveltejs/kit";
import { API_URL } from "$env/static/private";

export const load = async (event) => {
  // get the sessionId from the cookie
  const sessionId = event.cookies.get("sessionId");

  // if there is no sessionId, redirect to the sign-in page
  if (!sessionId) {
    throw redirect(301, "/sign-in");
  }

  // get the user's info from the server
  const res = await event.fetch(`${API_URL}/users/me`, {
    headers: {
      Authorization: `Bearer ${sessionId}`,
    },
  });

  // check the status
  if (res.ok) {
    const user = (await res.json()) as {
      Id: string;
      FirstName: string;
      LastName: string;
      Email: string;
    };

    return {
      props: {
        user,
      },
    };
  }
};
