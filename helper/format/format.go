package format

import (
	"fgd/app/config"
	"fmt"
	"strings"
)

func FormatImageLink(conf config.Config, fileNames ...*string) {
	for _, fileName := range fileNames {
		if fileName != nil && *fileName != "" && !strings.Contains(*fileName, conf.DOMAIN) {
			temp := *fileName
			*fileName = fmt.Sprintf("%s/%s%s", conf.DOMAIN, conf.STATIC_PATH, temp)
		}
	}
}
