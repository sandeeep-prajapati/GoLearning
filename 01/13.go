package main
import {
	"fmt"
	"os"
}
type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string){
	fmt.Println("[Console]", msg)
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filepath string) *FileLogger {
    f, _ := os.Create(filepath)
    return &FileLogger{file: f}
}

func (l *FileLogger) Log(msg string) {
    l.file.WriteString(msg + "\n")
}

type NoOpLogger struct{}

func (NoOpLogger) Log(msg string) {} // does nothing

func Process(logger Logger) {
    logger.Log("Task started")
}

func main() {
    Process(ConsoleLogger{})
    Process(NewFileLogger("log.txt"))
    Process(NoOpLogger{})
}