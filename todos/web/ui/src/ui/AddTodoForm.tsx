import React from "react";
import { Button, Form } from "react-bootstrap";
import { useForm } from "react-hook-form";
import classes from "./AddTodoForm.module.css";
import { todosService } from "../services/todos";

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
