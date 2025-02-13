package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const logDir = "./logs"

func cleanOldLogs() {
	env := NewEnv()

	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Printf("Error reading logs directory: %v", err)
		return
	}

	cutoffDate := time.Now().AddDate(0, 0, -env.LogRetention)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Parse the date from filename (assuming format: YYYY-MM-DD.log)
		fileDate, err := time.Parse("2006-01-02.log", file.Name())
		if err != nil {
			continue // Skip files that don't match our date format
		}

		if fileDate.Before(cutoffDate) {
			filePath := filepath.Join(logDir, file.Name())
			if err := os.Remove(filePath); err != nil {
				log.Printf("Error deleting old log file %s: %v", file.Name(), err)
			} else {
				log.Printf("Deleted old log file: %s", file.Name())
			}
		}
	}
}

func getLogFile() (*os.File, error) {
	env := NewEnv()

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Clean old log files before creating new one
	cleanOldLogs()

	var logPath string
	if env.LogStack == "daily" {
		currentTime := time.Now()
		fileName := fmt.Sprintf("%s.log", currentTime.Format("2006-01-02"))
		logPath = filepath.Join(logDir, fileName)
	} else {
		logPath = filepath.Join(logDir, "error.log")
	}

	return os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func SetupLogger() fiber.Handler {
	logFile, err := getLogFile()
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	return logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Format:     "\n" + `{"timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}", "error":"${error}", "response":[${resBody}]}`,
		Output:     logFile,
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() >= 400 {
				fmt.Println("printing to log file")
				logFile.Write(logString)
				defer logFile.Close()
			}
		},
	})
}
