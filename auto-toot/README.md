# Auto-toot

This is a simple Go app that checks the latest commit for new files added within the `/content/posts` directory. If there are new files, it will parse the Hugo blog post and automatically publish a Mastodon toot with the description, hashtags and link to the post.

### Environment Variables

| Variable                | Description                                                | Example                     |
| ----------------------- | ---------------------------------------------------------- | --------------------------- |
| `MASTODON_ACCESS_TOKEN` | The access token for your Mastodon account.                | `a1b2c3d4e5f6g7h8i9j0`      |
| `MASTODON_ORIGIN`       | The origin for your Mastodon instance.                     | `https://mas.to`            |
| `BLOG_ORIGIN`           | The origin used when generating the link to the blog post. | `https://www.tobyscott.dev` |
