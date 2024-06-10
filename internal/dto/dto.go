package dto

type CreateReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

type UpdateDraftedReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}
