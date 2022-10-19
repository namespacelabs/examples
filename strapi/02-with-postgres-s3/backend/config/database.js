const fs = require("fs");

module.exports = ({ env }) => {
  const nsConfigRaw = fs.readFileSync("/namespace/config/runtime.json");
  const nsConfig = JSON.parse(nsConfigRaw.toString());

  const postgresService = nsConfig.stack_entry
    .map((e) => e.service)
    .flat()
    .find((s) => s.name === "postgres");

  console.log(`Postgres endpoint: ${postgresService.endpoint}`);

  const [host, port] = postgresService.endpoint.split(":");

  return {
    connection: {
      client: "postgres",
      connection: {
        host: host,
        port: port,
        user: env("DATABASE_USERNAME", "postgres"),
        // Configured in server.cue
        database: env("DATABASE_NAME", ""),
        // Configured in server.cue
        password: env("DATABASE_PASSWORD", ""),
        ssl: env("DATABASE_SSL", false),
      },
      debug: false,
    },
  };
};
