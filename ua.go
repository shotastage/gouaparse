package gouaparse

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
