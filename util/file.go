package util

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

func ParseFile(f *xlsx.File) ([][]map[string]string, [][][]string, error) {
	sliceData, err := f.ToSlice()
	if err != nil {
		return [][]map[string]string{}, [][][]string{}, err
	}

	mapData := [][]map[string]string{}

	for _, sheet := range sliceData {
		if len(sheet) == 0 {
			break
		}

		oneSheet := []map[string]string{} // one sheet result

		header := sheet[0] // row contain tilte each column
		for _, row := range sheet[1:] {
			oneRow := map[string]string{} // contain one row in sheet
			for idx, cell := range row {
				oneRow[header[idx]] = cell
			}

			oneSheet = append(oneSheet, oneRow) // append row into sheet
		}

		mapData = append(mapData, oneSheet) // append sheet into result
	}

	return mapData, sliceData, nil
}

type ExcelFile struct {
	Name   string
	Sheets []*Sheet
}

type Sheet struct {
	Name                 string
	Headers              [][]string
	HeaderAttributes     []*Attribute
	Content              [][]string
	CommonRowAttribute   *Attribute         // map[column]Attribute
	SpecificRowAttribute map[int]*Attribute // map[row_index]Attribute
}

type Attribute struct {
	Aligment *Aligment
	Border   *Border
	Fill     *Fill
	Font     *Font
}

type Aligment struct {
	Horizontal string
	Vertical   string
	WrapText   bool
}

type Border struct {
	Left   string
	Right  string
	Top    string
	Bottom string
}

type Fill struct {
	PatternType string
	BgColor     string
	FgColor     string
}

type Font struct {
	Size  int
	Name  string
	Bold  bool
	Color string
}

func MakeExcelFile(file *ExcelFile) ([]byte, string, int64, error) {
	f := xlsx.NewFile()
	for _, sheet := range file.Sheets {
		s, err := f.AddSheet(sheet.Name)
		if err != nil {
			break
		}

		if len(sheet.Headers) == 0 {
			return []byte{}, "", 0, errors.New("Header should not null")
		}

		for idx, header := range sheet.Headers {
			var attr *Attribute
			if idx < len(sheet.HeaderAttributes) {
				attr = sheet.HeaderAttributes[idx]
				if attr == nil {
					attr = sheet.CommonRowAttribute
				}
			} else {
				attr = sheet.CommonRowAttribute
			}

			makeRow(s.AddRow(), header, attr)
		}

		var attr *Attribute
		idx := 1                                   // first row of sheet's content
		for _, rowContent := range sheet.Content { // foreach row
			attr = sheet.CommonRowAttribute

			// get attribute for row
			if specificAttr, ok := sheet.SpecificRowAttribute[idx]; ok {
				attr = specificAttr
			}

			makeRow(s.AddRow(), rowContent, attr)
			idx++
		}
	}

	excelFile, err := os.Create(file.Name)
	defer func() {
		excelFile.Close()
		if strings.Contains(file.Name, ".xlsx") {
			err := os.Remove(file.Name)
			if err != nil {
				logrus.WithError(err).Infof("Remove file fail, file name: %s", file.Name)
			}
		}
	}()

	if err != nil {
		logrus.WithError(err).Infof("Create file fail")
		return []byte{}, "", 0, err
	}

	err = f.Write(excelFile)
	if err != nil {
		logrus.WithError(err).Infof("Write file fail")
		return []byte{}, "", 0, err
	}

	_, err = excelFile.Seek(0, 0)
	if err != nil {
		logrus.WithError(err).Infof("Seek file fail")
		return []byte{}, "", 0, err
	}

	fileInfo, err := excelFile.Stat()
	if err != nil {
		return []byte{}, "", 0, err
	}

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	bytesFile, err := ioutil.ReadAll(excelFile)
	if err != nil {
		return []byte{}, "", 0, err
	}

	return bytesFile, fileName, fileSize, nil
}

func makeRow(row *xlsx.Row, content []string, attr *Attribute) {
	row.WriteSlice(&content, len(content))
	makeupRow(row, attr)
}

func makeupRow(row *xlsx.Row, attr *Attribute) {
	if attr != nil {
		style := xlsx.NewStyle()

		if attr.Aligment != nil {
			style.Alignment.WrapText = attr.Aligment.WrapText
			style.Alignment.Horizontal = attr.Aligment.Horizontal
			style.Alignment.Vertical = attr.Aligment.Vertical
			style.ApplyAlignment = true
		}

		if attr.Border != nil {
			style.Border = *xlsx.NewBorder(
				attr.Border.Left,
				attr.Border.Right,
				attr.Border.Top,
				attr.Border.Bottom)
			style.ApplyBorder = true
		}

		if attr.Fill != nil {
			style.Fill = *xlsx.NewFill(
				attr.Fill.PatternType,
				attr.Fill.FgColor,
				attr.Fill.BgColor)
			style.ApplyFill = true
		}

		if attr.Font != nil {
			style.Font = *xlsx.NewFont(
				attr.Font.Size,
				attr.Font.Name)
			style.Font.Bold = attr.Font.Bold
			style.Font.Color = attr.Font.Color
			style.ApplyFont = true
		}

		for _, cell := range row.Cells {
			cell.SetStyle(style)
		}
	}
}

func GetFileName(fileHeader *multipart.FileHeader) (string, string, string) {
	fileNameFull := fileHeader.Filename
	fileName := fileNameFull
	fileExtension := ""
	if index := strings.LastIndex(fileNameFull, "."); index >= 0 {
		fileName = fileNameFull[:index]
		fileExtension = fileNameFull[index:]
	}

	return fileNameFull, fileName, fileExtension
}
