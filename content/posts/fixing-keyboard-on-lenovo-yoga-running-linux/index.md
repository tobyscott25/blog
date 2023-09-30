---
date: 2023-01-14T10:28:45+11:00
title: "How to fix an unresponsive keyboard on Lenovo Yoga running Linux"
description: "Updating GRUB config to solve an unresponsive keyboard issue on a Lenovo Yoga Slim 7 Pro"
tags: ["Linux", "Ubuntu", "Lenovo", "Keyboard", "Fix", "GRUB", "unresponsive"]
# author: ["Toby Scott", "Other example contributor"]
hidden: false
draft: false
cover:
  image: "cover.jpg"
  alt: "Typing on a keyboard"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

I recently installed Ubuntu Linux on a Lenovo Yoga Slim 7 Pro, although I had to connect an external keyboard to go through the installation process because the built-in keyboard was completely unresponsive!

Luckily I came across [this Stack Overflow answer](https://askubuntu.com/questions/1352604/ubuntu-20-04-keyboard-not-working-on-lenovo-yoga-slim-7i-pro) which did the trick for me.

Once the installation was complete I needed to edit the GRUB config to fix this.

```bash
sudo vim /etc/default/grub
```

Add these parameters to `GRUB_CMDLINE_LINUX` like so:

```
GRUB_CMDLINE_LINUX="atkbd.reset=1 i8042.nomux=1 i8042.reset=1 i8042.nopnp=1 i8042.dumbkbd=1"
```

Once added, save and exit, then update your grub and reboot, like so:

```bash
sudo update-grub && reboot
```

I'm not sure if all of these paramters were required in my case, but I never bothered to try them one at a time or different combinations. If it works, it works.

The only thing I notice now is the light on the caps-lock key never turns on, but I can live with that 😉
