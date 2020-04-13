package items

import "fmt"

func getLabelBadges(labels map[string]string) string {
	badges := ""
	for key, value := range labels {
		badges = badges + fmt.Sprintf("<span class=\"badge badge-info\">%s=%s</span>", key, value)
	}
	return badges
}
