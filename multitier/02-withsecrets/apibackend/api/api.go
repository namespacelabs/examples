// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package api

type AddRequest struct {
	Name string `json:"name"`
}

type RemoveRequest struct {
	Id string `json:"id"`
}

type TodoItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
