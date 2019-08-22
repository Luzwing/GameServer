package serverlog

import (
	"log"
	"os"
	"time"
)

type ServerLog struct {
	fileName string
	logging  os.File
	logger   log.Logger
}

func (sl *ServerLog) LogInit() error {
	sl.fileName = "./" + time.Now().Format("2006-01-02") + ".log"
	tlogging, logErr := os.OpenFile(sl.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	sl.logging = *tlogging
	if logErr != nil {
		return logErr
	}
	sl.logger = *log.New(&sl.logging, "", log.LstdFlags|log.Lshortfile)

	return nil
}

func (sl *ServerLog) Log2file(format string) {
	sl.logger.Println(format)
}

func (sl *ServerLog) Log2filef(format string, v ...interface{}) {
	sl.logger.Printf(format, v)
}

func (sl *ServerLog) Close() {
	sl.logging.Close()
}

func (sl ServerLog) GetFileName() string {
	return sl.fileName
}

func (sl ServerLog) GetLogginFile() os.File {
	return sl.logging
}

func (sl ServerLog) GetLogger() log.Logger {
	return sl.logger
}

// func main() {
// 	var slog ServerLog
// 	logerr := slog.LogInit()
// 	if logerr != nil {
// 		os.Exit(-1)
// 	}
// 	slog.Log2file("hahha")
// 	defer slog.Close()
// }
