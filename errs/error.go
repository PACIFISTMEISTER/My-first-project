package errs

import (
	"log"
	"strconv"
	"strings"
)

func Printer(err error, position string) {
	log.Println("error happend in ", position, " error text is ", err.Error())
}

func CutID(idStr string, prefix string) int {

	idStr = strings.TrimSuffix(idStr, prefix)
	idStr = idStr[strings.LastIndexAny(idStr, "/"):]
	idStr = strings.Replace(idStr, "/", "", 1)
	log.Println("id str after", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {

		Printer(err, "heppend in CutID"+prefix)
	}

	return id
}
