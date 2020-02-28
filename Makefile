TARGET=$(shell basename $(PWD))
main:
	cd src && go build -o ../bin/$(TARGET)
clean:
	rm -r bin/$(TARGET)
