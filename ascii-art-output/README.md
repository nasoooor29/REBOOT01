# ascii-art

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

```console
student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$

student$ go run . "Hello There" | cat -e
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $


student$ go run . "Hello\n\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
 _______   _                           $
|__   __| | |                          $
   | |    | |__     ___   _ __    ___  $
   | |    |  _ \   / _ \ | '__|  / _ \ $
   | |    | | | | |  __/ | |    |  __/ $
   |_|    |_| |_|  \___| |_|     \___| $
                                       $
                                       $
student$
```

### License

Ascii-art is distributed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Credits

Ascii-art was created by:-

-   [nhussain](https://learn.reboot01.com/git/nhussain)
-   [yabuzuha](https://learn.reboot01.com/git/yabuzuha)
-   [etarada](https://learn.reboot01.com/git/etarada)

### About the License

Ascii-art is licensed under the [MIT License](LICENSE). You are free to use, modify, distribute, and sublicense the software. Refer to the license file for more information.
"# ascii-art-output" 
