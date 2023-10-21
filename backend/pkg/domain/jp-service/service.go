package jpservice

type LanguageProcessor interface {
	GetUnique(characters string) []rune
}

type JpProcessor struct {
}

func (jpprocessor *JpProcessor) GetUnique(characters string) []rune {
	panic("implement me") // TODO: Implement
}
