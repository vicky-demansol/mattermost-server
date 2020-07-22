// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClusterMessage(t *testing.T) {
	m := ClusterMessage{
		Event:    CLUSTER_EVENT_PUBLISH,
		SendType: CLUSTER_SEND_BEST_EFFORT,
		Data:     "hello",
	}
	json := m.ToJson()
	result := ClusterMessageFromJson(strings.NewReader(json))

	require.Equal(t, "hello", result.Data)

	badresult := ClusterMessageFromJson(strings.NewReader("junk"))

	require.Nil(t, badresult, "should not have parsed")
}
