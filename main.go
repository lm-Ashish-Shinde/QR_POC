package main

import (
	"fmt"
	// "bytes"
	// "image"
	// "image/draw"
	// "image/png"
	// "log"
	// "os"

	qrcode "github.com/yeqown/go-qrcode"
)

// func WithLogoImageFilePNG(f string) ImageOption {

// }

func main() {
	qrc, err := qrcode.New("https://play.google.com/store/apps/details?id=com.whatsapp&pcampaignid=web_share")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err := qrc.Save("./repo-qrcode.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}

}
