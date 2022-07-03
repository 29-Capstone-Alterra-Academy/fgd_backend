package format

import (
	"fgd/app/config"
	"fmt"
)

func FormatImageLink(conf config.Config, fileNames ...*string) {
	for _, fileName := range fileNames {
		if fileName != nil {
			temp := *fileName
			*fileName = fmt.Sprintf("%s:%s/public/asset/%s", conf.HOST, conf.PORT, temp)
		}
	}
}
