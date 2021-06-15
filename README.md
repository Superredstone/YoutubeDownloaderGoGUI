<p align="center">
  <img src="https://github.com/Superredstone/youtubeDownloaderGo/blob/master/Assets/Download.png" width=200></img> <br><br>
  <img src="https://img.shields.io/github/go-mod/go-version/Superredstone/YoutubeDownloaderGoGUI?label=Version&style=for-the-badge">
   <img src="https://img.shields.io/github/languages/top/Superredstone/YoutubeDownloaderGoGUI?style=for-the-badge">
  <img src="https://img.shields.io/github/license/Superredstone/YoutubeDownloaderGoGUI?color=Green&style=for-the-badge">
  <img src="https://img.shields.io/discord/821836676607115304?label=Discord&style=for-the-badge"></img>
<p>

# YoutubeDownloaderGoGUI
GUI Implementation of the YoutubeDownloaderGo library (https://github.com/Superredstone/youtubeDownloaderGo)
  
<div align="center">
  <img src="https://github.com/Superredstone/YoutubeDownloaderGoGUI/blob/main/Assets/Screenshot.png"></img>
</div>

# How to build 🛠️

*READ ME*
To build without having an heartattack you need [task](https://taskfile.dev/#/) which is a nice tool for 
task running like make but readable. You can compile also looking at the tasks inside [Taskfile.yml](""https://github.com/Superredstone/YoutubeDownloaderGoGUI/blob/main/Taskfile.yml)

## Linux
First download dependeces
```bash
sudo apt install gcc pkg-config libwayland-dev libx11-dev libx11-xcb-dev libxkbcommon-x11-dev libgles2-mesa-dev libegl1-mesa-dev libffi-dev libxcursor-dev
```
Then clone the repository and build
```bash
git clone https://github.com/Superredstone/YoutubeDownloaderGoGUI.git
cd YoutubeDownloaderGoGUI
mkdir build/
task build_windows
```

## Windows
Clone the repository and pray everything that builds on first try... (You can't build linux applications on Windows) 
```bash
git clone https://github.com/Superredstone/YoutubeDownloaderGoGUI.git

cd YoutubeDownloaderGoGUI
task build_windows_amd64
or 
task build_windows_arm
```

## MacOS
MacOS users, you need to have Xcode and then view this guide [here](https://gioui.org/doc/install#apple) or you can use the CLI application, you can find it [here](https://github.com/Superredstone/youtubeDownloaderGo)
  
## Android
```
git clone https://github.com/Superredstone/YoutubeDownloaderGoGUI.git
cd YoutubeDownloaderGoGUI

task build_android
```

# Donate :heart:
Every donation is appriciated <3 <br> <br>
<a href='https://ko-fi.com/A0A64PC0Y' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://cdn.ko-fi.com/cdn/kofi3.png?v=2' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>
