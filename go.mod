module namespacelabs.dev/examples

go 1.18

require (
	github.com/aws/aws-sdk-go-v2 v1.17.3
	github.com/aws/aws-sdk-go-v2/config v1.18.7
	github.com/aws/aws-sdk-go-v2/credentials v1.13.7
	github.com/aws/aws-sdk-go-v2/service/s3 v1.29.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx/v4 v4.16.1
	namespacelabs.dev/foundation v0.0.111-0.20221212125817-0ce5a3246be4
	namespacelabs.dev/go-ids v0.0.0-20220708072700-49c0a87d068b
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.20 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.7 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jxskiss/base62 v1.1.0 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/crypto v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20220613132600-b0d781184e0d // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace sigs.k8s.io/aws-iam-authenticator => namespacelabs.dev/aws-iam-authenticator v0.5.8-0.20220603141112-71d1e80cc0c1
