# Releasing

## Prerequisites

- [github-changelog-generator](https://github.com/github-changelog-generator/github-changelog-generator)

## Generating CHANGELOG

1. Change future-release= and since-tag in .github_changelog_generator

2. Run:

```bash
github_changelog_generator -u mongodb -p go-client-mongodb-atlas
```

And open a PR with your changes.

## Tag and Release

Feel free to use the github GUI to release, or run:

```bash
./scripts/release.sh
```

## Updating the library version

After releasing update the library version to its next desired version:
-  [.github_changelog_generator](.github_changelog_generator)
-  [Version](mongodbatlas/mongodbatlas.go)
