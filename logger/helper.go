package logger

import (
	"os"
)

type Helper struct {
	Logger
	fields map[string]interface{}
}

func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

func (h *Helper) Log(level Level, v ...interface{}) {
	h.Logger.Fields(h.fields).Log(level, v...)
}

func (h *Helper) Logf(level Level, format string, v ...interface{}) {
	h.Logger.Fields(h.fields).Logf(level, format, v...)
}

func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Log(InfoLevel, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Logf(InfoLevel, template, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Log(TraceLevel, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Logf(TraceLevel, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Log(DebugLevel, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Logf(DebugLevel, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Log(WarnLevel, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Logf(WarnLevel, template, args...)
}

func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Logf(ErrorLevel, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Log(FatalLevel, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Logf(FatalLevel, template, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}
