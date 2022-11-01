// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import type { NextPage } from "next";
import Head from "next/head";
import { TodoItem, todosService } from "../lib/todoservice";
import styles from "../styles/Home.module.css";

interface Props {
	todoList: TodoItem[];
}

const Home: NextPage<Props> = (props) => {
	return (
		<>
			<Head>
				<title>Namespace: Next.js example</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main className={styles.main}>
				<form className={styles.form} action="/api/add" method="post">
					<input type="text" name="name" />
					<button type="submit">Add</button>
				</form>
				<ul>
					{props.todoList.map((todo) => (
						<li key={todo.id}>{todo.name}</li>
					))}
				</ul>
			</main>
		</>
	);
};

export async function getServerSideProps() {
	const todoList = await todosService.list();

	return { props: { todoList } };
}

export default Home;
