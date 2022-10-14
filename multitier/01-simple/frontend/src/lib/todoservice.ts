import { nanoid } from "nanoid";
import axios from "axios";
import { Backends } from "../config/backends.ns";

export interface TodoItem {
	id: string;
	name: string;
}

export interface TodosService {
	list: () => Promise<TodoItem[]>;
	add: (todoItem: { name: string }) => Promise<void>;
}

class HttpTodosServiceImpl implements TodosService {
	readonly #baseUrl: string;

	constructor(baseUrl: string) {
		this.#baseUrl = baseUrl;
	}

	list = async () => {
		const response = await axios.post(`${this.#baseUrl}/list`);
		return response.data;
	};

	add = async (todoItem: { name: string }) => {
		await axios.post(`${this.#baseUrl}/add`, {
			id: nanoid(),
			name: todoItem.name,
		});
	};
}

export const todosService: TodosService = new HttpTodosServiceImpl(Backends["api"]["managed"]);
