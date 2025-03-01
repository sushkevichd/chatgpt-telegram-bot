package domain

type Response struct {
	Text  string // Text response from the AI (if applicable)
	Image *Image // Image data generated by the AI (if applicable)
}

type Image struct {
	ID   int64
	Data []byte
}
