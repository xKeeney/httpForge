package httpLogger

func (l *HttpLogger) Fatal(v ...any) {
	if l.logLevel >= 0 {
		l.fatalLogger.Fatal(v...)
	}
}

func (l *HttpLogger) Fatalf(format string, v ...any) {
	if l.logLevel >= 0 {
		l.fatalLogger.Fatalf(format, v...)
	}
}

func (l *HttpLogger) Fatalln(v ...any) {
	if l.logLevel >= 0 {
		l.fatalLogger.Fatalln(v...)
	}
}
