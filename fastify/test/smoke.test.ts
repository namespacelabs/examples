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
