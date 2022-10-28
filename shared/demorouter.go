// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package shared

import (
	"fmt"
	"net/http"
	"strings"

	"namespacelabs.dev/foundation/schema/runtime"
)

func WelcomePage(srv *runtime.Server) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		pkg := strings.TrimPrefix(srv.PackageName, srv.ModuleName+"/")
		fmt.Fprintf(rw, "%s server is up and running \\o/\n\n", pkg)
		fmt.Fprintf(rw, "Try running its integration test with `ns test %s`.", pkg)
	}
}
