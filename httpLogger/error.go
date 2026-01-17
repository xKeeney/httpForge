package httpLogger

func (l *HttpLogger) Error(v ...any) {
	if l.logLevel >= 1 {
		l.errorLogger.Print(v...)
	}
}

func (l *HttpLogger) Errorf(format string, v ...any) {
	if l.logLevel >= 1 {
		l.errorLogger.Printf(format, v...)
	}
}

func (l *HttpLogger) Errorln(v ...any) {
	if l.logLevel >= 1 {
		l.errorLogger.Println(v...)
	}
}
