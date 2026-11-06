package civilizationstate_test

import (
	"reflect"
	"testing"

	civilizationstate "civilizationstate"
)

func TestStateComposition(t *testing.T) {
	state := civilizationstate.State{
		Internalstate: civilizationstate.Internalstate{},
		Externalstate: civilizationstate.Externalstate{},
	}

	if state.Internalstate != (civilizationstate.Internalstate{}) {
		t.Error("State should retain its Internalstate")
	}
	if state.Externalstate != (civilizationstate.Externalstate{}) {
		t.Error("State should retain its Externalstate")
	}
}

func TestCivilizationstateComposition(t *testing.T) {
	civilization := civilizationstate.Civilizationstate{
		Peace: civilizationstate.Peace{},
		GaU: civilizationstate.GaU{
			Soch:     civilizationstate.Soch{},
			Kartabya: civilizationstate.Kartabya{},
			Karma:    civilizationstate.Karma{},
		},
		Triasta: civilizationstate.Triasta{},
	}

	if civilization.Peace != (civilizationstate.Peace{}) {
		t.Error("Civilizationstate should retain its Peace value")
	}
	if civilization.GaU.Soch != (civilizationstate.Soch{}) {
		t.Error("GaU should retain its Soch value")
	}
	if civilization.GaU.Kartabya != (civilizationstate.Kartabya{}) {
		t.Error("GaU should retain its Kartabya value")
	}
	if civilization.GaU.Karma != (civilizationstate.Karma{}) {
		t.Error("GaU should retain its Karma value")
	}
	if civilization.Triasta != (civilizationstate.Triasta{}) {
		t.Error("Civilizationstate should retain its Triasta value")
	}
}

func TestPlaceholderZeroValues(t *testing.T) {
	placeholders := []struct {
		name  string
		value any
	}{
		{name: "Others", value: civilizationstate.Others{}},
		{name: "Internalstate", value: civilizationstate.Internalstate{}},
		{name: "Externalstate", value: civilizationstate.Externalstate{}},
		{name: "Soch", value: civilizationstate.Soch{}},
		{name: "Kartabya", value: civilizationstate.Kartabya{}},
		{name: "Karma", value: civilizationstate.Karma{}},
		{name: "Peace", value: civilizationstate.Peace{}},
		{name: "Triasta", value: civilizationstate.Triasta{}},
	}

	for _, placeholder := range placeholders {
		t.Run(placeholder.name, func(t *testing.T) {
			if !reflect.ValueOf(placeholder.value).IsZero() {
				t.Errorf("%s zero value = %#v, want zero value", placeholder.name, placeholder.value)
			}
		})
	}
}
