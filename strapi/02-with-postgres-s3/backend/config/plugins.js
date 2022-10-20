const fs = require("fs");

const runtimeConfigFn = "/namespace/config/runtime.json";

module.exports = ({ env }) => {
  if (fs.existsSync(runtimeConfigFn)) {
    const nsConfigRaw = fs.readFileSync(runtimeConfigFn);
    const nsConfig = JSON.parse(nsConfigRaw.toString());

    const minioService = nsConfig.stack_entry
      .map((e) => e.service)
      .flat()
      .find((s) => s.name === "minioapi");

    console.log(`Minio endpoint: ${minioService.endpoint}`);

    const [host, port] = minioService.endpoint.split(":");

    return {
      upload: {
        // config: {
        //   provider: "aws-s3",
        //   providerOptions: {
        //     accessKeyId: env("S3_ACCESS_KEY_ID"),
        //     secretAccessKey: env("S3_SECRET_ACCESS_KEY"),
        //     endpoint: `http://${minioService.endpoint}`,
        //     params: {
        //       Bucket: env("S3_BUCKET"),
        //     },
        //     s3BucketEndpoint: true,
        //   },
        // },
        config: {
          provider: "strapi-provider-upload-minio-ce",
          providerOptions: {
            accessKey: env("S3_ACCESS_KEY_ID"),
            secretKey: env("S3_SECRET_ACCESS_KEY"),
            bucket: env("S3_BUCKET"),
            endPoint: host,
            port: port,
            useSSL: false,
          },
        },
      },
    };
  } else {
    // This happens during the build phase in `strapi build`, that shouldn't actualy
    // talk to any servers.
    return {};
  }
};
