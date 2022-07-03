package format

import (
	"fgd/app/config"
	"fmt"
)

func FormatImageLink(conf config.Config, fileNames ...*string) {
	for _, fileName := range fileNames {
		if fileName != nil && *fileName != "" {
			temp := *fileName
			*fileName = fmt.Sprintf("%s/%s%s", conf.DOMAIN, conf.STATIC_PATH, temp)
		}
	}
}
