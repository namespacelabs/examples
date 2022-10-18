import { useQuery, useQueryClient } from "react-query";
import styles from "./App.module.css";
import { todosService } from "./lib/todoservice";

function App() {
	const queryClient = useQueryClient();

	return (
		<div>
			<main className={styles.main}>
				<form
					className={styles.form}
					onSubmit={async (e) => {
						e.preventDefault();

						const nameInput = (e.target as any)["name"];
						await todosService.add({ name: nameInput["value"] });
						nameInput.value = "";
						queryClient.invalidateQueries("todoList");
					}}>
					<input type="text" name="name" />
					<button type="submit">Add</button>
				</form>
				<TodoList />
			</main>
		</div>
	);
}

function TodoList() {
	const { isLoading, error, data } = useQuery(["todoList"], todosService.list);

	if (error) {
		return <>An error has occurred: {(error as any)["message"]}</>;
	}

	if (isLoading) {
		return <>Loading...</>;
	}

	return (
		<ul>
			{data?.map((todo) => (
				<li key={todo.id}>{todo.name}</li>
			))}
		</ul>
	);
}

export default App;
