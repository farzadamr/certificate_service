package qr

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func Generate(certID int) (string, error) {
	url := fmt.Sprintf("http://localhost:8080/verify/%s", certID)
	path := fmt.Sprintf("tmp_qr_%s.png", certID)

	err := qrcode.WriteFile(url, qrcode.Medium, 256, path)
	return path, err
}
