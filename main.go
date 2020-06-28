package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
	"webdavGo/webdav"
)

func main() {
	uploadRateLimt, _ := strconv.ParseFloat(os.Getenv("UPLOAD_RATE_LIMIT"), 64)

	ser := &webdav.Handler{
		FileSystem: webdav.Dir("/webdav_files"),
		LockSystem: webdav.NewMemLS(),
		Logger: func(request *http.Request, err error) {
			if err != nil {
				fmt.Println(fmt.Sprintf("request fail:%s", err.Error()))
				return
			}
			uri, _ := url.QueryUnescape(request.RequestURI)
			fmt.Println(fmt.Sprintf(`%s method:%s uri:%s`,
				time.Now().Format("2006-01-02 15:04:05"),
				request.Method,
				uri,
			))
		},
		UploadRateLimit: uploadRateLimt,
	}

	err := http.ListenAndServe(":8080", ser)

	if err != nil {
		fmt.Println("webdav service setup fail:", err)
	} else {
		fmt.Println("close webdav service")
	}
}
