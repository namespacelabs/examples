tests: {
	simpleCurl: {
		integration: shellscript: {
			entrypoint: "./test.sh"
			requiredPackages: ["jq"]
		}
		serversUnderTest: [
			"namespacelabs.dev/examples/dockercompose/withnamespace",
		]
	}
}
