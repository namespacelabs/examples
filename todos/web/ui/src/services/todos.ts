import axios from "axios";
import { Backends } from "../../config/backends.fn.js";

export interface TodoItem {
	id: string;
	name: string;
}

interface StreamResponse {
	items: TodoItem[];
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
		const controller = new AbortController();

		fetch(`${this.#baseUrl}/api.todos.TodosService/StreamList`, {
			method: "POST",
			signal: controller.signal,
			keepalive: true,
		}).then(async (w) => {
			const reader = w.body?.getReader();
			if (reader) {
				const p = new ArrayParser<StreamResponse>((value) => {
					onItems(value.items);
				});
				while (true) {
					const { done, value } = await reader.read();
					if (done) break;
					if (value) p.write(value);
				}
				p.close();
			}
		});

		return () => {
			console.log("Stopping StreamList");
			controller.abort();
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

// Parser that is specific to our current Envoy behavior.
class ArrayParser<T> {
	readonly onValue: (value: T) => void;
	arraydepth: number;
	objectdepth: number;
	arrayBuffer: ArrayBuffer;
	cursor: number;
	instring: boolean;

	constructor(onValue: (value: T) => void) {
		this.onValue = onValue;
		this.arraydepth = 0;
		this.objectdepth = 0;
		this.arrayBuffer = new ArrayBuffer(1024 * 1024);
		this.cursor = 0;
		this.instring = false;
	}

	write(data: Uint8Array) {
		const x = new Uint8Array(this.arrayBuffer);

		for (let i = 0; i < data.length; i++) {
			if (this.arraydepth === 0) {
				if (data[i] == 91) {
					// '['
					this.arraydepth++;
				} else {
					throw new Error(`unexpected byte: ${data[i]}`);
				}
			} else {
				x[this.cursor++] = data[i];

				if (this.instring) {
					if (data[i] == 34 && this.cursor > 1 && x[this.cursor - 2] != 92) {
						// `"`, `\`
						this.instring = false;
					}
				} else {
					switch (data[i]) {
						case 123: // {
							this.objectdepth++;
							break;

						case 125: // }
							this.objectdepth--;

							if (this.objectdepth == 0) {
								const raw = new TextDecoder().decode(
									new Uint8Array(this.arrayBuffer, 0, this.cursor)
								);
								console.log({ raw }, JSON.parse(raw));
								this.cursor = 0;
								this.onValue(JSON.parse(raw));
							}
							break;

						case 91: // [
							this.arraydepth++;
							break;

						case 93: // ]
							this.arraydepth--;
							if (this.arraydepth === 0) {
								this.cursor--; // Ignore the closing array byte.
							}
							break;

						case 44: // ,
							if (this.arraydepth === 1) {
								this.cursor--; // Ignore the separating comma.
							}
							break;

						case 34: // "
							this.instring = true;
							break;
					}
				}
			}
		}
	}

	close() {
		if (this.cursor > 0) {
			throw new Error(
				`unconsumed data: ${new TextDecoder().decode(
					new Uint8Array(this.arrayBuffer, 0, this.cursor)
				)}`
			);
		}
	}
}
