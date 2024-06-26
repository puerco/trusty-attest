// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package trusty provides an evaluator that uses the trusty API
package trusty

import (
	"time"

	intoto "github.com/in-toto/in-toto-golang/in_toto"
)

type PredicateOpts struct {
	Package PackageInfo
}

func BuildPredicate(opts PredicateOpts, scores []PackageScore) (*Predicate, error) {
	t := time.Now()
	pred := &Predicate{
		Metadata: Metadata{
			Date:        &t,
			PackageInfo: opts.Package,
		},
		Packages: scores,
	}

	return pred, nil
}

func Attest(subjects []intoto.Subject, p *Predicate) (intoto.Statement, error) {
	return intoto.Statement{
		StatementHeader: intoto.StatementHeader{
			Type:          intoto.StatementInTotoV01,
			PredicateType: "http://trustypkg.dev",
			Subject:       []intoto.Subject{},
		},
		Predicate: p,
	}, nil
}
