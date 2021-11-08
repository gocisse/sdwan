# sdwan


This is an sdwan API client writting in golang 
Its use to fetch devices states ,  ipsec ,


Building Executables for Different Architectures

The go build command lets you build an executable file for any Go-supported target platform, on your platform. This means you can test, release and distribute your application without building those executables on the target platforms you wish to use.

Cross-compiling works by setting required environment variables that specify the target operating system and architecture. We use the variable GOOS for the target operating system, and GOARCH for the target architecture. To build an executable, the command would take this form:

    
	
	
	env GOOS=target-OS GOARCH=target-architecture go build package-import-path

 

The env command runs a program in a modified environment. This lets you use environment variables for the current command execution only. The variables are unset or reset after the command executes.

The following table shows the possible combinations of GOOS and GOARCH you can use:
GOOS - Target Operating System 	GOARCH - Target Platform
android 	arm
darwin 	386
darwin 	amd64
darwin 	arm
darwin 	arm64
dragonfly 	amd64
freebsd 	386
freebsd 	amd64
freebsd 	arm
linux 	386
linux 	amd64
linux 	arm
linux 	arm64
linux 	ppc64
linux 	ppc64le
linux 	mips
linux 	mipsle
linux 	mips64
linux 	mips64le
netbsd 	386
netbsd 	amd64
netbsd 	arm
openbsd 	386
openbsd 	amd64
openbsd 	arm
plan9 	386
plan9 	amd64
solaris 	amd64
windows 	386
windows 	amd64 

Warning: Cross-compiling executables for Android requires the Android NDK, and some additional setup which is beyond the scope of this tutorial.

Using the values in the table, we can build Caddy for Windows 64-bit like this:

    env GOOS=windows GOARCH=amd64 go build  -o appname.exe  github.com/gicisse/sdwan/main.go
	
	or download the repository and run : 
	
	cd sdwan 
	GOOS=windows GOARCH=386 go build -o appname.exe main.go

 

Once again, no output indicates that the operation was successful. The executable will be created in the current directory, using the package name as its name. However, since we built this executable for Windows, the name ends with the suffix .exe.

You should have appname.exe file in your current directory, which you can verify with the ls or dir command.

    dir   appname.exe
