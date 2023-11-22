package output

func Ls(path string, f func(string) string) error {
	return recurOutputAsTree("", &File{path: path}, f)
}
