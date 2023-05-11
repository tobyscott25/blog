# www.tobyscott.dev

A personal blog built with Hugo and served from S3

### Development

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
hugo --minify
```

### Deployment

Any pushes to the `main` branch will trigger a GitHub Action to:

- Build the static site assets
- Copy assets to S3 bucket
- Invalidate CloudFront distribution cache

The CloudFront distribution is configured to use the S3 bucket as its origin, with a custom domain (`www.tobyscott.dev`) and wildcard certificate (`*.tobyscott.dev`) from AWS Certificate Manager.

> **DNS configuration:**
> A CNAME record points `www.tobyscott.dev` to the CloudFront distribution, and the root domain (`tobyscott.dev`) forwards to `www.tobyscott.dev`.
