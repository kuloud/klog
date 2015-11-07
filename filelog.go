package klog

import (
	"fmt"
	"os"
	"time"
)

type FileLogWriter struct {
	rec chan *LogRecord
	rot chan bool

	fileName string
	file     *os.File

	maxLines int
	curLines int

	maxSize int
	curSize int

	currentDate int
}

func (w *FileLogWriter) LogWrite(rec *LogRecord) {
	w.rec <- rec
}

func (w *FileLogWriter) Close() {
	close(w.rec)
	w.file.Sync()
}

func NewFileLogWriter(fname string) *FileLogWriter {
	w := &FileLogWriter{
		rec:      make(chan *LogRecord, LogBufferLength),
		rot:      make(chan bool),
		maxLines: FILE_LOG_MAX_LINES,
		maxSize:  FILE_LOG_MAX_SIZE,
		fileName: fname,
	}

	if err := w.intRotate(); err != nil {
		fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.fileName, err)
		return nil
	}

	go func() {
		defer func() {
			if w.file != nil {
				w.file.Close()
			}
		}()

		for {
			select {
			case <-w.rot:
				if err := w.intRotate(); err != nil {
					fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.fileName, err)
					return
				}

			case rec, ok := <-w.rec:
				if !ok {
					return
				}
				now := time.Now()
				if (w.maxLines > 0 && w.curLines >= w.maxLines) ||
					(w.maxSize > 0 && w.curSize >= w.maxSize) ||
					(w.currentDate != now.Day()) {
					if err := w.intRotate(); err != nil {
						fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.fileName, err)
						return
					}
				}

				n, err := fmt.Fprint(w.file, rec.Created.Format("2006-01-02 15:04:05.999")+rec.Message)
				if err != nil {
					fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.fileName, err)
					return
				}

				// Update the counts
				w.curLines++
				w.curSize += n
			}
		}
	}()

	return w
}

// Request that the logs rotate
func (w *FileLogWriter) Rotate() {
	w.rot <- true
}

func (w *FileLogWriter) intRotate() error {
	if w.file != nil {
		w.file.Close()
	}

	_, err := os.Lstat(w.fileName)
	if err == nil { // file exists
		num := 1
		fname := ""
		if time.Now().Day() != w.currentDate {
			yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			for ; err == nil && num <= 999; num++ {
				fname = w.fileName + fmt.Sprintf("_%s_%03d", yesterday, num)
				_, err = os.Lstat(fname)
			}
		} else {
			for ; err == nil && num <= 999; num++ {
				today := time.Now().Format("2006-01-02")
				fname = w.fileName + fmt.Sprintf("_%s_%03d", today, num)
				_, err = os.Lstat(fname)
			}
		}
		if err == nil {
			return fmt.Errorf("Rotate: Cannot find free log number to rename %s\n", w.fileName)
		}
		w.file.Close()
		err = os.Rename(w.fileName, fname)
		if err != nil {
			return fmt.Errorf("Rotate: %s\n", err)
		}
	}

	fd, err := os.OpenFile(w.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	w.file = fd

	now := time.Now()
	w.currentDate = now.Day()
	w.curLines = 0
	w.curSize = 0
	return nil
}
