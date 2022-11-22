// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import { useQuery, useQueryClient } from "react-query";
import { Chrome } from "./Chrome";
import { todosService } from "./lib/todoservice";

function App() {
	const queryClient = useQueryClient();

	return (
		<Chrome
			description={
				<>
					<div>
						In this example Namespace deployed a multitier application that consists of:
						<ul className="list-disc list-inside px-5">
							<li>A Vite frontend</li>
							<li>An API server written in Go</li>
							<li>A Postgres instance for the persistence layer</li>
						</ul>
					</div>
					<p className="py-5">
						When you add an item, the api server will handle your request, and store the data in
						Postgres.
					</p>
					<p className="py-5 text-blue-300">
						Difference from the previous example: the Postgres server password is modeled as a
						Namespace secret. For this example, we use a generated secret, since we don't care about
						the actual content, but only want to ensure that Postgres and our API backend use the
						same.
					</p>
				</>
			}>
			<div className="flex flex-col bg-white text-black p-5">
				<form
					className="flex"
					onSubmit={async (e) => {
						e.preventDefault();

						const nameInput = (e.target as any)["name"];
						await todosService.add({ name: nameInput["value"] });
						nameInput.value = "";
						queryClient.invalidateQueries("todoList");
					}}>
					<input
						type="text"
						name="name"
						className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
						placeholder="Type the item description"
					/>
					<div className="mt-5 sm:mt-0 sm:ml-6 sm:flex sm:flex-shrink-0 sm:items-center">
						<button
							type="submit"
							className="inline-flex items-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:text-sm">
							Add item
						</button>
					</div>
				</form>
				<TodoList />
			</div>
		</Chrome>
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
		<div className="py-4">
			<div className="mt-8 flex flex-col">
				<div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
					<div className="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
						<div className="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
							<table className="min-w-full divide-y divide-gray-300">
								<thead className="bg-gray-50">
									<tr>
										<th
											scope="col"
											className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">
											Items
										</th>
									</tr>
								</thead>
								<tbody className="divide-y divide-gray-200 bg-white">
									{data?.map((todo) => (
										<tr key={todo.id}>
											<td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-normal text-gray-900 sm:pl-6">
												{todo.name}
											</td>
										</tr>
									))}
									{!data && (
										<tr>
											<td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-normal text-gray-500 sm:pl-6">
												No items yet
											</td>
										</tr>
									)}
								</tbody>
							</table>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}

export default App;
