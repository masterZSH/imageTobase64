package main

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const uploadPath = "./tmp"

// 上传处理，输出图片的base64编码
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	imagStr := toBase64(r)
	log.Println(imagStr)
	w.Write([]byte(imagStr))
}

// 对字符串进行md5加密
func md5(s string) string {
	md5Ctx := crypto.MD5.New()
	md5Ctx.Write([]byte(s))
	md5str := fmt.Sprintf("%x", md5Ctx.Sum(nil))
	return md5str
}

// 上传文件转base64
func toBase64(r *http.Request) string {
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	headerStr := "data:image/png;base64,"
	imgBaseStr := base64.StdEncoding.EncodeToString(content)
	imgBaseStr = strings.Join([]string{headerStr, imgBaseStr}, "")
	return imgBaseStr
}

// 保存图片返回保存到本地的图片的名称
func saveFile(r *http.Request) string {
	sFile, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	ext := path.Ext(header.Filename)
	nFileName := md5(header.Filename)
	log.Println(nFileName)
	nFileName = strings.Join([]string{nFileName, ext}, "")
	dest, err := os.Create(nFileName)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer dest.Close()
	_, err = io.Copy(dest, sFile)
	if err != nil {
		panic(err)
	}
	return nFileName
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
