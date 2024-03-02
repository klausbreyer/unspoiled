package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type YouTubeResponse struct {
	Items []struct {
		Id struct {
			VideoId string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt string `json:"publishedAt"`
		} `json:"snippet"`
	} `json:"items"`
}

var responseTemplate = template.Must(template.New("response").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Latest YouTube Video - Unspoiled</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
	* {
		font-size: 32px;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	}
	button {
		width: 100%;
		padding: 10px;
		margin: 5px;
		text-align: center;
	}
	</style>
</head>
<body>
	<p>Latest video link: <a href="{{.Link}}" target="_blank">{{.Link}}</a></p>
	<p>Published on: {{.PublishedAt}}</p>
	<button onclick="navigator.clipboard.writeText('{{.Link}}')">Copy to Clipboard</button>
	<button onclick="window.open('{{.Link}}', '_blank')">Open Link</button>
</body>
</html>
`))

var indexTemplate = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Welcome to Unspoiled</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
	* {
		font-size: 24px;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
	}
	</style>
</head>
<body>
	<h2>Welcome to Unspoiled - YouTube Latest Video Finder</h2>
	<p>This project, Unspoiled, helps you find the latest video for a specified YouTube channel. Default region is set to Germany (DE).</p>
	<p>To use, navigate to <a href="/latest?regionCode=DE&channelId=UCmaItsxNPLEQ-NIjv5gPScg">/latest?regionCode=DE&channelId=UCmaItsxNPLEQ-NIjv5gPScg</a> or replace the channel ID and region code to check another channel and region.</p>
</body>
</html>
`))

func getLatestVideo(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channelId")
	regionCode := r.URL.Query().Get("regionCode")
	if regionCode == "" {
		regionCode = "DE" // Set default regionCode to DE
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	apiUrl := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&channelId=%s&part=snippet,id&order=date&maxResults=1&regionCode=%s", apiKey, channelID, regionCode)
	response, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, "Failed to get response from YouTube API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read response from YouTube API", http.StatusInternalServerError)
		return
	}
	var ytResponse YouTubeResponse
	err = json.Unmarshal(responseData, &ytResponse)
	if err != nil {
		http.Error(w, "Failed to parse response from YouTube API", http.StatusInternalServerError)
		return
	}
	if len(ytResponse.Items) > 0 {
		videoId := ytResponse.Items[0].Id.VideoId
		videoLink := "https://www.youtube.com/watch?v=" + videoId
		publishedAt := ytResponse.Items[0].Snippet.PublishedAt
		responseTemplate.Execute(w, struct {
			Link        string
			PublishedAt string
		}{
			Link:        videoLink,
			PublishedAt: publishedAt,
		})
	} else {
		http.Error(w, "No videos found", http.StatusNotFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/latest", getLatestVideo)
	http.HandleFunc("/", index)
	fmt.Println("Starting Unspoiled server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
