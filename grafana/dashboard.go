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

package grafana

import (
	"encoding/json"
	"time"

	"github.com/utilitywarehouse/go-grafana/grafana/panel"
	panelQuery "github.com/utilitywarehouse/go-grafana/grafana/query"
)

type (
	// DashboardID is an ID type of Dashboard
	DashboardID    uint64
	dashboardStyle string
)

const (
	dashboardDarkStyle  dashboardStyle = "dark"
	dashboardLightStyle dashboardStyle = "light"
)

type Dashboard struct {
	Annotations struct {
		List []interface{} `json:"list"`
	} `json:"annotations"`
	Editable      bool            `json:"editable"`
	GraphTooltip  int             `json:"graphTooltip"`
	HideControls  bool            `json:"hideControls"`
	ID            DashboardID     `json:"-"`
	Links         []interface{}   `json:"links"`
	Refresh       interface{}     `json:"refresh"`
	Rows          []*DashboardRow `json:"rows"`
	SchemaVersion int             `json:"schemaVersion"`
	Style         dashboardStyle  `json:"style"`
	Tags          []string        `json:"tags"`
	Templating    struct {
		List []struct {
			Auto      bool   `json:"auto"`
			AutoCount int    `json:"auto_count"`
			AutoMin   string `json:"auto_min"`
			Current   struct {
				Text  string      `json:"text"`
				Value interface{} `json:"value"`
			} `json:"current"`
			Hide    int         `json:"hide"`
			Label   interface{} `json:"label"`
			Name    string      `json:"name"`
			Options []struct {
				Selected bool        `json:"selected"`
				Text     string      `json:"text"`
				Value    interface{} `json:"value"`
			} `json:"options"`
			Query   string `json:"query"`
			Refresh int    `json:"refresh"`
			Type    string `json:"type"`
		} `json:"list"`
	} `json:"templating"`
	Time struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"time"`
	Timepicker struct {
		RefreshIntervals []string `json:"refresh_intervals"`
		TimeOptions      []string `json:"time_options"`
	} `json:"timepicker"`
	Timezone string         `json:"timezone"`
	Title    string         `json:"title"`
	UID      string         `json:"uid"`
	Version  uint64         `json:"-"`
	Meta     *DashboardMeta `json:"meta"`
}

// NewDashboard creates new Dashboard.
func NewDashboard(title string) *Dashboard {
	return &Dashboard{
		Title:         title,
		Editable:      true,
		SchemaVersion: 14,
		Style:         dashboardDarkStyle,
		Tags:          []string{},
	}
}

type DashboardMeta struct {
	Type        string    `json:"type"`
	CanSave     bool      `json:"canSave"`
	CanEdit     bool      `json:"canEdit"`
	CanAdmin    bool      `json:"canAdmin"`
	CanStar     bool      `json:"canStar"`
	Slug        string    `json:"slug"`
	URL         string    `json:"url"`
	Expires     time.Time `json:"expires"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"-"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedBy   string    `json:"createdBy"`
	Version     int       `json:"version"`
	HasAcl      bool      `json:"hasAcl"`
	IsFolder    bool      `json:"isFolder"`
	FolderID    int       `json:"folderId"`
	FolderTitle string    `json:"folderTitle"`
	FolderURL   string    `json:"folderUrl"`
}

