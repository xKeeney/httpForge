package httpLogger

func (l *HttpLogger) Print(v ...any) {
	if l.logLevel >= 3 {
		l.logger.Print(v...)
	}
}

func (l *HttpLogger) Printf(format string, v ...any) {
	if l.logLevel >= 3 {
		l.logger.Printf(format, v...)
	}
}

func (l *HttpLogger) Println(v ...any) {
	if l.logLevel >= 3 {
		l.logger.Println(v...)
	}
}
