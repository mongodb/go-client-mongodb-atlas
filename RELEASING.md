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

## Post Release Actions for V2 SDK

Merge `master` branch changes into `sdkv2` branch:

```bash
git checkout master 
git pull origin master
git checkout sdkv2
git pull origin sdkv2
git merge master
```

Push changes to branch:
```
git push origin sdkv2 
```

## Updating the library version

After releasing update the library version to its next desired version:
-  [.github_changelog_generator](.github_changelog_generator)
-  [Version](mongodbatlas/mongodbatlas.go)
