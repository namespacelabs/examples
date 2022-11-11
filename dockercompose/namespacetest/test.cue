tests: {
	simpleCurl: {
		integration: shellscript: "./test.sh"

		env: {
			ENDPOINT: fromServiceEndpoint: "namespacelabs.dev/examples/dockercompose/withnamespace:web"
		}

		serversUnderTest: [
			"namespacelabs.dev/examples/dockercompose/withnamespace",
		]
	}
}
