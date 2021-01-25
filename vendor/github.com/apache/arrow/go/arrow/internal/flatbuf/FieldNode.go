// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// ----------------------------------------------------------------------
/// Data structures for describing a table row batch (a collection of
/// equal-length Arrow arrays)
/// Metadata about a field at some level of a nested type tree (but not
/// its children).
///
/// For example, a List<Int16> with values [[1, 2, 3], null, [4], [5, 6], null]
/// would have {length: 5, null_count: 2} for its List node, and {length: 6,
/// null_count: 0} for its Int16 node, as separate FieldNode structs
type FieldNode struct {
	_tab flatbuffers.Struct
}

func (rcv *FieldNode) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldNode) Table() flatbuffers.Table {
	return rcv._tab.Table
}

/// The number of value slots in the Arrow array at this level of a nested
/// tree
func (rcv *FieldNode) Length() int64 {
	return rcv._tab.GetInt64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
/// The number of value slots in the Arrow array at this level of a nested
/// tree
func (rcv *FieldNode) MutateLength(n int64) bool {
	return rcv._tab.MutateInt64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

/// The number of observed nulls. Fields with null_count == 0 may choose not
/// to write their physical validity bitmap out as a materialized buffer,
/// instead setting the length of the bitmap buffer to 0.
func (rcv *FieldNode) NullCount() int64 {
	return rcv._tab.GetInt64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
/// The number of observed nulls. Fields with null_count == 0 may choose not
/// to write their physical validity bitmap out as a materialized buffer,
/// instead setting the length of the bitmap buffer to 0.
func (rcv *FieldNode) MutateNullCount(n int64) bool {
	return rcv._tab.MutateInt64(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateFieldNode(builder *flatbuffers.Builder, length int64, nullCount int64) flatbuffers.UOffsetT {
	builder.Prep(8, 16)
	builder.PrependInt64(nullCount)
	builder.PrependInt64(length)
	return builder.Offset()
}
