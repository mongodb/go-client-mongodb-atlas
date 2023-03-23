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

## Release and sdkv2

1. When releasing changes we also need to manualy update `sdkv2` branch 
2. After release please make sure that we merge`master` branch changes into `sdkv2` branch.

## Updating the library version

After releasing update the library version to its next desired version:
-  [.github_changelog_generator](.github_changelog_generator)
-  [Version](mongodbatlas/mongodbatlas.go)
