package app

import (
	"image/png"
	"net/http"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	log "github.com/sirupsen/logrus"
)

func GeneratorHandler(w http.ResponseWriter, r *http.Request) {
	qrText := r.FormValue("text")
	qrSizeStr := r.FormValue("size")
	log.Info("Request", " text "+qrSizeStr, " size "+qrText)

	var qrSize int64
	if qrSizeStr == "" {
		qrSize = 200
	} else {
		v, err := strconv.ParseInt(qrSizeStr, 0, 16)
		qrSize = v
		if err != nil {
			log.Error(err)
			qrSize = 200
		}
	}

	qrCode, err := qr.Encode(qrText, qr.L, qr.Auto)
	if err != nil {
		log.Error(err)
	}

	qrCode, err = barcode.Scale(qrCode, int(qrSize), int(qrSize))
	if err != nil {

	}
	png.Encode(w, qrCode)
}
