# Change Log

## [v0.2.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.2.0) (2020-04-20)
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

## [v0.1.3](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.3) (2020-01-24)
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

## [v0.1.2](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.2) (2019-12-23)
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

## [v0.1.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.1) (2019-11-05)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.1.0...v0.1.1)

**Closed issues:**

- Provide Cluster Advanced configuration API [\#22](https://github.com/mongodb/go-client-mongodb-atlas/issues/22)
- Provide Alert Configurations API [\#18](https://github.com/mongodb/go-client-mongodb-atlas/issues/18)

**Merged pull requests:**

- Added missing comments and omitempty tags for the ProcessArgs struct [\#24](https://github.com/mongodb/go-client-mongodb-atlas/pull/24) ([PacoDw](https://github.com/PacoDw))
- Add cluster advaced configuration support [\#23](https://github.com/mongodb/go-client-mongodb-atlas/pull/23) ([marinsalinas](https://github.com/marinsalinas))
- Fix Linting issues. [\#21](https://github.com/mongodb/go-client-mongodb-atlas/pull/21) ([marinsalinas](https://github.com/marinsalinas))

## [v0.1.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.1.0) (2019-09-18)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.3...v0.1.0)

**Merged pull requests:**

- feat: add private ip mode service [\#19](https://github.com/mongodb/go-client-mongodb-atlas/pull/19) ([marinsalinas](https://github.com/marinsalinas))
- chore: update changelog [\#17](https://github.com/mongodb/go-client-mongodb-atlas/pull/17) ([marinsalinas](https://github.com/marinsalinas))

## [v0.0.3](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.3) (2019-08-14)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.2...v0.0.3)

**Merged pull requests:**

- Add AssignAPIKey struct to Assign request [\#16](https://github.com/mongodb/go-client-mongodb-atlas/pull/16) ([thetonymaster](https://github.com/thetonymaster))
- chore: add changelog for new release [\#15](https://github.com/mongodb/go-client-mongodb-atlas/pull/15) ([marinsalinas](https://github.com/marinsalinas))

## [v0.0.2](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.2) (2019-08-08)
[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.0.1...v0.0.2)

**Merged pull requests:**

- change implementation of the array whitelistAPIKeysReq [\#14](https://github.com/mongodb/go-client-mongodb-atlas/pull/14) ([PacoDw](https://github.com/PacoDw))
- Chore: Sort the services alphabetically [\#13](https://github.com/mongodb/go-client-mongodb-atlas/pull/13) ([thetonymaster](https://github.com/thetonymaster))
- Whitelist API Keys [\#11](https://github.com/mongodb/go-client-mongodb-atlas/pull/11) ([PacoDw](https://github.com/PacoDw))

## [v0.0.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.0.1) (2019-07-17)
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



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*
