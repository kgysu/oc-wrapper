package project

import (
	"fmt"
	"github.com/kgysu/oc-wrapper/config"
	"github.com/kgysu/oc-wrapper/util"
	"io"
)

func (op *OpenshiftProject) Save(w io.Writer) {
	err := util.SaveProjectToDisk(op)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		fmt.Fprintf(w, "saved project [%s] in [%s] \n", op.Name, config.GetRootFolderOrDefault())
	}
}
