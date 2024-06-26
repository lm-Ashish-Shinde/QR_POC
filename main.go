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
	path := "https://play.google.com/store/apps/details?id=com.whatsapp&pcampaignid=web_share"
	// crlevel := qrcode.ErrorCorrectionHighest
	// qrc, err := qrcode.New(path)
	// qrc, err := qrcode.NewWith(path, qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionMedium))
	qrc, err := qrcode.NewWith(path, qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest))

	// qr := qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// // save file
	// if err := qr.Save("./repo-qrcode.jpeg"); err != nil {
	// 	fmt.Printf("could not save image: %v", err)
	// }

	options := []standard.ImageOption{
		// as int value provided here is in pixels.
		// standard.WithHalftone("./a-img (1).png"),
		// standard.WithQRWidth(250),
		standard.WithQRWidth(21), // size here is in pixels, size will be handeled by fronend team. we give decent quality
		// standard.WithBgTransparent(),
		// standard.WithBgColorRGBHex("#95A5A6"),        //change colour of qr background
		// standard.WithFgColorRGBHex("#AF7AC5"),        // change colour of qr dots or square
		standard.WithBorderWidth(20),
		standard.WithLogoImageFilePNG("./a-img.png"), // here the size of image is - width-90pixels,Height-80pixels
		// standard.WithLogoImage()
	}
	filename := "./qr_code_high-width10.png"

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
