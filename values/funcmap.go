package values

import "html/template"

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"valYesNo": yesno,
		"valOnOff": onoff,
	}
}

// yesno returns "Yes" or "No" based on the boolean value
func yesno(b bool) string {
	if b {
		return "Yes"
	}

	return "No"
}

// onoff returns "On" or "Off" based on the boolean value
func onoff(b bool) string {
	if b {
		return "On"
	}

	return "Off"
}
