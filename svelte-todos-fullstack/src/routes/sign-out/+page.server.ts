import { redirect } from "@sveltejs/kit";

export const load = async (event) => {
  // remove the cookie
  event.cookies.set("auth_token", "");

  // redirect to the sign-in page
  throw redirect(301, "/sign-in");
};
