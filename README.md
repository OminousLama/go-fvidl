
<h1 align="center">
  <br>
  <a href="https://github.com/OminousLama/go-fvidl"><img src="./docs/res/fvidl-icon.svg" alt="fvidl" width="200"></a>
  <br>
  Find Videos by Length
  <br>
</h1>

<h4 align="center">CLI tool made to quickly locate video files based on duration.</h4>

![Demo](https://raw.githubusercontent.com/OminousLama/go-fvidl/dev/docs/res/demo-video.gif)

## Getting started

### Prerequisites

- ffmpeg
```bash
# Fedora
sudo dnf install ffmpeg

# Debian
sudo apt install ffmpeg
```

### Install / Download

Grab the latest release from the [release page](https://github.com/OminousLama/go-fvidl/releases/latest).

### Usage

Use `fvidl -help` to show a list of available options.

#### Examples

```bash
fvidl -min 5 -max 30 -d /my/directory/ -ft mp4,mov -r
```

This command will search for video files...

- (`-d`) ...in the directory `/my/directory/`
- (`-min`) ...longer than or equal to 5 seconds
- (`-max`) ...shorter or equal to 30 seconds
- (`-ft mp4,mov`) ...of type `mp4` or `mov`
- (`-r`) ...recursively

## Credits

This software uses the following open source packages:

- [Best README Template](https://github.com/othneildrew/Best-README-Template)
- [Go](https://github.com/golang)
- [Remix Icon](https://remixicon.com/)

## License

GPL-3.0 license
