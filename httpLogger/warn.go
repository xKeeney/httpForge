package httpLogger

func (l *HttpLogger) Warn(v ...any) {
	if l.logLevel >= 2 {
		l.warnLogger.Print(v...)
	}
}

func (l *HttpLogger) Warnf(format string, v ...any) {
	if l.logLevel >= 2 {
		l.warnLogger.Printf(format, v...)
	}
}

func (l *HttpLogger) Warnln(v ...any) {
	if l.logLevel >= 2 {
		l.warnLogger.Println(v...)
	}
}
