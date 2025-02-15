/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hessian2

import (
	"bytes"
	"context"

	"github.com/kitex-contrib/codec-hessian2/pkg/iface"
)

type BaseSerializer struct {
	Protocol iface.BaseProtocol
}

type BaseStruct interface {
	Write(p iface.BaseProtocol) error
	Read(p iface.BaseProtocol) error
}

func NewBaseSerializer() *BaseSerializer {
	protocol := NewBinaryProtocol(&bytes.Buffer{})
	return &BaseSerializer{protocol}
}

func (b *BaseSerializer) Write(ctx context.Context, msg BaseStruct) (bt []byte, err error) {
	if err = msg.Write(b.Protocol); err != nil {
		return
	}

	return
}
