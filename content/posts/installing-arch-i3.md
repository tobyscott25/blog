---
date: 2021-11-18T22:14:38+11:00
title: "Installing Arch Linux and i3WM"
description: 'Installing the famously "bare-bones" Arch Linux along with the i3 Window Manager'
tags: ["linux", "archlinux", "installation", "x display server", "i3 window manager", "networking", "bluetooth", "pacman", "foss"]
---


## 1. Installing the base system

Boot from the installation media and run the installer with the following command:
```
# archinstall guided
```
(Fun fact: This is a python script so `# python -m archinstall guided` works too)
- *When prompted to select a network interface to configure*, choose `Use NetworkManager to control and manage your internet connection`.
- *When prompted to reboot*, shutdown your computer, remove the installation media, and turn on the PC.

Login to your new system and update all packages using this command:
```
sudo pacman -Syyu
```


## 2. Set the timezone and locale (to Melbourne, Australia)

Set the timezone with the following command
```
sudo ln -sf /usr/share/zoneinfo/Australia/Melbourne /etc/localtime
```
Run `hwclock` to generate `/etc/adjtime`
```
sudo hwclock --systohc
```
Install a CLI text editor
```
sudo pacman -S vim
```
Now we tell the system what locales it needs to generate by editing `/etc/locale.gen` (as superuser)
```
sudo vim /etc/locale.gen
```
**Uncomment the following lines** in `/etc/locale.gen`
```
en-AU.UTF-8 UTF-8
en-AU ISO-8859-1
```
Now generate the locales from `/etc/locale.gen`
```
sudo locale-gen
```
List the available locales to check they were generated correctly:
```
localectl list-locales
```
Set the system locale. To do this, write the `LANG` variable to `/etc/locale.conf` where `en_US.UTF-8` belongs to the first column of an uncommented entry in `/etc/locale.gen`. This command will do the trick, nice and easy:
```
sudo localectl set-locale LANG=en_AU.UTF-8
```
Once the system and user `locale.conf` files have been created or edited, their new values will only take effect in the next session. But to have the current session use the new settings, unset `LANG` and then source `/etc/profile.d/locale.sh`
```
unset LANG
source /etc/profile.d/locale.sh
```
Finally, check the current locale
```
locale
```


## 3. Installing an AUR Helper

The official Arch repo is great, but it doesn't have everything... So we have the *Arch User Repository (AUR)* which is a community-driven repository for Arch users.

Read more here: https://wiki.archlinux.org/title/Arch_User_Repository

Install basic CLI packages required for the next steps:
```
sudo pacman -S base-devel git
```
Install `paru`, an AUR helper:
```
git clone https://aur.archlinux.org/paru.git
cd paru
makepkg -si
```
Note: when you run `paru` without any options, it will automatically add the `-Syu` flags and update your system. This is equivalent to `sudo pacman -Syu`


## 4. Installing a GUI

