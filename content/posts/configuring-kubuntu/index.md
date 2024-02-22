---
date: 2023-11-07T10:35:52+11:00
title: "Configuring Kubuntu"
description: "Until recently, Arch Linux has been my daily driver. Here's how I configure my new Kubuntu installation."
tags:
  [
    "Linux",
    "Kubuntu",
    "Ubuntu",
    "KDE Plasma",
    "Flatpak",
    "Flathub",
    "Discover",
    "Defaults",
    "Editor",
    "Vim",
    "Nano",
    "Shell",
    "Bash",
    "Zsh",
    "OhMyZsh",
  ]
# author: ["Toby Scott", "Other example contributor"]
hidden: false
draft: false
---

Until recently, Arch Linux has been my daily driver. I love the rolling release model and you really can't fault the AUR. But always getting the latest features and patches is a double edged sword because along with those benefits come the latest bugs and issues. I considered the benefits to outweigh the short comings up until one dreadful day...

```bash
# Updating my system with an AUR helper as per normal
paru

# Things started playing up so I restarted my machine
sudo reboot now
```

And my system was bricked. Yes, I could've chroot-ed in to find and fix the problem, but I just want a machine that works, so I called it there and installed Kubuntu 23.10.

## Why Kubuntu?

I've had a very stable experience in the past with Ubuntu and although the GNOME desktop is stunning, I prefer the KDE Plasma desktop for a couple of reasons:

- The Wayland session fixes all scaling issues I've ever faced running a 4k monitor next to a 1080p monitor 🎉
- I am a fan of the Windows-like start menu and the Mac-like global menu. Plasma ships both of these out of the box!

## Setting up fresh installation of Kubuntu 23.10

#### Installing some basics

Let's install ZSH and some fundamental CLI tools:

```bash
sudo apt install zsh curl vim ranger neofetch htop
```

#### Setting the default editor to Vim

I don't know why Canonical chose Nano as the default editor but Vim is far surperior (IMO), so let's set that as the default editor.

```bash
sudo update-alternatives --config editor
```

#### Installing Oh My Zsh + Plugins

Once you've used [Oh My Zsh](https://github.com/ohmyzsh/ohmyzsh), you really can't go back to plain ol' bash.

```bash
# Install OhMyZsh, follow it's prompts to make zsh the default shell
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Clone some plugins to the zsh plugins directory
git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting $ZSH_CUSTOM/plugins/zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-history-substring-search $ZSH_CUSTOM/plugins/zsh-history-substring-search
```

Now let's configure zsh, edit `~/.zshrc`

1. Find this line and add the missing plugins, some are there by default:

```
plugins=(git zsh-autosuggestions zsh-syntax-highlighting history-substring-search history aliases sudo themes docker nmap kubectl)
```

2. Change `ZSH_THEME` from `bobbyrussel` to `af-magic`:

```env
ZSH_THEME="af-magic"
```

#### Installing Flatpak

Canonical is trying to push their Snapstore, but Flatpaks are adopted by most other distributions and I prefer to use something that's more widely used.

I'll still use `.deb` packages (installed with apt/dpkg) for CLI tools, but Flatpak is now undeniably the better option for GUI desktop programs.

```bash
# Install Flatpak
sudo apt install flatpak

# Install the Flatpak backend for Discover (KDE Plasma's Software Store)
sudo apt install plasma-discover-backend-flatpak

# Add the main Flathub repository. Flathub is where most companies publish their official Flatpak releases (Discord, for one example)
flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo

# Restart your system
sudo reboot now
```

Once your system has restarted, open Discover, go to Settings and make Flatpak the default source.

Next time you're looking for a GUI program such as Spotify, Discord, VLC, etc you can just search for it in Discover and install it from there. You'll get a break down of all the permissions that Flatpak app requires, and more. You can also see all this information (including whether or not the Flatpak is verified as the official Flatpak published by the official company) on the Flathub website: https://flathub.org/

#### Installing software in a fresh Kubuntu installation

I usually install Flatpaks (and Snaps in very rare cases) through Discover. But I keep this list of install commands handy to quickly install everything I need after a fresh system install:

```bash
# Daily use
flatpak install flathub com.google.Chrome -y
flatpak install flathub com.brave.Browser -y
flatpak install flathub com.spotify.Client -y
flatpak install flathub com.discordapp.Discord -y
flatpak install flathub org.videolan.VLC -y

# Productivity
flatpak install flathub org.mozilla.Thunderbird -y
flatpak install flathub md.obsidian.Obsidian -y
flatpak install flathub org.gimp.GIMP -y
flatpak install flathub org.kde.krita -y
flatpak install flathub com.obsproject.Studio -y

# Development
sudo snap install code --classic # Snap, ew
flatpak install flathub com.getpostman.Postman -y
flatpak install flathub io.dbeaver.DBeaverCommunity -y

# Gaming
flatpak install flathub com.valvesoftware.Steam -y
flatpak install flathub org.DolphinEmu.dolphin-emu -y

# System tools
flatpak install flathub org.flameshot.Flameshot -y

```
