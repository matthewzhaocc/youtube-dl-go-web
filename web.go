// golang web imple
package main

// dependencies
import (
	//	"fmt"
	"io"

	//	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	youtube "github.com/kkdai/youtube/v2"
	"github.com/sirupsen/logrus"
)

// static configurations
const (
	DownloadDirectory = "downloads/"
)
// other global vars
var (
	log = logrus.New()
)
// downloader
func DownloadVideo(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"path": r.URL.Path,
		"XFF Origin": r.Header.Get("X-Forwarded-For"),
		"Host Header": r.Host,
	}).Info("just received a request for the download path")
	// parse URL
	var VideoId string
	UrlPath := strings.Split(r.URL.Path, "/")
	if len(UrlPath) != 3 {
		http.Error(w, "bad url", 400)
		return
	} else {
		VideoId = UrlPath[2]
	}
	// get Client instance
	client := youtube.Client{}
	video, err := client.GetVideo(VideoId)
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	// get the video itself
	resp, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	defer resp.Body.Close()

	// creates the file
	file, err := os.Create(path.Join("downloads/video.mp4"))
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	// saves file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	// serves the file
	http.ServeFile(w, r, "downloads/video.mp4")
}

// initializes some stuff
func init() {
	log.SetOutput(os.Stdout)
}
// starts the site
func main() {
	http.HandleFunc("/download/", DownloadVideo)
	log.WithFields(logrus.Fields{
		"event":"server started",
		"port": "6443",
	}).Info("started listening on port 6443")
	http.ListenAndServe(":6443", nil)
}
