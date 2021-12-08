// Copyright 2017 by the contributors.
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

package checks

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestTCPDialCheck(t *testing.T) {
	assert.NoError(t, TCPDial("heptio.com:80", 5*time.Second)())
	assert.Error(t, TCPDial("heptio.com:25327", 5*time.Second)())
}

func TestHTTPGetCheck(t *testing.T) {
	assert.NoError(t, HTTPGet("https://heptio.com", 5*time.Second)())
	assert.Error(t, HTTPGet("http://heptio.com", 5*time.Second)(), "redirect should fail")
	assert.Error(t, HTTPGet("https://heptio.com/nonexistent", 5*time.Second)(), "404 should fail")
}

func TestDatabasePingCheck(t *testing.T) {
	assert.Error(t, DatabasePing(nil, 1*time.Second)(), "nil DB should fail")

	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	assert.NoError(t, DatabasePing(db, 1*time.Second)(), "ping should succeed")
}

func TestDNSResolveCheck(t *testing.T) {
	assert.NoError(t, DNSResolve("heptio.com", 5*time.Second)())
	assert.Error(t, DNSResolve("nonexistent.heptio.com", 5*time.Second)())
}

func TestGoroutineCountCheck(t *testing.T) {
	assert.NoError(t, GoroutineCount(1000)())
	assert.Error(t, GoroutineCount(0)())
}

func TestGCMaxPauseCheck(t *testing.T) {
	runtime.GC()
	assert.NoError(t, GCMaxPause(1*time.Second)())
	assert.Error(t, GCMaxPause(0)())
}