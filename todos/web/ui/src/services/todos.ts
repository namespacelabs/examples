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
	streamList: (onItems: (items: TodoItem[]) => void) => () => void;
	add: (todoItem: { name: string }) => Promise<void>;
	remove: (id: string) => Promise<void>;
	getRelatedData: (id: string) => Promise<TodoRelatedData>;
}

export class HttpTodosServiceImpl implements TodosService {
	readonly #baseUrl: string;

	constructor(baseUrl: string) {
		this.#baseUrl = baseUrl;
	}

	public async list(): Promise<TodoItem[]> {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/List`);
		console.log(`Server response: ${JSON.stringify(response.data)}`);
		return response.data["items"] as TodoItem[];
	}

	// TODO: handle reconnects.
	public streamList(onItems: (items: TodoItem[]) => void): () => void {
		const xhr = new XMLHttpRequest();
		xhr.open("POST", `${this.#baseUrl}/api.todos.TodosService/StreamList`, true);

		let previousPtr = 0;
		xhr.onprogress = () => {
			while (true) {
				const newPtr = xhr.responseText.substring(previousPtr).indexOf("\n");
				if (newPtr < 0) {
					return;
				}
				const newLine = xhr.responseText.substring(previousPtr, previousPtr + newPtr);
				previousPtr += newPtr + "\n".length;

				console.log(`StreamList: server response: ${newLine}`);

				const parsed = JSON.parse(newLine);
				onItems(parsed.result?.items as TodoItem[]);
			}
		};

		xhr.send(null);
		console.log(`Started StreamList`);

		return () => {
			console.log("Stopping StreamList");
			xhr.abort();
		};
	}

	public async add(todoItem: { name: string }) {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/Add`, todoItem);
		console.log(`Server response: ${JSON.stringify(response.data)}`);
	}

	public async remove(id: string) {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/Remove`, { id: id });
		console.log(`Server response: ${JSON.stringify(response.data)}`);
	}

	public async getRelatedData(id: string): Promise<TodoRelatedData> {
		const response = await axios.post(`${this.#baseUrl}/api.todos.TodosService/GetRelatedData`, {
			id: id,
		});
		console.log(`Server response: ${JSON.stringify(response.data)}`);
		return response.data;
	}
}

console.log(`Backends: ${JSON.stringify(Backends)}`);

export const todosService: TodosService = new HttpTodosServiceImpl(Backends.apiBackend.managed);
