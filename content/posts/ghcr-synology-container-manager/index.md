---
date: 2024-06-29T12:38:40+10:00
title: "Pulling images from GitHub Container Registry in Synology Container Manager"
description: "The proper work-around to pull images from GitHub's Container Registry and use them in Synology's Container Manager, with SSH and Docker CLI."
tags:
  [
    "synology",
    "nas",
    "docker",
    "container-manager",
    "ghcr",
    "github",
    "containers",
    "registry",
    "ssh",
    "cli",
    "homelab",
    "self-hosting",
    "serverless",
    "privacy",
  ]
cover:
  image: "synology.jpg"
  alt: "Sybnology logo banner"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
# author: ["Toby Scott", "Other example contributor"]
hidden: false
draft: false
---

I have a Synology NAS which I've recently started using to run Docker containers using Synology's Container Manager.

Until now, the images for all the containers I wanted to run were available on Docker Hub. I'd use Synology's Container Manager UI to pull the image and start the container and it's been quite a nice experience until I needed to run a container from an image hosted in GitHub's Container Registry (https://ghcr.io).

When I tried adding the GitHub Container Registry to Synology Container Manager, I would get a prompt saying "Registry returned bad result", with no further information...

After digging through documentation (and eventually reaching out to Synology support to confirm my suspicion), it appears that **Synology Container Manager does not support the GitHub Container Registry's token authentication**. 🤦‍♂

Given Synology Container Manager uses Docker under the hood, the solution seems to be using the Docker CLI to pull the image from the registry. So you'll need to enable SSH on your Synology NAS for your admin user (Control Panel > Terminal & SNMP > Enable SSH service) and SSH into your NAS to run the Docker CLI command.

SSH into your NAS with your admin account's username and password:

```bash
ssh toby@synologynas
```

Once you're logged into your NAS you'll need to run Docker CLI with sudo to pull the image you require:

```bash
sudo docker pull ghcr.io/open-webui/open-webui:main
```

You can now exit the SSH session and close and reopen Synology DSM and Container Manager and you should now see the image you just pulled in the Container Manager UI.
