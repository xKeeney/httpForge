package httpLogger

func (l *HttpLogger) Debug(v ...any) {
	if l.logLevel >= 4 {
		l.debugLogger.Print(v...)
	}
}

func (l *HttpLogger) Debugf(format string, v ...any) {
	if l.logLevel >= 4 {
		l.debugLogger.Printf(format, v...)
	}
}

func (l *HttpLogger) Debugln(v ...any) {
	if l.logLevel >= 4 {
		l.debugLogger.Println(v...)
	}
}
