//This is focused on Experimental Framework for Quantum Civilization State space using Classical Civilization State space.
//This is a simplified model for the purpose of demonstrating the structure of a civilization state.
// All remaining work is kept for future development and expansion of the model.

package civilizationstate

// this Others can be inserted into any struct to allow for future expansion of the model without changing the existing structure.
// Others is also a placeholder for development.
// Instead of Others there can be any other operational or structural elements.
type Others struct {
	others string
}

type State struct {
	Internalstate
	Externalstate
}
type Internalstate struct {
	others Others
}
type Externalstate struct {
	others Others
}

type Civilizationstate struct {
	Peace   Peace
	GaU     GaU
	Triasta Triasta
}
type Soch struct {
	others Others
}
type Kartabya struct {
	others Others
}
type Karma struct {
	others Others
}

type GaU struct {
	Soch     Soch
	Kartabya Kartabya
	Karma    Karma
}
type Peace struct {
	others Others
}
type Triasta struct {
	others Others
}
