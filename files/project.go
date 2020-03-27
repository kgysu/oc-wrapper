package files

type OpenshiftProject struct {
	Name       string
	ConfigFile string
	Items      []OpenshiftItem
}

type OpenshiftItem struct {
	Name string
	File string
	Data string
}

func (op OpenshiftProject) loadItemsData() {

}
