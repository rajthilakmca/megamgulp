package bind

import "fmt"

// EnvVar represents a environment variable for an app.
type EnvVar struct {
	Name         string
	Value        string
	Public       bool
	InstanceName string
}

func (e *EnvVar) String() string {
	var value, suffix string
	if e.Public {
		value = e.Value
	} else {
		value = "***"
		suffix = " (private variable)"
	}
	return fmt.Sprintf("%s%s", value, suffix)
}

type Unit interface {
	// GetIp returns the unit ip.
	GetIp() string
}

type App interface {
	// GetIp returns the app ip.
	GetIp() string

	// GetName returns the app name.
	GetName() string

	// GetUnits returns the app units.
	GetUnits() []Unit

	// InstanceEnv returns the app enviroment variables.
	InstanceEnv(string) map[string]EnvVar

	// SetEnvs adds enviroment variables in the app.
	SetEnvs([]EnvVar, bool) error

	// UnsetEnvs removes the given enviroment variables from the app.
	UnsetEnvs([]string, bool) error
}

type Binder interface {
	// BindApp makes the bind between the binder and an app.
	BindApp(App) error

	// BindUnit makes the bind between the binder and an unit.
	BindUnit(App, Unit) (map[string]string, error)

	// UnbindApp makes the unbind between the binder and an app.
	UnbindApp(App) error

	// UnbindUnit makes the unbind between the binder and an unit.
	UnbindUnit(Unit) error
}


