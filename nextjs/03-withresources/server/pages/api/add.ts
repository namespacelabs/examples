// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

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
