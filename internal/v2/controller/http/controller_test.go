//
// Copyright (c) 2020 Intel Corporation
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
//

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	contractsV2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingRquest(t *testing.T) {
	target := NewV2Controller(logger.NewMockClient())

	req, err := http.NewRequest(http.MethodGet, contractsV2.ApiPingRoute, nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(target.Ping)

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t, clients.ContentTypeJSON, recorder.HeaderMap.Get(clients.ContentType))
	assert.NotEmpty(t, recorder.HeaderMap.Get(clients.CorrelationHeader))

	require.NotEmpty(t, recorder.Body.String())

	actual := common.PingResponse{}
	err = json.Unmarshal(recorder.Body.Bytes(), &actual)
	require.NoError(t, err)

	_, err = time.Parse(time.UnixDate, actual.Timestamp)
	assert.NoError(t, err)

	require.Equal(t, contractsV2.ApiVersion, actual.Versionable.ApiVersion)
}
