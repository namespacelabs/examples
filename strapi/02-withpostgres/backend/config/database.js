// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

const fs = require("fs");

const runtimeConfigFn = "/namespace/config/runtime.json";

module.exports = ({ env }) => {
  if (fs.existsSync(runtimeConfigFn)) {
    const nsConfigRaw = fs.readFileSync(runtimeConfigFn);
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
  } else {
    // This happens during the build phase in `strapi build`, that shouldn't actualy
    // require a DB connection.
    return {};
  }
};
