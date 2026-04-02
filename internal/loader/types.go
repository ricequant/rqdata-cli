package loader

type Command struct {
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	APIMethod         string                 `json:"api_method"`
	Validator         string                 `json:"validator"`
	InjectDefaults    map[string]interface{} `json:"inject_defaults"`
	APISupportsFields bool                   `json:"api_supports_fields"`
}

type Group struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Commands    []Command `json:"commands"`
	Groups      []Group   `json:"groups"`
}

type Config struct {
	Groups []Group `json:"groups"`
}

var (
	config Config
	schema map[string]interface{}
)
