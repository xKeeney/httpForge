package httpLogger

func (l *HttpLogger) Trace(v ...any) {
	if l.logLevel >= 5 {
		l.traceLogger.Print(v...)
	}
}

func (l *HttpLogger) Tracef(format string, v ...any) {
	if l.logLevel >= 5 {
		l.traceLogger.Printf(format, v...)
	}
}

func (l *HttpLogger) Traceln(v ...any) {
	if l.logLevel >= 5 {
		l.traceLogger.Println(v...)
	}
}
