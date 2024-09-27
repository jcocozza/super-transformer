package formattransformers

import "path/filepath"

func GetFileType(path string) string {
    return filepath.Ext(path)
}
