package botinfoUsecases

import "fmt"

type IBotinfoUsecase interface {
	Feature(msg string) string
}

type botinfoUsecase struct{}

func NewBotinfoUsecase() IBotinfoUsecase {
	return &botinfoUsecase{}
}

func (u *botinfoUsecase) Feature(msg string) string {
	return fmt.Sprintf("```Just an example bot: %v```", msg)
}
