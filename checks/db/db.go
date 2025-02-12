// Copyright 2021 by the contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gsdenys/healthcheck/checks"
)

// Ping returns a Check that validates connectivity to a
// database/sql.DB using Ping().
func Ping(database *sql.DB, timeout time.Duration) checks.Check {
	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if database == nil {
			return fmt.Errorf("database is nil")
		}

		return database.PingContext(ctx)
	}
}
