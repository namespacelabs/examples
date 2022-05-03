import React, { useState } from "react";
import { ListGroup } from "react-bootstrap";
import { useInterval } from "usehooks-ts";
import { TodoItem, todosService } from "../services/todos";
import classes from "./TodoList.module.css";

export function TodoList(props: {
	selectedItem?: TodoItem;
	setSelectedItem: (id: TodoItem | undefined) => void;
}) {
	const [todoList, setTodoList] = useState<TodoItem[] | undefined>();
	const { selectedItem, setSelectedItem } = props;

	// Periodically poll the list of items.
	useInterval(async () => {
		const list = await todosService.list();
		setTodoList(list);
		if (!list.find((i) => i.id === selectedItem?.id)) {
			setSelectedItem(undefined);
		}
	}, 200 /* ms */);

	return todoList ? (
		<div>
			{todoList.length ? (
				<>
					<div className={classes.itemsTitle}>Items:</div>
					<ListGroup>
						{todoList.map((todoItem) => (
							<TodoItem
								item={todoItem}
								key={todoItem.id}
								isSelected={todoItem.id === selectedItem?.id}
								onClick={() => {
									const newSelectedItem = todoItem.id === selectedItem?.id ? undefined : todoItem;
									setSelectedItem(newSelectedItem);
								}}></TodoItem>
						))}
					</ListGroup>
				</>
			) : (
				<span>No items</span>
			)}
		</div>
	) : (
		<span>Loading data from the server...</span>
	);
}

function TodoItem(props: { item: TodoItem; isSelected: boolean; onClick?: () => void }) {
	return (
		<ListGroup.Item className={classes.todoItem} active={props.isSelected} action>
			<div onClick={props.onClick}>{props.item.name}</div>

			<div
				className={classes.removeButton}
				onClick={() => {
					todosService.remove(props.item.id);
				}}></div>
		</ListGroup.Item>
	);
}
