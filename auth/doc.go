// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package auth provides a way to follow a Device Authorization Grant https://datatracker.ietf.org/doc/html/rfc8628.

Usage

	import "go.mongodb.org/atlas/auth"

Construct a new client Config, then use the various methods to complete a flow.
For example:

	config := auth.NewConfigWithOptions(nil, auth.SetClientID("my-client-id"), auth.SetScopes([]string{"openid"}))

	code, _, err := config.RequestCode(ctx)
	if err!= nil {
		panic(err)
	}
	token, _, err := config.PollToken(ctx, code)
	if err!= nil {
		panic(err)
	}
	fmt.PrintLn(accessToken.AccessToken)

NOTE: Using the https://godoc.org/context package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then context.Background()
can be used as a starting point.
*/
package auth
