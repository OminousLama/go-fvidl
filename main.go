package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var version string

func main() {
	//#region CLI param definitions
	dir := flag.String("dir", ".", "The directory to search")
	//recursive := flag.Bool("recursive", false, "Search recursively")
	filetypes := flag.String("filetypes", "mp4,mkv,avi,mov", "Comma-separated list of filetypes")
	maxLengthInput := flag.String("maxLength", "30", "Maximum video length")
	minLengthInput := flag.String("minLength", "0", "Minimum video length")
	//#endregion

	//#region Parse command-line flags
	flag.Parse()

	//#region Parse CLI params
	//#region Max Length
	maxLength := 30

	if *maxLengthInput == "" {
		fmt.Printf("Using default maxLength value: %d\n", maxLength)
	} else {
		if i, err := strconv.Atoi(*maxLengthInput); err == nil {
			maxLength = i
		} else {
			fmt.Printf("Error: The value for maxLength is not a valid integer.", err)
		}
	}
	//#endregion
	//#region Min Length
	minLength := 0

	if *minLengthInput == "" {
		fmt.Printf("Using default minLength value: %d\n", minLength)
	} else {
		if i, err := strconv.Atoi(*minLengthInput); err == nil {
			minLength = i
		} else {
			fmt.Printf("Error: The value for minLength is not a valid integer.", err)
		}
	}
	//#endregion
	//#endregion

	// Split filetypes into a slice
	filetypeSlice := strings.Split(*filetypes, ",")

	// Walk through the directory
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Check if it's a regular file
		if info.Mode().IsRegular() {
			// Check if the file has a valid video extension
			if hasValidExtension(info.Name(), filetypeSlice) {
				currentDuration := getVideoDuration(path)

				if currentDuration >= minLength && currentDuration <= maxLength {
					fmt.Printf("%s - %s seconds\n", path, strconv.Itoa(currentDuration))
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func hasValidExtension(filename string, validExtensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, validExt := range validExtensions {
		if ext == "."+validExt {
			return true
		}
	}
	return false
}

func getVideoDuration(filePath string) int {

	// Run FFmpeg command
	cmd := exec.Command("ffprobe", "-i", filePath, "-show_entries", "format=duration", "-v", "quiet")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	pattern := `duration=([0-9.]+)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(string(output))

	duration := 0
	if len(match) >= 2 {
		num, err := strconv.Atoi(strings.Split(match[1], ".")[0])
		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}
		duration = num
	} else {
		fmt.Println("Duration information not found.")
	}

	return duration
}

func formatDuration(duration time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
}
