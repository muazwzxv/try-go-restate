package dto

type CreateUserDto struct {
	UUID      string
	Name      string
	Email     string
	Status    string
	CreatedBy string
	UpdatedBy string
}
