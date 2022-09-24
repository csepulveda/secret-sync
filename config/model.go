package config

type Config struct {
	Secrets []*Secret
}

type Secret struct {
	Provider  string `json:"provider"`
	Source    string `json:"source"`
	Dest      string `json:"dest"`
	Namespace string `json:"namespace"`
}
