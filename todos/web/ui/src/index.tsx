import React, { useState } from "react";
import ReactDOM from "react-dom";
import { TodoList } from "./ui/TodoList";
import classes from "./index.module.css";
import { AddTodoForm } from "./ui/AddTodoForm";
import { Card } from "react-bootstrap";
import { ItemInfo } from "./ui/ItemInfo";
import { TodoItem } from "./services/todos";

ReactDOM.render(
	<React.StrictMode>
		<App />
	</React.StrictMode>,
	document.getElementById("app")
);

function App() {
	const [selectedItem, setSelectedItem] = useState<TodoItem | undefined>();

	return (
		<div className={classes.root}>
			<Card className={classes.card}>
				<AddTodoForm></AddTodoForm>
				<TodoList onSelectedChanged={setSelectedItem}></TodoList>
			</Card>
			{selectedItem ? (
				<Card className={classes.card}>
					<ItemInfo item={selectedItem}></ItemInfo>
				</Card>
			) : (
				""
			)}
		</div>
	);
}
