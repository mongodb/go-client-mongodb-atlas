# Releasing

## Prerequisites

- [github-changelog-generator](https://github.com/github-changelog-generator/github-changelog-generator)

## Generating CHANGELOG

Run:

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
