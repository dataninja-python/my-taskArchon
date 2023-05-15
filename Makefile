.PHONY:	setup build clean

# Setup taskwarrior project
setup:	go mod init taskwarrior

# Build taskwarrior
build:	go build -o taskwarrior main.go task.go taskstorage.go pomodoro.go

# Clean taskwarrior build artifacts
clean:	rm -f taskwarrior
