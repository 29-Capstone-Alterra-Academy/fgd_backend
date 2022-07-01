package format

import (
	"fgd/app/config"
	"fmt"
)

func FormatImageLink(s *string, conf config.Config) {
	fileName := *s
	*s = fmt.Sprintf("%s:%s/public/asset/%s", conf.HOST, conf.PORT, fileName)
}
