package main

import (
	"fmt"
	// "bytes"
	// "image"
	// "image/draw"
	// "image/png"
	// "log"
	// "os"

	// qrcode "github.com/yeqown/go-qrcode"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

// func WithLogoImageFilePNG(f string) ImageOption {

// }

func main() {
	qrc, err := qrcode.New("https://play.google.com/store/apps/details?id=com.whatsapp&pcampaignid=web_share")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// // save file
	// if err := qrc.Save("./repo-qrcode.jpeg"); err != nil {
	// 	fmt.Printf("could not save image: %v", err)
	// }

	options := []standard.ImageOption{
		// standard.WithHalftone("./a-img (1).png"),
		// standard.WithQRWidth(250),
		standard.WithQRWidth(10),

		standard.WithLogoImageFilePNG("./a-img.png"), // here the size of image is - width-90pixels,Height-80pixels
		// standard.WithLogoImage()
	}
	filename := "./qr_code.png"

	// if *transparent {
	// 	options = append(
	// 		options,
	// 		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	// 		standard.WithBgTransparent(),
	// 	)
	// 	filename = "./halftone-qr-transparent.png"
	// }

	w0, err := standard.New(filename, options...)
	if err != nil {
		fmt.Printf("error %v:", err)
	}
	// handleErr(err)

	err = qrc.Save(w0)
	if err != nil {
		fmt.Printf("error %v:", err)
	}
	// handleErr(err)

}
