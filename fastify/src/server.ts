// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import Fastify from "fastify";
import { readFileSync } from "fs";

const nsConfig = JSON.parse(readFileSync("/namespace/config/runtime.json").toString());

const PORT = nsConfig.current.port.find((s: any) => s.name === "webapi").port;
const HOST = "0.0.0.0";

const fastify = Fastify({
	logger: true,
});

fastify.post("/echo", function (request, reply) {
	reply.send({ message: `Hello, ${(request.body as any)["name"]}!` });
});

// Run the server!
fastify.listen({ port: PORT, host: HOST }, function (err, _) {
	if (err) {
		fastify.log.error(err);
		process.exit(1);
	}
});
