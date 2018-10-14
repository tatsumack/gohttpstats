package parsers

type Parser interface {
	Read() (map[string]string, error)
}
