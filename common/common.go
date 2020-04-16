package common

type BaseTheme struct{}

const Version = "v0.0.31"

func (BaseTheme) GetVersion() string {
	return Version
}

func (BaseTheme) GetRequirements() []string {
	return []string{"v1.2.9"}
}
