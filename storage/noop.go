// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"github.com/prometheus/prometheus/pkg/labels"
)

type noopQuerier struct{}

// NoopQuerier is a Querier that does nothing.
func NoopQuerier() Querier {
	return noopQuerier{}
}

func (noopQuerier) Select(bool, *SelectHints, ...*labels.Matcher) (SeriesSet, Warnings, error) {
	return NoopSeriesSet(), nil, nil
}

func (noopQuerier) LabelValues(string) ([]string, Warnings, error) {
	return nil, nil, nil
}

func (noopQuerier) LabelNames() ([]string, Warnings, error) {
	return nil, nil, nil
}

func (noopQuerier) Close() error {
	return nil
}

type noopChunkedQuerier struct{}

// NoopChunkedQuerier is a ChunkQuerier that does nothing.
func NoopChunkedQuerier() ChunkQuerier {
	return noopChunkedQuerier{}
}

func (noopChunkedQuerier) Select(bool, *SelectHints, ...*labels.Matcher) (ChunkSeriesSet, Warnings, error) {
	return NoopChunkedSeriesSet(), nil, nil
}

func (noopChunkedQuerier) LabelValues(string) ([]string, Warnings, error) {
	return nil, nil, nil
}

func (noopChunkedQuerier) LabelNames() ([]string, Warnings, error) {
	return nil, nil, nil
}

func (noopChunkedQuerier) Close() error {
	return nil
}

type noopSeriesSet struct{}

// NoopSeriesSet is a SeriesSet that does nothing.
func NoopSeriesSet() SeriesSet {
	return noopSeriesSet{}
}

func (noopSeriesSet) Next() bool { return false }

func (noopSeriesSet) At() Series { return nil }

func (noopSeriesSet) Err() error { return nil }

type noopChunkedSeriesSet struct{}

// NoopChunkedSeriesSet is a ChunkSeriesSet that does nothing.
func NoopChunkedSeriesSet() ChunkSeriesSet {
	return noopChunkedSeriesSet{}
}

func (noopChunkedSeriesSet) Next() bool { return false }

func (noopChunkedSeriesSet) At() ChunkSeries { return nil }

func (noopChunkedSeriesSet) Err() error { return nil }
