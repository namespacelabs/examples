import React from "react";
import { Button, Form } from "react-bootstrap";
import { useForm } from "react-hook-form";
import classes from "./add_todo_form.module.css";
import { todosService } from "./todos_service";

export function AddTodoForm() {
	const { register, handleSubmit, reset } = useForm();
	const onSubmit = async (data: { [key: string]: any }) => {
		await todosService.add({ name: data["name"] });
		reset();
	};

	return (
		<Form className={classes.addTodoForm} onSubmit={handleSubmit(onSubmit)}>
			<Form.Control
				required
				placeholder="What needs to be done?"
				autoComplete="off"
				{...register("name")}
			/>

			<Button variant="primary" type="submit">
				Add
			</Button>
		</Form>
	);
}
