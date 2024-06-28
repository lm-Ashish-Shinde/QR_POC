package main

import (
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

// func WithLogoImageFilePNG(f string) ImageOption {

// }

func main() {
	path := "https://play.google.com/store/apps/details?id=com.whatsapp&pcampaignid=web_share"

	// qrc, err := qrcode.New(path) // generate qr code

	qrc, err := qrcode.NewWith(path, qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest)) // generate qr code with correction level

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// // save file
	// if err := qr.Save("./repo-qrcode.jpeg"); err != nil {
	// 	fmt.Printf("could not save image: %v", err)
	// }

	options := []standard.ImageOption{
		// as int value provided here is in pixels.
		// standard.WithHalftone("./a-img (1).png"), // used to create Halftone qr code
		standard.WithQRWidth(21), // size here is in pixels, size will be handeled by frontend team. we give decent quality
		// standard.WithBgTransparent(),	// used to create qr code with transparent background
		// standard.WithBgColorRGBHex("#95A5A6"),        // change colour of qr background
		// standard.WithFgColorRGBHex("#AF7AC5"),        // change colour of qr dots or square
		// standard.WithBorderWidth(20),	// used to give border width, size is in pixels
		standard.WithLogoImageFilePNG("./a-img.png"), // used to add png format logo inside qr, here the size of image is - width-90pixels,Height-80pixels
		// standard.WithLogoImage() // used to add logo inside the qr
	}
	filename := "./qr_code_high-21.jpeg" // save file with name

	w0, err := standard.New(filename, options...)
	if err != nil {
		fmt.Printf("error %v:", err)
	}

	err = qrc.Save(w0)
	if err != nil {
		fmt.Printf("error %v:", err)
	}

}
