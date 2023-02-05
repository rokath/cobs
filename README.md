# COBS

<!-- PROJECT LOGO -->

<div align="center">
  <a href="https://github.com/rokath/cobs">
    <img src="./logo.png" alt="Logo" width="800" height="80">
  </a>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>

<!-- vscode-markdown-toc -->
- [COBS](#cobs)
  - [1. About The project](#1-about-the-project)
  - [2. COBS Specification](#2-cobs-specification)
  - [3. Folder Information](#3-folder-information)
  - [4. Getting Started](#4-getting-started)
    - [4.1. Prerequisites](#41-prerequisites)
    - [4.2. Installation](#42-installation)
    - [4.3. Roadmap](#43-roadmap)
  - [5. Contributing](#5-contributing)
  - [6. License](#6-license)
  - [7. Contact](#7-contact)
    - [7.1. Acknowledgments](#71-acknowledgments)
<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc --><div id="top"></div>

  </ol>
</details>

<!--
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/rokath/cobs/goreleaser)
![GitHub All Releases](https://img.shields.io/github/downloads/rokath/cobs/total)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/rokath/cobs)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/rokath/cobs/latest)
-->

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/rokath/cobs)
![GitHub watchers](https://img.shields.io/github/watchers/rokath/cobs?label=watch)
[![Go Report Card](https://goreportcard.com/badge/github.com/rokath/cobs)](https://goreportcard.com/report/github.com/rokath/cobs)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![test](https://github.com/shogo82148/actions-goveralls/workflows/test/badge.svg?branch=main)](https://coveralls.io/github.com/rokath/cobs)
[![Coverage Status](https://coveralls.io/repos/github.com/rokath/cobs/badge.svg?branch=master)](https://coveralls.io/github.com/rokath/cobs?branch=master)
![GitHub issues](https://img.shields.io/github/issues/rokath/cobs)


<!-- ABOUT THE PROJECT -->
##  1. <a name='AboutTheproject'></a>About The project

This is a [COBS](https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing) implementation in **Go** and **C** with tests.

* The **C** code is copied from Wikipedia and adapted. Some other implementations did not pass all tests.
* The **Go** code was written from scratch after reviewing other implementations.

<p align="right">(<a href="#top">back to top</a>)</p>

##  2. <a name='COBSSpecification'></a>COBS Specification

* See [https://pythonhosted.org/cobs/cobs.cobs.html](https://pythonhosted.org/cobs/cobs.cobs.html).

<p align="right">(<a href="#top">back to top</a>)</p>

##  3. <a name='FolderInformation'></a>Folder Information

| Folder | Content |
| - | - |
| src | C sources ready to be used in any C project |
| go  | **Go** source |
| cgo | **Go** project for testing |


<!-- GETTING STARTED -->
##  4. <a name='GettingStarted'></a>Getting Started

* Clone the project with `git clone git@github.com:rokath/cobs.git` for example.
* Run `go vet ./...` to check build environment.
* Run `go test ./...` to perform tests. Can take a minute.

* Add [cobs.c](./cobs.c) to your embedded project and use function `COBSEncode` to convert a buffer into COBS format. Or use function `COBSDecode` for the other direction.
* After transmitting one (or more) COBS package(s) transmit a 0-delimiter byte.
* Encoding is currently not implemented in **Go**, but its is no big deal to write an encoder in **Go** or an other language of your choice using the documentation.
* Anyway you can encode inside **Go** using function CEncode, which is a **Go** wrapper for the **C** COBSEncode function .
* Contributions are appreciated.

###  4.1. <a name='Prerequisites'></a>Prerequisites

* Just a **C** compiler and, for testing or, if using, a **Go** installation.

###  4.2. <a name='Installation'></a>Installation

* To use COBS with **Go** execute `go get github.com/rokath/cobs`
* In your **Go** file: `include( cobs github.com/rokath/cobs/go )`, when using the **Go** code (decoding only).
* In your **Go** file: `include( cobs github.com/rokath/cobs/c )`, when using the **CGO** code (encoding and decoding).

<!-- ROADMAP -->

###  4.3. <a name='Roadmap'></a>Roadmap

* [x] Add Changelog
* [x] Add back to top links
* [x] Add **Go** Reader & Writer interface
* [ ] Add Additional Templates w/ Examples

See the [open issues](https://github.com/rokath/cobs/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->
##  5. <a name='Contributing'></a>Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

 **Working on your first Pull Request?** You can learn how from this *free* series [How to Contribute to an Open Source Project on GitHub](https://kcd.im/pull-request) 

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
##  6. <a name='License'></a>License

Distributed under the MIT License. See [LICENSE.txt](./LICENSE.txt) for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
##  7. <a name='Contact'></a>Contact

Thomas Höhenleitner - <!-- [@twitter_handle](https://twitter.com/twitter_handle) - --> th@seerose.net
Project Link: [https://github.com/rokath/cobs](https://github.com/rokath/cobs)

<!-- ACKNOWLEDGMENTS -->
###  7.1. <a name='Acknowledgments'></a>Acknowledgments

* [https://pypi.org/project/cobs/](https://pypi.org/project/cobs/)
* [https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing](https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- PROJECT LOGO -->

<div align="center">
  <a href="https://github.com/rokath/cobs">
    <img src="./logo.png" alt="Logo" width="800" height="80">
  </a>

<h3 align="center">COBS</h3>

  <p align="center">
    Common Object Byte Stuffing 
    <br />
    <a href="https://pypi.org/project/cobs/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rokath/cobs/blob/master">View Code</a>
    ·
    <a href="https://github.com/rokath/cobs/issues">Report Bug</a>
    ·
    <a href="https://github.com/rokath/cobs/issues">Request Feature</a>
  </p>
</div>
