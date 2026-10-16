package persistence

import "embed"

//go:embed repository/**/sql/*.sql
var SQLFiles embed.FS

func LoadFile(path string) string {
	b, err := SQLFiles.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
