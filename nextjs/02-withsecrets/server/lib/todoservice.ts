import { nanoid } from "nanoid";
import dbConn from "./db";

export interface TodoItem {
	id: string;
	name: string;
}

export interface TodosService {
	list: () => Promise<TodoItem[]>;
	add: (todoItem: { name: string }) => Promise<void>;
}

export class SqlTodosServiceImpl implements TodosService {
	public async list(): Promise<TodoItem[]> {
		const query = "SELECT Id, Name FROM todos_table;";
		const result = await (await dbConn).query(query);
		return result.rows;
	}

	public async add(todoItem: { name: string }): Promise<void> {
		const query = "INSERT INTO todos_table (Id, Name) VALUES ($1, $2);";
		const values = [nanoid(), todoItem.name];
		await (await dbConn).query(query, values);
	}
}

export const todosService: TodosService = new SqlTodosServiceImpl();
