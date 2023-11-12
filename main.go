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
)

var version = "undefined"
var metaBuildTime = "undefined"
var metaBuilderOS = "undefined"
var metaBuilderArch = "undefined"

func main() {
	//#region CLI param definitions
	showVersion := flag.Bool("v", false, "Show version information")
	flag.BoolVar(showVersion, "version", false, "Show version information")

	dir := flag.String("d", ".", "The directory to search")
	recursive := flag.Bool("r", false, "Search recursively")
	filetypes := flag.String("ft", "mp4,mkv,avi,mov", "Comma-separated list of filetypes")
	maxLength := flag.Int("max", 30, "Maximum video length")
	minLength := flag.Int("min", 0, "Minimum video length")
	//#endregion

	//#region Parse command-line flags
	flag.Parse()

	//#region Parse CLI params
	//#region Version info
	if *showVersion {
		fmt.Println("fvidl version info:")
		fmt.Println("- Version:", version)
		fmt.Println("- Build time:", metaBuildTime)
		fmt.Println("- Builder OS:", metaBuilderOS)
		fmt.Println("- Builder Arch:", metaBuilderArch)
	}
	//#endregion
	//#endregion

	//#region Directory walker
	filetypeSlice := strings.Split(*filetypes, ",")

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !*recursive && info.IsDir() && path != *dir {
			return filepath.SkipDir
		}

		// Check if it's a regular file
		if info.Mode().IsRegular() {
			// Check if the file has a valid video extension
			if hasValidExtension(info.Name(), filetypeSlice) {
				currentDuration := getVideoDuration(path)

				if currentDuration >= *minLength && currentDuration <= *maxLength {
					fmt.Printf("%s - %s seconds\n", path, strconv.Itoa(currentDuration))
				}
			}
		}

		return nil
	})
	//#endregion

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
