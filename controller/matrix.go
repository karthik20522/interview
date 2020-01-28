package controller

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Echo godoc
// @Summary Echo File Data
// @ID file.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "1,2,3\n4,5,6\n7,8,9"
// @Failure 500 {string} string "Internal Server Error"
// @Router /echo [post]
func Echo(ctx *gin.Context) {
	records, err := getCSVData(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	ctx.String(http.StatusOK, fmt.Sprintf(response))
}

func getCSVData(c *gin.Context) ([][]string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return [][]string{}, err
	}

	src, _ := file.Open()
	defer src.Close()

	records, err := csv.NewReader(src).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
