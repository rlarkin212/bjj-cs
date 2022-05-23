package spliturl

import (
	"errors"
	"fmt"
	"strings"
)

func SplitUrl(url string) (download, watch string, err error) {
	if url == "" {
		return "", "", errors.New("url empty")
	}

	split := strings.Split(url, ".mp4")

	download = fmt.Sprintf("%s%s", split[0], ".mp4?m=y")
	watch = fmt.Sprintf("%s%s", split[0], ".mp4")

	return download, watch, nil
}
