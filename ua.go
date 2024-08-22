package gouaparse

import (
	"regexp"
	"strings"
)

type UserAgent struct {
	BrowserName            string
	BrowserVersion         string
	OperatingSystem        string
	OperatingSystemVersion string
	DeviceType             string
	DeviceModel            string
	RenderingEngine        string
	RenderingEngineVersion string
	IsBot                  bool
	FullString             string
}

func ParseUserAgent(uaString string) UserAgent {
	ua := UserAgent{FullString: uaString}

	// Browser and version
	browserRegex := regexp.MustCompile(`(Firefox|Chrome|Safari|Opera|MSIE|Trident)\/?\s*([\d\.]+)`)
	if matches := browserRegex.FindStringSubmatch(uaString); len(matches) > 2 {
		ua.BrowserName = matches[1]
		ua.BrowserVersion = matches[2]
	}

	// Operating System
	osRegex := regexp.MustCompile(`\((.*?)\)`)
	if matches := osRegex.FindStringSubmatch(uaString); len(matches) > 1 {
		osParts := strings.Split(matches[1], ";")
		if len(osParts) > 0 {
			ua.OperatingSystem = strings.TrimSpace(osParts[0])
		}
	}

	// Device Type
	if strings.Contains(strings.ToLower(uaString), "mobile") {
		ua.DeviceType = "Mobile"
	} else if strings.Contains(strings.ToLower(uaString), "tablet") {
		ua.DeviceType = "Tablet"
	} else {
		ua.DeviceType = "Desktop"
	}

	// Rendering Engine
	engineRegex := regexp.MustCompile(`(Gecko|WebKit|Presto|Trident)\/?\s*([\d\.]+)?`)
	if matches := engineRegex.FindStringSubmatch(uaString); len(matches) > 2 {
		ua.RenderingEngine = matches[1]
		ua.RenderingEngineVersion = matches[2]
	}

	// Is Bot
	ua.IsBot = strings.Contains(strings.ToLower(uaString), "bot") ||
		strings.Contains(strings.ToLower(uaString), "crawler") ||
		strings.Contains(strings.ToLower(uaString), "spider")

	return ua
}
