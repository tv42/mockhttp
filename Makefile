include $(GOROOT)/src/Make.inc

TARG=github.com/tv42/mockhttp.go

GOFILES=\
	mockresponsewriter.go\
	mockrequest.go\

include $(GOROOT)/src/Make.pkg
