package configs

import _ "embed"

//go:embed commands.json
var EmbeddedCommands []byte

//go:embed schema.json
var EmbeddedSchema []byte
