# Webdog - Lightweight and customizable web service

## What is webdog？
Webdog is a lightweight and customizable web service. It allows us to configure routes via yaml, implement controllers externally.You can implement an external controller in any language like bash, lua, python ... It can also be a simple static resources web service.

## Installation
### Download Releases
Chick here [releases](https://github.com/edboffical/webdog/releases/) to download the executable file corressponding to your system
### Building
```bash
git clone https://github.com/edboffical/webdog
cd webdog
# for mac
make build-mac
# for linux
make build-linux
# for windows
make build-win
```

## Documentation
[Documentation](./DOCUMENT.md)

## Dependencies
- [httprouter](https://github.com/julienschmidt/httprouter)
- [fsnotify](https://github.com/fsnotify/fsnotify)