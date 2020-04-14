package items

import "fmt"

func createLabelBadges(labels map[string]string) string {
	badges := ""
	for key, value := range labels {
		badges = badges + fmt.Sprintf("<span class=\"badge badge-info\">%s=%s</span> ", key, value)
	}
	return badges
}

func createInfo(kind string, name string) string {
	return fmt.Sprintf("<b>%s: %s</b> ", kind, name)
}

func createStatusButton(status, content string) string {
	return fmt.Sprintf(`<button type="button" class="btn btn-sm btn-%s float-right">%s</button> `, status, content)
}

func createBadge(color string, content string) string {
	return fmt.Sprintf("<span class=\"badge badge-%s\">%s</span> ", color, content)
}
