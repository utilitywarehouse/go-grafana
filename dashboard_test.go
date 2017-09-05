package grafana

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDashboard_Tags(t *testing.T) {
	d := Dashboard{tags: []string{"tag1", "tag2"}}
	expected := []string{"tag1", "tag2"}
	got := d.Tags()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Dashboard{tags: %v}.Tags(): expected %v, got %v", d.tags, expected, got)
	}
}

func TestDashboard_SetTags(t *testing.T) {
	ts := []struct {
		initial   []string
		tagsToSet []string
		expected  []string
	}{
		{[]string{"tag1"}, []string{"tag2", "tag3"}, []string{"tag2", "tag3"}},
		{[]string{"tag1"}, []string{}, []string{}},
	}

	for _, tt := range ts {
		d := Dashboard{tags: tt.initial}
		d.SetTags(tt.tagsToSet...)
		if !reflect.DeepEqual(d.tags, tt.expected) {
			t.Errorf("Dashboard{tags: %v}.SetTags(%v): expected %v, got %v", tt.initial, tt.tagsToSet, tt.expected, d.tags)
		}
	}
}
func TestDashboard_AddTags(t *testing.T) {
	ts := []struct {
		initial   []string
		tagsToAdd []string
		expected  []string
	}{
		{[]string{"tag1"}, []string{"tag2", "tag3"}, []string{"tag1", "tag2", "tag3"}},
		{[]string{"tag1"}, []string{"tag1", "tag2"}, []string{"tag1", "tag2"}},
	}

	for _, tt := range ts {
		d := Dashboard{tags: tt.initial}
		d.AddTags(tt.tagsToAdd...)
		if !reflect.DeepEqual(d.tags, tt.expected) {
			t.Errorf("Dashboard{tags: %v}.AddTags(%v): expected %v, got %v", tt.initial, tt.tagsToAdd, tt.expected, d.tags)
		}
	}
}

func TestDashboard_RemoveTags(t *testing.T) {
	ts := []struct {
		initial      []string
		tagsToRemove []string
		expected     []string
	}{
		{[]string{"tag1"}, []string{"tag1"}, []string{}},
		{[]string{"tag1", "tag2"}, []string{"tag2"}, []string{"tag1"}},
		{[]string{"tag1", "tag2"}, []string{"tag3"}, []string{"tag1", "tag2"}},
	}

	for _, tt := range ts {
		d := Dashboard{tags: tt.initial}
		d.RemoveTags(tt.tagsToRemove...)
		if !reflect.DeepEqual(d.tags, tt.expected) {
			t.Errorf("Dashboard{tags: %v}.RemoveTags(%v): expected %v, got %v", tt.initial, tt.tagsToRemove, tt.expected, d.tags)
		}
	}
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
