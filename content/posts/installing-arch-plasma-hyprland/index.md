---
date: 2024-03-25T18:04:33+11:00
title: "Installing Arch Linux with KDE Plasma and Hyprland"
description: "..."
tags:
  [
    "linux",
    "archlinux",
    "installation",
    "wayland",
    "hyprland",
    "window manager",
    "kde",
    "plasma",
    "desktop",
    "networking",
    "bluetooth",
    "pacman",
    "foss",
  ]
# author: ["Toby Scott", "Other example contributor"]
hidden: true
draft: true
cover:
  image: "cover.png"
  alt: "Arch Linux logo"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

## Archinstall script

Use archinstall script.

Install with KDE profile, Pipewire, Network Manager, Grub, BTRFS.

## Install Zsh and Oh My Zsh

```bash
# ...
```

## Package management

I prefer to use Pacman for system packages, Homebrew for user packages, and Flatpak for sandboxed GUI applications. Sometimes I may need a package that doesn't exist in any of the aforementioned package managers, so I also install Paru to install AUR packages.

#### Install Flatpak

For sandboxed GUI applications.

```bash
sudo pacman -S flatpak
```

Open Discover, you should be able to install flatpaks now.

#### Install Homebrew

For user CLI utilities.

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

#### Install Paru

For sytem packages only available in AUR.

```bash
git clone https://aur.archlinux.org/paru.git
cd paru
makepkg -si
cd ../
rm -rf paru
```

## Installing Docker

```bash
sudo pacman -S docker docker-compose
```

then the post-installation instructions on Docker's website.

## System Utilities

Install `bluez` and `bluez-utils`, then start and enable `bluetooth.service`.

Install `cups` for printers

Install `ufw` for firewall

Install `fwupd` from pacman for firmware security checker in Info Centre

Japanese font stuff
`sudo pacman -S noto-fonts-cjk noto-fonts-emoji noto-fonts`

Install `man`
`sudo pacman -S man-db`

Install brightness controller
`sudo pacman -S brightnessctl`

## Hyprland

```bash
sudo pacman -S waybar
```

for waybar:

```bash
sudo pacman -S ttf-font-awesome
```

install nerd fonts group to pre-emptively fix issues as lots of things use nerd fonts:

```bash
sudo pacman -S nerd-fonts
```
