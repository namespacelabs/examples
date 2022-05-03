import axios from "axios";
import { Backends } from "../../config/backends.fn.js";

export interface TodoItem {
	id: string;
	name: string;
}

export interface TodoRelatedData {
	popularity: number;
}

export interface TodosService {
	list: () => Promise<TodoItem[]>;
	add: (todoItem: { name: string }) => Promise<void>;
	remove: (id: string) => Promise<void>;
	getRelatedData: (id: string) => Promise<TodoRelatedData>;
}

export class HttpTodosServiceImpl implements TodosService {
	readonly #baseUrl: string;

	constructor(baseUrl: string) {
		this.#baseUrl = baseUrl;
	}

	list = async (): Promise<TodoItem[]> => {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/List`);
		console.log(`Server response: ${JSON.stringify(response.data)}`);
		return response.data["items"] as TodoItem[];
	};

	add = async (todoItem: { name: string }) => {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/Add`, todoItem);
		console.log(`Server response: ${JSON.stringify(response.data)}`);
	};

	remove = async (id: string) => {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/Remove`, { id: id });
		console.log(`Server response: ${JSON.stringify(response.data)}`);
	};

	getRelatedData = async (id: string): Promise<TodoRelatedData> => {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/GetRelatedData`, {
			id: id,
		});
		console.log(`Server response: ${JSON.stringify(response.data)}`);
		return response.data;
	};
}

console.log(`Backends: ${JSON.stringify(Backends)}`);

export const todosService: TodosService = new HttpTodosServiceImpl(Backends.apiBackend.managed);
