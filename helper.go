package sportsearch

import(
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)


func ReadJson(p string, d interface{}) {

	if len(p) == 0 || p == "" {
		return
	}

	f, err := os.ReadFile(filepath.Join(KEYWORDS_DIR, p))

	if err != nil {
		log.Println(err)
	} else {

		err := json.Unmarshal(f, &d)

		if err != nil {
			log.Println(err)
		}

	}

} // ReadJson
