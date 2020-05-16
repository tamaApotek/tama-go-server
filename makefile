all: watch
.PHONY: all

watch:
	CompileDaemon -directory="." -command="./tama-go-server"

clean:
	rm tama-go-server
