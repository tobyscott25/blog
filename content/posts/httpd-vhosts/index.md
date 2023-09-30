---
date: 2021-11-03T21:23:54+11:00
title: "Configuring subdomains in XAMPP"
description: 'How to configure subdomains - or "virtual hosts" - in an XAMPP development environment.'
tags:
  [
    "xampp",
    "apache",
    "httpd",
    "virtual hosts",
    "subdomains",
    "web development",
    "dev environment",
    "configuration",
    "macos",
    "localhost",
  ]
cover:
  image: "cover.webp"
  alt: "XAMPP Dashboard"
  relative: false # To use relative path for cover image, used in hugo Page-bundles
---

# My environemnt

At the time of writing, I am running [XAMPP](https://www.apachefriends.org/) 8.0.3 on macOS 11.4

# 1. Enable Virtual Hosts in Apache

Go to `/Applications/XAMPP/xamppfiles/etc/httpd.conf` and uncomment this line:

```
#Include etc/extra/httpd-vhosts.conf
```

# 2. Configuring a new virtual host

1. Navigate to XAMPP's `htdocs` and create a directory named `subdomains`.

2. Create a directory inside `subdomains` to be the root directory of the new subdomain. You may call it whatever you like, for this example we will call it `mysubdomain`.

3. Replace the code inside `/Applications/XAMPP/xamppfiles/etc/extra/httpd-vhosts.conf` to configure the new virtual host (subdomain).

```
# localhost
<VirtualHost *:80>
    DocumentRoot "/Applications/XAMPP/xamppfiles/htdocs"
    ServerName localhost
    <Directory  "/Applications/XAMPP/xamppfiles/htdocs">
       Require all granted
    </Directory>
</VirtualHost>
<VirtualHost *:443>
    DocumentRoot "/Applications/XAMPP/xamppfiles/htdocs"
    ServerName localhost

    SSLEngine On

    SSLCertificateFile "etc/ssl.crt/server.crt"
    SSLCertificateKeyFile "etc/ssl.key/server.key"
    <Directory  "/Applications/XAMPP/xamppfiles/htdocs">
        Options Indexes FollowSymLinks
        AllowOverride All
        Require all granted
    </Directory>
</VirtualHost>

# mysubdomain.localhost
<VirtualHost *:80>
    DocumentRoot "/Applications/XAMPP/xamppfiles/htdocs/subdomains/mysubdomain"
    ServerName mysubdomain.localhost
    <Directory  "/Applications/XAMPP/xamppfiles/htdocs/subdomains/mysubdomain">
        Options Indexes FollowSymLinks
        AllowOverride All
        Require all granted
    </Directory>
</VirtualHost>
<VirtualHost *:443>
    DocumentRoot "/Applications/XAMPP/xamppfiles/htdocs/subdomains/mysubdomain"
    ServerName mysubdomain.localhost

    SSLEngine On

    SSLCertificateFile "etc/ssl.crt/server.crt"
    SSLCertificateKeyFile "etc/ssl.key/server.key"
    <Directory  "/Applications/XAMPP/xamppfiles/htdocs/subdomains/mysubdomain">
        Options Indexes FollowSymLinks
        AllowOverride All
        Require all granted
    </Directory>
</VirtualHost>

# Repeat as above to add more
```

# 3. Edit your hosts file

Edit `/etc/hosts` as superuser and add your new subdomain to the list.

```
sudo vim /etc/hosts
```

```
127.0.0.1 localhost
127.0.0.1 mysubdomain.localhost
```

This will point the hostname `mysubdomain.localhost` to your local machine, Apache will then handle any HTTP/HTTPS requests to it.

# Final notes

- Restart XAMPP for the changes to apply.
- Remember to ALWAYS make backups of the original config files.
- You may repeat this process to add as many subdomains as you like.
- If you have a `.htaccess` in the root directory of `htdocs`, double check the rules you're applying with it. You may have to add a `.htaccess` file to your subdomain folder to override the rules set by the one in the root folder since subdomains folder is still a child of `htdocs`, otherwise you may be faced with an annnoying `500 Internal Server` error like I was.
