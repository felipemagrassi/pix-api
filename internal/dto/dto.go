package dto

type CreateReceiverInput struct {
	Name        string
	Document    string
	Email       string
	PixKeyValue string
	PixKeyType  string
}

type CreateReceiverOutput struct{}

type EditReceiverInput struct {
	Id string
}

type EditReceiverOutput struct{}
