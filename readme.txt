POC for QR code Generation in GO.


Packages for QR code generation in golang: 

1. yeqown/go-qrcode:- 
        i.  This packages doesnt have many forks and stars in github,
        ii. This package is well maintained and have new features
        iii.Latest Version of this package is is major updagrade and is not backward compatible.
        iv. New Version redesigned the API, and it is more flexible and powerful. 
        v. Features are split into different modules (according to functionality).


2. skip2/go-qr :- 
        i.   This package is old, and does have any latest updates.
        ii.  The last update for this was 6 years ago.
        iii. This is commonly used package 


3. boombuler/barcode:-
        i. This package has 1.6k stars in github, 
        ii.Last updated 2 years ago, 
        iii.Has many other formats like, 2 of 5,Aztec Code,Codabar,Code 128,
            Code 39,Code 93,Datamatrix,EAN 13,EAN 8,PDF 417, and QR code
        iv. This could be a good alternative.



Features of yeqown/go-qrcode:- 
    1. Specifying cell shape allowably with WithCustomShape, WithCircleShape (default is rectangle).
    2. Specifying output file's format with WithBuiltinImageEncoder, WithCustomImageEncoder (default is JPEG).
    3. Not only shape of cell, but also color of QR Code background and foreground color.
    4. WithLogoImage, WithLogoImageFilePNG, WithLogoImageFileJPEG help you add an icon at the central of QR Code.
    5. WithBorderWidth allows to specify any width of 4 sides around the qrcode.
    6. WebAssembly support.
    7. support Halftone QR Codes.


Viewing Distance for QR code :- 
        1.The size of qr code, its correction level, device camera quality affects the scaning Distance of qr code.
        2.The general rule of thumb is a 10:1 distance-to-size ratio, meaning that the QR code size should be 
           1/10th of the distance from which it will be scanned.
        3.For example, if you anticipate users scanning a QR code from a distance of 1 meter (approximately 39 inches), 
          a QR code size of around 10-15 cm (4-6 inches) in physical dimensions or around 400-600 pixels in each dimension 
          could be appropriate. This provides enough detail and clarity for reliable scanning without requiring users to get too close to the QR code.
        

    Factor affecting the readibility of qr code:
        i. QR code Size:- Larger qr code is more readable
        ii. QR code resolution and module size:-QR codes are composed of modules (black and white squares). 
        iii. Higher resolution and larger module sizes improve readability.
        iv. Scaning device quality.

    Practical Considerations:
        i. QR code should have sufficient contrast between the code and its background, opt for high contrast colors incase using colourful qr.
        ii.Test qr code before deploying under the similar conditions we are using them.
        iii. Add sufficient correction level.


Error correction:- 
        1. This allows the code to still be readable even if it's slightly damaged or obscured or even blured. 
        2. It has 4 levels of the correction Low (L), Medium (M), Quartile (Q), and High (H).
        3. The higher the error correction level, the more data redundancy is added to the QR code, making it larger but more resilient to errors.

    disadvantages of error correction:-
        1. QR codes with higher error correction levels may take slightly longer to scan compared to codes with lower error correction levels.
        2. Depending too heavily on error correction can lead to complacency in QR code production and placement.
        3. While error correction aims to recover data accurately, in some cases, a severely damaged QR code may still be misinterpreted due to limitations
           in error correction algorithms or scanning devices.
        4. Different QR code readers and software implementations may interpret error correction levels differently or have varying levels of support for 
           high-error correction codes.


Halftone Qr code :-
    1. Halftone QR codes are a specialized type of QR code that uses halftone patterns instead of solid black and white squares (modules) to encode information.
    2. This technique allows QR codes to be visually integrated into complex or colorful backgrounds without sacrificing scanning reliability.
    3. The image should be provide to create a halftone patten.


Logo/Image Inside the Qr code:-
    1. Image or logo can be added in the qr code.
    2. A common rule of thumb is to limit the logo size to be ~25% of the width and height of the QR code to ensure scannability.
    3. The logo size can be up to 1/3 of the edge length of the QR code

Color of Qr:
    1. The colour of QR code and its background can be changed.
    2. The both colours should have enough contrast between them to ensure better readability.
    3. The colours can be used to make it unique, attractive and represent brand identity.
    4. The default black and white colours gives the best contrast and has highest readability.

Capacity of QR code:-
    1.The maximum capacity of a QR Code varies according to the content encoded and the error recovery level. 
    2.The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.


Points to remember as discussed with team:-
    1. The use of yeqown/go-qrcode/v2 package is considered after discussion and comparision with other packages
    2. Size can be changed from frontend team Depending on usecase.
    3. We have to generate the QR code of decent size and quality which can be scaled if required.
    4. Should consider decent size for qr code that will be readabile from distance (1:10 ratio  bet qr code size and scanning distance).
    5. The Halftone image cannot be used in the qr code as the logo size is larger.
    6. The small logo can be placed inside the qr code.
    7. The size of logo image which is placed inside the qr code should be reduced according to the size of qr code.
    8. Colour of qr code should be kept default(black and white) as it provides high contrast ratio which is more readable.
    9. Error correction for qr should be implemented.    



