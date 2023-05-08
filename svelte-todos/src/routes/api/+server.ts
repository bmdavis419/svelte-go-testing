import { error, json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const GET = (({ url }) => {
  console.log(url.searchParams);
  // get a random number
  const random = Math.random();

  if (random < 0.5) {
    throw error(500, "The number is too small");
  }

  return json({ random });
}) satisfies RequestHandler;
