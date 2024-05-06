package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func export2Excel(headerRow []interface{}, dataIterator func() ([]interface{}, bool)) (fileContent []byte, err error) {
	f := excelize.NewFile()
	// 以流式创建写入工作表，避免写入工作表数据量太大导致内存oom
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return []byte{}, err
	}

	// 设置数据行的样式
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center", // 水平居中
			Vertical:   "center", // 垂直居中
		},
	})

	// 写表格标题
	_ = streamWriter.SetColWidth(1, 20, 20)
	err = streamWriter.SetRow("A1", headerRow, excelize.RowOpts{
		StyleID: style,
	})
	if err != nil {
		return []byte{}, err
	}

	// 第一行写标题，从第二行开始写数据
	rowIndex := 2
	for {
		rowData, hasNext := dataIterator()
		if !hasNext {
			break
		}

		err = streamWriter.SetRow("A"+cast.ToString(rowIndex), rowData,
			excelize.RowOpts{StyleID: style})
		if err != nil {
			return []byte{}, err
		}

		rowIndex++
	}
	// flush完后不能再追加数据
	_ = streamWriter.Flush()

	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Export(ctx *gin.Context) {
	// 写表格标题
	headerRow := []interface{}{
		"姓名",
		"年龄",
		"邮件",
	}

	var data = []*Person{
		{Name: "张三", Age: 18, Email: "zhangsan@example.com"},
		{Name: "李四", Age: 20, Email: "lisi@example.com"},
		{Name: "王五", Age: 22, Email: "wangwu@example.com"},
	}

	index := 0
	dataIterator := func() ([]interface{}, bool) {
		if index < len(data) {
			rowData := []interface{}{
				data[index].Name,
				data[index].Age,
				data[index].Email,
			}
			index++
			return rowData, true
		}
		return nil, false
	}

	fileStream, err := export2Excel(headerRow, dataIterator)
	if err != nil {
		fmt.Println("err:", err)
	}

	r := bytes.NewReader(fileStream)
	// 设置文件名
	fileName := "test_" + time.Now().Format("20060102150405") + ".xlsx"
	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// 返回文件流
	http.ServeContent(ctx.Writer, ctx.Request, fileName, time.Now(), r)
}
