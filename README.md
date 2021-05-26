[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/Gyro7/">
    <img src="_assets/mangodl.png">
  </a>
</p>
<h3 align="center">GoFetch</h3>

  <p align="center">
    Show off your Go information with this cool command-line tool!
    <br />
    <br />
    <a href="https://github.com/Gyro7/gofetch/issues">Report Bug</a> || 
    <a href="https://github.com/Gyro7/gofetch/pulls">Request Feature</a>
  </p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

-   [About the Project](#about-the-project)
    -   [Built With](#built-with)
-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
    -   [Installation](#installation)
-   [Usage](#usage)
-   [Roadmap](#roadmap)
-   [Contributing](#contributing)
-   [License](#license)
-   [Contact](#contact)
-   [Acknowledgements](#acknowledgements)

## About The Project

<br>
<p align="center">A pretty command-line "Go and System information" tool written in Go
  <br>
  <br>
<img src="https://i.imgur.com/Vm9gENO.png" alt="example" width="800">
</p>

### Built With

-   [Go](https://golang.org)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

-   golang
-   linux, android (Termux) or windows (not available for mac at the moment)

### Installation

#### Linux
```sh
# clone and go into repo
git clone https://github.com/Gyro7/gofetch.git
cd gofetch/
# install
sudo install -m755 gofetch /usr/bin/gofetch
# go back and remove the download
cd ..
rm -rf gofetch/
# run
gofetch
```
#### Windows
```sh
Open start menu,
1. Type Edit environment variables
2. Open the option Edit the system environment variables
3. Click Environment variables... button
4. There you see two boxes, in System Variables box find path variable
5. Click Edit
6. a window pops up, click New
7. Type the Directory path of gofetch.exe (Directory means exclude the file name from path)
8. Click Ok on all open windows and restart the command prompt.
```
#### Android
The only way to use gofetch for android if not having root access is building from source.  
If you have root access just follow the [Linux Installation](#linux)
```sh
# clone and go into repo
git clone https://github.com/Gyro7/gofetch.git
cd gofetch/
# remove the linux executable
rm gofetch
# build
go build
# run
./gofetch
```
## Usage

If you followed the previous steps, you just have to run the program with

```sh
gofetch # start the main program
```

<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/Gyro7/gofetch/issues) for a list of proposed features (and known issues).

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

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

gyro - [gyro@sach1.tk](mailto:gyro@sach1.tk)
Project Link: [https://github.com/Gyro7/gofetch](https://github.com/Gyro7/gofetch)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

-   [Myself for doing everything.](https://github.com/Gyro7)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/Gyro7/gofetch.svg?style=flat-square
[contributors-url]: https://github.com/Gyro7/gofetch/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Gyro7/gofetch.svg?style=flat-square
[forks-url]: https://github.com/Gyro7/gofetch/network/members
[stars-shield]: https://img.shields.io/github/stars/Gyro7/gofetch.svg?style=flat-square
[stars-url]: https://github.com/Gyro7/gofetch/stargazers
[issues-shield]: https://img.shields.io/github/issues/Gyro7/gofetch.svg?style=flat-square
[issues-url]: https://github.com/Gyro7/gofetch/issues
[license-shield]: https://img.shields.io/github/license/Gyro7/gofetch.svg?style=flat-square
[license-url]: https://github.com/Gyro7/gofetch/blob/master/LICENSE

