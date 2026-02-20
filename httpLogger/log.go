package httpLogger

func (l *HttpLogger) Print(v ...any) {
	if l.logLevel >= 3 {
		l.Logger.Print(v...)
	}
}

func (l *HttpLogger) Printf(format string, v ...any) {
	if l.logLevel >= 3 {
		l.Logger.Printf(format, v...)
	}
}

func (l *HttpLogger) Println(v ...any) {
	if l.logLevel >= 3 {
		l.Logger.Println(v...)
	}
}
