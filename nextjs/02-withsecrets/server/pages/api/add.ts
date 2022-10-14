import type { NextApiRequest, NextApiResponse } from "next";
import { TodoItem, todosService } from "../../lib/todoservice";

export default async function handler(req: NextApiRequest, res: NextApiResponse<TodoItem>) {
	try {
		await todosService.add({ name: req.body.name });
		res.redirect("/");
	} catch (error) {
		console.log(error);
	}

	res.end();
}
