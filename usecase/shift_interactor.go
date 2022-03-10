package usecase

type ShiftRepository interface {

}

type ShiftInteractor struct {
    ShiftRepository ShiftRepository
}