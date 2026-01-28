package certificate

type Certificate struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Date    string `json:"date"`
	Hash    string `json:"hash"`
	PDFPath string `json:"pdf_path"`
}
