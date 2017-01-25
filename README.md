"For testing go-github api"

Configuration
---

It looks for environment variables to tune things to your needs.

| Variable | Purpose | Default |
| -------- | ------- | ------- |
| CHECK_INTERVAL | Number of hours to wait between checks | 5 hours |
| MAX_AGE | Number of seconds since the last update before calling attention to this pull request | 86400 seconds |
| GITHUB_URL | The URL to GitHub, useful for GitHub Enterprise users | https://github.com |
| GITHUB_ORG | The GitHub organization or owner to scan | eeddins |
| GITHUB_REPOS | A common-separated list of repos in GITHUB_ORG to check | client |