Using the [CLI](https://en.wikipedia.org/wiki/Command-line_interface) can get boring quick, not to mention you'll surely want to run GUI programs, to view/edit images, watch videos, play games, etc, right?

When it comes to a GUI environment, you can choose between:
- A [desktop environment](https://wiki.archlinux.org/title/Desktop_environment) (such as [GNOME](https://www.gnome.org/), [KDE Plasma](https://kde.org/plasma-desktop/), [XFCE](https://xfce.org/))
- or a [window manager](https://wiki.archlinux.org/title/window_manager) (such as [i3wm](https://i3wm.org/), [dwm](https://dwm.suckless.org/), [awesome](https://awesomewm.org/)).

Desktop environments and window managers then require a [display server](https://en.wikipedia.org/wiki/Windowing_system#Display_server) (such as [Xorg](https://wiki.archlinux.org/title/xorg) or [Wayland](https://wiki.archlinux.org/title/wayland)) to provides the basic framework for a GUI environment: drawing and moving windows on the display device and interacting with a mouse and keyboard.

<!-- Then there are other bits and pieces such as a compositor, etc -->

Install i3wm
```
sudo pacman -S i3-gaps i3status ttf-dejavu
```
 - `i3-gaps` A maintained fork of i3wm, which I prefer for its added features
 - `i3status` Contents for the status bar so it works and doesn't show an error
 - `ttf-dejavu` A font so the text will display correctly on screen

Install the X Display Server
```
sudo pacman -S xorg xorg-xinit xterm
```

Install the drivers for your hardware. You may not need these at all, it's possible that your system will work out-of-the-box.
```sh
sudo pacman -S nvidia nvidia-utils      # NVIDIA 
sudo pacman -S xf86-video-amdgpu mesa   # AMD
sudo pacman -S xf86-video-intel mesa    # Intel
```

If you are using VirtualBox, you may want to install the VirtualBox guest utilities. **This is completely uncessary on a physical machine.**
```
sudo pacman -S virtualbox-guest-utils
```

Create and edit `~/.xinitrc` so we can make changes to the default config
```
cp /etc/X11/xinit/xinitrc ~/.xinitrc
vim ~/.xinitrc
```

and tell it to execute i3 when the X Server is launched by commenting out the following lines and adding `exec i3`
```sh
#twm &
#xclock -geometry 50x50-1+1 &
#xterm -geometry 80x50+494+51 &
#xterm -geometry 80x20+494-0 &
#exec xterm -geometry 80x66+0+0 -name login
exec i3
```

Exit vim and start the X Server, in turn, starting i3.
```
startx
```

## 5. Connecting to the internet

Running an ethernet cable should work out of the box. But wifi can be a little trickier. The first thing to consider is **whether or not your wifi card is supported by the kernel**.


### Checking the wireless driver status

Run one of these commands to check the kernel has the required drivers
```sh
lspci -k 		# If card is connected via PCI(e)
lsusb -v 		# If card is connected via USB
```
You should see the *network controller* in the list (with the drivers noted)
```
05:00.0	Network controller: Intel Corporation Wi-Fi 6 AX200 (rev 1a)
		Subsystem: Intel Corporation Wi-Fi 6 AX200NGW
		Kernel driver in use: iwlwifi
		Kernel modules: iwlwifi
```
Now run this command to check a wireless interface was created
```
ip link
```

If the kernel doesn't support your wifi card then you'll have to install the drivers yourself. That can get a bit tricky and is out of the scope of this tutorial but [this page](https://wiki.archlinux.org/title/Network_configuration/Wireless#Device_driver) has all the info you will need. Good luck!

### Getting connected with NetworkManager

Assuming the kernel supports your wifi card (which is more than likely the case) or you have successfully installed the required drivers, let's get connected!

List the nearby wifi networks
```
nmcli device wifi list
```
Connect to your network (Replace the text in CAPS)
```
nmcli device wifi connect SSID_or_BSSID password PASSWORD
```
Congratulations, you now have wifi working on Arch Linux! Test your connection with
```
ping archlinux.org
```
For more NetworkManager commands [check out this page](https://wiki.archlinux.org/title/NetworkManager#nmcli_examples). You can also install a GUI network managing tool such as [nm-connection-editor](https://archlinux.org/packages/extra/x86_64/nm-connection-editor/) or [network-manager-applet](https://archlinux.org/packages/extra/x86_64/network-manager-applet/).

## 6. Connecting [Bluetooth](https://wiki.archlinux.org/title/bluetooth) devices

Install Bluetooth and the CLI utilities
```
sudo pacman -S bluez bluez-utils
```

Check the Bluetooth driver is loaded. If it appears in the list then it is loaded.
```
lsmod | grep btusb
```
If `btusb` isn't loaded, then you want to load it by running
```
modprobe btusb
```

Start the Bluetooth service and then enable it so it automatically starts next time you boot up your computer.
```
sudo systemctl start bluetooth.service
sudo systemctl enable bluetooth.service
```

Now lets connect a Bluetooth device. Check you Bluetooth adapter is not blocked by running
```
sudo rfkill list
```

If it is blocked, then unblock it
```
sudo rfkill unblock bluetooth
```

Start the interactive command with
```
bluetoothctl
```

Power it on
```
[bluetooth]# power on
```

Turn on the agent. This will automatically connect any trusted Bluetooth devices when you restart your computer to save you manually reconnecting everything
```
[bluetooth]# agent on
[bluetooth]# default-agent
```

Start scanning for nearby Bluetooth devices
```
[bluetooth]# scan on
```

You can see the list of nearby devices with their MAC address and name by running this command
```
[bluetooth]# devices
```

Trust the device you want to connect. *(Tip: You can press [Tab] to auto-complete the MAC address after typing the first few characters)*
```
[bluetooth]# trust MAC_ADDRESS
```

Now pair the device you want to connect. You may be prompted for a passkey, depending on the device.
```
[bluetooth]# pair MAC_ADDRESS
```

Once you've trusted and paired all devices you need, turn the scan off
```
[bluetooth]# scan off
```

There is one thing left to do. Your system is not going to automatically enable the Bluetooth device when it's found. So let's change that.
```
sudo vim /etc/bluetooth/main.conf
```
Uncomment the the line `#AutoEnable=false` and change it to `true`
```
AutoEnable=true
```


## Well done, you now have a complete installation of the famously "bare-bones" Arch Linux!

You can go on to customise your installation now! Two of the first things I always do is install ZSH, Oh My Zsh, and the GitHub CLI.

### Install Zsh and two plugins (zsh-autosuggestions & zsh-syntax-highlighting)

```
sudo pacman -S zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
git clone https://github.com/zsh-users/zsh-autosuggestions.git $ZSH_CUSTOM/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git $ZSH_CUSTOM/plugins/zsh-syntax-highlighting
```

Enable these two plugins by editing `.zshrc`

Add the two new plugins after git like so:

```
plugins=(git zsh-autosuggestions zsh-syntax-highlighting)
```

### Installing and authenticating GitHub CLI
```
paru -S github-cli
git config --global user.name "Full Name"
git config --global user.email "username@domain.com"
gh auth login
```