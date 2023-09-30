---
date: 2021-11-22T16:05:14+11:00
title: "Increasing VirtualBox's video memory limit"
description: "How to increase VirtualBox's default video memory limit of 128Mb"
tags: ["virtualbox", "vram", "video memory", "virtualisation", "vm"]
cover:
  image: "cover.jpg"
  alt: "Virtualisation"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

VirtualBox seems to have a tiny limit on VRAM, luckily there's a way around it!

Open your terminal and run the following, with your VM name in the quotation marks.

```sh
vboxmanage modifyvm "Arch Linux VM" --vram 256
```

Now open VirtualBox and you will see the amount of video memory has been increased to 256Mb, double the default limit. You can go as high as your hardware allows :)
