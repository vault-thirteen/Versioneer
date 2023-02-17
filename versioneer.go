package ver

import (
	"fmt"
	"log"
	"runtime"

	"github.com/kr/pretty"
	pi "github.com/vault-thirteen/Versioneer/ProgramInfo"
)

const (
	IntroTextShort = "%s, ver. %s. Go language: %s."
	IntroTextFull  = "%s %s, ver. %s. Go language: %s."
)

// Versioneer is an extended version of the ProgramInfo class.
type Versioneer struct {
	programInfo *pi.ProgramInfo
}

func New() (v *Versioneer, err error) {
	v = new(Versioneer)

	v.programInfo, err = pi.NewProgramInfo()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ShowIntroText shows introductory text about the program.
// The 'product' parameter is optional. If it is set, it is printed after the
// program name. It is used mostly for showing server and client variants of a
// product.
func (v *Versioneer) ShowIntroText(product string) {
	if len(product) == 0 {
		fmt.Println(
			fmt.Sprintf(IntroTextShort,
				v.programInfo.ProgramName(),
				v.programInfo.ProgramVersion(),
				runtime.Version(),
			),
		)
	} else {
		fmt.Println(
			fmt.Sprintf(IntroTextFull,
				v.programInfo.ProgramName(),
				product,
				v.programInfo.ProgramVersion(),
				runtime.Version(),
			),
		)
	}
}

func (v *Versioneer) ShowComponentsInfoText() {
	fmt.Println(v.programInfo.DependenciesText())
}

func (v *Versioneer) ShowComponentsInfoList() {
	_, err := pretty.Println(v.programInfo.DependenciesList())
	if err != nil {
		log.Println(err)
	}
}

func (v *Versioneer) ProgramName() (programName string) {
	return v.programInfo.ProgramName()
}

func (v *Versioneer) ProgramVersion() (programVersion string) {
	return v.programInfo.ProgramVersion()
}

func (v *Versioneer) DependenciesList() (list []*pi.Dependency) {
	return v.programInfo.DependenciesList()
}

func (v *Versioneer) DependenciesText() (txt string) {
	return v.programInfo.DependenciesText()
}
