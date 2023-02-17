package pi

import (
	"errors"
	"path"
	"runtime/debug"
)

// Dependency is information about program's dependency.
type Dependency struct {
	name    string
	version string
}

func NewDependency(m *debug.Module) (dep *Dependency, err error) {
	dep = &Dependency{
		name:    NameUnknown,
		version: VersionUnknown,
	}

	if m == nil {
		return dep, errors.New(ErrDependencyInfoIsNotAvailable)
	}

	var tmp string
	tmp = path.Base(m.Path)
	if len(tmp) == 0 {
		dep.name = NameUnknown
	} else {
		dep.name = tmp
	}

	tmp = path.Base(m.Version)
	if len(tmp) == 0 {
		dep.version = VersionUnknown
	} else {
		dep.version = tmp
	}

	return dep, nil
}
