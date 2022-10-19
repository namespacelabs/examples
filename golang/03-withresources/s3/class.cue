resourceClasses: {
	"Bucket": {
		intent: {
			type:   "examples.golang.withresources.s3.BucketIntent"
			source: "./api.proto"
		}
		produces: {
			type:   "examples.golang.withresources.s3.BucketInstance"
			source: "./api.proto"
		}
	}
}
