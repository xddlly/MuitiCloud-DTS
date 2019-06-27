package log

import (
	"os"
	"errors"
	"log"
)

type Logger struct {
    level    string   `error base info`
    logfile  string   `log file`
    logcount int      `log count`
}


func InitLog(logname string,level string) *Logger {
	err := CreatePath("runtime")
	if(err != nil) {
		panic(err)
	}
    logger := &Logger{}
    if(level != "info" && level != "base" && level != "error" ) {
		panic(errors.New("Wrong log level with " + level))
	}
    logger.level    = level
	f, err := os.OpenFile("./runtime/" + logname, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm)
	if(err != nil) {
		panic(errors.New("Can not open or create file /runtime/" + logname))
	}
	defer  f.Close()
    logger.logfile = "./runtime/" + logname
    return logger
}

func (logger * Logger) Info(errStr string) {
    if(logger.level == "info") {
		f, err := os.OpenFile(logger.logfile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm)
		if(err != nil) {
			panic(errors.New("Can not open or create file "  + logger.logfile))
		}
		defer  f.Close()
		loggerWrite := log.New(f, "info: ", log.LstdFlags|log.Lshortfile)
		loggerWrite.Println(errStr)
	}
}

func (logger * Logger) Base(errStr string) {
	if(logger.level == "info" || logger.level == "base") {
		f, err := os.OpenFile(logger.logfile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm)
		if(err != nil) {
			panic(errors.New("Can not open or create file "  + logger.logfile))
		}
		defer  f.Close()
		loggerWrite := log.New(f, "base: ", log.LstdFlags|log.Lshortfile)
		loggerWrite.Println(errStr)
	}
}

func (logger * Logger) Error(errStr string) {
	f, err := os.OpenFile(logger.logfile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm)
	if(err != nil) {
		panic(errors.New("Can not open or create file "  + logger.logfile))
	}
	defer  f.Close()
	loggerWrite := log.New(f, "base: ", log.LstdFlags|log.Lshortfile)
	loggerWrite.Println(errStr)
}

func (logger * Logger) log(errStr string) {
	f, err := os.OpenFile(logger.logfile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm)
	if(err != nil) {
		panic(errors.New("Can not open or create file "  + logger.logfile))
	}
	defer  f.Close()
	loggerWrite := log.New(f, "error: ", log.LstdFlags|log.Lshortfile)
	loggerWrite.Println(errStr)
}

func CreatePath(path string) (error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0666)
		if err != nil {
			return errors.New("没有权限创建目录" + path)
		}
	} else {
		return nil
	}
	return nil
}