import { relations } from "drizzle-orm";
import {
  boolean,
  int,
  mysqlTable,
  serial,
  text,
  varchar,
} from "drizzle-orm/mysql-core";

export const usersTable = mysqlTable("users", {
  id: serial("id").primaryKey(),
  first_name: varchar("first_name", { length: 255 }).notNull(),
  last_name: varchar("last_name", { length: 255 }).notNull(),
  email: varchar("email", { length: 255 }).notNull(),
  password: varchar("password", { length: 255 }).notNull(),
});

export const todosTable = mysqlTable("todos", {
  id: serial("id").primaryKey(),
  title: varchar("title", { length: 255 }).notNull(),
  description: text("description"),
  completed: boolean("completed").notNull(),
  user_id: int("user_id").notNull(),
});

export const usersTableRelations = relations(usersTable, ({ many }) => ({
  todos: many(todosTable),
}));

export const todosTableRelations = relations(todosTable, ({ one }) => ({
  user: one(usersTable, {
    fields: [todosTable.user_id],
    references: [usersTable.id],
  }),
}));
