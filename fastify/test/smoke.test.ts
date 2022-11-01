// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

import { readFileSync } from "fs";
import axios from "axios";

let serviceEndpoint: string;

beforeAll(() => {
	const nsConfig = JSON.parse(readFileSync("/namespace/config/runtime.json").toString());

	serviceEndpoint = nsConfig.stack_entry[0].service[0].endpoint;
});

test("Smoke", async () => {
	const response = await axios.post(`http://${serviceEndpoint}/echo`, { name: "World" });

	expect(response.data.message).toBe("Hello, World!");
});
