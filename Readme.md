# Project Title

Desktop app to convert image files from a format to another.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Usage

You can find the latest windows executable on the [Latest Tag][latest-tag]. For other architectures you may need to (clone and build the aource code)[#development-prerequisites].

The usage is simple, you choose one or more files to convert to the selected format, and then choose where to save.
If you choose multiple files, an zip with all the results will be saved.

#### Allowed source formats

- JPEG
- PNG
- Bitmap (.bmp)
- WebP
- Tiff

#### Allowed destination formats

- JPEG
- PNG
- Bitmap
- Tiff
- WebP

> I'll try to add WebP too

### Development Prerequisites

Before you get a running example of this project, you'll need to install [Wails](https://wails.io) on your machine

```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Then check if you have Wails dependencies installed and your system is ready to use it

```
wails doctor
```

### Installing

If you already have Wails CLI installed, all you need to do is clone the project and run it

```
git clone git@github.com:Hilson-Alex/image_conversor.git
wails dev
```

### Building

To build the artifact you just need to run the build command

```
wails build
```

## Built With

* [Wails](https://wails.io) - Tool to create desktop apps with Go and Web technologies

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Hilson-Alex/image_conversor/releases). 

## Author

- [Hilson A. W. Junior][Hilson-Alex] - *Initial work*


## License

This project is under the MIT License - for more info read the [LICENSE](LICENSE) file

[Hilson-Alex]: https://github.com/Hilson-Alex
[latest-tag]: https://github.com/Hilson-Alex/image_conversor/releases/tag/1.1.0
