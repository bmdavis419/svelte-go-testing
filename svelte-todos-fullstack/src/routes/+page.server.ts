import { db } from "$lib/server/db.js";
import { verifyAuthJWT } from "$lib/server/jwt.js";
import { todosTable, usersTable } from "$lib/server/schema.js";
import { redirect } from "@sveltejs/kit";
import { eq } from "drizzle-orm";

export const load = async ({ cookies, fetch }) => {
  // fetch the current user's todos from the server
  const token = cookies.get("auth_token");

  if (!token) {
    throw redirect(301, "/sign-in");
  }

  const userPayload = await verifyAuthJWT(token);

  const todos = await db
    .select({
      completed: todosTable.completed,
      description: todosTable.description,
      title: todosTable.title,
      id: todosTable.id,
    })
    .from(todosTable)
    .where(eq(todosTable.user_id, userPayload.id));

  return { todos };
};

export const actions = {
  delete: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const id = formData.get("id") || "";

    // ensure the user is logged in
    const token = cookies.get("auth_token");
    if (!token) {
      throw redirect(301, "/sign-in");
    }

    await verifyAuthJWT(token);

    await db
      .delete(todosTable)
      .where(eq(todosTable.id, parseInt(id.toString())));

    return { success: true };
  },
  complete: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const id = formData.get("id") || "";

    // ensure the user is logged in
    const token = cookies.get("auth_token");
    if (!token) {
      throw redirect(301, "/sign-in");
    }

    await verifyAuthJWT(token);

    await db
      .update(todosTable)
      .set({ completed: true })
      .where(eq(todosTable.id, parseInt(id.toString())));

    return { success: true };
  },
  create: async ({ request, fetch, cookies }) => {
    // prepare request body
    const formData = await request.formData();
    const title = formData.get("title") || "";
    const description = formData.get("description") || "";

    // ensure the user is logged in
    const token = cookies.get("auth_token");
    if (!token) {
      throw redirect(301, "/sign-in");
    }

    const userPayload = await verifyAuthJWT(token);

    await db.insert(todosTable).values({
      title: title.toString(),
      description: description.toString(),
      user_id: userPayload.id,
      completed: false,
    });

    return { success: true };
  },
};
