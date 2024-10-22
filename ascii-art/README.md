# ascii-art-fs

### Objectives

Ascii-art is a program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the `string` received using ASCII characters, as you can see in the example below:

-   This project should handle an input with numbers, letters, spaces, special characters and `\n`.
-   Take a look at the ASCII manual.

### fonts

-   the available fonts are:
    -   [shadow](shadow.txt)
    -   [standard](standard.txt)
    -   [thinkertoy](thinkertoy.txt)

### Usage

$ go run . "hello" standard | cat -e
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/


$ go run . "Hello There!" shadow | cat -e

_|    _|          _| _|             _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|          _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|        _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|        _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|          _|     _|    _|   _|_|_| _|         _|_|_| _|


$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $

### License

Ascii-art is distributed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Credits

Ascii-art-fs was created by:-

-   [nhussain](https://learn.reboot01.com/git/nhussain)
-   [yabuzuha](https://learn.reboot01.com/git/yabuzuha)
-   [etarada](https://learn.reboot01.com/git/etarada)

### About the License

Ascii-art is licensed under the [MIT License](LICENSE). You are free to use, modify, distribute, and sublicense the software. Refer to the license file for more information.
"# ascii-art-fs" 