type DashboardRow struct {
	Collapse bool        `json:"collapse"`
	Height   interface{} `json:"height"`
	Panels   []struct {
		AliasColors struct {
		} `json:"aliasColors"`
		Bars       bool   `json:"bars"`
		DashLength int    `json:"dashLength"`
		Dashes     bool   `json:"dashes"`
		Datasource string `json:"datasource"`
		Editable   bool   `json:"editable"`
		Error      bool   `json:"error"`
		Fill       int    `json:"fill"`
		ID         int    `json:"id"`
		Legend     struct {
			Avg     bool `json:"avg"`
			Current bool `json:"current"`
			Max     bool `json:"max"`
			Min     bool `json:"min"`
			Show    bool `json:"show"`
			Total   bool `json:"total"`
			Values  bool `json:"values"`
		} `json:"legend"`
		Lines           bool          `json:"lines"`
		Linewidth       int           `json:"linewidth"`
		Links           []interface{} `json:"links"`
		NullPointMode   string        `json:"nullPointMode"`
		Percentage      bool          `json:"percentage"`
		Pointradius     int           `json:"pointradius"`
		Points          bool          `json:"points"`
		Renderer        string        `json:"renderer"`
		SeriesOverrides []interface{} `json:"seriesOverrides"`
		SpaceLength     int           `json:"spaceLength"`
		Span            float64       `json:"span"`
		Stack           bool          `json:"stack"`
		SteppedLine     bool          `json:"steppedLine"`
		Targets         []struct {
			Expr           string `json:"expr"`
			Format         string `json:"format"`
			Interval       string `json:"interval"`
			IntervalFactor int    `json:"intervalFactor"`
			LegendFormat   string `json:"legendFormat"`
			Metric         string `json:"metric"`
			RefID          string `json:"refId"`
			Step           int    `json:"step"`
		} `json:"targets"`
		Thresholds interface{} `json:"thresholds"`
		TimeFrom   interface{} `json:"timeFrom"`
		TimeShift  interface{} `json:"timeShift"`
		Title      string      `json:"title"`
		Tooltip    struct {
			MsResolution bool   `json:"msResolution"`
			Shared       bool   `json:"shared"`
			Sort         int    `json:"sort"`
			ValueType    string `json:"value_type"`
		} `json:"tooltip"`
		Type  string `json:"type"`
		Xaxis struct {
			Buckets interface{}   `json:"buckets"`
			Mode    string        `json:"mode"`
			Name    interface{}   `json:"name"`
			Show    bool          `json:"show"`
			Values  []interface{} `json:"values"`
		} `json:"xaxis"`
		Yaxes []struct {
			Format  string      `json:"format"`
			Label   string      `json:"label"`
			LogBase int         `json:"logBase"`
			Max     interface{} `json:"max"`
			Min     interface{} `json:"min"`
			Show    bool        `json:"show"`
		} `json:"yaxes"`
	} `json:"panels"`
	Repeat          interface{} `json:"repeat"`
	RepeatIteration interface{} `json:"repeatIteration"`
	RepeatRowID     interface{} `json:"repeatRowId"`
	ShowTitle       bool        `json:"showTitle"`
	Title           string      `json:"title"`
	TitleSize       string      `json:"titleSize"`
}

//Panel represents Dashboard's panel
type Panel interface {
	GeneralOptions() *panel.GeneralOptions
}

type panelType string

const (
	textPanelType       panelType = "text"
	singlestatPanelType panelType = "singlestat"
	graphPanelType      panelType = "graph"
)

type probePanel struct {
	ID   uint      `json:"id"`
	Type panelType `json:"type"`

	panel Panel
}

func (p *probePanel) GeneralOptions() *panel.GeneralOptions {
	return p.panel.GeneralOptions()
}

