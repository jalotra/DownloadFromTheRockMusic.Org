# Project Title

This repository lets you to download all the songs that are present in therockmusic.org/albums/aAlbumName recursively and quite fast as the code is already compiled, written in golang it runs as fast as C++.

## Getting Started

You can just download the precompiled binaries that I have build for Linux arm64 and windows amd64 or you can download the source code and compile it on your own.For compiling on your own version,https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04 this website may prove to be a starting point.

### Prerequisites

For Running the compiled versions you don't need anything just a cli and a keyboard and ctrl c + ctrl v command putting will.
For compiling the source code, you would need the go language ofcourse with some packages that I am not leaving to append later here,  

```
Give examples
# Using the linux precompiled Binary:
./ForLinuxArm64 therockmusic.orgAlbumURL foldername
Example : ./ForLinuxArm64 http://therockmusic.org/albums/promises/index.html promises 

# Using the windows precompiled Binary:
./ForWindowsAmd64.exe therockmusic.orgAlbumURL foldername
Example : ./ForWindowsAmd64.exe http://therockmusic.org/albums/promises/index.html promises 

```

### Installing
#Getting the precompiled binaries:
open https://github.com/jalotra/DownloadFromTheRockMusic.Org.
Point to ./src/PrecompiledBinaries/
Make a folder to save the file. 
Download the file according to your architecture and os.
Run the command described above.

If you run all these steps a *songs folder will be made in the ../ it will have the <--foldername> folder in it that will contain all your songs.Enjoy!!! 


## Running the tests

Nothing written till now.


## Contributing

Not wriiten but I would like to add a few features.
1. To print all the songs available and to download what songs that I like to download.
2. Code-Review, because I started learning Go 3 DAYS back and this project is the staring stone.



## Authors

* **Shivam Jalotra** - (https://shivamjalotra.me)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* cheggaaa for his downloader. 
* Inspiration was me/myself, also google for writing a great language. 
* Last but not the least, listen to this song to get into Go.Must Watch :: https://www.youtube.com/watch?v=LJvEIjRBSDA