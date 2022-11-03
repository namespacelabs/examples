// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

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
