# Changelog

## [v0.23.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.23.1) (March 01, 2023)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.23.0...v0.23.1)

**Breaking changes:**

- fix: make org owner optional [\#393](https://github.com/mongodb/go-client-mongodb-atlas/pull/393) ([gssbzn](https://github.com/gssbzn))

## [v0.23.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.23.0) (February 23, 2023)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.22.0...v0.23.0)

**Closed issues:**

- CloudQuery Source Plugin? [\#375](https://github.com/mongodb/go-client-mongodb-atlas/issues/375)

**Merged pull requests:**

- feat: add org creation endpoint [\#387](https://github.com/mongodb/go-client-mongodb-atlas/pull/387) ([gssbzn](https://github.com/gssbzn))
- task: update linter version [\#384](https://github.com/mongodb/go-client-mongodb-atlas/pull/384) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-156187: change owners team to apix-1 [\#367](https://github.com/mongodb/go-client-mongodb-atlas/pull/367) ([wtrocki](https://github.com/wtrocki))

## [v0.22.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.22.0) (January 30, 2023)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.21.0...v0.22.0)

**Closed issues:**

- Issue calling CloudProviderSnapshotBackupPolicies Service [\#349](https://github.com/mongodb/go-client-mongodb-atlas/issues/349)

**Merged pull requests:**

- test: clean up shadow variables for constants [\#351](https://github.com/mongodb/go-client-mongodb-atlas/pull/351) ([gssbzn](https://github.com/gssbzn))
- feat: add outage simulation [\#350](https://github.com/mongodb/go-client-mongodb-atlas/pull/350) ([gssbzn](https://github.com/gssbzn))
- feat: add test failover method [\#343](https://github.com/mongodb/go-client-mongodb-atlas/pull/343) ([gssbzn](https://github.com/gssbzn))
- chore\(deps\): bump golangci/golangci-lint-action from 3.3.1 to 3.4.0 [\#342](https://github.com/mongodb/go-client-mongodb-atlas/pull/342) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v0.21.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.21.0) (January 09, 2023)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.20.0...v0.21.0)

**Breaking changes:**

- fix: fixes ldap config to allow PATCH [\#340](https://github.com/mongodb/go-client-mongodb-atlas/pull/340) ([gssbzn](https://github.com/gssbzn))

**Closed issues:**

- Unable to set the "authenticationEnabled" = false using the client.LDAPConfigurations.Save\(\) [\#339](https://github.com/mongodb/go-client-mongodb-atlas/issues/339)

**Merged pull requests:**

- chore\(deps\): bump actions/stale from 6 to 7 [\#338](https://github.com/mongodb/go-client-mongodb-atlas/pull/338) ([dependabot[bot]](https://github.com/apps/dependabot))
- Add global cluster support for advanced cluster [\#337](https://github.com/mongodb/go-client-mongodb-atlas/pull/337) ([Sugar-pack](https://github.com/Sugar-pack))
- INTMDB-400: Empty struct to delete backup distribution configuration [\#336](https://github.com/mongodb/go-client-mongodb-atlas/pull/336) ([martinstibbe](https://github.com/martinstibbe))
- CLOUDP-149874: export jobs should use exportID as per docs [\#335](https://github.com/mongodb/go-client-mongodb-atlas/pull/335) ([matt-condon](https://github.com/matt-condon))
- INTMDB-454: Add support for root query to get org\_id from API Key used in Terraform [\#333](https://github.com/mongodb/go-client-mongodb-atlas/pull/333) ([martinstibbe](https://github.com/martinstibbe))

## [v0.20.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.20.0) (December 16, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.19.0...v0.20.0)

**Merged pull requests:**

- Added missing fields 'Links', 'CopySettings' and 'DeleteCopiedBackups… [\#332](https://github.com/mongodb/go-client-mongodb-atlas/pull/332) ([sangramPeerI](https://github.com/sangramPeerI))
- CLOUDP-149874: export jobs should use exportJobID [\#331](https://github.com/mongodb/go-client-mongodb-atlas/pull/331) ([matt-condon](https://github.com/matt-condon))
- chore\(deps\): bump github.com/go-test/deep from 1.0.8 to 1.1.0 [\#330](https://github.com/mongodb/go-client-mongodb-atlas/pull/330) ([dependabot[bot]](https://github.com/apps/dependabot))
- fix: change wrong variable name on cluster update test to avoid  misu… [\#328](https://github.com/mongodb/go-client-mongodb-atlas/pull/328) ([atileren](https://github.com/atileren))
- chore\(deps\): bump golangci/golangci-lint-action from 3.3.0 to 3.3.1 [\#327](https://github.com/mongodb/go-client-mongodb-atlas/pull/327) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v0.19.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.19.0) (November 11, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.18.0...v0.19.0)

**Breaking changes:**

- task: drop support for 1.17 [\#323](https://github.com/mongodb/go-client-mongodb-atlas/pull/323) ([gssbzn](https://github.com/gssbzn))

**Merged pull requests:**

- INTMDB-444 Add support for terminationProtectionEnabled flag [\#325](https://github.com/mongodb/go-client-mongodb-atlas/pull/325) ([martinstibbe](https://github.com/martinstibbe))
- CLOUDP-143634: Add analyticsAutoScaling to advanced cluster atlas go client [\#324](https://github.com/mongodb/go-client-mongodb-atlas/pull/324) ([andreaangiolillo](https://github.com/andreaangiolillo))
- chore\(deps\): bump golangci/golangci-lint-action from 3.2.0 to 3.3.0 [\#322](https://github.com/mongodb/go-client-mongodb-atlas/pull/322) ([dependabot[bot]](https://github.com/apps/dependabot))
- chore\(deps\): bump github.com/openlyinc/pointy from 1.1.2 to 1.2.0 [\#321](https://github.com/mongodb/go-client-mongodb-atlas/pull/321) ([dependabot[bot]](https://github.com/apps/dependabot))
- INTMDB-373: Add support for Webhook and MicrosoftTeamsWebhookURL Notifications [\#320](https://github.com/mongodb/go-client-mongodb-atlas/pull/320) ([martinstibbe](https://github.com/martinstibbe))
- CLOUDP-141434: Add RootCertType to the atlas go client [\#319](https://github.com/mongodb/go-client-mongodb-atlas/pull/319) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-79100: Autoclose stale issues and PRs [\#318](https://github.com/mongodb/go-client-mongodb-atlas/pull/318) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.18.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.18.0) (October 17, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.17.0...v0.18.0)

**Closed issues:**

- "CloudProviderSnapshots.GetOneCloudProviderSnapshot\(\)" method returns already deleted snapshot [\#315](https://github.com/mongodb/go-client-mongodb-atlas/issues/315)

**Merged pull requests:**

- CLOUDP-139997: added OplogMinRetentionHours to ProcessArgs [\#316](https://github.com/mongodb/go-client-mongodb-atlas/pull/316) ([andreaangiolillo](https://github.com/andreaangiolillo))
- INTMDB-364: \[Terraform\] Add support for serverless private endpoints [\#314](https://github.com/mongodb/go-client-mongodb-atlas/pull/314) ([martinstibbe](https://github.com/martinstibbe))

## [v0.17.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.17.0) (September 27, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.16.0...v0.17.0)

**Breaking changes:**

- task: drop support for whitelist [\#306](https://github.com/mongodb/go-client-mongodb-atlas/pull/306) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-127482: Add filter by name on list all groups by org [\#302](https://github.com/mongodb/go-client-mongodb-atlas/pull/302) ([fmenezes](https://github.com/fmenezes))
- CLOUDP-110989: Adds OnResponseProcessed [\#298](https://github.com/mongodb/go-client-mongodb-atlas/pull/298) ([fmenezes](https://github.com/fmenezes))
- feat: add pagination to listing x509 users [\#294](https://github.com/mongodb/go-client-mongodb-atlas/pull/294) ([gssbzn](https://github.com/gssbzn))

**Closed issues:**

- UpdateTeamRoles should take a groupID as a parameter instead of an orgID [\#303](https://github.com/mongodb/go-client-mongodb-atlas/issues/303)
- `X509AuthDBUsers.GetUserCertificates` does not support pagination  [\#293](https://github.com/mongodb/go-client-mongodb-atlas/issues/293)

**Merged pull requests:**

- feat: add GET process [\#311](https://github.com/mongodb/go-client-mongodb-atlas/pull/311) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-132512: \[Atlas CLI\] Add support for updating a serverless cluster [\#310](https://github.com/mongodb/go-client-mongodb-atlas/pull/310) ([gssbzn](https://github.com/gssbzn))
- Revert CLOUDP-132967: add the ability to use yaml file for advance cluster  [\#309](https://github.com/mongodb/go-client-mongodb-atlas/pull/309) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Update to Go v1.18 [\#308](https://github.com/mongodb/go-client-mongodb-atlas/pull/308) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-132967: add the ability to use yaml file for advance cluster [\#307](https://github.com/mongodb/go-client-mongodb-atlas/pull/307) ([andreaangiolillo](https://github.com/andreaangiolillo))
- task: add tests against 1.19 [\#305](https://github.com/mongodb/go-client-mongodb-atlas/pull/305) ([gssbzn](https://github.com/gssbzn))
- Fix: Fixed naming of parameter in UpdateTeamRoles [\#304](https://github.com/mongodb/go-client-mongodb-atlas/pull/304) ([marcoriopel](https://github.com/marcoriopel))
- CLOUDP-126267: Add 'Upgrade One Shared Tier Cluster' endpoint to atlas go client [\#301](https://github.com/mongodb/go-client-mongodb-atlas/pull/301) ([Razeer123](https://github.com/Razeer123))
- CLOUDP-120670 client for registration config endpoint [\#299](https://github.com/mongodb/go-client-mongodb-atlas/pull/299) ([tibulca](https://github.com/tibulca))
- CLOUDP-110989: Capture Raw response body [\#297](https://github.com/mongodb/go-client-mongodb-atlas/pull/297) ([fmenezes](https://github.com/fmenezes))
- INTMDB-301: Feature add: Add support for management of federated authentication configuration [\#296](https://github.com/mongodb/go-client-mongodb-atlas/pull/296) ([martinstibbe](https://github.com/martinstibbe))
- chore\(deps\): bump golangci/golangci-lint-action from 3.1.0 to 3.2.0 [\#295](https://github.com/mongodb/go-client-mongodb-atlas/pull/295) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v0.16.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.16.0) (May 10, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.15.0...v0.16.0)

**Merged pull requests:**

- CLOUDP-118714 add IsTimeoutErr func [\#291](https://github.com/mongodb/go-client-mongodb-atlas/pull/291) ([tibulca](https://github.com/tibulca))
- feat: add useOrgAndGroupNamesInExportPrefix to snapshot shcedule [\#290](https://github.com/mongodb/go-client-mongodb-atlas/pull/290) ([gssbzn](https://github.com/gssbzn))
- feat: add support to export snapshots on backup policies [\#289](https://github.com/mongodb/go-client-mongodb-atlas/pull/289) ([trunet](https://github.com/trunet))
- chore\(deps\): bump actions/setup-go from 2 to 3 [\#288](https://github.com/mongodb/go-client-mongodb-atlas/pull/288) ([dependabot[bot]](https://github.com/apps/dependabot))
- CLOUDP-118028: \[atlascli\] Allow to list and edit project settings [\#287](https://github.com/mongodb/go-client-mongodb-atlas/pull/287) ([andreaangiolillo](https://github.com/andreaangiolillo))
- feat: add ready for cutover field to live migrations [\#286](https://github.com/mongodb/go-client-mongodb-atlas/pull/286) ([gssbzn](https://github.com/gssbzn))
- task: add go 1.18 to the test matrix [\#285](https://github.com/mongodb/go-client-mongodb-atlas/pull/285) ([gssbzn](https://github.com/gssbzn))
- INTMDB-311: Feature Add: Prometheus and Microsoft Team to the Third Party Integration Settings [\#284](https://github.com/mongodb/go-client-mongodb-atlas/pull/284) ([martinstibbe](https://github.com/martinstibbe))
- chore\(deps\): bump actions/cache from 2 to 3 [\#283](https://github.com/mongodb/go-client-mongodb-atlas/pull/283) ([dependabot[bot]](https://github.com/apps/dependabot))
- chore\(deps\): bump actions/checkout from 2 to 3 [\#282](https://github.com/mongodb/go-client-mongodb-atlas/pull/282) ([dependabot[bot]](https://github.com/apps/dependabot))
- CLOUDP-114669: Automatically add trailing slash to the URL [\#281](https://github.com/mongodb/go-client-mongodb-atlas/pull/281) ([igor-karpukhin](https://github.com/igor-karpukhin))
- chore\(deps\): bump golangci/golangci-lint-action from 2 to 3.1.0 [\#280](https://github.com/mongodb/go-client-mongodb-atlas/pull/280) ([dependabot[bot]](https://github.com/apps/dependabot))
- INTMDB-299: Added snapshot export jobs support [\#279](https://github.com/mongodb/go-client-mongodb-atlas/pull/279) ([abner-dou](https://github.com/abner-dou))

## [v0.15.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.15.0) (February 09, 2022)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.14.0...v0.15.0)

**Merged pull requests:**

- feat: add the ability to refresh tokens [\#277](https://github.com/mongodb/go-client-mongodb-atlas/pull/277) ([gssbzn](https://github.com/gssbzn))
- INTMDB-299: Added snapshot export bucket support [\#276](https://github.com/mongodb/go-client-mongodb-atlas/pull/276) ([abner-dou](https://github.com/abner-dou))
- feat: add the ability to revoke tokens [\#275](https://github.com/mongodb/go-client-mongodb-atlas/pull/275) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-110641: Add device grant  flow to atlas client [\#274](https://github.com/mongodb/go-client-mongodb-atlas/pull/274) ([gssbzn](https://github.com/gssbzn))
- INTMDB-288: Added support for restore backup job [\#273](https://github.com/mongodb/go-client-mongodb-atlas/pull/273) ([abner-dou](https://github.com/abner-dou))
- INTMDB-5: added parameter team name for alert configurations [\#272](https://github.com/mongodb/go-client-mongodb-atlas/pull/272) ([coderGo93](https://github.com/coderGo93))
- INTMDB-263: Added Private\_link\_endpoint\_adl resource to go-client [\#271](https://github.com/mongodb/go-client-mongodb-atlas/pull/271) ([abner-dou](https://github.com/abner-dou))
- INTMDB-276: Added parameter versionReleaseSystem in clusters [\#270](https://github.com/mongodb/go-client-mongodb-atlas/pull/270) ([coderGo93](https://github.com/coderGo93))
- fix: fixes release smoke tests [\#269](https://github.com/mongodb/go-client-mongodb-atlas/pull/269) ([gssbzn](https://github.com/gssbzn))

## [v0.14.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.14.0) (November 18, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.13.0...v0.14.0)

**Breaking changes:**

- custom\_db\_role resource use \*string [\#264](https://github.com/mongodb/go-client-mongodb-atlas/pull/264) ([Freyert](https://github.com/Freyert))

**Closed issues:**

- undefined: io.Discard and io.ReadAll in latest code push [\#267](https://github.com/mongodb/go-client-mongodb-atlas/issues/267)
- Can't provide customRole.actions\[\*\].resources\[\*\].collection as empty string [\#263](https://github.com/mongodb/go-client-mongodb-atlas/issues/263)

**Merged pull requests:**

- CLOUDP-105826: Update golangci-lint to 1.43 [\#266](https://github.com/mongodb/go-client-mongodb-atlas/pull/266) ([gssbzn](https://github.com/gssbzn))
- INTMDB-270: Added  'withDefaultAlertsSettings' property to Project struct [\#265](https://github.com/mongodb/go-client-mongodb-atlas/pull/265) ([abner-dou](https://github.com/abner-dou))
- chore\(deps\): bump github.com/go-test/deep from 1.0.7 to 1.0.8 [\#262](https://github.com/mongodb/go-client-mongodb-atlas/pull/262) ([dependabot[bot]](https://github.com/apps/dependabot))
- INTMDB-239: added fields for cloud provider snapshot [\#261](https://github.com/mongodb/go-client-mongodb-atlas/pull/261) ([coderGo93](https://github.com/coderGo93))
- INTMDB-260: Added for gcp features in private endpoints [\#260](https://github.com/mongodb/go-client-mongodb-atlas/pull/260) ([coderGo93](https://github.com/coderGo93))

## [v0.13.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.13.0) (October 07, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.12.0...v0.13.0)

**Breaking changes:**

- Adding projectOwnerId to Projects.Create [\#251](https://github.com/mongodb/go-client-mongodb-atlas/pull/251) ([andreaangiolillo](https://github.com/andreaangiolillo))

**Closed issues:**

- ClusterID attribute to Processes.List returns 400 error response  [\#254](https://github.com/mongodb/go-client-mongodb-atlas/issues/254)
- Wrong type for cluster diskSizeGB [\#140](https://github.com/mongodb/go-client-mongodb-atlas/issues/140)

**Merged pull requests:**

- INTMDB-216: Added parameter version release system for advanced cluster [\#258](https://github.com/mongodb/go-client-mongodb-atlas/pull/258) ([coderGo93](https://github.com/coderGo93))
- task: run go mod tidy [\#256](https://github.com/mongodb/go-client-mongodb-atlas/pull/256) ([gssbzn](https://github.com/gssbzn))
- docs: clarify processes ClusterID [\#255](https://github.com/mongodb/go-client-mongodb-atlas/pull/255) ([gssbzn](https://github.com/gssbzn))
- INTMDB-211: Added shard keys properties to Managed Namespace in global configurat… [\#253](https://github.com/mongodb/go-client-mongodb-atlas/pull/253) ([abner-dou](https://github.com/abner-dou))
- CLOUDP-97442: update performance advisor with new endpoints -\> atlas go client [\#252](https://github.com/mongodb/go-client-mongodb-atlas/pull/252) ([andreaangiolillo](https://github.com/andreaangiolillo))
- INTMDB-252: Added two parameters for Process args [\#250](https://github.com/mongodb/go-client-mongodb-atlas/pull/250) ([coderGo93](https://github.com/coderGo93))
- INTMDB-230: Added autoDeferOnceEnabled property to MaintenanceWindow struct [\#249](https://github.com/mongodb/go-client-mongodb-atlas/pull/249) ([abner-dou](https://github.com/abner-dou))
- CLOUDP-82836: Parse service version [\#248](https://github.com/mongodb/go-client-mongodb-atlas/pull/248) ([fmenezes](https://github.com/fmenezes))

## [v0.12.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.12.0) (August 30, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.11.0...v0.12.0)

**Breaking changes:**

- CLOUDP-98424: Do not parse parts of the request in atlas search [\#244](https://github.com/mongodb/go-client-mongodb-atlas/pull/244) ([fmenezes](https://github.com/fmenezes))

**Merged pull requests:**

- CLOUDP-98424: Add unit tests for atlas search [\#246](https://github.com/mongodb/go-client-mongodb-atlas/pull/246) ([fmenezes](https://github.com/fmenezes))
- CLOUDP-98679: Add support for synonyms [\#245](https://github.com/mongodb/go-client-mongodb-atlas/pull/245) ([fmenezes](https://github.com/fmenezes))
- feat: add support for creation of cloudgov only projects [\#243](https://github.com/mongodb/go-client-mongodb-atlas/pull/243) ([gssbzn](https://github.com/gssbzn))
- feat: update supported go versions [\#242](https://github.com/mongodb/go-client-mongodb-atlas/pull/242) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-97443: add access tracking to the atlas go client [\#241](https://github.com/mongodb/go-client-mongodb-atlas/pull/241) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Renamed LiveMigrationOp to LiveMigrationServiceOp [\#240](https://github.com/mongodb/go-client-mongodb-atlas/pull/240) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-97549: Add cloud migration endpoint to the atlas go client [\#239](https://github.com/mongodb/go-client-mongodb-atlas/pull/239) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.11.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.11.0) (August 10, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.10.1...v0.11.0)

**Merged pull requests:**

- Adding AutoDefer method to automatically defer any maintenance [\#237](https://github.com/mongodb/go-client-mongodb-atlas/pull/237) ([vgarcia-te](https://github.com/vgarcia-te))
- CLOUDP-96138: Add serverless to the atlas go client [\#236](https://github.com/mongodb/go-client-mongodb-atlas/pull/236) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.10.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.10.1) (July 09, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.10.0...v0.10.1)

**Breaking changes:**

- Renamed AdvancedRegionSpec to AdvancedReplicationSpec [\#230](https://github.com/mongodb/go-client-mongodb-atlas/pull/230) ([andreaangiolillo](https://github.com/andreaangiolillo))

**Merged pull requests:**

- Bug: The atlas user create was calling the wrong endpoint [\#234](https://github.com/mongodb/go-client-mongodb-atlas/pull/234) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Fix project invitation update [\#233](https://github.com/mongodb/go-client-mongodb-atlas/pull/233) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Fixed Invite New User endpoints for organizations and projects [\#232](https://github.com/mongodb/go-client-mongodb-atlas/pull/232) ([andreaangiolillo](https://github.com/andreaangiolillo))
- task: release 0.10.0 prep [\#231](https://github.com/mongodb/go-client-mongodb-atlas/pull/231) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.10.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.10.0) (June 30, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.9.0...v0.10.0)

**Breaking changes:**

- Renamed AdvancedRegionSpec to AdvancedReplicationSpec [\#230](https://github.com/mongodb/go-client-mongodb-atlas/pull/230) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.9.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.9.0) (June 29, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.8.0...v0.9.0)

**Breaking changes:**

- INTMDB-17: Update search [\#223](https://github.com/mongodb/go-client-mongodb-atlas/pull/223) ([abner-dou](https://github.com/abner-dou))
- CLOUDP-91945: Add support for endpoints v1.5 [\#212](https://github.com/mongodb/go-client-mongodb-atlas/pull/212) ([andreaangiolillo](https://github.com/andreaangiolillo))
  - This only affects users setting a custom base URL which now should not include the `api/{version}}` segment 
- feat: remove project whitelist [\#206](https://github.com/mongodb/go-client-mongodb-atlas/pull/206) ([gssbzn](https://github.com/gssbzn))

**Merged pull requests:**

- CLOUDP-93382: Add Programmatic Invite Management for project to the atlas go client [\#228](https://github.com/mongodb/go-client-mongodb-atlas/pull/228) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-93381: Add Programmatic Invite Management for organization to the atlas go client [\#227](https://github.com/mongodb/go-client-mongodb-atlas/pull/227) ([andreaangiolillo](https://github.com/andreaangiolillo))
- tests: enable ifshort [\#226](https://github.com/mongodb/go-client-mongodb-atlas/pull/226) ([gssbzn](https://github.com/gssbzn))
- task: enable thelper and errorlint [\#225](https://github.com/mongodb/go-client-mongodb-atlas/pull/225) ([gssbzn](https://github.com/gssbzn))
- added custom analyzer structs [\#224](https://github.com/mongodb/go-client-mongodb-atlas/pull/224) ([abner-dou](https://github.com/abner-dou))
- lint: replace golint with revive [\#222](https://github.com/mongodb/go-client-mongodb-atlas/pull/222) ([gssbzn](https://github.com/gssbzn))
- lint: enable godot [\#221](https://github.com/mongodb/go-client-mongodb-atlas/pull/221) ([gssbzn](https://github.com/gssbzn))
- docs: update README [\#220](https://github.com/mongodb/go-client-mongodb-atlas/pull/220) ([gssbzn](https://github.com/gssbzn))
- Addded autoscaling to the advance cluster [\#219](https://github.com/mongodb/go-client-mongodb-atlas/pull/219) ([andreaangiolillo](https://github.com/andreaangiolillo))
- task: check dependencies on Tuesday [\#218](https://github.com/mongodb/go-client-mongodb-atlas/pull/218) ([gssbzn](https://github.com/gssbzn))
- fix: fix alert configurations path [\#217](https://github.com/mongodb/go-client-mongodb-atlas/pull/217) ([gssbzn](https://github.com/gssbzn))
- added update-all-analyzers method to search [\#216](https://github.com/mongodb/go-client-mongodb-atlas/pull/216) ([abner-dou](https://github.com/abner-dou))
- Fixed failing tests [\#215](https://github.com/mongodb/go-client-mongodb-atlas/pull/215) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Bug Fix: Add CloudURL [\#214](https://github.com/mongodb/go-client-mongodb-atlas/pull/214) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-92196: add advanced cluster endpoints in the atlas go client [\#213](https://github.com/mongodb/go-client-mongodb-atlas/pull/213) ([andreaangiolillo](https://github.com/andreaangiolillo))
- INTMDB-145: Support delete function for Cloud Backup Schedules [\#210](https://github.com/mongodb/go-client-mongodb-atlas/pull/210) ([coderGo93](https://github.com/coderGo93))
- INTMDB-19: Added fields in datalake databases and stores [\#209](https://github.com/mongodb/go-client-mongodb-atlas/pull/209) ([coderGo93](https://github.com/coderGo93))
- chore\(deps\): bump actions/cache from 2.1.5 to 2.1.6 [\#208](https://github.com/mongodb/go-client-mongodb-atlas/pull/208) ([dependabot[bot]](https://github.com/apps/dependabot))
- Adding missing args to do a custom online archive [\#207](https://github.com/mongodb/go-client-mongodb-atlas/pull/207) ([leofigy](https://github.com/leofigy))
- chore\(deps\): bump actions/checkout from 2 to 2.3.4 [\#205](https://github.com/mongodb/go-client-mongodb-atlas/pull/205) ([dependabot[bot]](https://github.com/apps/dependabot))
- chore\(deps\): bump actions/cache from 2 to 2.1.5 [\#204](https://github.com/mongodb/go-client-mongodb-atlas/pull/204) ([dependabot[bot]](https://github.com/apps/dependabot))
- task: add CODEOWNERS [\#203](https://github.com/mongodb/go-client-mongodb-atlas/pull/203) ([gssbzn](https://github.com/gssbzn))
- feat: update golancgci and enable tagliatelle [\#202](https://github.com/mongodb/go-client-mongodb-atlas/pull/202) ([gssbzn](https://github.com/gssbzn))
- ldap group is admin db [\#201](https://github.com/mongodb/go-client-mongodb-atlas/pull/201) ([leofigy](https://github.com/leofigy))

## [v0.8.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.8.0) (April 20, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.7.2...v0.8.0)

**Closed issues:**

- Delete user does not work for IAM based users [\#189](https://github.com/mongodb/go-client-mongodb-atlas/issues/189)

**Merged pull requests:**

- feat: move cloud provider regions method to cluster [\#198](https://github.com/mongodb/go-client-mongodb-atlas/pull/198) ([gssbzn](https://github.com/gssbzn))
- fix: fix username path escape for db username [\#199](https://github.com/mongodb/go-client-mongodb-atlas/pull/199) ([gssbzn](https://github.com/gssbzn))
- docs fix godoc rendering [\#196](https://github.com/mongodb/go-client-mongodb-atlas/pull/196) ([gssbzn](https://github.com/gssbzn))
- task: update golangci-lint to 1.39.0 [\#197](https://github.com/mongodb/go-client-mongodb-atlas/pull/197) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-86937: Add support to the Atlas Go Client to load sample data into a cluster [\#195](https://github.com/mongodb/go-client-mongodb-atlas/pull/195) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-86337: Update atlas region service to use new parameters - atlas go client [\#193](https://github.com/mongodb/go-client-mongodb-atlas/pull/193) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-85933: Add service to get the public ip to the atlas go client [\#191](https://github.com/mongodb/go-client-mongodb-atlas/pull/191) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-85749: Add new available region endpoint to the atlas go client [\#190](https://github.com/mongodb/go-client-mongodb-atlas/pull/190) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-83026: support for cluster status endpoint [\#188](https://github.com/mongodb/go-client-mongodb-atlas/pull/190) ([antonlisovenko](https://github.com/antonlisovenko))
- feat: add regions to GCP containers [\#186](https://github.com/mongodb/go-client-mongodb-atlas/pull/186) ([gssbzn](https://github.com/gssbzn))
- docs: update contributing guidelines [\#185](https://github.com/mongodb/go-client-mongodb-atlas/pull/185) ([gssbzn](https://github.com/gssbzn))
- build: add go1.16 tests [\#184](https://github.com/mongodb/go-client-mongodb-atlas/pull/184) ([gssbzn](https://github.com/gssbzn))
- lint: update GOLANGCI_VERSION to 1.37.0 [\#183](https://github.com/mongodb/go-client-mongodb-atlas/pull/183) ([gssbzn](https://github.com/gssbzn))

## [v0.7.2](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.7.2) (February 02, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.7.1...v0.7.2)

**Merged pull requests:**

- INTMDB-175: Added azure status in endpoint service [\#179](https://github.com/mongodb/go-client-mongodb-atlas/pull/179) ([coderGo93](https://github.com/coderGo93))

## [v0.7.1](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.7.1) (January 27, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.7.0...v0.7.1)

**Closed issues:**

- Missing `components` in `CloudProviderSnapshotRestoreJob` [\#171](https://github.com/mongodb/go-client-mongodb-atlas/issues/171)

**Merged pull requests:**

- INTMDB-173: Update the path of AccessListAPIKeys service [\#178](https://github.com/mongodb/go-client-mongodb-atlas/pull/178) ([andreaangiolillo](https://github.com/andreaangiolillo))
- INTMDB-158: changed field type to another object [\#177](https://github.com/mongodb/go-client-mongodb-atlas/pull/177) ([coderGo93](https://github.com/coderGo93))
- Apply License header to source files [\#175](https://github.com/mongodb/go-client-mongodb-atlas/pull/175) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-81043: add Regionalized Private Endpoint Status to atlas go client [\#174](https://github.com/mongodb/go-client-mongodb-atlas/pull/174) ([andreaangiolillo](https://github.com/andreaangiolillo))
- INTMDB-167: URL encode private enpoint ID [\#173](https://github.com/mongodb/go-client-mongodb-atlas/pull/173) ([andreaangiolillo](https://github.com/andreaangiolillo))
- Add support for components in `CloudProviderSnapshotRestoreJob` [\#172](https://github.com/mongodb/go-client-mongodb-atlas/pull/172) ([Stovoy](https://github.com/Stovoy))
- Add a helper method to group client options [\#170](https://github.com/mongodb/go-client-mongodb-atlas/pull/170) ([gssbzn](https://github.com/gssbzn))
- INTMDB-165: Update connection strings for cluster [\#169](https://github.com/mongodb/go-client-mongodb-atlas/pull/169) ([gssbzn](https://github.com/gssbzn))
- INTMDB-164: Update data lake to support new required fields for creation [\#168](https://github.com/mongodb/go-client-mongodb-atlas/pull/168) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-80413: \[go-client\] add service to get default mongodb major version [\#167](https://github.com/mongodb/go-client-mongodb-atlas/pull/167) ([gssbzn](https://github.com/gssbzn))

## [v0.7.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.7.0) (January 12, 2021)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.6.0...v0.7.0)

**Merged pull requests:**

- INTMDB-157: Add deprecation notice [\#165](https://github.com/mongodb/go-client-mongodb-atlas/pull/165) ([themantissa](https://github.com/themantissa))
- Fix broken link in api documentation for CloudProviderSnapshotRestoreJobsService. [\#164](https://github.com/mongodb/go-client-mongodb-atlas/pull/164) ([Stovoy](https://github.com/Stovoy))
- INTMDB-139: Add access list for programmatic API keys [\#163](https://github.com/mongodb/go-client-mongodb-atlas/pull/163) ([coderGo93](https://github.com/coderGo93))
- INTMDB-144: deleted omitempty for scopes of database users [\#162](https://github.com/mongodb/go-client-mongodb-atlas/pull/162) ([coderGo93](https://github.com/coderGo93))
- INTMDB-138: Post release version bump [\#161](https://github.com/mongodb/go-client-mongodb-atlas/pull/161) ([gssbzn](https://github.com/gssbzn))

## [v0.6.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.6.0) (December 14, 2020)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.5.0...v0.6.0)

**Breaking changes:**

- CLOUDP-78002: adds support for Atlas Error codes [\#157](https://github.com/mongodb/go-client-mongodb-atlas/pull/157) ([antonlisovenko](https://github.com/antonlisovenko))
- fix: update online archive structs [\#155](https://github.com/mongodb/go-client-mongodb-atlas/pull/155) ([gssbzn](https://github.com/gssbzn))
- INTMDB-122: Private Endpoints [\#148](https://github.com/mongodb/go-client-mongodb-atlas/pull/148) ([coderGo93](https://github.com/coderGo93))
- add OrganizationsListOptions for filtering organization list results [\#144](https://github.com/mongodb/go-client-mongodb-atlas/pull/144) ([sunny-b](https://github.com/sunny-b))

**Closed issues:**

- Add "Name" ListOption for filtering list results [\#143](https://github.com/mongodb/go-client-mongodb-atlas/issues/143)

**Merged pull requests:**

- INTMDB-134: Added parameter roleId and valid for AWS KMS [\#158](https://github.com/mongodb/go-client-mongodb-atlas/pull/158) ([coderGo93](https://github.com/coderGo93))
- fix: improve connection closing [\#156](https://github.com/mongodb/go-client-mongodb-atlas/pull/156) ([gssbzn](https://github.com/gssbzn))
- test: move from travis to github actions [\#154](https://github.com/mongodb/go-client-mongodb-atlas/pull/154) ([gssbzn](https://github.com/gssbzn))
- Provider access integration [\#153](https://github.com/mongodb/go-client-mongodb-atlas/pull/153) ([leofigy](https://github.com/leofigy))
- Add Raw field to atlas response [\#152](https://github.com/mongodb/go-client-mongodb-atlas/pull/152) ([svagner](https://github.com/svagner))
- Fix private endpoints [\#151](https://github.com/mongodb/go-client-mongodb-atlas/pull/151) ([coderGo93](https://github.com/coderGo93))
- NOREF: Include userAlias for processes [\#150](https://github.com/mongodb/go-client-mongodb-atlas/pull/150) ([howdoicomputer](https://github.com/howdoicomputer))
- INTMDB-132: Add support for Cloud Provider Access [\#147](https://github.com/mongodb/go-client-mongodb-atlas/pull/147) ([gssbzn](https://github.com/gssbzn))
- docs: fix AlertConfigurationsService docs [\#146](https://github.com/mongodb/go-client-mongodb-atlas/pull/146) ([gssbzn](https://github.com/gssbzn))
- letting 0 be as value in the threshold [\#145](https://github.com/mongodb/go-client-mongodb-atlas/pull/145) ([leofigy](https://github.com/leofigy))
- chore: update versions to next desired [\#142](https://github.com/mongodb/go-client-mongodb-atlas/pull/142) ([gssbzn](https://github.com/gssbzn))

## [v0.5.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.5.0) (October 01, 2020)

[Full Changelog](https://github.com/mongodb/go-client-mongodb-atlas/compare/v0.4.0...v0.5.0)

**Closed issues:**

- Wrong import path? [\#119](https://github.com/mongodb/go-client-mongodb-atlas/issues/119)

**Merged pull requests:**

- feat: add access list API [\#139](https://github.com/mongodb/go-client-mongodb-atlas/pull/139) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-73368: fixing the Operation struct in the atlas go client [\#138](https://github.com/mongodb/go-client-mongodb-atlas/pull/138) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-72746: Add Performance Advisor API call to atlas go client [\#137](https://github.com/mongodb/go-client-mongodb-atlas/pull/137) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-72727: Fix LDAPConfigurationsService.Save to pass LDAPConfiguration as input [\#136](https://github.com/mongodb/go-client-mongodb-atlas/pull/136) ([andreaangiolillo](https://github.com/andreaangiolillo))
- CLOUDP-72643: Add LDAP Configurations to atlas go client [\#135](https://github.com/mongodb/go-client-mongodb-atlas/pull/135) ([andreaangiolillo](https://github.com/andreaangiolillo))
- feat: add IsDeleted to organizations [\#134](https://github.com/mongodb/go-client-mongodb-atlas/pull/134) ([gssbzn](https://github.com/gssbzn))
- Setting the right auth db name  [\#133](https://github.com/mongodb/go-client-mongodb-atlas/pull/133) ([leofigy](https://github.com/leofigy))
- doc: update and fix a lot of the go doc references [\#132](https://github.com/mongodb/go-client-mongodb-atlas/pull/132) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-70872: Adding Third-Party Integration Settings [\#131](https://github.com/mongodb/go-client-mongodb-atlas/pull/131) ([andreaangiolillo](https://github.com/andreaangiolillo))
- chore: update supported go versions [\#130](https://github.com/mongodb/go-client-mongodb-atlas/pull/130) ([gssbzn](https://github.com/gssbzn))
- CLOUDP-70755: Add DELETE /groups/{GROUP-ID}/users/{USER-ID} to atlas go client [\#129](https://github.com/mongodb/go-client-mongodb-atlas/pull/129) ([andreaangiolillo](https://github.com/andreaangiolillo))
- feature: add scope support for db users [\#128](https://github.com/mongodb/go-client-mongodb-atlas/pull/128) ([gssbzn](https://github.com/gssbzn))
- Added threshold attribute [\#127](https://github.com/mongodb/go-client-mongodb-atlas/pull/127) ([PacoDw](https://github.com/PacoDw))
- Custom db roles [\#126](https://github.com/mongodb/go-client-mongodb-atlas/pull/126) ([PacoDw](https://github.com/PacoDw))
- CLOUDP-69426: Atlas go client - Update PeersService.List to use providerName [\#125](https://github.com/mongodb/go-client-mongodb-atlas/pull/125) ([andreaangiolillo](https://github.com/andreaangiolillo))

## [v0.4.0](https://github.com/mongodb/go-client-mongodb-atlas/tree/v0.4.0) (August 10, 2020)

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
