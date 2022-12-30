// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import { Pool, PoolConfig } from "pg";
import { readFileSync } from "fs";

// Lazy load the connection pool as Next.js does not support startup hooks:
// https://github.com/vercel/next.js/discussions/11686
const dbConn = new Promise<Pool>(async (resolve, reject) => {
	try {
		const [host, port] = process.env.POSTGRES_ADDRESS.split(":");

		const config: PoolConfig = {
			user: "postgres",
			password: process.env.POSTGRES_PASSWORD,
			database: process.env.POSTGRES_DB,
			host: host,
			port: Number.parseInt(port),
		};

		console.log(`PoolConfig: ${JSON.stringify(config, null, 2)}`);

		const conn = new Pool(config);

		// schema is already applied.

		resolve(conn);
	} catch (error) {
		reject(error);
	}
});

export default dbConn;
