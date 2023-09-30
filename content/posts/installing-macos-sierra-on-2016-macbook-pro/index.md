---
date: 2023-01-01T23:11:06+11:00
title: "How to extract an old macOS installer from PKG on a newer version of macOS"
description: "I have an old MacBook running Ubuntu and I decided to reinstall macOS on it. Should be easy right?"
tags:
  [
    "mac",
    "macos",
    "apple",
    "macbook",
    "intel",
    "m1",
    "sierra",
    "bootable",
    "recovery",
    "pkgutil",
    "appstore",
  ]
hidden: false
cover:
  image: "cover.png"
  alt: "MacOS Sierra"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

## Background

I have an old 2016 Intel MacBook Pro that I've had running Ubuntu for a while but could never get the sound to work in Linux so I decided to reinstall macOS on it. Should be easy right? Wrong...

Instead of setting up a dual boot situation, I _cleverly_ decided to wipe the whole hard drive in favour of Ubuntu, recovery partitions and all. But no matter, Intel Macs can boot into an internet recovery system without the recovery partition by holding ⌘ R on power. So that I did, I wiped the hard drive with Disk Utility, and started a fresh install of macOS.

## Installing macOS Sierra via Internet Recovery failing

The installation preparation would take forever, the time remaining would go into the negatives (`-16 seconds remaining`) and finally it would fail. I tried this over and over and again, with the same result each time.

Checking the install log (with ⌘ L) I noticed this message repeated hundreds of times with the percentage slowly incrementing by tiny, tiny bits...

```
Dec 29 16:10:32 MacBook-Pro storedownloadd[497]: Commerce:sending status (macOS Sierra): 0.000126% (474623.067036)
Dec 29 16:10:34 MacBook-Pro storedownloadd[497]: Commerce:sending status (macOS Sierra): 0.000126% (474623.067036)
Dec 29 16:10:35 MacBook-Pro storedownloadd[497]: Commerce:sending status (macOS Sierra): 0.000132% (474623.067036)
Dec 29 16:10:37 MacBook-Pro storedownloadd[497]: Commerce:sending status (macOS Sierra): 0.000132% (474623.067036)
```

After some digging I found that I was not the only one having this issue installing macOS Sierra (10.12) via Internet Recovery, so I decided to make a bootable USB and install from that.

I found [this official article on creating a bootable installer for macOS](https://support.apple.com/en-us/HT201372) which instructs you to download the installer you require from [this official download page](https://support.apple.com/en-gb/HT211683), and use the `createinstallmedia` binary inside the installer, like so:

```bash
sudo /Applications/Install\ macOS\ High\ Sierra.app/Contents/Resources/createinstallmedia --volume /Volumes/MyVolume
```

Except to actually use the `createinstallmedia` utility, you have to extract the actual installer from the downloaded file first.

## Extracting the older installer

So I downloaded macOS Sierra (10.12) but I cannot believe the complexity of it... the download is actually a `.dmg` file, that then contains a `.pkg` file, which you _THEN_ need to run (with Installer.app), which will _THEN_ copy the _ACTUAL_ installer into your Applications folder, but _ONLY IF_ you are running a version of macOS that is _OLDER_ than the version you downloaded! Honestly, Apple, what the heck...

So trying to run this `.pkg` file on my 2021 M1 MacBook Pro, I get this message: `This package will run a program to determine if the software can be installed.` Followed by this error: `This version of macOS 10.12.6 cannot be installed on this computer.`. But I don't want to _install_ it on this computer, I just want the installer itself!!

I figured the installer I need must exist in this `.pkg` file somewhere so I tried to inspect the package contents, but the usual 'Show package contents' option of the context menu was missing! Alrighty then, let's try this manually..

```bash
pkgutil --expand ~/Desktop/InstallOS.pkg ~/Desktop/sierrapkg
```

Cool, that worked, and in the contents there's a file named `Distribution` which I noticed is just an XML file (no file name extension) with some JavaScript in it that looks like it defines some machine and OS version requirements...

So where it defines an error code and returns `false` in some spots, I added `return true;` before it did so, and finally rebuilt it into a `.pkg` file again, like so:

```bash
pkgutil --flatten ~/Desktop/sierrapkg ~/Desktop/ModifiedInstallOS.pkg
```

Now when I run this modified `.pkg` file (with Installer.app) on my 2021 M1 MacBook Pro, it happily copies the full `Install macOS Sierra.app` file into my Applications folder! 🎉

## Creating the bootable installation media

Once I had the Sierra installer, I went back to [creating a bootable installer for macOS](https://support.apple.com/en-us/HT201372) but the command for macOS Sierra (10.12) is missing! We have the command for macOS El Capitan (10.11):

```bash
sudo /Applications/Install\ OS\ X\ El\ Capitan.app/Contents/Resources/createinstallmedia --volume /Volumes/MyVolume --applicationpath /Applications/Install\ OS\ X\ El\ Capitan.app
```

And for macOS **_High_** Sierra (10.13) and onwards:

```bash
sudo /Applications/Install\ macOS\ High\ Sierra.app/Contents/Resources/createinstallmedia --volume /Volumes/MyVolume
```

But macOS Sierra (10.12) is just missing. 😲 That's it. I'm giving up on Sierra (10.12) and skipping straight to **_High_** Sierra (10.13).

So I went back to the [downloads page](https://support.apple.com/en-gb/HT211683) to realise that High Sierra (10.13) onwards aren't direct downloads, they're all just links to the Mac App Store.

> Interesting note: Lion (10.7), Mountain Lion (10.8), Yosemite (10.10), El Capitan (10.11) and Sierra (10.12) are available as direct downloads, but Mavericks (10.9) is missing from the list of direct downloads 🤔

No problem, I'll just download High Sierra Installer from the Mac App Store. Oh wait, it doesn't let me, and gives me nothing more than a useless message that just states "Not compatible with this device" 😡

I figured it could be due to the chip architecture because this OS version was released before the Apple Silicon chip, so I tried opening this link on my brother's 2017 Intel MacBook Air (running macOS Monterey), and it let me download it! Lucky guess. I plugged my thumbdrive into his laptop and ran the command to create the bootable installation media from there.

```bash
sudo /Applications/Install\ macOS\ High\ Sierra.app/Contents/Resources/createinstallmedia --volume /Volumes/MyVolume
```

## Conclusion

I was able to boot the 2016 Intel MacBook from the bootable High Sierra Installation media and succefully install a fresh version of macOS High Sierra. It also re-created the recovery partitions, etc. And I was then able to update it to the latest macOS version it could run (macOS Monterey). 🎉

I'm lost as to why Apple makes it so damn hard to execute a simple task like this. I know that in theory, the installation over Internet Recovery would've worked, but it didn't.

I understand preventing people from installing macOS on third party hardware (although I disagree with it), but come on Apple, why do you need to make the downloading and creation of bootable installation medias so ridiculous! Prevent it from running on third party hardware all you like, but just release some simple `.iso` files that can be flashed to a thumb drive, like every other operating system out there!
