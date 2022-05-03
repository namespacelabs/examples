import React, { useEffect, useState } from "react";
import { TodoItem, TodoRelatedData, todosService } from "../services/todos";
import classes from "./ItemInfo.module.css";

export function ItemInfo(props: { item: TodoItem }) {
	const [todoRelatedData, setTodoRelatedData] = useState<TodoRelatedData | undefined>();

	useEffect(() => {
		todosService.getRelatedData(props.item.id).then(setTodoRelatedData);
	}, [props.item]);

	return todoRelatedData ? (
		<div>
			<div className={classes.title}>{props.item.name}</div>
			<div>
				Popularity:{" "}
				<span className={classes.star}>
					{Array.from({ length: todoRelatedData.popularity }).map((_) => "★")}
				</span>
			</div>
		</div>
	) : (
		<div>Loading related data...</div>
	);
}
