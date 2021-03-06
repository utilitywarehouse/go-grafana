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

import "encoding/json"

type valueMappingType uint

const (
	_ valueMappingType = iota
	ValueToTextType
	RangeToTextType
)

// Singlestat  represents Singlestat panel.
type Singlestat struct {

	// Options. Value.
	ValueName       string `json:"valueName"`     // Stat: min/max/avg/current/total/name/first/delta/diff/range
	ValueFontSize   string `json:"valueFontSize"` // 0%-100% TODO: validation
	Postfix         string `json:"postfix"`
	PostfixFontSize string `json:"postfixFontSize"` // 0%-100% TODO: validation
	Prefix          string `json:"prefix"`
	PrefixFontSize  string `json:"prefixFontSize"` // 0%-100% TODO: validation
	Format          string `json:"format"`         // Unit option. TODO: make a custom type with constants

	// Options. Coloring.
	// Colorize background or not
	ColorBackground bool `json:"colorBackground"`
	// Colorize value or not
	ColorValue bool     `json:"colorValue"`
	Thresholds string   `json:"thresholds"` // comma separated values "x,x". TODO: validation
	Colors     []string `json:"colors"`     // array of 3 colors, ie. rgba(50, 172, 45, 0.97)

	// Options. Spark lines.
	SparkLine struct {
		Show       bool   `json:"show"`
		FullHeight bool   `json:"full"`
		LineColor  string `json:"lineColor"` // ie. rgba(50, 172, 45, 0.97)
		FillColor  string `json:"fillColor"` // ie. rgba(50, 172, 45, 0.97)
	} `json:"sparkline"`

	// Options. Gauge
	Gauge struct {
		Show             bool `json:"show"`
		MaxValue         int  `json:"maxValue"`
		MinValue         int  `json:"minValue"`
		ThresholdLabels  bool `json:"thresholdLabels"`
		ThresholdMarkers bool `json:"thresholdMarkers"`
	} `json:"gauge"`

	ValueMappings
	TimeRangeOptions

	generalOptions GeneralOptions
	queries        []Query
}

// NewSinglestat creates new "Singlestat" panel.
func NewSinglestat() *Singlestat {
	return &Singlestat{}
}

// GeneralOptions implements grafana.Panel interface
func (p *Singlestat) GeneralOptions() *GeneralOptions {
	return &p.generalOptions
}

// Queries implements Queryable interface
func (p *Singlestat) Queries() *[]Query {
	return &p.queries
}

type ValueMappings struct {
	Type        valueMappingType     `json:"mappingType"`
	ValueToText []ValueToTextMapping `json:"valueMaps"`
	RangeToText []RangeToTextMapping `json:"rangeMaps"`
}

type ValueToTextMapping struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

// MarshalJSON implements json.Marshaler interface
func (m ValueToTextMapping) MarshalJSON() ([]byte, error) {
	type JSONMapping ValueToTextMapping
	jm := struct {
		Op operator `json:"op"`
		JSONMapping
	}{
		Op:          EqualSignOp,
		JSONMapping: JSONMapping(m),
	}

	return json.Marshal(jm)
}

type RangeToTextMapping struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
