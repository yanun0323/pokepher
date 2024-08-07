package usecase

type Stage interface {
	View

	RestoreGameTime() int
}
