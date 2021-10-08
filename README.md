[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/Gyro7/">
    <img src="assets/logo.png" alt="Logo" width="200">
  </a>
</p>
<h1 align="center">mangodl</h1>

  <p align="center">
Download and search manga right from the terminal!
    <br />
    <br />
    <a href="https://github.com/Gyro7/mangodl/issues">Report Bug</a> || 
    <a href="https://github.com/Gyro7/mangodl/pulls">Request Feature</a>
  </p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

-   [About the Project](#about-the-project)
    -   [Built With](#built-with)
-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
    -   [Installation](#installation)
        -   [Linux](#linux)
        -   [Linux (Build from Source)](#linux-build-from-source)
        -   [Arch Linux and derivatives](#arch-linux-and-derivatives)
        -   [Windows](#windows)
        -   [macOS](#macos)
        -   [macOS (alternative)](#macos-alternative)
-   [Usage](#usage)
-   [Reading](#reading)
-   [Roadmap](#roadmap)
-   [Contributing](#contributing)
-   [License](#license)
-   [Contact](#contact)
-   [Acknowledgements](#acknowledgements)

## About The Project

<br>
<p align="center">
   An easy-to-use cli tool for downloading manga 
  <br>
  <br>
<img src="assets/example.gif" alt="example" width="800">
</p>

### Built With

-   [Go](https://golang.org)
-   [Goquery](https://github.com/PuerkitoBio/goquery)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

-   Go compiler (if you want to build from source)
-   Linux, Windows or Mac

### Installation

#### Linux
Download the files mangodl and install.sh from the latest Linux version in the [Releases](https://github.com/Gyro7/mangodl/releases)
```sh
# run the installation script
chmod +x install.sh
./install.sh
```

#### Linux (build from source)
```sh
# clone and go into the repository
git clone https://github.com/Gyro7/mangodl.git
cd mangodl

# NOW JUST OPEN THE INSTALL.SH SCRIPT AND UNCOMMENT THE COMMENTED LINES
# then run the installation script
chmod +x install.sh
./install.sh
```
#### Arch Linux and derivatives
An AUR package is now available.

Just `yay -S mangodl` or, if you use paru `paru -S mangodl`
#### Windows
Download the executable (mangodl.exe) from the latest Windows version in the [Releases](https://github.com/Gyro7/mangodl/releases)  
If you just want to use it without installing it, just run mangodl.exe everytime and skip these steps below
```sh
Open start menu,
1. Type Edit environment variables
2. Open the option Edit the system environment variables
3. Click Environment variables... button
4. There you see two boxes, in System Variables box find path variable
5. Click Edit
6. a window pops up, click New
7. Type the Directory path of mangodl.exe (Directory means exclude the file name from path)
8. Click Ok on all open windows and restart the command prompt.
```

#### macOS
If you haven't already given the terminal access to the disk, then do it, for further help see <b>[THIS](https://osxdaily.com/2018/10/09/fix-operation-not-permitted-terminal-error-macos/) </b>  
Installing via brew (assuming that `/usr/local/bin` is already in the $PATH variable):
```bash
brew tap Gyro7/mangodl
brew install --build-from-source mangodl
```

#### macOS (alternative)
Download the executable mangodl-darwin from the latest macOS version in the [Releases](https://github.com/Gyro7/mangodl/releases)
```sh
# rename the executable
mv mangodl-darwin mangodl
chmod +x mangodl
# move the executable to the /usr/local/bin/ path, be aware of not deleting the directory!
sudo mv mangodl /usr/local/bin/mangodl
```
## Usage
Usage: mangodl [FLAGS]...

Arguments and flags:

	-h, --help			shows this message and exit

	Needed (one of them):
	-D, --download			downloads the manga specified after -D (e.g. mangodl -D jojo will search for 5 manga with that name and ask you which one to download)
	-S, --search			searches for the manga specified after this flag (e.g. mangodl -S "kanojo x kanojo" will search and display the manga found with that name)
	-Q, --query			show downloaded manga
	-Dir, --directory		sets the default directory to download manga (e.g. mangodl -Dir "$HOME/Documents/manga/"), otherwise the default one would be "$HOME/Downloaded Manga/" and the Desktop for Windows
	
	Optional:
	For -D:
	-c, --chapter			used to specify the chapter to download (if omitted it will download them all)
	-cr, --chapterrange		used to specify a range of chapters to download (e.g. mangodl -D "Martial Peak" -cr 1 99 will download chapters from 1 to 99 (included)
	-o, --output			used to specify the file output of the pages (img or pdf), e.g. mangodl -D "Tokyo Revengers" -o pdf will create a pdf for every chapter. By default, it's images.
	
	For -S:
	-n, --noplot		do not print the plot of searched manga	
<!-- ROADMAP -->

## Reading
To read the downloaded pages, I really suggest this free piece of software, which is lightweight and flexible:
### <u>[OpenComic](https://github.com/ollm/OpenComic) </u>
<img src="assets/opencomic.png" alt="OpenComic Demonstration" width="800">

You can simply add the folder "Downloaded Manga" to OpenComic, which is the most recommended thing to do.
<br>
And then it'll open all your manga, divided into chapters.
## Roadmap

See the [open issues](https://github.com/Gyro7/mangodl/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the GPL 3.0 License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Me - [gyro@sach1.tk](mailto:gyro@sach1.tk)

Project Link: [https://github.com/Gyro7/mangodl](https://github.com/Gyro7/gofetch)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

-   [Myself for doing everything.](https://github.com/Gyro7)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/Gyro7/mangodl.svg?style=flat-square
[contributors-url]: https://github.com/Gyro7/mangodl/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Gyro7/mangodl.svg?style=flat-square
[forks-url]: https://github.com/Gyro7/mangodl/network/members
[stars-shield]: https://img.shields.io/github/stars/Gyro7/mangodl.svg?style=flat-square
[stars-url]: https://github.com/Gyro7/mangodl/stargazers
[issues-shield]: https://img.shields.io/github/issues/Gyro7/mangodl.svg?style=flat-square
[issues-url]: https://github.com/Gyro7/mangodl/issues
[license-shield]: https://img.shields.io/github/license/Gyro7/mangodl.svg?style=flat-square
[license-url]: https://github.com/Gyro7/mangodl/blob/main/LICENSE

