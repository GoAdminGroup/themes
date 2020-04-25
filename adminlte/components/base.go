package components

type Base struct {
}

func (b Base) GetAssetList() []string {
	return make([]string, 0)
}

func (b Base) GetAsset(string) ([]byte, error) {
	return nil, nil
}
