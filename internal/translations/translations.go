package translations

import (
	_ "embed"
	"encoding/json"
	"github.com/yoonaowo/discord_verifier/internal/utils"
	"sync"
)

// We cannot access subdirectories w/o using embed.FS
//
//go:embed translations.json
var jsonData []byte

var (
	langs map[string]string
	once  sync.Once
)

func onceFunc() {
	err := json.Unmarshal(jsonData, &langs)

	if err != nil {
		utils.Logger().Fatalln("Failed to setup translations")
	}

}

func Get(text string) string {
	once.Do(onceFunc)

	res, found := langs[text]

	if !found {
		return text
	}

	return res
}
