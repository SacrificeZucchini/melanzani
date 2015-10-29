Melanzani
=========

Melanzani Converts USB Guitar inputs to MIDI signals. Use a software synthesizer like fluidsynth as backend.


Dependencies
------------

This software requires [Go's wrapper of the portmidi library](http://github.com/rakyll/portmidi).
Before you can install it, you need to install portmidi itself.

### Installing portmidi

You need to build portmidi from source. First, download it from [here](http://sourceforge.net/projects/portmedia/files/portmidi/217/portmidi-src-217.zip/download). Follow the instructions of the README file in the `pm_linux` subfolder. After you have configured everything correctly with ccmake and generated the Makefiles, you have to tweak one of the Makefiles.

  cd cd pm_java/CMakeFiles/pmdefaults_target.dir
  sed -i 's/pm_java\/pm_java/pm_java/' build.make

Afterwards, you can proceed with running `make` as described in the README.

### The Go wrapper

Install the Go wrapper via `go get` as usual.

  go get github.com/rakyll/portmidi

If there is a problem with this command, most likely your portmidi installation is faulty.



Building Melanzani from source
------------------------------

Melanzani is built via `make`.


Starting and stopping Melanzani
------------------

Start melanzani like this: `./melanzani --guitar`. You can pass the device's name after the `--guitar` switch. If you omit it, the name defaults to "Guitar". Other switches are not implemented yet. Stop melanzani by pressing `Ctrl+C`.


Useful tools
------------

Use `evtest` to debug your guitar's input signals (usually requires root privileges). This way, you can also find out your guitar's name.

