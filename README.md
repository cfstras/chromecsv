ChromeCSV [![Build Status](https://travis-ci.org/cfstras/chromecsv.svg?branch=master)](https://travis-ci.org/cfstras/chromecsv)
=========

ChromeCSV is a simple tool for exporting the passwords you saved in Google Chrome / Chromium to a plain text CSV file. You can then import the contents of that file into your [PassDeposit][passdep] account or do whatever you like to it!


[passdep]: https://www.passdeposit.com/

Features
--------

- Works on Windows and Linux
- Decrypts passwords if necessary
- Save logins to CSV file

Usage
-----

No need to Compile or Install! Simply grab the latest release from here:
[Downloads & Releases][releases]

First, close all Chrome windows. Yes, all of them. Even this one. Then, open a console where you downloaded it and type the following, then hit <kbd>Enter</kbd>.

```bash
# on linux
chromecsv ~/.config/google-chrome/
```

```bat
:: on windows
chromecsv "%LOCALAPPDATA%\Google\Chrome\User Data"
```
If you don't want to memorize the path, you can open the console and type the command, then close Chrome and _then_ press <kbd>Enter</kbd>.

Now, you should have a new file named `passwords.csv`. Open it (when noone is around :smile:) and cheer in happiness!

[releases]: https://github.com/cfstras/chromecsv/releases

Compiling from Source
---------------------

You'll need a working [Go][golang] environment and libsqlite3-dev.

On Ubuntu, for example, you can get that up and running like this:

```bash
sudo apt-get install libsqlite3-dev golang
```

Now for the fun part:

```bash
go get github.com/mattn/go-sqlite3 # get sqlite3-bindings for golang
go get github.com/cfstras/chromecsv # get and install chromecsv
```

[golang]: http://golang.org/

Details
-------

Chrome saves login data into an SQLite3 database called _Login Data_, which is stored in your `%LOCALAPPDATA%\Google\Chrome\User Data\Default\` (on Windows) or `~/.config/google-chrome/Default/`.

On Windows, the passwords are encrypted through the [CryptProtectData][protect] WinApi function, which derives a key from your logon credentials so only you can decrypt it again. ChromeCSV then uses [CryptUnprotectData][unprotect] to get back your passwords.
(Ironically, the documentation uses the word _typically_ whenever explaining _who_ in particular can decrypt the data.)

On Linux, per Default, no encryption/decryption is done, so no decryption is necessary. (See [this ticket][masterpw] or this fun [comment in the code][code].) You can force it to use Gnome Wallet or KDEWallet with a command-line flag, but this won't migrate your data.


[unprotect]: http://msdn.microsoft.com/en-us/library/windows/desktop/aa380882(v=vs.85).aspx
[protect]: http://msdn.microsoft.com/en-us/library/windows/desktop/aa380261(v=vs.85).aspx
[masterpw]: https://code.google.com/p/chromium/issues/detail?id=53
[code]: https://code.google.com/p/chromium/codesearch#chromium/src/chrome/browser/password_manager/login_database_posix.cc&l=9

License
------- 

This software is released under the MIT license. For details, see [LICENSE.md][license]

[license]: https://github.com/cfstras/chromecsv/blob/master/LICENSE.md
