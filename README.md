# tobyscott.dev

## Guide

Download themes on fresh clone (git submodules may not get cloned automatically)
```bash
git submodule update --init --recursive
```

Update themes (by updating the git submodules)
```bash
git submodule update --remote --merge
```

Create a new blog post
```bash
hugo new posts/title-of-the-post.md
```

Start the development server (with drafts enabled)
```bash
hugo server -D
```

Build for production
```bash
hugo
```

Build Docker image:
```bash
docker build -t tobyscott.dev .
```

Run Docker container:
```bash
docker run -p 8080:80 -d tobyscott.dev
```

Deploy manually to NGINX:
```bash
# Within the project
hugo
zip -r public.zip public
scp public.zip tobyscott.dev:~
rm -r public && rm public.zip

# Now SSH into the server
ssh tobyscott.dev

# Within the server
sudo mv public.zip /var/www/tobyscott.dev
cd /var/www/tobyscott.dev
sudo mv html html_old
sudo unzip public.zip
sudo mv public html
# Clean up once confirmed the new deployment is working
sudo rm public.zip
sudo rm -rf html_old
```