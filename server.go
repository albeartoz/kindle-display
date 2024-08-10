package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const screenshotPath = "/tmp/screenshot.jpg"
const timeStamp = true

/*
	Grab the screenshot from the POST request
	Save it to {screenshotPath}
	Use eips to display the screenshot
*/
func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(screenshotPath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("eips", "-g", screenshotPath)
	stdout, err := cmd.Output()
	if err != nil {
		http.Error(w, "Unable to display file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if timeStamp {
		fmt.Fprintln(w, "File uploaded and saved successfully at " + time.Now().String())
	} else {
		fmt.Fprintln(w, "File uploaded and saved successfully")
	}
	
	fmt.Println(string(stdout))
}

func main() {
	exec.Command("eips", "-c") // Wipe screen

	// Start server on port 4545
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server listening on port 4545...")
	err := http.ListenAndServe(":4545", nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
