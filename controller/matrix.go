package controller

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	MatrixOps "../operations"
	Utils "../utils"
	"github.com/gin-gonic/gin"
)

// Echo godoc
// @Summary Echo
// @Description Return the matrix as a string in matrix format.
// @ID echo.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "1,2,3<br />4,5,6<br />7,8,9"
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

// InvertMatrix godoc
// @Summary Invert
// @Description Return the matrix as a string in matrix format where the columns and rows are inverted
// @ID invert.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "1,2,3<br />4,5,6<br />7,8,9"
// @Failure 500 {string} string "Internal Server Error"
// @Router /invert [post]
func InvertMatrix(ctx *gin.Context) {
	records, err := getCSVData(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
		return
	}

	matrixArray := Utils.ConvertCSVToMatrix(records)
	ctx.String(http.StatusOK, fmt.Sprintf(Utils.Matrix2String(MatrixOps.Invert(matrixArray))))
}

// FlattenMatrix godoc
// @Summary Flatten
// @Description Return the matrix as a 1 line string, with values separated by commas.
// @ID flatten.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "1,2,3,4,5,6,7,8,9"
// @Failure 500 {string} string "Internal Server Error"
// @Router /flatten [post]
func FlattenMatrix(ctx *gin.Context) {
	records, err := getCSVData(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf(MatrixOps.Flatten(records)))
}

// SumOfMatrix godoc
// @Summary Sum
// @Description Return the sum of the integers in the matrix
// @ID sum.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "45"
// @Failure 500 {string} string "Internal Server Error"
// @Router /sum [post]
func SumOfMatrix(ctx *gin.Context) {
	records, err := getCSVData(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
		return
	}
	matrixArray := Utils.ConvertCSVToMatrix(records)
	matrixSum := MatrixOps.Sum(matrixArray)

	ctx.String(http.StatusOK, fmt.Sprintf(strconv.Itoa(matrixSum)))
}

// MultiplyMatrix godoc
// @Summary Multiply
// @Description Return the product of the integers in the matrix
// @ID multiply.upload
// @Accept  multipart/form-data
// @Param   file formData file true  "Comma Seperated Matrix data"
// @Success 200 {string} string "362880"
// @Failure 500 {string} string "Internal Server Error"
// @Router /multiply [post]
func MultiplyMatrix(ctx *gin.Context) {
	records, err := getCSVData(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
		return
	}
	matrixArray := Utils.ConvertCSVToMatrix(records)
	matrixMultiply := MatrixOps.Multiply(matrixArray)

	ctx.String(http.StatusOK, fmt.Sprintf(strconv.Itoa(matrixMultiply)))
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
