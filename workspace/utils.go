package workspace

import "strings"

func GetFileName(raw string) string {
	//if strings.HasSuffix(raw, ".md") {
	//	return raw[:len(raw)-3]
	//}
	//return raw
	if !strings.HasSuffix(raw, ".md") {
		return raw + ".md"
	}
	return raw
}

func GetFilePath(fileName string) string {
	if !strings.HasPrefix(fileName, path) {
		return path + fileName
	} else {
		return fileName
	}
}
