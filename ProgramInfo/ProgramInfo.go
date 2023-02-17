package pi

import (
	"errors"
	"path"
	"runtime/debug"
	"strings"
)

const (
	DependenciesTextPrefix = "Dependencies: "
)

// ProgramInfo is information about program and its dependencies.
type ProgramInfo struct {
	name             string
	version          string
	dependencies     []*Dependency
	dependenciesText string
}

func NewProgramInfo() (info *ProgramInfo, err error) {
	info = &ProgramInfo{
		name:         NameUnknown,
		version:      VersionUnknown,
		dependencies: make([]*Dependency, 0),
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return info, errors.New(ErrBuildInfoIsNotAvailable)
	}

	var tmp string
	tmp = path.Base(bi.Main.Path)
	if len(tmp) == 0 {
		info.name = NameUnknown
	} else {
		info.name = tmp
	}

	tmp = path.Base(bi.Main.Version)
	if len(tmp) == 0 {
		info.version = VersionUnknown
	} else {
		info.version = tmp
	}

	var dep *Dependency
	for _, d := range bi.Deps {
		dep, err = NewDependency(d)
		if err != nil {
			return nil, err
		}

		info.dependencies = append(info.dependencies, dep)
	}

	info.initDependenciesText()

	return info, nil
}

func (pi *ProgramInfo) initDependenciesText() {
	var sb = new(strings.Builder)
	sb.WriteString(DependenciesTextPrefix)

	for _, dep := range pi.dependencies {
		sb.WriteString("[" + dep.name + " " + dep.version + "] ")
	}

	pi.dependenciesText = sb.String()
}

func (pi *ProgramInfo) ProgramName() (name string) {
	return pi.name
}

func (pi *ProgramInfo) ProgramVersion() (version string) {
	return pi.version
}

func (pi *ProgramInfo) DependenciesList() (list []*Dependency) {
	return pi.dependencies
}

func (pi *ProgramInfo) DependenciesText() (txt string) {
	return pi.dependenciesText
}
