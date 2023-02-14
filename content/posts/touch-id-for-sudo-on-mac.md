---
date: 2023-02-08T15:04:17+11:00
title: "How to configure Touch ID to authorise sudo commands on macOS"
description: "Configuring Touch ID to authorise sudo commands on macOS (at a system-wide scope)"
tags: ["macos", "sudo", "touch id", "macintosh", "apple", "unix", "configuration", "terminal", "shell", "command line", "cli"]
hidden: false
---

The other day I was watching one of [NetworkChuck's recent videos](https://www.youtube.com/watch?v=qOrlYzqXPa8) where he was running through a bunch of shell commands for Mac, and one stuck out to me in particular... The ability to configure Touch ID to authorise `sudo` commands! 🤯

Edit the following file:

```bash
sudo vim /etc/pam.d/sudo
```

Add this line `auth sufficient pam_tid.so` to the top of the file, below the comment. It should look something like:

```conf
# sudo: auth account password session

auth sufficient pam_tid.so

auth       sufficient     pam_smartcard.so
auth       required       pam_opendirectory.so
account    required       pam_permit.so
password   required       pam_deny.so
session    required       pam_permit.so
```

Now close that Terminal, open a fresh one, run a `sudo` command and you'll be prompted to use your Touch ID! 🎉

I was concerned this may cause issues when running `sudo` commands on another server via SSH, but I tested it and was pleased to find it didn't interfere!

Although, one cavieat I noticed was this wouldn't work while docked to a Lenovo DisplayLink, which luckily, I was able to fix with one command from a [StackExchange answer](https://apple.stackexchange.com/a/444202/465731)!

```bash
defaults write com.apple.security.authorization ignoreArd -bool TRUE
```

Now I'm used to this, I really notice having to type my password for sudo commands on my desktop! 😂
