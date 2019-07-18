package downloadfiles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/cheggaaa/pb/v3"
)

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string, filepath string) error {
	startTime := time.Now()
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writing code to start a Progress bar using the cheggaaa code
	var fileSize int64 = resp.ContentLength
	bar := pb.New(int(fileSize))
	bar.Start()
	rd := bar.NewProxyReader(resp.Body)
	io.Copy(out, rd)
	// Write the body to file
	// fileSize , err = io.Copy(out, resp.Body)
	// if err != nil {
	// 	return err
	// }

	endTime := time.Now()
	fmt.Printf("This operation took %s \n", endTime.Sub(startTime).String())

	return nil
}
