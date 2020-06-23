<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/FrancoRivera/sample-go-gin-tailwind">
    <img src="assets/static/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Sample Go App</h3>

  <p align="center">
    This is a sample Go (golang) app that uses the Gin router, pq, and Gorm to do a basic search
    <br />
    <a href="https://github.com/FrancoRivera/sample-go-gin-tailwind"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/FrancoRivera/sample-go-gin-tailwind">View Demo</a>
    ·
    <a href="https://github.com/FrancoRivera/sample-go-gin-tailwind/issues">Report Bug</a>
    ·
    <a href="https://github.com/FrancoRivera/sample-go-gin-tailwind/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

[![Sample Go Gin Tailwind Screenshot][product-screenshot]](https://example.com)

### Built With

* [Go](https://golang.org)
* [Tailwind CSS](https://tailwindcss.com)

##### Go Packages

* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://github.com/go-gorm/gorm)
* [pq](https://github.com/lib/pq)


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Go
* PostgreSQL

```sh
# MAC OS X (after installing homebrew)
brew install golang
brew install postgresql
```

### Installation

1. Clone the sample-go-gin-tailwind repo
```sh
git clone https://github.com/FrancoRivera/sample-go-gin-tailwind.git && cd sample-go-gin-tailwind
```
2. Install dependencies
```sh
go install
```

3. Setup database
```sh
createuser -P -d -R -S --interactive
# username: houses
# password: chose_a_password
createdb houses -o houses
psql -U houses -f script.sql
```

<!-- USAGE EXAMPLES -->
## Usage

Run the webserver with

```sh
PG_USER=db_username PG_PASSWORD=db_password go run main.go
```

Compile binaries for all architectures with 

```sh
chmod +x go-executable-build.bash	# make script runnable
./go-executable-build.bash 	  	# run script
```



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/FrancoRivera/sample-go-gin-tailwind/issues) for a list of proposed features (and known issues).



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

Franco R - [@francodalima](https://twitter.com/francodalima) - email

Project Link: [https://github.com/FrancoRivera/sample-go-gin-tailwind](https://github.com/FrancoRivera/sample-go-gin-tailwind)



<!-- ACKNOWLEDGEMENTS -->
<!-- ## Acknowledgements -->

<!-- * [Me]() -->



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/FrancoRivera/sample-go-gin-tailwind.svg?style=flat-square
[contributors-url]: https://github.com/FrancoRivera/sample-go-gin-tailwind/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/FrancoRivera/sample-go-gin-tailwind.svg?style=flat-square
[forks-url]: https://github.com/FrancoRivera/sample-go-gin-tailwind/network/members
[stars-shield]: https://img.shields.io/github/stars/FrancoRivera/sample-go-gin-tailwind.svg?style=flat-square
[stars-url]: https://github.com/FrancoRivera/sample-go-gin-tailwind/stargazers
[issues-shield]: https://img.shields.io/github/issues/FrancoRivera/sample-go-gin-tailwind.svg?style=flat-square
[issues-url]: https://github.com/FrancoRivera/sample-go-gin-tailwind/issues
[license-shield]: https://img.shields.io/github/license/FrancoRivera/sample-go-gin-tailwind.svg?style=flat-square
[license-url]: https://github.com/FrancoRivera/sample-go-gin-tailwind/blob/master/LICENSE.txt
[product-screenshot]: assets/static/screenshot.png
