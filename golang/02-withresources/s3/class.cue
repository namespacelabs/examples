resourceClasses: {
	"S3": {
		intent: {
			type:   "examples.golang.withresources.s3.S3Intent"
			source: "./api.proto"
		}
		produces: {
			type:   "examples.golang.withresources.s3.S3Instance"
			source: "./api.proto"
		}
	}
}
