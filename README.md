# 🔑 Master Password
The Master Password generates unique passwords in a consistent way, unlike traditional password managers. Passwords are not saved but generated every time based on information entered by the user, such as their name, master password, and a unique identifier for the service the password is used for (URL). 

![Golang](https://img.shields.io/badge/language-Go-blue) [![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FStepashka20%2Fmaster-password&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

## ⁉ Introduction

Master Password is a password generator that generates unique passwords in a consistent way, unlike traditional password managers. The generated passwords are not saved, but instead generated every time based on information entered by the user, such as their name, master password, and a unique identifier for the service the password is used for (URL).

## 🖼️ Screenshots
 
| ![image](https://user-images.githubusercontent.com/40739871/217006677-89c70913-e249-4f30-86b3-5e34a9bd9bf3.png) | ![image](https://user-images.githubusercontent.com/40739871/217001331-7889cf84-8c72-409e-a169-79b8d0de1df8.png) |
| --- | --- |
| ![image](https://user-images.githubusercontent.com/40739871/217007026-afc65bb7-568e-4404-9fc6-f2a4983b9a88.png) | ![image](https://user-images.githubusercontent.com/40739871/217007130-25aeb37c-c13c-4e18-9e52-fae12e050490.png) |


## Features

- Consistent password generation
- No saved passwords
- Customizable password generation
- Easy to use UI
- Secure algorithms
- Lightweight (7MB exe, 3MB with UPX)
- Open source

## ⚒ Building
### 1. Prerequisites for build
Linux: 
```
apt install -y gcc xorg-dev libgtk-3-dev libgl1-mesa-dev libglu1-mesa
```
Windows: A C compiler, ideally TDM-GCC or MinGW-w64

### 2. Build From Source
```
git clone https://github.com/Stepashka20/master-password
cd master-password
```

- Windows: `go build -ldflags="-s -w -H=windowsgui -extldflags=-static"`
- Linux: `go build -ldflags="-s -w"`

## 📁 License
This project is licensed under the MIT License
