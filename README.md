# Windows Wlan Password Recovery
Simple tool to recovery wifi password from Windows OS

# Download
The executable file can be downloaded [here](https://github.com/looCiprian/windows-wlan-password-recovery/releases/tag/v1.0)

**The executable file can be targeted as malicious by browsers and anti-virus, please allow download and temporally disable the anti-virus**

# Self compiling (from Windows OS)
```
git clone https://github.com/looCiprian/windows-wlan-password-recovery.git
cd windows-wlan-password-recovery
go build wwpr.go
```

# Cross compiling
```
git clone https://github.com/looCiprian/windows-wlan-password-recovery.git
cd windows-wlan-password-recovery
GOOS=windows GOARCH=386 go build -o wwpr.exe wwpr.go
```
