# COBS

<!-- PROJECT LOGO -->

<div align="center">
  <a href="https://github.com/rokath/COBS">
    <img src="./images/logo.png" alt="Logo" width="800" height="80">
  </a>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>

<!-- vscode-markdown-toc -->

- [COBS](#cobs)
  - [1. <a name='AboutTheproject'></a>About The project](#1-about-the-project)
  - [2. <a name='COBSSpecification'></a>COBS Specification](#2-cobs-specification)
  - [3. <a name='COCScode'></a>COBS Implementation](#3-cobs-implementation)
  - [4. <a name='GettingStarted'></a>Getting Started](#4-getting-started)
    - [4.1. <a name='Prerequisites'></a>Prerequisites](#41-prerequisites)
    - [4.2. <a name='Installation'></a>Installation](#42-installation)
    - [4.3. <a name='Roadmap'></a>Roadmap](#43-roadmap)
  - [5. <a name='Contributing'></a>Contributing](#5-contributing)
  - [6. <a name='License'></a>License](#6-license)
  - [7. <a name='Contact'></a>Contact](#7-contact)
    - [7.1. <a name='Acknowledgments'></a>Acknowledgments](#71-acknowledgments)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc --><div id="top"></div>

  </ol>
</details>

<!-- ABOUT THE PROJECT -->
##  1. <a name='AboutTheproject'></a>About The project

This is a [COBS](https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing) implementation in **Go** and **C** with tests.

* The **C** code is copied from Wikipedia and adapted. Some other implementations did not pass all tests.
* The **Go** code was written from scratch after reviewing other implementations.

##  2. <a name='COBSSpecification'></a>COBS Specification

* See [https://pythonhosted.org/cobs/cobs.cobs.html](https://pythonhosted.org/cobs/cobs.cobs.html).

<p align="right">(<a href="#top">back to top</a>)</p>

##  3. <a name='COCScode'></a>COBS Implementation
  
* See [./pkg/cobs/](./pkg/cobs/) folder

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
##  4. <a name='GettingStarted'></a>Getting Started

* Run `go vet ./...` to check build environment.
* Run `go test ./...` to perform tests. Can take a few seconds.

* Add [./pkg/cobs/cobs.c](./pkg/cobs/cobs.c) to your embedded project and use function `COBSEncode` to convert a buffer into COBS format.
* After transmitting one (or more) COBS package(s) transmit a 0-delimiter byte.
* Encoding is currently not implemented in **Go**, but its is no big deal to write an encoder in **Go** or an other language of your choice using the documentation.
* Anyway you can encode inside **Go** using function CEncode, which is a **Go** wrapper for the **C** COBSEncode function .
* Contributions are appreciated.

<p align="right">(<a href="#top">back to top</a>)</p>

###  4.1. <a name='Prerequisites'></a>Prerequisites

* Just a **C** compiler and, for testing, a **Go** installation.

<p align="right">(<a href="#top">back to top</a>)</p>

###  4.2. <a name='Installation'></a>Installation

* To use COBS with **Go** execute `go get github.com/rokath/cobs`

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->
###  4.3. <a name='Roadmap'></a>Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [ ] Add Additional Templates w/ Examples

See the [open issues](https://github.com/rokath/COBS/issues) for a full list of proposed features (and known issues).

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

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
##  6. <a name='License'></a>License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
##  7. <a name='Contact'></a>Contact

Thomas Höhenleitner - <!-- [@twitter_handle](https://twitter.com/twitter_handle) - --> th@seerose.net
Project Link: [https://github.com/rokath/COBS](https://github.com/rokath/COBS)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
###  7.1. <a name='Acknowledgments'></a>Acknowledgments

* [https://pypi.org/project/cobs/](https://pypi.org/project/cobs/)
* [https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing](https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- PROJECT LOGO -->

<div align="center">
  <a href="https://github.com/rokath/COBS">
    <img src="./images/logo.png" alt="Logo" width="800" height="80">
  </a>

<h3 align="center">COBS</h3>

  <p align="center">
    Common Object Byte Stuffing 
    <br />
    <a href="https://pypi.org/project/cobs/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rokath/COBS/blob/master/pkg/cobs">View Code</a>
    ·
    <a href="https://github.com/rokath/COBS/issues">Report Bug</a>
    ·
    <a href="https://github.com/rokath/COBS/issues">Request Feature</a>
  </p>
</div>
