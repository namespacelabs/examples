// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

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
