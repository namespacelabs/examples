server: {
    name: "django-todo"
    integration: dockerfile: src: "Dockerfile"
    args: ["./scripts/runserver.sh"]
    env: {
        NS_ENV_NAME: $env.name
        NS_ENV_PURPOSE: $env.purpose
        DJANGO_APPLICATION_SETTINGS: "todo.settings.dev"
        DJANGO_DATABASE: fromResourceField: {
            resource: ":db"
            fieldRef: "connection_uri"
        }
        DJANGO_SECRET_KEY: "secret"
        DJANGO_DEBUG: "True"
        AWS_ACCESS_KEY_ID: fromResourceField: {
			resource: ":staticBucket"
			fieldRef: "accessKey"
		}
		AWS_SECRET_ACCESS_KEY: fromResourceField: {
			resource: ":staticBucket"
			fieldRef: "secretAccessKey"
		}
		AWS_STORAGE_BUCKET_NAME: fromResourceField: {
			resource: ":staticBucket"
			fieldRef: "bucketName"
		}
		AWS_S3_ENDPOINT_URL: fromResourceField: {
			resource: ":staticBucket"
			fieldRef: "url"
		}
        DJANGO_STATIC_BASE_URL: "http://localhost:4566"
        // DJANGO_STATIC_URL_BASE: fromServiceIngress: "namespacelabs.dev/examples/django/01-original/infra/localstack:api"
        // MY_INGRESS: fromServiceIngress: ":app"
    }
    services: {
        app: {
            port: 8000
            kind: "http"
            ingress: true
        }
    }
    
    resources: {
        db: {
            class:    "namespacelabs.dev/foundation/library/database/postgres:Database"
            provider: "namespacelabs.dev/foundation/library/oss/postgres"

            intent: {
                name: "djangotodo"
            }

            resources: {
                "cluster": ":postgresCluster"
            }
        }
        staticBucket: {
            class:    "namespacelabs.dev/foundation/library/storage/s3:Bucket"
			provider: "namespacelabs.dev/foundation/library/oss/localstack"

			intent: {
				bucketName: "ns-django-example-static"
			}

            resources: {
                "cluster": ":localstackCluster"
            }
        }
    }
    // requires: [
    //     "namespacelabs.dev/examples/django/01-original/infra/localstack"
    // ]
}

resources: {
    postgresCluster: {
        class:    "namespacelabs.dev/foundation/library/database/postgres:Cluster"
        provider: "namespacelabs.dev/foundation/library/oss/postgres"

        intent: {
            server: packageName: "namespacelabs.dev/examples/django/01-base/infra/postgres"
        }
    }
    localstackCluster: {
        class:  "namespacelabs.dev/foundation/library/oss/localstack:Cluster"
        provider: "namespacelabs.dev/foundation/library/oss/localstack"

        intent: {
            server: packageName: "namespacelabs.dev/examples/django/01-base/infra/localstack"
        }
    }
}