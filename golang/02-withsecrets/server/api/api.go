// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package api

type PutRequest struct {
	Key  string `json:"key"`
	Body []byte `json:"body"`
}

type GetRequest struct {
	Key string `json:"key"`
}

type GetResponse struct {
	Body []byte `json:"body"`
}
