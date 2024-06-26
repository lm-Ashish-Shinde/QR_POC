For generating the QR code, packages in golang inclued, skip2/go-qr, yeqown/go-qrcode

1. yeqown/go-qrcode:-  This packages doesnt have many forks and stars in github,
                    This package is well maintained and have new features


2. skip2/go-qr :- This package is old, and does have any latest updates.
                The last update for this was 6 years ago.
                This is commonly used package



3. boombuler/barcode:- this package has 1.6k stars in github, 
                    last updated 2 years ago, 
                    has many other formats like, 2 of 5,Aztec Code,Codabar,Code 128,
                    Code 39,Code 93,Datamatrix,EAN 13,EAN 8,PDF 417, and QR code


so here we are using go-qrcode.v2 which is updagrade to go-qrcode

go-qrcode.v2 :- is a major upgrade from v1, and it is not backward compatible. 
v2 redesigned the API, and it is more flexible and powerful. 
Features are split into different modules (according to functionality).



error correction:- This allows the code to still be readable even if it's slightly damaged or obscured. 
                    Higher error correction levels can increase the size of the QR code slightly.

                    It has 4 levels of the correction low, medium, high, highest

disadvantages of error correction:-
                1. higher correction level has low readibility.
                2. big size
                3. more complex
                4. choose as per requirement.


As per discussed:-
    size can be set from frontend
    we have to generate the qr code of decent size and quality which can be scaled if required.
