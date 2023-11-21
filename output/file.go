package output

import (
	. "design/interfaces"
	"os"
)

//var path string

type File struct {
	//entry fs.DirEntry
	//entry fs.FileInfo
	path string
	name string
}

func (file *File) Name() string {
	//fileInfo, err := os.Stat(file.path)
	//if err != nil {
	//	return ""
	//}
	//return fileInfo.Name()
	//return file.entry.Name()
	return file.name
}

func (file *File) GetChildren() []TreeOut {
	fileInfo, err := os.Stat(file.path)
	if err != nil {
		return nil
	}
	if !fileInfo.IsDir() {
		return nil
	}
	//if !file.entry.IsDir() {
	//	return nil
	//}
	//dirEntries, err := os.ReadDir(filepath.Join(file.entry., file.Name()))
	//filePath := filepath.Join(file.path, file.Name())
	filePath := file.path
	dirEntries, err := os.ReadDir(filePath)
	if err != nil {
		return nil
	}
	children := make([]TreeOut, len(dirEntries))
	for i, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			return nil
		}
		child := &File{path: filePath + "/" + info.Name(), name: info.Name()}
		children[i] = child
	}
	return children
}

//if path == "./" {
//	path, err = os.Getwd()
//	if err != nil {
//		return err
//	}
//}
//fileInfo, err := os.Stat(path)
//if err != nil {
//	return err
//}

//dirEntry := fs.FileInfoToDirEntry(fileInfo)

//dirEntry, ok := fileInfo.(fs.DirEntry)
//if !ok {
//	return errors.New("not a directory")
//}

//entry, err := os.ReadDir(path)
//if err != nil {
//	return err
//}
