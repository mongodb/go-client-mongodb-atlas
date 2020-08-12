# Change Log

## [0.4.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.4.0) (August 10, 2020)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.3.0...0.4.0)

**Implemented enhancements:**

- Updated Database user for AWS IAM [\#114](https://github.com/mongodb/go-client-mongodb-atlas/pull/114) ([coderGo93](https://github.com/coderGo93)
- Added Custom AWS DNS support [\#112](https://github.com/mongodb/go-client-mongodb-atlas/pull/112) ([gssbzn](https://github.com/gssbzn))
- Implemented Atlas Search support [\#111](https://github.com/mongodb/go-client-mongodb-atlas/pull/111) ([gssbzn](https://github.com/gssbzn))
- Added Atlas Data Lakes support [\#108](https://github.com/mongodb/go-client-mongodb-atlas/pull/108) ([robcarlan-mlab](https://github.com/robcarlan-mlab))
- Added Online Archive support [\#107](https://github.com/mongodb/go-client-mongodb-atlas/pull/107) ([gssbzn](https://github.com/gssbzn))
- Added Organization API [\#103](https://github.com/mongodb/go-client-mongodb-atlas/pull/103) ([gssbzn](https://github.com/gssbzn))
- Added vanity URL [\#87](https://github.com/mongodb/go-client-mongodb-atlas/pull/87) ([gssbzn](https://github.com/gssbzn))

**Closed issues:**

- CloudProviderSnapshotBackupPolicy: Zero Values are ignored  [\#93](https://github.com/mongodb/go-client-mongodb-atlas/issues/93)

**Merged pull requests:**

- Create SECURITY.md [\#118](https://github.com/mongodb/go-client-mongodb-atlas/pull/118) ([gssbzn](https://github.com/gssbzn))
- fix: add more fields to the search field [\#117](https://github.com/mongodb/go-client-mongodb-atlas/pull/117) ([gssbzn](https://github.com/gssbzn))
- Updated the references to the digest library [\#116](https://github.com/mongodb/go-client-mongodb-atlas/pull/116) ([MihaiBojin](https://github.com/MihaiBojin))
- Add a hook to ensure the client's version is updated before pushing a tag/release [\#115](https://github.com/mongodb/go-client-mongodb-atlas/pull/115) ([MihaiBojin](https://github.com/MihaiBojin))
- Added field for aws iam type for database user [\#114](https://github.com/mongodb/go-client-mongodb-atlas/pull/114) ([coderGo93](https://github.com/coderGo93))
- Remove libraryVersion from user agent [\#113](https://github.com/mongodb/go-client-mongodb-atlas/pull/113) ([gssbzn](https://github.com/gssbzn))
- Add custom AWS DNS support [\#112](https://github.com/mongodb/go-client-mongodb-atlas/pull/112) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-67056: Implement search API in the go client [\#111](https://github.com/mongodb/go-client-mongodb-atlas/pull/111) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-66612: fix json marshaling [\#110](https://github.com/mongodb/go-client-mongodb-atlas/pull/110) ([robcarlan-mlab](https://github.com/robcarlan-mlab))
- Update GNUmakefile [\#109](https://github.com/mongodb/go-client-mongodb-atlas/pull/109) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-66607: Atlas Data Lakes [\#108](https://github.com/mongodb/go-client-mongodb-atlas/pull/108) ([robcarlan-mlab](https://github.com/robcarlan-mlab))
- CLOUDP-66608: Implement Online Archive in the Go client [\#107](https://github.com/mongodb/go-client-mongodb-atlas/pull/107) ([gssbzn](https://github.com/gssbzn))
- Add the ability to filter processes by cluster ID [\#105](https://github.com/mongodb/go-client-mongodb-atlas/pull/105) ([gssbzn](https://github.com/gssbzn))
- feat: Add missing method to get all containers [\#104](https://github.com/mongodb/go-client-mongodb-atlas/pull/104) ([gssbzn](https://github.com/gssbzn))
- feat: adds organization API [\#103](https://github.com/mongodb/go-client-mongodb-atlas/pull/103) ([gssbzn](https://github.com/gssbzn))
- task: Update and fix linting [\#102](https://github.com/mongodb/go-client-mongodb-atlas/pull/102) ([gssbzn](https://github.com/gssbzn))
- Improve docs [\#101](https://github.com/mongodb/go-client-mongodb-atlas/pull/101) ([gssbzn](https://github.com/gssbzn))
- Updated Changelog for v0.3.0 release [\#99](https://github.com/mongodb/go-client-mongodb-atlas/pull/99) ([PacoDw](https://github.com/PacoDw))
- Use the new vanity url [\#87](https://github.com/mongodb/go-client-mongodb-atlas/pull/87) ([gssbzn](https://github.com/gssbzn))

## [v0.3.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.3.0) (June 01, 2020)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.2.0...v0.3.0)

**Fixed bugs:**

- Fix cluster struct to allow pasue/start a cluster [\#84](https://github.com/mongodb/go-client-mongodb-atlas/pull/84) ([gssbzn](https://github.com/gssbzn))

**Closed issues:**

- Support pagination [\#95](https://github.com/mongodb/go-client-mongodb-atlas/issues/95)
- CloudProviderSnapshotBackupPolicy: Zero Values are ignored  [\#93](https://github.com/mongodb/go-client-mongodb-atlas/issues/93)

**Merged pull requests:**

- feat: support the pagination of WhitelistAPIKeysService.List [\#98](https://github.com/mongodb/go-client-mongodb-atlas/pull/98) ([suzuki-shunsuke](https://github.com/suzuki-shunsuke))
- feat: support the pagination at CloudProviderSnapshotRestoreJobsService.List [\#97](https://github.com/mongodb/go-client-mongodb-atlas/pull/97) ([suzuki-shunsuke](https://github.com/suzuki-shunsuke))
- feat: support the pagination at GetAllCloudProviderSnapshots [\#96](https://github.com/mongodb/go-client-mongodb-atlas/pull/96) ([suzuki-shunsuke](https://github.com/suzuki-shunsuke))
- fix: fixes \#93 CloudProviderSnapshotBackupPolicy zero values are ignored [\#94](https://github.com/mongodb/go-client-mongodb-atlas/pull/94) ([gmlp](https://github.com/gmlp))
- Cluster autoscaling compute [\#92](https://github.com/mongodb/go-client-mongodb-atlas/pull/92) ([coderGo93](https://github.com/coderGo93))
- Three new fields for Snapshot restore jobs options [\#91](https://github.com/mongodb/go-client-mongodb-atlas/pull/91) ([coderGo93](https://github.com/coderGo93))
- CLOUDP-62220: Update the AcknowledgedUntil field to be a pointer in the atlas client [\#90](https://github.com/mongodb/go-client-mongodb-atlas/pull/90) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Fix and update pointy dependency [\#89](https://github.com/mongodb/go-client-mongodb-atlas/pull/89) ([acroca](https://github.com/acroca))
- Fix projects lists [\#88](https://github.com/mongodb/go-client-mongodb-atlas/pull/88) ([gssbzn](https://github.com/gssbzn))
- Fix event list options [\#86](https://github.com/mongodb/go-client-mongodb-atlas/pull/86) ([gssbzn](https://github.com/gssbzn))
- Make the measurements param an array [\#85](https://github.com/mongodb/go-client-mongodb-atlas/pull/85) ([gssbzn](https://github.com/gssbzn))
- Validate base url trailing slash [\#83](https://github.com/mongodb/go-client-mongodb-atlas/pull/83) ([gssbzn](https://github.com/gssbzn))
- Update CHANGELOG.md [\#82](https://github.com/mongodb/go-client-mongodb-atlas/pull/82) ([marinsalinas](https://github.com/marinsalinas))
- Improve docs [\#61](https://github.com/mongodb/go-client-mongodb-atlas/pull/61) ([gssbzn](https://github.com/gssbzn))

## [v0.2.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.2.0) (April 20, 2020)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.1.3...v0.2.0)

**Merged pull requests:**

- Add database measurements Endpoint [\#80](https://github.com/mongodb/go-client-mongodb-atlas/pull/80) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-60178: Implement logs API, Atlas [\#79](https://github.com/mongodb/go-client-mongodb-atlas/pull/79) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-60177: Implement Get Measurements of a Disk for a MongoDB Process, Atlas [\#78](https://github.com/mongodb/go-client-mongodb-atlas/pull/78) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Connection strings private [\#77](https://github.com/mongodb/go-client-mongodb-atlas/pull/77) ([coderGo93](https://github.com/coderGo93))
- CLOUDP-60176: Implement Get Available Databases for a MongoDB Process [\#76](https://github.com/mongodb/go-client-mongodb-atlas/pull/76) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-60028: Implement index creation in atlas client [\#75](https://github.com/mongodb/go-client-mongodb-atlas/pull/75) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-60022: List available disks for a process, Atlas [\#74](https://github.com/mongodb/go-client-mongodb-atlas/pull/74) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-59968: implement list measurements on atlas go client [\#73](https://github.com/mongodb/go-client-mongodb-atlas/pull/73) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-57848: Implement events resource in the atlas go client [\#72](https://github.com/mongodb/go-client-mongodb-atlas/pull/72) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Add get processes in a group [\#71](https://github.com/mongodb/go-client-mongodb-atlas/pull/71) ([mmb](https://github.com/mmb))
- CLOUDP-59516: List available matchers for alerts in the Atlas client [\#70](https://github.com/mongodb/go-client-mongodb-atlas/pull/70) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-59435: Update alert list to use status in Atlas [\#69](https://github.com/mongodb/go-client-mongodb-atlas/pull/69) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Connection strings [\#68](https://github.com/mongodb/go-client-mongodb-atlas/pull/68) ([coderGo93](https://github.com/coderGo93))
- Cloud Provider Snapshot Backup Policy [\#67](https://github.com/mongodb/go-client-mongodb-atlas/pull/67) ([PacoDw](https://github.com/PacoDw))
- fix: added field roles in notifications for alert configurations, add… [\#66](https://github.com/mongodb/go-client-mongodb-atlas/pull/66) ([coderGo93](https://github.com/coderGo93))
- CLOUDP-58634: Add support for checkpoints in the Atlas client [\#65](https://github.com/mongodb/go-client-mongodb-atlas/pull/65) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-58969: Create restore job returns a list and not a single job [\#64](https://github.com/mongodb/go-client-mongodb-atlas/pull/64) ([gssbzn](https://github.com/gssbzn))
- Improve Travis [\#63](https://github.com/mongodb/go-client-mongodb-atlas/pull/63) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-58633: Add Restore Jobs support to atlas client [\#62](https://github.com/mongodb/go-client-mongodb-atlas/pull/62) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-58632: Add Continuous Backup Snapshots support for atlas go client [\#60](https://github.com/mongodb/go-client-mongodb-atlas/pull/60) ([gssbzn](https://github.com/gssbzn))
- Make the ServiceOp depend on an interface [\#59](https://github.com/mongodb/go-client-mongodb-atlas/pull/59) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-57632: Implementing Alerts resource [\#58](https://github.com/mongodb/go-client-mongodb-atlas/pull/58) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Fix incorrect json tag for groupId [\#56](https://github.com/mongodb/go-client-mongodb-atlas/pull/56) ([gssbzn](https://github.com/gssbzn))
- fix Database Users: added the reflect missing dependency [\#55](https://github.com/mongodb/go-client-mongodb-atlas/pull/55) ([PacoDw](https://github.com/PacoDw))
- x509AuthDBUsers: Fixed the id type [\#54](https://github.com/mongodb/go-client-mongodb-atlas/pull/54) ([PacoDw](https://github.com/PacoDw))
- feat: add pitEnabled attribute un clusters service [\#53](https://github.com/mongodb/go-client-mongodb-atlas/pull/53) ([marinsalinas](https://github.com/marinsalinas))
- chore: add changelog file for 0.1.3 version [\#52](https://github.com/mongodb/go-client-mongodb-atlas/pull/52) ([marinsalinas](https://github.com/marinsalinas))
- X509 auth db users [\#51](https://github.com/mongodb/go-client-mongodb-atlas/pull/51) ([PacoDw](https://github.com/PacoDw))
- fix license [\#50](https://github.com/mongodb/go-client-mongodb-atlas/pull/50) ([themantissa](https://github.com/themantissa))
- Database Users [\#49](https://github.com/mongodb/go-client-mongodb-atlas/pull/49) ([PacoDw](https://github.com/PacoDw))
- Private Endpoints [\#48](https://github.com/mongodb/go-client-mongodb-atlas/pull/48) ([PacoDw](https://github.com/PacoDw))
- Simplify tests [\#46](https://github.com/mongodb/go-client-mongodb-atlas/pull/46) ([gssbzn](https://github.com/gssbzn))

## [v0.1.3](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.3) (January 24, 2020)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.1.2...v0.1.3)

**Implemented enhancements:**

- Label argument missing on Database User [\#28](https://github.com/mongodb/go-client-mongodb-atlas/issues/28)

**Closed issues:**

- Missing awsSecurityGroup property for ProjectIPWhitelist [\#40](https://github.com/mongodb/go-client-mongodb-atlas/issues/40)

**Merged pull requests:**

- add DefaultDiskSizeGB for each provider [\#47](https://github.com/mongodb/go-client-mongodb-atlas/pull/47) ([marinsalinas](https://github.com/marinsalinas))
- Custom DB Roles [\#45](https://github.com/mongodb/go-client-mongodb-atlas/pull/45) ([PacoDw](https://github.com/PacoDw))
- Teams and projects resource were refactored due to issues [\#44](https://github.com/mongodb/go-client-mongodb-atlas/pull/44) ([PacoDw](https://github.com/PacoDw))
- Added the label attribute to the cluster struct [\#43](https://github.com/mongodb/go-client-mongodb-atlas/pull/43) ([PacoDw](https://github.com/PacoDw))
- Added Label struct and array inside the database users  [\#42](https://github.com/mongodb/go-client-mongodb-atlas/pull/42) ([PacoDw](https://github.com/PacoDw))
- changed the AddUsersToTeam function to allow a string array of users id [\#41](https://github.com/mongodb/go-client-mongodb-atlas/pull/41) ([PacoDw](https://github.com/PacoDw))
- Alert Configuration [\#39](https://github.com/mongodb/go-client-mongodb-atlas/pull/39) ([PacoDw](https://github.com/PacoDw))
- Enable golint and fix errors [\#38](https://github.com/mongodb/go-client-mongodb-atlas/pull/38) ([gssbzn](https://github.com/gssbzn))
- Udpate changelog for v0.1.2 release [\#37](https://github.com/mongodb/go-client-mongodb-atlas/pull/37) ([marinsalinas](https://github.com/marinsalinas))
- Whitelist [\#35](https://github.com/mongodb/go-client-mongodb-atlas/pull/35) ([PacoDw](https://github.com/PacoDw))

## [v0.1.2](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.2) (December 23, 2019)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.1.1...v0.1.2)

**Merged pull requests:**

- add providerName option in listoptions for containers service [\#36](https://github.com/mongodb/go-client-mongodb-atlas/pull/36) ([marinsalinas](https://github.com/marinsalinas))
- Improve Makefile [\#34](https://github.com/mongodb/go-client-mongodb-atlas/pull/34) ([gssbzn](https://github.com/gssbzn))
- Global Clusters Service [\#32](https://github.com/mongodb/go-client-mongodb-atlas/pull/32) ([marinsalinas](https://github.com/marinsalinas))
- fix: linting issue custom db roles [\#31](https://github.com/mongodb/go-client-mongodb-atlas/pull/31) ([marinsalinas](https://github.com/marinsalinas))
- Auditing [\#30](https://github.com/mongodb/go-client-mongodb-atlas/pull/30) ([PacoDw](https://github.com/PacoDw))
- Atlas Users Service [\#29](https://github.com/mongodb/go-client-mongodb-atlas/pull/29) ([marinsalinas](https://github.com/marinsalinas))
- chore: added default disk size for each provider tier [\#27](https://github.com/mongodb/go-client-mongodb-atlas/pull/27) ([PacoDw](https://github.com/PacoDw))
- Teams Service [\#26](https://github.com/mongodb/go-client-mongodb-atlas/pull/26) ([marinsalinas](https://github.com/marinsalinas))
- Maintance windows [\#25](https://github.com/mongodb/go-client-mongodb-atlas/pull/25) ([PacoDw](https://github.com/PacoDw))
- Service for manipulating Custom MongoDB Roles [\#20](https://github.com/mongodb/go-client-mongodb-atlas/pull/20) ([mpaluchowski](https://github.com/mpaluchowski))

## [v0.1.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.1) (November 05, 2019)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.1.0...v0.1.1)

**Closed issues:**

- Provide Cluster Advanced configuration API [\#22](https://github.com/mongodb/go-client-mongodb-atlas/issues/22)
- Provide Alert Configurations API [\#18](https://github.com/mongodb/go-client-mongodb-atlas/issues/18)

**Merged pull requests:**

- Added missing comments and omitempty tags for the ProcessArgs struct [\#24](https://github.com/mongodb/go-client-mongodb-atlas/pull/24) ([PacoDw](https://github.com/PacoDw))
- Add cluster advaced configuration support [\#23](https://github.com/mongodb/go-client-mongodb-atlas/pull/23) ([marinsalinas](https://github.com/marinsalinas))
- Fix Linting issues. [\#21](https://github.com/mongodb/go-client-mongodb-atlas/pull/21) ([marinsalinas](https://github.com/marinsalinas))

## [v0.1.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.0) (September 18, 2019)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.3...v0.1.0)

**Merged pull requests:**

- feat: add private ip mode service [\#19](https://github.com/mongodb/go-client-mongodb-atlas/pull/19) ([marinsalinas](https://github.com/marinsalinas))
- chore: update changelog [\#17](https://github.com/mongodb/go-client-mongodb-atlas/pull/17) ([marinsalinas](https://github.com/marinsalinas))

## [v0.0.3](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.3) (August 14, 2019)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.2...v0.0.3)

**Merged pull requests:**

- Add AssignAPIKey struct to Assign request [\#16](https://github.com/mongodb/go-client-mongodb-atlas/pull/16) ([thetonymaster](https://github.com/thetonymaster))
- chore: add changelog for new release [\#15](https://github.com/mongodb/go-client-mongodb-atlas/pull/15) ([marinsalinas](https://github.com/marinsalinas))

## [v0.0.2](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.2) (August 08, 2019)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.1...v0.0.2)

**Merged pull requests:**

- change implementation of the array whitelistAPIKeysReq [\#14](https://github.com/mongodb/go-client-mongodb-atlas/pull/14) ([PacoDw](https://github.com/PacoDw))
- Chore: Sort the services alphabetically [\#13](https://github.com/mongodb/go-client-mongodb-atlas/pull/13) ([thetonymaster](https://github.com/thetonymaster))
- Whitelist API Keys [\#11](https://github.com/mongodb/go-client-mongodb-atlas/pull/11) ([PacoDw](https://github.com/PacoDw))

## [v0.0.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.1) (July 17, 2019)
**Merged pull requests:**

- chore: add changelog file [\#10](https://github.com/mongodb/go-client-mongodb-atlas/pull/10) ([marinsalinas](https://github.com/marinsalinas))
- ProjectAPIKeys.Create test [\#9](https://github.com/mongodb/go-client-mongodb-atlas/pull/9) ([thetonymaster](https://github.com/thetonymaster))
- Project api keys [\#8](https://github.com/mongodb/go-client-mongodb-atlas/pull/8) ([thetonymaster](https://github.com/thetonymaster))
- Update Readme.md [\#7](https://github.com/mongodb/go-client-mongodb-atlas/pull/7) ([marinsalinas](https://github.com/marinsalinas))
-  Encryption At Rest Service [\#6](https://github.com/mongodb/go-client-mongodb-atlas/pull/6) ([PacoDw](https://github.com/PacoDw))
- fix: remove content-type before perform containders DELETE request [\#5](https://github.com/mongodb/go-client-mongodb-atlas/pull/5) ([marinsalinas](https://github.com/marinsalinas))
- Containers Service [\#4](https://github.com/mongodb/go-client-mongodb-atlas/pull/4) ([marinsalinas](https://github.com/marinsalinas))
- Peers Service [\#3](https://github.com/mongodb/go-client-mongodb-atlas/pull/3) ([marinsalinas](https://github.com/marinsalinas))
- Chore: TravisCI integration [\#2](https://github.com/mongodb/go-client-mongodb-atlas/pull/2) ([marinsalinas](https://github.com/marinsalinas))
- Cloud Provider Snapshot Restore Job Support [\#1](https://github.com/mongodb/go-client-mongodb-atlas/pull/1) ([PacoDw](https://github.com/PacoDw))



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
