// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import { Pool, PoolConfig } from "pg";
import { readFileSync } from "fs";

// Lazy load the connection pool as Next.js does not support startup hooks:
// https://github.com/vercel/next.js/discussions/11686
const dbConn = new Promise<Pool>(async (resolve, reject) => {
	try {
		const [host, port] = dbUrlFromNsConfig().split(":");

		const config: PoolConfig = {
			user: "postgres",
			password: process.env.POSTGRES_PASSWORD,
			database: process.env.POSTGRES_DB,
			host: host,
			port: Number.parseInt(port),
		};

		console.log(`PoolConfig: ${JSON.stringify(config, null, 2)}`);

		const conn = new Pool(config);

		// Injecting the schema.
		const schemaFile = readFileSync("schema.sql").toString();
		await conn.query(schemaFile);

		resolve(conn);
	} catch (error) {
		reject(error);
	}
});

function dbUrlFromNsConfig(): string {
	const nsConfigRaw = readFileSync("/namespace/config/runtime.json");
	const nsConfig = JSON.parse(nsConfigRaw.toString());
	console.log(`Namespace config: ${JSON.stringify(nsConfig, null, 2)}`);

	const dbService = nsConfig.stack_entry
		.map((e: any) => e.service)
		.flat()
		.find((s: any) => s.name === "postgres");

	console.log(`dbService: ${JSON.stringify(dbService, null, 2)}`);

	return dbService.endpoint;
}

export default dbConn;