func (p *probePanel) UnmarshalJSON(data []byte) error {
	type JSONPanel probePanel
	jp := struct {
		*JSONPanel
	}{
		JSONPanel: (*JSONPanel)(p),
	}
	if err := json.Unmarshal(data, &jp); err != nil {
		return err
	}

	var pp Panel
	switch jp.Type {
	case textPanelType:
		pp = new(panel.Text)
	case singlestatPanelType:
		pp = new(panel.Singlestat)
	case graphPanelType:
		pp = new(panel.Graph)
	default:
		return nil
	}

	if err := json.Unmarshal(data, pp); err != nil {
		return err
	}

	// Unmarshal general options
	var generalOptions panel.GeneralOptions
	if err := json.Unmarshal(data, &generalOptions); err != nil {
		return err
	}
	gOpts := pp.GeneralOptions()
	*gOpts = generalOptions

	// Unmarshal queries
	var queriesOpts queriesOptions
	if err := json.Unmarshal(data, &queriesOpts); err != nil {
		return err
	}
	if queryablePanel, ok := pp.(QueryablePanel); ok {
		queriesPtr := queryablePanel.Queries()
		newQueries := []panel.Query{}
		for _, q := range queriesOpts.Queries {
			if q.query == nil {
				continue
			}
			newQueries = append(newQueries, q.query)
		}
		*queriesPtr = newQueries
	}

	p.panel = pp
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (p *probePanel) MarshalJSON() ([]byte, error) {
	type JSONPanel probePanel
	jp := struct {
		*JSONPanel

		*panel.Text
		*panel.Singlestat
		*panel.Graph

		*panel.GeneralOptions
		*queriesOptions
	}{
		JSONPanel:      (*JSONPanel)(p),
		GeneralOptions: p.GeneralOptions(),
	}

	switch v := p.panel.(type) {
	case *panel.Text:
		jp.Text = v
		jp.Type = textPanelType
	case *panel.Singlestat:
		jp.Singlestat = v
		jp.Type = singlestatPanelType
	case *panel.Graph:
		jp.Graph = v
		jp.Type = graphPanelType
	}

	if qp, ok := p.panel.(QueryablePanel); ok {
		// Determine do each query uses its own datassource or not
		isOwnDatasource := false
		queries := *qp.Queries()
		for i := 0; i < len(queries)-1; i++ {
			if queries[i].Datasource() != queries[i+1].Datasource() {
				isOwnDatasource = true
			}
		}

		probeQueries := make([]probeQuery, len(queries))
		for i, q := range queries {
			pq := probeQuery{
				RefID: makeRefID(i),
				query: q,
			}

			if isOwnDatasource {
				pq.Datasource = q.Datasource()
			}
			probeQueries[i] = pq
		}

		var datasource string
		if isOwnDatasource {
			datasource = mixedDatasource
		} else {
			if len(queries) > 0 {
				datasource = queries[0].Datasource()
			}
		}

		jp.queriesOptions = &queriesOptions{
			Queries:    probeQueries,
			Datasource: datasource,
		}
	}

	return json.Marshal(jp)
}

// QueryablePanel is interface for panels that supports quering metrics from datasources.
type QueryablePanel interface {
	Queries() *[]panel.Query
}

const mixedDatasource = "-- Mixed --"

type queriesOptions struct {
	Datasource string       `json:"datasource,omitempty"`
	Queries    []probeQuery `json:"targets"`
}

// probeQuery is an auxiliary entity thats purpose to manage marshaling and unmarshal of panel's query into concrete
// types.
type probeQuery struct {
	RefID      string `json:"refid"`
	Datasource string `json:"datasource,omitempty"`

	query panel.Query
}

// UnmarshalJSON implements json.Unmarshaler interface
func (q *probeQuery) UnmarshalJSON(data []byte) error {
	type JSONQuery probeQuery
	jq := struct {
		*JSONQuery

		// Prometheus query fields
		IntervalFactor *uint   `json:"intervalFactor"`
		Expression     *string `json:"expr"`

		// Graphite queryfields
		Target *string `json:"target"`
	}{
		JSONQuery: (*JSONQuery)(q),
	}
	if err := json.Unmarshal(data, &jq); err != nil {
		return err
	}

	// There is no any information about query type in Grafana's JSON object. Further more, some queries uses the same
	// fields. Thus we need to use some heurisitcs to map json fields into our query types properly.
	// This heurisitcs based on searching specific for query type fields in JSON data.
	var query panel.Query
	if jq.Expression != nil && jq.IntervalFactor != nil {
		query = new(panelQuery.Prometheus)
	} else if jq.Target != nil {
		query = new(panelQuery.Graphite)
	}

	// TODO: Initialize Unknown query here instead
	if query == nil {
		return nil
	}

	if err := json.Unmarshal(data, &query); err != nil {
		return err
	}

	q.query = query
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (q *probeQuery) MarshalJSON() ([]byte, error) {
	type JSONQuery probeQuery
	jq := struct {
		*JSONQuery
		*panelQuery.Prometheus
		*panelQuery.Graphite
	}{
		JSONQuery: (*JSONQuery)(q),
	}

	switch v := q.query.(type) {
	case *panelQuery.Prometheus:
		jq.Prometheus = v
	case *panelQuery.Graphite:
		jq.Graphite = v
	}

	return json.Marshal(jq)
}

// makeRefID returns symbolic ID for given index.
// TODO: It has very rough implementation. Needs refactoring.
func makeRefID(index int) string {
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	var id string
	if index >= len(letters) {
		id += makeRefID(index % len(letters))
	} else {
		id = string(letters[index])
	}

	var result string
	for _, v := range id {
		result = string(v) + result
	}
	return result
}
