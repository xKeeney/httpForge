package httpLogger

func (l *HttpLogger) Info(v ...any) {
	if l.logLevel >= 3 {
		l.infoLogger.Print(v...)
	}
}

func (l *HttpLogger) Infof(format string, v ...any) {
	if l.logLevel >= 3 {
		l.infoLogger.Printf(format, v...)
	}
}

func (l *HttpLogger) Infoln(v ...any) {
	if l.logLevel >= 3 {
		l.infoLogger.Println(v...)
	}
}
