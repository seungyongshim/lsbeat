// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by "stringer -linecomment -type LoadMode"; DO NOT EDIT.

package idxmgmt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LoadModeUnset-0]
	_ = x[LoadModeDisabled-1]
	_ = x[LoadModeEnabled-2]
	_ = x[LoadModeOverwrite-3]
	_ = x[LoadModeForce-4]
}

const _LoadMode_name = "unsetdisabledenabledoverwriteforce"

var _LoadMode_index = [...]uint8{0, 5, 13, 20, 29, 34}

func (i LoadMode) String() string {
	if i >= LoadMode(len(_LoadMode_index)-1) {
		return "LoadMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoadMode_name[_LoadMode_index[i]:_LoadMode_index[i+1]]
}
