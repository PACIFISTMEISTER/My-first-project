package Host

import (
	"ShopProject/errs"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func HandleImage(w http.ResponseWriter, r *http.Request) {

	pic := r.URL.String()
	pic = pic[strings.LastIndexAny(pic, "/"):]
	pic = strings.Replace(pic, "/", "", 1)
	ext := pic[strings.LastIndexAny(pic, "."):]
	if ext == ".jpeg" {
		existingImageFile, err := os.Open("./images/" + pic)
		if err != nil {
			errs.Printer(err, "HandleImage1")
		}
		defer existingImageFile.Close()
		w.Header().Set("Content-Type", "image/jpeg")
		loadedImage, err := jpeg.Decode(existingImageFile)
		if err != nil {
			errs.Printer(err, "HandleImage2")
		}
		err = png.Encode(w, loadedImage)
		if err != nil {
			errs.Printer(err, "HandleImage3")
		}
	} else {
		if ext == ".png" {
			existingImageFile, err := os.Open("./images/" + pic)
			if err != nil {
				errs.Printer(err, "HandleImage4")
			}
			defer existingImageFile.Close()
			w.Header().Set("Content-Type", "image/png")
			loadedImage, err := png.Decode(existingImageFile)
			if err != nil {
				errs.Printer(err, "HandleImage2")
			}
			err = png.Encode(w, loadedImage)
			if err != nil {
				errs.Printer(err, "HandleImage3")
			}

		}
	}
}
