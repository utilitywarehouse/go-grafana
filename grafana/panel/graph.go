// Copyright 2017 Sergey Safonov
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

package panel

import (
	"encoding/json"
	"errors"

	"github.com/guregu/null"
	"github.com/spoof/go-grafana/pkg/field"
)

type graphXAxisMode string

const (
	graphXAxisHistogram graphXAxisMode = "histogram"
	graphXAxisSeries    graphXAxisMode = "series"
	graphXAxisTime      graphXAxisMode = "time"
)

// Graph represents Graph panel
type Graph struct {
	YAxes GraphYaxesOptions `json:"yaxes"`
	XAxis struct {
		Buckets null.Int       `json:"buckets,omitempty"`
		Mode    graphXAxisMode `json:"mode"`           // TODO: add validaition here: histogram/series/time
		Name    *string        `json:"name,omitempty"` // it's seems that it's not used anymore
		Show    bool           `json:"show"`
		Values  []string       `json:"values"` // TODO: actually it's only single value here. Need custom type
	} `json:"xaxis"`

	generalOptions GeneralOptions
	//queriesOptions QueriesOptions
	queries []Query
}

// NewGraph creates new Graph panel.
func NewGraph() *Graph {
	return &Graph{}
}

// GeneralOptions implements grafana.Panel interface
func (p *Graph) GeneralOptions() *GeneralOptions {
	return &p.generalOptions
}

// Queries implements Queryable interface
func (p *Graph) Queries() *[]Query {
	return &p.queries
}

type GraphYaxesOptions struct {
	Left  GraphYAxis
	Right GraphYAxis
}

func (y *GraphYaxesOptions) MarshalJSON() ([]byte, error) {
	axes := []GraphYAxis{y.Left, y.Right}
	return json.Marshal(axes)
}
func (y *GraphYaxesOptions) UnmarshalJSON(data []byte) error {
	var axes []GraphYAxis
	if err := json.Unmarshal(data, &axes); err != nil {
		return err
	}
	if len(axes) < 2 {
		return errors.New("Axes should be 2")
	}

	y.Left = axes[0]
	y.Right = axes[1]

	return nil
}

type GraphYAxis struct {
	Format  string             `json:"format"` // TODO: replace with custom type with default value "short".
	Label   string             `json:"label,omitempty"`
	LogBase int                `json:"logBase"` // TODO: default value should be 1 (linear)
	Max     *field.ForceString `json:"max"`
	Min     *field.ForceString `json:"min"`
	Show    bool               `json:"show"`
}
