package common

type BaseTheme struct{}

const Version = "0.0.30"

func (BaseTheme) GetVersion() string {
	return Version
}

func (BaseTheme) GetRequirements() []string {
	return []string{"v1.2.8"}
}
