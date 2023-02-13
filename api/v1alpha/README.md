# Go API client
## Documentation for API Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AWSClustersDNSApi* | [**ReturnOneCustomDnsConfigurationForAtlasClustersOnAws**](docs/AWSClustersDNSApi.md#returnonecustomdnsconfigurationforatlasclustersonaws) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS
*AWSClustersDNSApi* | [**ToggleOneStateOfOneCustomDnsConfigurationForAtlasClustersOnAws**](docs/AWSClustersDNSApi.md#toggleonestateofonecustomdnsconfigurationforatlasclustersonaws) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS
*AccessTrackingApi* | [**ReturnDatabaseAccessHistoryForOneClusterUsingItsClusterName**](docs/AccessTrackingApi.md#returndatabaseaccesshistoryforoneclusterusingitsclustername) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster using Its Cluster Name
*AccessTrackingApi* | [**ReturnDatabaseAccessHistoryForOneClusterUsingItsHostname**](docs/AccessTrackingApi.md#returndatabaseaccesshistoryforoneclusterusingitshostname) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster using Its Hostname
*AlertConfigurationsApi* | [**CreateOneAlertConfigurationInOneProject**](docs/AlertConfigurationsApi.md#createonealertconfigurationinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project
*AlertConfigurationsApi* | [**RemoveOneAlertConfigurationFromOneProject**](docs/AlertConfigurationsApi.md#removeonealertconfigurationfromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project
*AlertConfigurationsApi* | [**ReturnAlertConfigMatchersFieldNames**](docs/AlertConfigurationsApi.md#returnalertconfigmatchersfieldnames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Get All Alert Configuration Matchers Field Names
*AlertConfigurationsApi* | [**ReturnAllAlertConfigurationsForOneProject**](docs/AlertConfigurationsApi.md#returnallalertconfigurationsforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations for One Project
*AlertConfigurationsApi* | [**ReturnOneAlertConfigurationFromOneProject**](docs/AlertConfigurationsApi.md#returnonealertconfigurationfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project
*AlertConfigurationsApi* | [**ToggleOneStateOfOneAlertConfigurationInOneProject**](docs/AlertConfigurationsApi.md#toggleonestateofonealertconfigurationinoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle One State of One Alert Configuration in One Project
*AlertConfigurationsApi* | [**UpdateOneAlertConfigurationForOneProject**](docs/AlertConfigurationsApi.md#updateonealertconfigurationforoneproject) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration for One Project
*AlertsApi* | [**AcknowledgeOneAlertFromOneProject**](docs/AlertsApi.md#acknowledgeonealertfromoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project
*AlertsApi* | [**ReturnAllAlertsFromOneProject**](docs/AlertsApi.md#returnallalertsfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project
*AlertsApi* | [**ReturnAllOpenAlertsForAlertConfiguration**](docs/AlertsApi.md#returnallopenalertsforalertconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertConfigId}/alerts | Return All Open Alerts for Alert Configuration
*AlertsApi* | [**ReturnOneAlertFromOneProject**](docs/AlertsApi.md#returnonealertfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project
*AtlasSearchApi* | [**CreateOneAtlasSearchIndex**](docs/AtlasSearchApi.md#createoneatlassearchindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
*AtlasSearchApi* | [**RemoveOneAtlasSearchIndex**](docs/AtlasSearchApi.md#removeoneatlassearchindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
*AtlasSearchApi* | [**ReturnAllAtlasSearchIndexesForOneCollection**](docs/AtlasSearchApi.md#returnallatlassearchindexesforonecollection) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
*AtlasSearchApi* | [**ReturnAllUserDefinedAnalyzersForOneCluster**](docs/AtlasSearchApi.md#returnalluserdefinedanalyzersforonecluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/analyzers | Return All User-Defined Analyzers for One Cluster
*AtlasSearchApi* | [**ReturnOneAtlasSearchIndex**](docs/AtlasSearchApi.md#returnoneatlassearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
*AtlasSearchApi* | [**UpdateAllUserDefinedAnalyzersForOneCluster**](docs/AtlasSearchApi.md#updatealluserdefinedanalyzersforonecluster) | **Put** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/analyzers | Update All User-Defined Analyzers for One Cluster
*AtlasSearchApi* | [**UpdateOneAtlasSearchIndex**](docs/AtlasSearchApi.md#updateoneatlassearchindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index
*AuditingApi* | [**ReturnOneAuditingConfigurationForOneProject**](docs/AuditingApi.md#returnoneauditingconfigurationforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/auditLog | Return the Auditing Configuration for One Project
*AuditingApi* | [**UpdateAuditingConfigurationForOneProject**](docs/AuditingApi.md#updateauditingconfigurationforoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/auditLog | Update Auditing Configuration for One Project
*CloudBackupsApi* | [**CancelOneRestoreJobOfOneCluster**](docs/CloudBackupsApi.md#cancelonerestorejobofonecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
*CloudBackupsApi* | [**ChangeExpirationDateForOneCloudBackup**](docs/CloudBackupsApi.md#changeexpirationdateforonecloudbackup) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup
*CloudBackupsApi* | [**CreateOneCloudBackupSnapshotExportJob**](docs/CloudBackupsApi.md#createonecloudbackupsnapshotexportjob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [**GrantAccessToAwsS3BucketForCloudBackupSnapshotExports**](docs/CloudBackupsApi.md#grantaccesstoawss3bucketforcloudbackupsnapshotexports) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**RemoveAllCloudBackupSchedules**](docs/CloudBackupsApi.md#removeallcloudbackupschedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
*CloudBackupsApi* | [**RemoveOneReplicaSetCloudBackup**](docs/CloudBackupsApi.md#removeonereplicasetcloudbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
*CloudBackupsApi* | [**RemoveOneShardedClusterCloudBackup**](docs/CloudBackupsApi.md#removeoneshardedclustercloudbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [**RestoreOneSnapshotOfOneCluster**](docs/CloudBackupsApi.md#restoreonesnapshotofonecluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
*CloudBackupsApi* | [**RestoreOneSnapshotOfOneCluster1**](docs/CloudBackupsApi.md#restoreonesnapshotofonecluster1) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
*CloudBackupsApi* | [**ReturnAllAwsS3BucketsUsedForCloudBackupSnapshotExports**](docs/CloudBackupsApi.md#returnallawss3bucketsusedforcloudbackupsnapshotexports) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**ReturnAllCloudBackupSnapshotExportJobs**](docs/CloudBackupsApi.md#returnallcloudbackupsnapshotexportjobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Cloud Backup Snapshot Export Jobs
*CloudBackupsApi* | [**ReturnAllReplicaSetCloudBackups**](docs/CloudBackupsApi.md#returnallreplicasetcloudbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
*CloudBackupsApi* | [**ReturnAllRestoreJobsForOneCluster**](docs/CloudBackupsApi.md#returnallrestorejobsforonecluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
*CloudBackupsApi* | [**ReturnAllRestoreJobsForOneServerlessInstance**](docs/CloudBackupsApi.md#returnallrestorejobsforoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
*CloudBackupsApi* | [**ReturnAllShardedClusterCloudBackups**](docs/CloudBackupsApi.md#returnallshardedclustercloudbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
*CloudBackupsApi* | [**ReturnAllSnapshotsOfOneServerlessInstance**](docs/CloudBackupsApi.md#returnallsnapshotsofoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
*CloudBackupsApi* | [**ReturnOneAwsS3BucketUsedForCloudBackupSnapshotExports**](docs/CloudBackupsApi.md#returnoneawss3bucketusedforcloudbackupsnapshotexports) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**ReturnOneCloudBackupSchedule**](docs/CloudBackupsApi.md#returnonecloudbackupschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
*CloudBackupsApi* | [**ReturnOneCloudBackupSnapshotExportJob**](docs/CloudBackupsApi.md#returnonecloudbackupsnapshotexportjob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [**ReturnOneReplicaSetCloudBackup**](docs/CloudBackupsApi.md#returnonereplicasetcloudbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
*CloudBackupsApi* | [**ReturnOneRestoreJobForOneServerlessInstance**](docs/CloudBackupsApi.md#returnonerestorejobforoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
*CloudBackupsApi* | [**ReturnOneRestoreJobOfOneCluster**](docs/CloudBackupsApi.md#returnonerestorejobofonecluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster
*CloudBackupsApi* | [**ReturnOneShardedClusterCloudBackup**](docs/CloudBackupsApi.md#returnoneshardedclustercloudbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [**ReturnOneSnapshotOfOneServerlessInstance**](docs/CloudBackupsApi.md#returnonesnapshotofoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
*CloudBackupsApi* | [**RevokeAccessToAwsS3BucketForCloudBackupSnapshotExports**](docs/CloudBackupsApi.md#revokeaccesstoawss3bucketforcloudbackupsnapshotexports) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**TakeOneOnDemandSnapshot**](docs/CloudBackupsApi.md#takeoneondemandsnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot
*CloudBackupsApi* | [**UpdateCloudBackupBackupPolicyForOneCluster**](docs/CloudBackupsApi.md#updatecloudbackupbackuppolicyforonecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster
*CloudMigrationServiceApi* | [**CreateLinkToken**](docs/CloudMigrationServiceApi.md#createlinktoken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token
*CloudMigrationServiceApi* | [**CreatePushMigration**](docs/CloudMigrationServiceApi.md#createpushmigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas
*CloudMigrationServiceApi* | [**CutoverOneMigration**](docs/CloudMigrationServiceApi.md#cutoveronemigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over the Migrated Cluster
*CloudMigrationServiceApi* | [**DeleteLinkToken**](docs/CloudMigrationServiceApi.md#deletelinktoken) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token
*CloudMigrationServiceApi* | [**ListProjects**](docs/CloudMigrationServiceApi.md#listprojects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration
*CloudMigrationServiceApi* | [**ReturnOnePushMigration**](docs/CloudMigrationServiceApi.md#returnonepushmigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job
*CloudMigrationServiceApi* | [**ReturnOneValidationJob**](docs/CloudMigrationServiceApi.md#returnonevalidationjob) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job
*CloudMigrationServiceApi* | [**ValidateOneMigration**](docs/CloudMigrationServiceApi.md#validateonemigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request
*CloudProviderAccessApi* | [**AuthorizeOneCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#authorizeonecloudprovideraccessrole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [**CreateOneCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#createonecloudprovideraccessrole) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
*CloudProviderAccessApi* | [**DeauthorizeOneCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#deauthorizeonecloudprovideraccessrole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [**ReturnAllCloudProviderAccessRoles**](docs/CloudProviderAccessApi.md#returnallcloudprovideraccessroles) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles
*CloudProviderAccessApi* | [**ReturnCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#returncloudprovideraccessrole) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return specified Cloud Provider Access Role
*ClusterOutageSimulationApi* | [**EndOutageSimulation**](docs/ClusterOutageSimulationApi.md#endoutagesimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | End an Outage Simulation
*ClusterOutageSimulationApi* | [**GetOutageSimulation**](docs/ClusterOutageSimulationApi.md#getoutagesimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation
*ClusterOutageSimulationApi* | [**StartOutageSimulation**](docs/ClusterOutageSimulationApi.md#startoutagesimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start an Outage Simulation
*ClustersApi* | [**CheckStatusOfClusterSampleDatasetRequest**](docs/ClustersApi.md#checkstatusofclustersampledatasetrequest) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Check Status of Cluster Sample Dataset Request
*ClustersApi* | [**LoadSampleDatasetRequestIntoCluster**](docs/ClustersApi.md#loadsampledatasetrequestintocluster) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset Request into Cluster
*ClustersApi* | [**ReturnAllAuthorizedClustersInAllProjects**](docs/ClustersApi.md#returnallauthorizedclustersinallprojects) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects
*ClustersApi* | [**ReturnAllCloudProviderRegions**](docs/ClustersApi.md#returnallcloudproviderregions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions
*ClustersApi* | [**ReturnOneAdvancedConfigurationOptionsForOneCluster**](docs/ClustersApi.md#returnoneadvancedconfigurationoptionsforonecluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return One Advanced Configuration Options for One Cluster
*ClustersApi* | [**ReturnOneStatusOfAllClusterOperations**](docs/ClustersApi.md#returnonestatusofallclusteroperations) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations
*ClustersApi* | [**UpdateAdvancedConfigurationOptionsForOneCluster**](docs/ClustersApi.md#updateadvancedconfigurationoptionsforonecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster
*ClustersApi* | [**UpgradeOneTenantCluster**](docs/ClustersApi.md#upgradeonetenantcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-tier Cluster
*CustomDatabaseRolesApi* | [**CreateOneCustomRole**](docs/CustomDatabaseRolesApi.md#createonecustomrole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role
*CustomDatabaseRolesApi* | [**RemoveOneCustomRoleFromOneProject**](docs/CustomDatabaseRolesApi.md#removeonecustomrolefromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
*CustomDatabaseRolesApi* | [**ReturnAllCustomRolesInOneProject**](docs/CustomDatabaseRolesApi.md#returnallcustomrolesinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
*CustomDatabaseRolesApi* | [**ReturnOneCustomRoleInOneProject**](docs/CustomDatabaseRolesApi.md#returnonecustomroleinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
*CustomDatabaseRolesApi* | [**UpdateOneCustomRoleInOneProject**](docs/CustomDatabaseRolesApi.md#updateonecustomroleinoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project
*DataFederationApi* | [**CreateOneDataFederationPrivateEndpointForOneProject**](docs/DataFederationApi.md#createonedatafederationprivateendpointforoneproject) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project
*DataFederationApi* | [**CreateOneFederatedDatabaseInOneProject**](docs/DataFederationApi.md#createonefederateddatabaseinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project
*DataFederationApi* | [**DownloadQueryLogsForOneFederatedDatabase**](docs/DataFederationApi.md#downloadquerylogsforonefederateddatabase) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance
*DataFederationApi* | [**RemoveOneDataFederationPrivateEndpointFromOneProject**](docs/DataFederationApi.md#removeonedatafederationprivateendpointfromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project
*DataFederationApi* | [**RemoveOneFederatedDatabaseFromOneProject**](docs/DataFederationApi.md#removeonefederateddatabasefromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project
*DataFederationApi* | [**ReturnAllDataFederationPrivateEndpointsInOneProject**](docs/DataFederationApi.md#returnalldatafederationprivateendpointsinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project
*DataFederationApi* | [**ReturnAllFederatedDatabasesInOneProject**](docs/DataFederationApi.md#returnallfederateddatabasesinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project
*DataFederationApi* | [**ReturnOneDataFederationPrivateEndpointInOneProject**](docs/DataFederationApi.md#returnonedatafederationprivateendpointinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project
*DataFederationApi* | [**ReturnOneFederatedDatabaseInOneProject**](docs/DataFederationApi.md#returnonefederateddatabaseinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project
*DataFederationApi* | [**UpdateOneFederatedDatabaseInOneProject**](docs/DataFederationApi.md#updateonefederateddatabaseinoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project
*DataLakePipelinesApi* | [**CreateOneDataLakePipeline**](docs/DataLakePipelinesApi.md#createonedatalakepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline
*DataLakePipelinesApi* | [**DeletePipelineRunDataset**](docs/DataLakePipelinesApi.md#deletepipelinerundataset) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete Pipeline Run Dataset
*DataLakePipelinesApi* | [**PauseOnePipelineInOneProject**](docs/DataLakePipelinesApi.md#pauseonepipelineinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline
*DataLakePipelinesApi* | [**RemoveOneDataLakePipeline**](docs/DataLakePipelinesApi.md#removeonedatalakepipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline
*DataLakePipelinesApi* | [**ResumeOnePipelineInOneProject**](docs/DataLakePipelinesApi.md#resumeonepipelineinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline
*DataLakePipelinesApi* | [**ReturnAllDataLakePipelineRunsFromOneProject**](docs/DataLakePipelinesApi.md#returnalldatalakepipelinerunsfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs from One Project
*DataLakePipelinesApi* | [**ReturnAllDataLakePipelinesFromOneProject**](docs/DataLakePipelinesApi.md#returnalldatalakepipelinesfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines from One Project
*DataLakePipelinesApi* | [**ReturnAvailableSchedulesForPipeline**](docs/DataLakePipelinesApi.md#returnavailableschedulesforpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return Available Ingestion Schedules for One Data Lake Pipeline
*DataLakePipelinesApi* | [**ReturnAvailableSnapshotsForPipeline**](docs/DataLakePipelinesApi.md#returnavailablesnapshotsforpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return Available Backup Snapshots for One Data Lake Pipeline
*DataLakePipelinesApi* | [**ReturnOnePipelineInOneProject**](docs/DataLakePipelinesApi.md#returnonepipelineinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline
*DataLakePipelinesApi* | [**ReturnOnePipelineRunInOneProject**](docs/DataLakePipelinesApi.md#returnonepipelineruninoneproject) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run
*DataLakePipelinesApi* | [**TriggerOneOnDemandSnapshotIngestion**](docs/DataLakePipelinesApi.md#triggeroneondemandsnapshotingestion) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger on demand snapshot ingestion
*DataLakePipelinesApi* | [**UpdateOneDataLakePipeline**](docs/DataLakePipelinesApi.md#updateonedatalakepipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline
*DatabaseUsersApi* | [**CreateDatabaseUser**](docs/DatabaseUsersApi.md#createdatabaseuser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project
*DatabaseUsersApi* | [**DeleteDatabaseUser**](docs/DatabaseUsersApi.md#deletedatabaseuser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project
*DatabaseUsersApi* | [**GetDatabaseUser**](docs/DatabaseUsersApi.md#getdatabaseuser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project
*DatabaseUsersApi* | [**ListDatabaseUsers**](docs/DatabaseUsersApi.md#listdatabaseusers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users from One Project
*DatabaseUsersApi* | [**UpdateDatabaseUser**](docs/DatabaseUsersApi.md#updatedatabaseuser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [**ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject**](docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#returnoneconfigurationforencryptionatrestusingcustomermanagedkeysforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [**UpdateConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject**](docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateconfigurationforencryptionatrestusingcustomermanagedkeysforoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EventsApi* | [**ReturnAllEventsFromOneOrganization**](docs/EventsApi.md#returnalleventsfromoneorganization) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return All Events from One Organization
*EventsApi* | [**ReturnAllEventsFromOneProject**](docs/EventsApi.md#returnalleventsfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/events | Return All Events from One Project
*EventsApi* | [**ReturnOneEventFromOneOrganization**](docs/EventsApi.md#returnoneeventfromoneorganization) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
*EventsApi* | [**ReturnOneEventFromOneProject**](docs/EventsApi.md#returnoneeventfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project
*FederatedAuthenticationApi* | [**CreateRoleMapping**](docs/FederatedAuthenticationApi.md#createrolemapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
*FederatedAuthenticationApi* | [**DeleteFederationApp**](docs/FederatedAuthenticationApi.md#deletefederationapp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete the federation settings instance.
*FederatedAuthenticationApi* | [**DeleteRoleMapping**](docs/FederatedAuthenticationApi.md#deleterolemapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
*FederatedAuthenticationApi* | [**GetAllConnectedOrgConfigs**](docs/FederatedAuthenticationApi.md#getallconnectedorgconfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from the Federation
*FederatedAuthenticationApi* | [**GetAllIdentityProviders**](docs/FederatedAuthenticationApi.md#getallidentityproviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return all identity providers from the specified federation.
*FederatedAuthenticationApi* | [**GetFederationSettings**](docs/FederatedAuthenticationApi.md#getfederationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
*FederatedAuthenticationApi* | [**GetIdentityProvider**](docs/FederatedAuthenticationApi.md#getidentityprovider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return one identity provider from the specified federation.
*FederatedAuthenticationApi* | [**GetIdentityProviderMetadata**](docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the metadata of one identity provider in the specified federation.
*FederatedAuthenticationApi* | [**GetOrganizationConfig**](docs/FederatedAuthenticationApi.md#getorganizationconfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**GetRoleMapping**](docs/FederatedAuthenticationApi.md#getrolemapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
*FederatedAuthenticationApi* | [**ListRoleMappings**](docs/FederatedAuthenticationApi.md#listrolemappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
*FederatedAuthenticationApi* | [**RemoveOrganizationConfig**](docs/FederatedAuthenticationApi.md#removeorganizationconfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**UpdateIdentityProvider**](docs/FederatedAuthenticationApi.md#updateidentityprovider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update the identity provider.
*FederatedAuthenticationApi* | [**UpdateOrganizationConfig**](docs/FederatedAuthenticationApi.md#updateorganizationconfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**UpdateRoleMapping**](docs/FederatedAuthenticationApi.md#updaterolemapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization
*GlobalClustersApi* | [**DeleteAllCustomZoneMappings**](docs/GlobalClustersApi.md#deleteallcustomzonemappings) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster
*GlobalClustersApi* | [**DeleteManagedNamespace**](docs/GlobalClustersApi.md#deletemanagednamespace) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Multi-Cloud Cluster
*InvoicesApi* | [**ReturnAllInvoicesForOneOrganization**](docs/InvoicesApi.md#returnallinvoicesforoneorganization) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization
*InvoicesApi* | [**ReturnAllPendingInvoicesForOneOrganization**](docs/InvoicesApi.md#returnallpendinginvoicesforoneorganization) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
*InvoicesApi* | [**ReturnOneOrganizationInvoice**](docs/InvoicesApi.md#returnoneorganizationinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Organization Invoice
*InvoicesApi* | [**ReturnOneOrganizationInvoiceAsCsv**](docs/InvoicesApi.md#returnoneorganizationinvoiceascsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Organization Invoice as CSV
*LDAPConfigurationApi* | [**ReturnOneCurrentLdapConfiguration**](docs/LDAPConfigurationApi.md#returnonecurrentldapconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
*LDAPConfigurationApi* | [**ReturnOneStatusOfOneVerifyLdapConfigurationRequest**](docs/LDAPConfigurationApi.md#returnonestatusofoneverifyldapconfigurationrequest) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
*LDAPConfigurationApi* | [**SaveOneLdapConfiguration**](docs/LDAPConfigurationApi.md#saveoneldapconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
*LDAPConfigurationApi* | [**VerifyOneLdapConfigurationInOneProject**](docs/LDAPConfigurationApi.md#verifyoneldapconfigurationinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project
*LegacyBackupApi* | [**ChangeOneLegacyBackupSnapshotExpiration**](docs/LegacyBackupApi.md#changeonelegacybackupsnapshotexpiration) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Change One Legacy Backup Snapshot Expiration
*LegacyBackupApi* | [**CreateOneLegacyBackupRestoreJob**](docs/LegacyBackupApi.md#createonelegacybackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job
*LegacyBackupApi* | [**RemoveOneLegacyBackupSnapshot**](docs/LegacyBackupApi.md#removeonelegacybackupsnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot
*LegacyBackupApi* | [**ReturnAllLegacyBackupCheckpoints**](docs/LegacyBackupApi.md#returnalllegacybackupcheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints
*LegacyBackupApi* | [**ReturnAllLegacyBackupRestoreJobs**](docs/LegacyBackupApi.md#returnalllegacybackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs
*LegacyBackupApi* | [**ReturnAllLegacyBackupSnapshots**](docs/LegacyBackupApi.md#returnalllegacybackupsnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots
*LegacyBackupApi* | [**ReturnOneLegacyBackupCheckpoint**](docs/LegacyBackupApi.md#returnonelegacybackupcheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint
*LegacyBackupApi* | [**ReturnOneLegacyBackupRestoreJob**](docs/LegacyBackupApi.md#returnonelegacybackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job
*LegacyBackupApi* | [**ReturnOneLegacyBackupSnapshot**](docs/LegacyBackupApi.md#returnonelegacybackupsnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot
*LegacyBackupApi* | [**ReturnOneSnapshotSchedule**](docs/LegacyBackupApi.md#returnonesnapshotschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule
*LegacyBackupApi* | [**UpdateSnapshotScheduleForOneCluster**](docs/LegacyBackupApi.md#updatesnapshotscheduleforonecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster
*MaintenanceWindowsApi* | [**ResetOneMaintenanceWindowForOneProject**](docs/MaintenanceWindowsApi.md#resetonemaintenancewindowforoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project
*MaintenanceWindowsApi* | [**ReturnOneMaintenanceWindowForOneProject**](docs/MaintenanceWindowsApi.md#returnonemaintenancewindowforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project
*MaintenanceWindowsApi* | [**UpdateMaintenanceWindowForOneProject**](docs/MaintenanceWindowsApi.md#updatemaintenancewindowforoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Update Maintenance Window for One Project
*MongoDBCloudUsersApi* | [**CreateOneCloudUser**](docs/MongoDBCloudUsersApi.md#createoneclouduser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
*MongoDBCloudUsersApi* | [**ReturnOneCloudUserUsingItsId**](docs/MongoDBCloudUsersApi.md#returnoneclouduserusingitsid) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User using Its ID
*MongoDBCloudUsersApi* | [**ReturnOneCloudUserUsingItsUsername**](docs/MongoDBCloudUsersApi.md#returnoneclouduserusingitsusername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User using Their Username
*MonitoringAndLogsApi* | [**GetHostLogs**](docs/MonitoringAndLogsApi.md#gethostlogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Multi-Cloud Cluster Host in One Project
*MonitoringAndLogsApi* | [**GetIndexMetrics**](docs/MonitoringAndLogsApi.md#getindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Specified Namespace
*MonitoringAndLogsApi* | [**GetMeasurements**](docs/MonitoringAndLogsApi.md#getmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics
*MonitoringAndLogsApi* | [**ListIndexMetrics**](docs/MonitoringAndLogsApi.md#listindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace
*MonitoringAndLogsApi* | [**ListMetricTypes**](docs/MonitoringAndLogsApi.md#listmetrictypes) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process
*MonitoringAndLogsApi* | [**ReturnAllMongodbProcessesInOneProject**](docs/MonitoringAndLogsApi.md#returnallmongodbprocessesinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project
*MonitoringAndLogsApi* | [**ReturnAvailableDatabasesForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returnavailabledatabasesforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process
*MonitoringAndLogsApi* | [**ReturnAvailableDisksForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returnavailabledisksforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process
*MonitoringAndLogsApi* | [**ReturnDatabasesForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returndatabasesforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for a MongoDB Process
*MonitoringAndLogsApi* | [**ReturnMeasurementsForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returnmeasurementsforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process
*MonitoringAndLogsApi* | [**ReturnMeasurementsOfOneDatabaseForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returnmeasurementsofonedatabaseforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements of One Database for One MongoDB Process
*MonitoringAndLogsApi* | [**ReturnMeasurementsOfOneDisk**](docs/MonitoringAndLogsApi.md#returnmeasurementsofonedisk) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements of One Disk
*MonitoringAndLogsApi* | [**ReturnMeasurementsOfOneDiskForOneMongodbProcess**](docs/MonitoringAndLogsApi.md#returnmeasurementsofonediskforonemongodbprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process
*MonitoringAndLogsApi* | [**ReturnOneMongodbProcessById**](docs/MonitoringAndLogsApi.md#returnonemongodbprocessbyid) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID
*MultiCloudClustersApi* | [**CreateCluster**](docs/MultiCloudClustersApi.md#createcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters | Create One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**DeleteCluster**](docs/MultiCloudClustersApi.md#deletecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Remove One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**GetCluster**](docs/MultiCloudClustersApi.md#getcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Return One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**ListClusters**](docs/MultiCloudClustersApi.md#listclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters | Return All Multi-Cloud Clusters from One Project
*MultiCloudClustersApi* | [**UpdateCluster**](docs/MultiCloudClustersApi.md#updatecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Modify One Multi-Cloud Cluster from One Project
*NetworkPeeringApi* | [**CreateOneNewNetworkPeeringConnection**](docs/NetworkPeeringApi.md#createonenewnetworkpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One New Network Peering Connection
*NetworkPeeringApi* | [**CreateOneNewNetworkPeeringContainer**](docs/NetworkPeeringApi.md#createonenewnetworkpeeringcontainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One New Network Peering Container
*NetworkPeeringApi* | [**DisableConnectViaPeeringOnlyModeForOneProject**](docs/NetworkPeeringApi.md#disableconnectviapeeringonlymodeforoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering Only Mode for One Project
*NetworkPeeringApi* | [**RemoveOneExistingNetworkPeeringConnection**](docs/NetworkPeeringApi.md#removeoneexistingnetworkpeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Existing Network Peering Connection
*NetworkPeeringApi* | [**RemoveOneNetworkPeeringContainer**](docs/NetworkPeeringApi.md#removeonenetworkpeeringcontainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
*NetworkPeeringApi* | [**ReturnAllNetworkPeeringConnectionsInOneProject**](docs/NetworkPeeringApi.md#returnallnetworkpeeringconnectionsinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project
*NetworkPeeringApi* | [**ReturnAllNetworkPeeringContainersInOneProject**](docs/NetworkPeeringApi.md#returnallnetworkpeeringcontainersinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
*NetworkPeeringApi* | [**ReturnAllNetworkPeeringContainersInOneProjectForOneCloudProvider**](docs/NetworkPeeringApi.md#returnallnetworkpeeringcontainersinoneprojectforonecloudprovider) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
*NetworkPeeringApi* | [**ReturnOneNetworkPeeringConnectionInOneProject**](docs/NetworkPeeringApi.md#returnonenetworkpeeringconnectioninoneproject) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
*NetworkPeeringApi* | [**ReturnOneNetworkPeeringContainer**](docs/NetworkPeeringApi.md#returnonenetworkpeeringcontainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
*NetworkPeeringApi* | [**UpdateOneNetworkPeeringContainer**](docs/NetworkPeeringApi.md#updateonenetworkpeeringcontainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container
*NetworkPeeringApi* | [**UpdateOneNewNetworkPeeringConnection**](docs/NetworkPeeringApi.md#updateonenewnetworkpeeringconnection) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One New Network Peering Connection
*NetworkPeeringApi* | [**VerifyConnectViaPeeringOnlyModeForOneProject**](docs/NetworkPeeringApi.md#verifyconnectviapeeringonlymodeforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering Only Mode for One Project
*OnlineArchiveApi* | [**CreateOneOnlineArchive**](docs/OnlineArchiveApi.md#createoneonlinearchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
*OnlineArchiveApi* | [**DownloadOnlineArchiveQueryLogs**](docs/OnlineArchiveApi.md#downloadonlinearchivequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
*OnlineArchiveApi* | [**RemoveOneOnlineArchive**](docs/OnlineArchiveApi.md#removeoneonlinearchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
*OnlineArchiveApi* | [**ReturnAllOnlineArchivesForOneCluster**](docs/OnlineArchiveApi.md#returnallonlinearchivesforonecluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
*OnlineArchiveApi* | [**ReturnOneOnlineArchive**](docs/OnlineArchiveApi.md#returnoneonlinearchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
*OnlineArchiveApi* | [**UpdateOneOnlineArchive**](docs/OnlineArchiveApi.md#updateoneonlinearchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive
*OrganizationsApi* | [**CreateOrganizationInvitation**](docs/OrganizationsApi.md#createorganizationinvitation) | **Post** /api/atlas/v2/orgs/{orgId}/invites | Invite One MongoDB Cloud User to Join One Atlas Organization
*OrganizationsApi* | [**DeleteOrganization**](docs/OrganizationsApi.md#deleteorganization) | **Delete** /api/atlas/v2/orgs/{orgId} | Remove One Organization
*OrganizationsApi* | [**DeleteOrganizationInvitation**](docs/OrganizationsApi.md#deleteorganizationinvitation) | **Delete** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Cancel One Organization Invitation
*OrganizationsApi* | [**GetOrganization**](docs/OrganizationsApi.md#getorganization) | **Get** /api/atlas/v2/orgs/{orgId} | Return One Organization
*OrganizationsApi* | [**GetOrganizationInvitation**](docs/OrganizationsApi.md#getorganizationinvitation) | **Get** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Return One Organization Invitation
*OrganizationsApi* | [**ListOrganizationInvitations**](docs/OrganizationsApi.md#listorganizationinvitations) | **Get** /api/atlas/v2/orgs/{orgId}/invites | Return All Organization Invitations
*OrganizationsApi* | [**ListOrganizationProjects**](docs/OrganizationsApi.md#listorganizationprojects) | **Get** /api/atlas/v2/orgs/{orgId}/groups | Return One or More Projects in One Organization
*OrganizationsApi* | [**ListOrganizationUsers**](docs/OrganizationsApi.md#listorganizationusers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
*OrganizationsApi* | [**ListOrganizations**](docs/OrganizationsApi.md#listorganizations) | **Get** /api/atlas/v2/orgs | Return All Organizations
*OrganizationsApi* | [**RenameOrganization**](docs/OrganizationsApi.md#renameorganization) | **Patch** /api/atlas/v2/orgs/{orgId} | Rename One Organization
*OrganizationsApi* | [**UpdateOrganizationInvitation**](docs/OrganizationsApi.md#updateorganizationinvitation) | **Patch** /api/atlas/v2/orgs/{orgId}/invites | Update One Organization Invitation
*OrganizationsApi* | [**UpdateOrganizationInvitationById**](docs/OrganizationsApi.md#updateorganizationinvitationbyid) | **Patch** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Update One Organization Invitation by Invitation ID
*PerformanceAdvisorApi* | [**ReturnAllNamespacesForOneHost**](docs/PerformanceAdvisorApi.md#returnallnamespacesforonehost) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
*PerformanceAdvisorApi* | [**ReturnSlowQueries**](docs/PerformanceAdvisorApi.md#returnslowqueries) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
*PerformanceAdvisorApi* | [**ReturnSuggestedIndexes**](docs/PerformanceAdvisorApi.md#returnsuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes
*PrivateEndpointServicesApi* | [**CreateOnePrivateEndpointForOneProvider**](docs/PrivateEndpointServicesApi.md#createoneprivateendpointforoneprovider) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**CreateOnePrivateEndpointServiceForOneProvider**](docs/PrivateEndpointServicesApi.md#createoneprivateendpointserviceforoneprovider) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**RemoveOnePrivateEndpointForOneProvider**](docs/PrivateEndpointServicesApi.md#removeoneprivateendpointforoneprovider) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**RemoveOnePrivateEndpointServiceForOneProvider**](docs/PrivateEndpointServicesApi.md#removeoneprivateendpointserviceforoneprovider) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**ReturnAllPrivateEndpointServicesForOneProvider**](docs/PrivateEndpointServicesApi.md#returnallprivateendpointservicesforoneprovider) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider
*PrivateEndpointServicesApi* | [**ReturnOnePrivateEndpointForOneProvider**](docs/PrivateEndpointServicesApi.md#returnoneprivateendpointforoneprovider) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**ReturnOnePrivateEndpointServiceForOneProvider**](docs/PrivateEndpointServicesApi.md#returnoneprivateendpointserviceforoneprovider) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**ReturnRegionalizedPrivateEndpointStatus**](docs/PrivateEndpointServicesApi.md#returnregionalizedprivateendpointstatus) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status
*PrivateEndpointServicesApi* | [**ToggleRegionalizedPrivateEndpointStatus**](docs/PrivateEndpointServicesApi.md#toggleregionalizedprivateendpointstatus) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status
*ProgrammaticAPIKeysApi* | [**AssignOneOrganizationApiKeyToOneProject**](docs/ProgrammaticAPIKeysApi.md#assignoneorganizationapikeytooneproject) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [**CreateAccessListEntriesForOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#createaccesslistentriesforoneorganizationapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [**CreateAndAssignOneOrganizationApiKeyToOneProject**](docs/ProgrammaticAPIKeysApi.md#createandassignoneorganizationapikeytooneproject) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [**CreateOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#createoneorganizationapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key
*ProgrammaticAPIKeysApi* | [**RemoveOneAccessListEntryForOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#removeoneaccesslistentryforoneorganizationapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [**RemoveOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#removeoneorganizationapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key
*ProgrammaticAPIKeysApi* | [**ReturnAllAccessListEntriesForOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#returnallaccesslistentriesforoneorganizationapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [**ReturnAllOrganizationApiKeys**](docs/ProgrammaticAPIKeysApi.md#returnallorganizationapikeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys
*ProgrammaticAPIKeysApi* | [**ReturnAllOrganizationApiKeysAssignedToOneProject**](docs/ProgrammaticAPIKeysApi.md#returnallorganizationapikeysassignedtooneproject) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project
*ProgrammaticAPIKeysApi* | [**ReturnOneAccessListEntryForOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#returnoneaccesslistentryforoneorganizationapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [**ReturnOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#returnoneorganizationapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key
*ProgrammaticAPIKeysApi* | [**UnassignOneOrganizationApiKeyFromOneProject**](docs/ProgrammaticAPIKeysApi.md#unassignoneorganizationapikeyfromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project
*ProgrammaticAPIKeysApi* | [**UpdateOneOrganizationApiKey**](docs/ProgrammaticAPIKeysApi.md#updateoneorganizationapikey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key
*ProgrammaticAPIKeysApi* | [**UpdateRoles**](docs/ProgrammaticAPIKeysApi.md#updateroles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Roles of One Organization API Key to One Project
*ProjectIPAccessListApi* | [**CreateProjectIpAccessList**](docs/ProjectIPAccessListApi.md#createprojectipaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List
*ProjectIPAccessListApi* | [**DeleteProjectIpAccessList**](docs/ProjectIPAccessListApi.md#deleteprojectipaccesslist) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List
*ProjectIPAccessListApi* | [**GetProjectIpAccessListStatus**](docs/ProjectIPAccessListApi.md#getprojectipaccessliststatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry
*ProjectIPAccessListApi* | [**GetProjectIpList**](docs/ProjectIPAccessListApi.md#getprojectiplist) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry
*ProjectIPAccessListApi* | [**ListProjectIpAccessLists**](docs/ProjectIPAccessListApi.md#listprojectipaccesslists) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return Project IP Access List
*ProjectsApi* | [**CancelOneProjectInvitation**](docs/ProjectsApi.md#canceloneprojectinvitation) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Cancel One Project Invitation
*ProjectsApi* | [**CreateOneProject**](docs/ProjectsApi.md#createoneproject) | **Post** /api/atlas/v2/groups | Create One Project
*ProjectsApi* | [**GetAllProjectLimits**](docs/ProjectsApi.md#getallprojectlimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project
*ProjectsApi* | [**GetGroupSettings**](docs/ProjectsApi.md#getgroupsettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return One Project Settings
*ProjectsApi* | [**GetOneProjectLimit**](docs/ProjectsApi.md#getoneprojectlimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project
*ProjectsApi* | [**InviteOneMongoDBCloudUserToJoinOneProject**](docs/ProjectsApi.md#inviteonemongodbcloudusertojoinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/invites | Invite One MongoDB Cloud User to Join One Project
*ProjectsApi* | [**PatchGroupSettings**](docs/ProjectsApi.md#patchgroupsettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update One Project Settings
*ProjectsApi* | [**RemoveOneProject**](docs/ProjectsApi.md#removeoneproject) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project
*ProjectsApi* | [**RemoveOneProjectLimit**](docs/ProjectsApi.md#removeoneprojectlimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit
*ProjectsApi* | [**ReturnAllProjectInvitations**](docs/ProjectsApi.md#returnallprojectinvitations) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Project Invitations
*ProjectsApi* | [**ReturnAllProjects**](docs/ProjectsApi.md#returnallprojects) | **Get** /api/atlas/v2/groups | Return All Projects
*ProjectsApi* | [**ReturnAllUsersInOneProject**](docs/ProjectsApi.md#returnallusersinoneproject) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All Users in One Project
*ProjectsApi* | [**ReturnOneProject**](docs/ProjectsApi.md#returnoneproject) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project
*ProjectsApi* | [**ReturnOneProjectInvitation**](docs/ProjectsApi.md#returnoneprojectinvitation) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Project Invitation
*ProjectsApi* | [**ReturnOneProjectUsingItsName**](docs/ProjectsApi.md#returnoneprojectusingitsname) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project using Its Name
*ProjectsApi* | [**SetOneProjectLimit**](docs/ProjectsApi.md#setoneprojectlimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit
*ProjectsApi* | [**UpdateOneProjectInvitation**](docs/ProjectsApi.md#updateoneprojectinvitation) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Project Invitation
*ProjectsApi* | [**UpdateOneProjectInvitationByInvitationId**](docs/ProjectsApi.md#updateoneprojectinvitationbyinvitationid) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Project Invitation by Invitation ID
*RootApi* | [**GetSystemStatus**](docs/RootApi.md#getsystemstatus) | **Get** /api/atlas/v2 | Return the status of this MongoDB application
*ServerlessInstancesApi* | [**CreateOneServerlessInstanceInOneProject**](docs/ServerlessInstancesApi.md#createoneserverlessinstanceinoneproject) | **Post** /api/atlas/v2/groups/{groupId}/serverless | Create One Serverless Instance in One Project
*ServerlessInstancesApi* | [**RemoveOneServerlessInstanceFromOneProject**](docs/ServerlessInstancesApi.md#removeoneserverlessinstancefromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project
*ServerlessInstancesApi* | [**ReturnAllServerlessInstancesFromOneProject**](docs/ServerlessInstancesApi.md#returnallserverlessinstancesfromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances from One Project
*ServerlessInstancesApi* | [**ReturnOneServerlessInstanceFromOneProject**](docs/ServerlessInstancesApi.md#returnoneserverlessinstancefromoneproject) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project
*ServerlessInstancesApi* | [**UpdateOneServerlessInstanceInOneProject**](docs/ServerlessInstancesApi.md#updateoneserverlessinstanceinoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project
*ServerlessPrivateEndpointsApi* | [**CreateOnePrivateEndpointForOneServerlessInstance**](docs/ServerlessPrivateEndpointsApi.md#createoneprivateendpointforoneserverlessinstance) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**RemoveOnePrivateEndpointFromOneServerlessInstance**](docs/ServerlessPrivateEndpointsApi.md#removeoneprivateendpointfromoneserverlessinstance) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**ReturnAllPrivateEndpointsForOneServerlessInstance**](docs/ServerlessPrivateEndpointsApi.md#returnallprivateendpointsforoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**ReturnOnePrivateEndpointForOneServerlessInstance**](docs/ServerlessPrivateEndpointsApi.md#returnoneprivateendpointforoneserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**UpdateOnePrivateEndpointForOneServerlessInstance**](docs/ServerlessPrivateEndpointsApi.md#updateoneprivateendpointforoneserverlessinstance) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance
*SharedTierRestoreJobsApi* | [**CreateOneRestoreJobFromOneM2OrM5Cluster**](docs/SharedTierRestoreJobsApi.md#createonerestorejobfromonem2orm5cluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job from One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [**ReturnAllRestoreJobsForOneM2OrM5Cluster**](docs/SharedTierRestoreJobsApi.md#returnallrestorejobsforonem2orm5cluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [**ReturnOneRestoreJobForOneM2OrM5Cluster**](docs/SharedTierRestoreJobsApi.md#returnonerestorejobforonem2orm5cluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [**DownloadOneM2OrM5ClusterSnapshot**](docs/SharedTierSnapshotsApi.md#downloadonem2orm5clustersnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download | Download One M2 or M5 Cluster Snapshot
*SharedTierSnapshotsApi* | [**ReturnAllSnapshotsForOneM2OrM5Cluster**](docs/SharedTierSnapshotsApi.md#returnallsnapshotsforonem2orm5cluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots | Return All Snapshots for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [**ReturnOneSnapshotForOneM2OrM5Cluster**](docs/SharedTierSnapshotsApi.md#returnonesnapshotforonem2orm5cluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId} | Return One Snapshot for One M2 or M5 Cluster
*TeamsApi* | [**AddOneOrMoreTeamsToOneProject**](docs/TeamsApi.md#addoneormoreteamstooneproject) | **Post** /api/atlas/v2/groups/{groupId}/teams | Add One or More Teams to One Project
*TeamsApi* | [**AssignOneOrganizationUserToOneTeam**](docs/TeamsApi.md#assignoneorganizationusertooneteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users from One Organization to One Team
*TeamsApi* | [**CreateOneTeamInOneOrganization**](docs/TeamsApi.md#createoneteaminoneorganization) | **Post** /api/atlas/v2/orgs/{orgId}/teams | Create One Team in One Organization
*TeamsApi* | [**RemoveOneTeamFromOneOrganization**](docs/TeamsApi.md#removeoneteamfromoneorganization) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization
*TeamsApi* | [**RemoveOneTeamFromOneProject**](docs/TeamsApi.md#removeoneteamfromoneproject) | **Delete** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Remove One Team from One Project
*TeamsApi* | [**RenameOneTeam**](docs/TeamsApi.md#renameoneteam) | **Patch** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Rename One Team
*TeamsApi* | [**ReturnAllMongoDBCloudUsersAssignedToOneTeam**](docs/TeamsApi.md#returnallmongodbcloudusersassignedtooneteam) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
*TeamsApi* | [**ReturnAllTeams**](docs/TeamsApi.md#returnallteams) | **Get** /api/atlas/v2/groups/{groupId}/teams | Return All Teams in One Project
*TeamsApi* | [**ReturnAllTeamsInOneOrganization**](docs/TeamsApi.md#returnallteamsinoneorganization) | **Get** /api/atlas/v2/orgs/{orgId}/teams | Return All Teams in One Organization
*TeamsApi* | [**ReturnOneTeamUsingItsId**](docs/TeamsApi.md#returnoneteamusingitsid) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Return One Team using its ID
*TeamsApi* | [**ReturnOneTeamUsingItsName**](docs/TeamsApi.md#returnoneteamusingitsname) | **Get** /api/atlas/v2/orgs/{orgId}/teams/byName/{teamName} | Return One Team using its Name
*TeamsApi* | [**UpdateTeamRolesInOneProject**](docs/TeamsApi.md#updateteamrolesinoneproject) | **Patch** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project
*TestApi* | [**VersionedExample**](docs/TestApi.md#versionedexample) | **Get** /api/atlas/v2/example/info | Example resource info for versioning of the Atlas API
*ThirdPartyIntegrationsApi* | [**ConfigureOneThirdPartyServiceIntegration**](docs/ThirdPartyIntegrationsApi.md#configureonethirdpartyserviceintegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**RemoveOneThirdPartyServiceIntegration**](docs/ThirdPartyIntegrationsApi.md#removeonethirdpartyserviceintegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**ReturnAllActiveThirdPartyServiceIntegrations**](docs/ThirdPartyIntegrationsApi.md#returnallactivethirdpartyserviceintegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations
*ThirdPartyIntegrationsApi* | [**ReturnOneThirdPartyServiceIntegration**](docs/ThirdPartyIntegrationsApi.md#returnonethirdpartyserviceintegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**UpdateOneThirdPartyServiceIntegration**](docs/ThirdPartyIntegrationsApi.md#updateonethirdpartyserviceintegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration
*X509AuthenticationApi* | [**CreateDatabaseUserCertificate**](docs/X509AuthenticationApi.md#createdatabaseusercertificate) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One MongoDB User
*X509AuthenticationApi* | [**DisableCustomerManagedX509**](docs/X509AuthenticationApi.md#disablecustomermanagedx509) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509
*X509AuthenticationApi* | [**ListDatabaseUserCertificates**](docs/X509AuthenticationApi.md#listdatabaseusercertificates) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One MongoDB User


## Documentation For Models

 - [AWSAutoScaling](docs/AWSAutoScaling.md)
 - [AWSCloudProviderContainer](docs/AWSCloudProviderContainer.md)
 - [AWSCloudProviderContainerAllOf](docs/AWSCloudProviderContainerAllOf.md)
 - [AWSComputeAutoScaling](docs/AWSComputeAutoScaling.md)
 - [AWSCustomDNSEnabledView](docs/AWSCustomDNSEnabledView.md)
 - [AWSHardwareSpec](docs/AWSHardwareSpec.md)
 - [AWSInterfaceEndpoint](docs/AWSInterfaceEndpoint.md)
 - [AWSKMS](docs/AWSKMS.md)
 - [AWSPeerVpc](docs/AWSPeerVpc.md)
 - [AWSPeerVpcRequest](docs/AWSPeerVpcRequest.md)
 - [AWSPrivateLinkConnection](docs/AWSPrivateLinkConnection.md)
 - [AWSProviderSettings](docs/AWSProviderSettings.md)
 - [AWSProviderSettingsAllOf](docs/AWSProviderSettingsAllOf.md)
 - [AWSRegionConfig](docs/AWSRegionConfig.md)
 - [AccessListItemView](docs/AccessListItemView.md)
 - [AlertAuditType](docs/AlertAuditType.md)
 - [AlertAuditView](docs/AlertAuditView.md)
 - [AlertConfigAuditType](docs/AlertConfigAuditType.md)
 - [AlertConfigAuditView](docs/AlertConfigAuditView.md)
 - [AlertView](docs/AlertView.md)
 - [AlertViewForGroup](docs/AlertViewForGroup.md)
 - [ApiAddUserToTeamView](docs/ApiAddUserToTeamView.md)
 - [ApiAlertConfigView](docs/ApiAlertConfigView.md)
 - [ApiApiUserView](docs/ApiApiUserView.md)
 - [ApiAppUserView](docs/ApiAppUserView.md)
 - [ApiAtlasContainerPeerViewRequest](docs/ApiAtlasContainerPeerViewRequest.md)
 - [ApiAtlasFTSAnalyzersViewManual](docs/ApiAtlasFTSAnalyzersViewManual.md)
 - [ApiAtlasFTSAnalyzersViewManualCharFiltersInner](docs/ApiAtlasFTSAnalyzersViewManualCharFiltersInner.md)
 - [ApiAtlasFTSAnalyzersViewManualTokenFiltersInner](docs/ApiAtlasFTSAnalyzersViewManualTokenFiltersInner.md)
 - [ApiAtlasFTSAnalyzersViewManualTokenizer](docs/ApiAtlasFTSAnalyzersViewManualTokenizer.md)
 - [ApiAtlasFTSMappingsViewManual](docs/ApiAtlasFTSMappingsViewManual.md)
 - [ApiAtlasNetPeerRequestBase](docs/ApiAtlasNetPeerRequestBase.md)
 - [ApiAtlasServerlessBackupRestoreJobViewManual](docs/ApiAtlasServerlessBackupRestoreJobViewManual.md)
 - [ApiAtlasServerlessBackupSnapshotViewManual](docs/ApiAtlasServerlessBackupSnapshotViewManual.md)
 - [ApiAvailableRegionView](docs/ApiAvailableRegionView.md)
 - [ApiBSONTimestampView](docs/ApiBSONTimestampView.md)
 - [ApiCheckpointPartView](docs/ApiCheckpointPartView.md)
 - [ApiCreateApiKeyView](docs/ApiCreateApiKeyView.md)
 - [ApiDatabaseView](docs/ApiDatabaseView.md)
 - [ApiDatadogView](docs/ApiDatadogView.md)
 - [ApiDatadogViewAllOf](docs/ApiDatadogViewAllOf.md)
 - [ApiDeleteCopiedBackupsView](docs/ApiDeleteCopiedBackupsView.md)
 - [ApiDiskPartitionView](docs/ApiDiskPartitionView.md)
 - [ApiError](docs/ApiError.md)
 - [ApiExportStatusView](docs/ApiExportStatusView.md)
 - [ApiFTSMetricView](docs/ApiFTSMetricView.md)
 - [ApiFTSMetricsView](docs/ApiFTSMetricsView.md)
 - [ApiGroupInvitationRequestView](docs/ApiGroupInvitationRequestView.md)
 - [ApiGroupInvitationUpdateRequestView](docs/ApiGroupInvitationUpdateRequestView.md)
 - [ApiGroupInvitationView](docs/ApiGroupInvitationView.md)
 - [ApiHipChatView](docs/ApiHipChatView.md)
 - [ApiHostViewAtlas](docs/ApiHostViewAtlas.md)
 - [ApiIndexRequestView](docs/ApiIndexRequestView.md)
 - [ApiInstanceSizeView](docs/ApiInstanceSizeView.md)
 - [ApiIntegerThresholdView](docs/ApiIntegerThresholdView.md)
 - [ApiIntegrationView](docs/ApiIntegrationView.md)
 - [ApiIntegrationViewManual](docs/ApiIntegrationViewManual.md)
 - [ApiInvoiceView](docs/ApiInvoiceView.md)
 - [ApiKeyView](docs/ApiKeyView.md)
 - [ApiLimitView](docs/ApiLimitView.md)
 - [ApiLineItemView](docs/ApiLineItemView.md)
 - [ApiMatcherView](docs/ApiMatcherView.md)
 - [ApiMeasurementView](docs/ApiMeasurementView.md)
 - [ApiMeasurementViewAtlas](docs/ApiMeasurementViewAtlas.md)
 - [ApiMeasurementsGeneralViewAtlas](docs/ApiMeasurementsGeneralViewAtlas.md)
 - [ApiMeasurementsIndexesView](docs/ApiMeasurementsIndexesView.md)
 - [ApiMeasurementsNonIndexView](docs/ApiMeasurementsNonIndexView.md)
 - [ApiMetricDataPointView](docs/ApiMetricDataPointView.md)
 - [ApiMetricDataPointViewAtlas](docs/ApiMetricDataPointViewAtlas.md)
 - [ApiMetricThresholdView](docs/ApiMetricThresholdView.md)
 - [ApiMicrosoftTeamsView](docs/ApiMicrosoftTeamsView.md)
 - [ApiMicrosoftTeamsViewAllOf](docs/ApiMicrosoftTeamsViewAllOf.md)
 - [ApiMongoDBAccessLogsListView](docs/ApiMongoDBAccessLogsListView.md)
 - [ApiMongoDBAccessLogsView](docs/ApiMongoDBAccessLogsView.md)
 - [ApiNamespaceObjView](docs/ApiNamespaceObjView.md)
 - [ApiNamespacesView](docs/ApiNamespacesView.md)
 - [ApiNewRelicView](docs/ApiNewRelicView.md)
 - [ApiNewRelicViewAllOf](docs/ApiNewRelicViewAllOf.md)
 - [ApiNotificationView](docs/ApiNotificationView.md)
 - [ApiOpsGenieView](docs/ApiOpsGenieView.md)
 - [ApiOpsGenieViewAllOf](docs/ApiOpsGenieViewAllOf.md)
 - [ApiOrganizationInvitationRequestView](docs/ApiOrganizationInvitationRequestView.md)
 - [ApiOrganizationInvitationUpdateRequestView](docs/ApiOrganizationInvitationUpdateRequestView.md)
 - [ApiOrganizationInvitationView](docs/ApiOrganizationInvitationView.md)
 - [ApiOrganizationView](docs/ApiOrganizationView.md)
 - [ApiPagerDutyView](docs/ApiPagerDutyView.md)
 - [ApiPagerDutyViewAllOf](docs/ApiPagerDutyViewAllOf.md)
 - [ApiPaymentView](docs/ApiPaymentView.md)
 - [ApiPerformanceAdvisorIndexView](docs/ApiPerformanceAdvisorIndexView.md)
 - [ApiPerformanceAdvisorResponseView](docs/ApiPerformanceAdvisorResponseView.md)
 - [ApiPerformanceAdvisorShapeView](docs/ApiPerformanceAdvisorShapeView.md)
 - [ApiPerformanceAdvisorSlowQueryListView](docs/ApiPerformanceAdvisorSlowQueryListView.md)
 - [ApiPerformanceAdvisorSlowQueryView](docs/ApiPerformanceAdvisorSlowQueryView.md)
 - [ApiPolicyItemView](docs/ApiPolicyItemView.md)
 - [ApiPolicyView](docs/ApiPolicyView.md)
 - [ApiPrometheusView](docs/ApiPrometheusView.md)
 - [ApiPrometheusViewAllOf](docs/ApiPrometheusViewAllOf.md)
 - [ApiRefundView](docs/ApiRefundView.md)
 - [ApiRestoreJobDeliveryView](docs/ApiRestoreJobDeliveryView.md)
 - [ApiRestoreJobFileHashView](docs/ApiRestoreJobFileHashView.md)
 - [ApiRestoreJobView](docs/ApiRestoreJobView.md)
 - [ApiRoleAssignmentView](docs/ApiRoleAssignmentView.md)
 - [ApiSlackView](docs/ApiSlackView.md)
 - [ApiSlackViewAllOf](docs/ApiSlackViewAllOf.md)
 - [ApiSnapshotPartView](docs/ApiSnapshotPartView.md)
 - [ApiSnapshotView](docs/ApiSnapshotView.md)
 - [ApiSystemStatusView](docs/ApiSystemStatusView.md)
 - [ApiTeamResponseView](docs/ApiTeamResponseView.md)
 - [ApiTeamRoleView](docs/ApiTeamRoleView.md)
 - [ApiTeamView](docs/ApiTeamView.md)
 - [ApiUserAccessListView](docs/ApiUserAccessListView.md)
 - [ApiUserEventTypeForGroupAndOrg](docs/ApiUserEventTypeForGroupAndOrg.md)
 - [ApiUserEventViewForGroupAndOrg](docs/ApiUserEventViewForGroupAndOrg.md)
 - [ApiUserRoleAssignment](docs/ApiUserRoleAssignment.md)
 - [ApiVictorOpsView](docs/ApiVictorOpsView.md)
 - [ApiVictorOpsViewAllOf](docs/ApiVictorOpsViewAllOf.md)
 - [ApiWebhookView](docs/ApiWebhookView.md)
 - [ApiWebhookViewAllOf](docs/ApiWebhookViewAllOf.md)
 - [AppServiceAlertView](docs/AppServiceAlertView.md)
 - [AppServiceEventType](docs/AppServiceEventType.md)
 - [AppServiceEventView](docs/AppServiceEventView.md)
 - [AuditLog](docs/AuditLog.md)
 - [AutoExportPolicyView](docs/AutoExportPolicyView.md)
 - [AutoScaling](docs/AutoScaling.md)
 - [AutoScalingV15](docs/AutoScalingV15.md)
 - [AutomationConfigEventType](docs/AutomationConfigEventType.md)
 - [AutomationConfigEventView](docs/AutomationConfigEventView.md)
 - [AvailableDeploymentView](docs/AvailableDeploymentView.md)
 - [AvailableProjectView](docs/AvailableProjectView.md)
 - [AzureAutoScaling](docs/AzureAutoScaling.md)
 - [AzureCloudProviderContainer](docs/AzureCloudProviderContainer.md)
 - [AzureCloudProviderContainerAllOf](docs/AzureCloudProviderContainerAllOf.md)
 - [AzureComputeAutoScaling](docs/AzureComputeAutoScaling.md)
 - [AzureHardwareSpec](docs/AzureHardwareSpec.md)
 - [AzureKeyVault](docs/AzureKeyVault.md)
 - [AzurePeerNetwork](docs/AzurePeerNetwork.md)
 - [AzurePeerNetworkRequest](docs/AzurePeerNetworkRequest.md)
 - [AzurePrivateEndpoint](docs/AzurePrivateEndpoint.md)
 - [AzurePrivateLinkConnection](docs/AzurePrivateLinkConnection.md)
 - [AzureProviderSettings](docs/AzureProviderSettings.md)
 - [AzureProviderSettingsAllOf](docs/AzureProviderSettingsAllOf.md)
 - [AzureRegionConfig](docs/AzureRegionConfig.md)
 - [BackupAlertView](docs/BackupAlertView.md)
 - [BackupEventType](docs/BackupEventType.md)
 - [BackupEventTypeForGroup](docs/BackupEventTypeForGroup.md)
 - [BackupEventViewForGroup](docs/BackupEventViewForGroup.md)
 - [BiConnector](docs/BiConnector.md)
 - [BiConnectorEventType](docs/BiConnectorEventType.md)
 - [BiConnectorEventView](docs/BiConnectorEventView.md)
 - [BillingEventTypeForGroup](docs/BillingEventTypeForGroup.md)
 - [BillingEventTypeForOrg](docs/BillingEventTypeForOrg.md)
 - [BillingEventViewForGroup](docs/BillingEventViewForGroup.md)
 - [BillingEventViewForOrg](docs/BillingEventViewForOrg.md)
 - [CharFilterhtmlStrip](docs/CharFilterhtmlStrip.md)
 - [CharFiltericuNormalize](docs/CharFiltericuNormalize.md)
 - [CharFiltermapping](docs/CharFiltermapping.md)
 - [CharFiltermappingMappings](docs/CharFiltermappingMappings.md)
 - [CharFilterpersian](docs/CharFilterpersian.md)
 - [Checkpoint](docs/Checkpoint.md)
 - [CloudProviderAccess](docs/CloudProviderAccess.md)
 - [CloudProviderAccessAWSIAMRole](docs/CloudProviderAccessAWSIAMRole.md)
 - [CloudProviderAccessDataLakeFeatureUsage](docs/CloudProviderAccessDataLakeFeatureUsage.md)
 - [CloudProviderAccessDataLakeFeatureUsageAllOf](docs/CloudProviderAccessDataLakeFeatureUsageAllOf.md)
 - [CloudProviderAccessEncryptionAtRestFeatureUsage](docs/CloudProviderAccessEncryptionAtRestFeatureUsage.md)
 - [CloudProviderAccessEncryptionAtRestFeatureUsageAllOf](docs/CloudProviderAccessEncryptionAtRestFeatureUsageAllOf.md)
 - [CloudProviderAccessExportSnapshotFeatureUsage](docs/CloudProviderAccessExportSnapshotFeatureUsage.md)
 - [CloudProviderAccessExportSnapshotFeatureUsageAllOf](docs/CloudProviderAccessExportSnapshotFeatureUsageAllOf.md)
 - [CloudProviderAccessFeatureUsage](docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsageDataLakeFeatureId](docs/CloudProviderAccessFeatureUsageDataLakeFeatureId.md)
 - [CloudProviderAccessFeatureUsageExportSnapshotFeatureId](docs/CloudProviderAccessFeatureUsageExportSnapshotFeatureId.md)
 - [CloudProviderAccessRole](docs/CloudProviderAccessRole.md)
 - [CloudProviderContainer](docs/CloudProviderContainer.md)
 - [ClusterAlertView](docs/ClusterAlertView.md)
 - [ClusterDescriptionConnectionStrings](docs/ClusterDescriptionConnectionStrings.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpoint](docs/ClusterDescriptionConnectionStringsPrivateEndpoint.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpointEndpoint](docs/ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ClusterDescriptionProcessArgs](docs/ClusterDescriptionProcessArgs.md)
 - [ClusterDescriptionV15](docs/ClusterDescriptionV15.md)
 - [ClusterEventType](docs/ClusterEventType.md)
 - [ClusterEventView](docs/ClusterEventView.md)
 - [ClusterOutageSimulation](docs/ClusterOutageSimulation.md)
 - [ClusterOutageSimulationOutageFilter](docs/ClusterOutageSimulationOutageFilter.md)
 - [ClusterProviderSettings](docs/ClusterProviderSettings.md)
 - [ClusterStatus](docs/ClusterStatus.md)
 - [ClusterView](docs/ClusterView.md)
 - [Collation](docs/Collation.md)
 - [ComputeAutoScaling](docs/ComputeAutoScaling.md)
 - [ComputeAutoScalingV15](docs/ComputeAutoScalingV15.md)
 - [ConnectedOrgConfigView](docs/ConnectedOrgConfigView.md)
 - [ContainerPeer](docs/ContainerPeer.md)
 - [CreateAWSEndpointRequest](docs/CreateAWSEndpointRequest.md)
 - [CreateAWSEndpointRequestAllOf](docs/CreateAWSEndpointRequestAllOf.md)
 - [CreateAzureEndpointRequest](docs/CreateAzureEndpointRequest.md)
 - [CreateAzureEndpointRequestAllOf](docs/CreateAzureEndpointRequestAllOf.md)
 - [CreateEndpointServiceRequest](docs/CreateEndpointServiceRequest.md)
 - [CreateGCPEndpointGroupRequest](docs/CreateGCPEndpointGroupRequest.md)
 - [CreateGCPEndpointGroupRequestAllOf](docs/CreateGCPEndpointGroupRequestAllOf.md)
 - [CreateGCPForwardingRuleRequest](docs/CreateGCPForwardingRuleRequest.md)
 - [CreateOneNewNetworkPeeringConnection200Response](docs/CreateOneNewNetworkPeeringConnection200Response.md)
 - [CreateOnePrivateEndpointForOneProviderRequest](docs/CreateOnePrivateEndpointForOneProviderRequest.md)
 - [CriteriaView](docs/CriteriaView.md)
 - [CustomCriteriaView](docs/CustomCriteriaView.md)
 - [CustomCriteriaViewAllOf](docs/CustomCriteriaViewAllOf.md)
 - [CustomDBRole](docs/CustomDBRole.md)
 - [CustomerX509](docs/CustomerX509.md)
 - [DBAction](docs/DBAction.md)
 - [DBResource](docs/DBResource.md)
 - [DLSIngestionSink](docs/DLSIngestionSink.md)
 - [DLSIngestionSinkAllOf](docs/DLSIngestionSinkAllOf.md)
 - [DailyScheduleView](docs/DailyScheduleView.md)
 - [DailyScheduleViewAllOf](docs/DailyScheduleViewAllOf.md)
 - [DataExplorerAccessedEventType](docs/DataExplorerAccessedEventType.md)
 - [DataExplorerAccessedEventView](docs/DataExplorerAccessedEventView.md)
 - [DataFederationQueryLimit](docs/DataFederationQueryLimit.md)
 - [DataLakeAWSCloudProviderConfig](docs/DataLakeAWSCloudProviderConfig.md)
 - [DataLakeAtlasStore](docs/DataLakeAtlasStore.md)
 - [DataLakeAtlasStoreAllOf](docs/DataLakeAtlasStoreAllOf.md)
 - [DataLakeAtlasStoreReadPreference](docs/DataLakeAtlasStoreReadPreference.md)
 - [DataLakeAtlasStoreReadPreferenceTag](docs/DataLakeAtlasStoreReadPreferenceTag.md)
 - [DataLakeCloudProviderConfig](docs/DataLakeCloudProviderConfig.md)
 - [DataLakeDataProcessRegion](docs/DataLakeDataProcessRegion.md)
 - [DataLakeDatabase](docs/DataLakeDatabase.md)
 - [DataLakeDatabaseCollection](docs/DataLakeDatabaseCollection.md)
 - [DataLakeDatabaseDataSource](docs/DataLakeDatabaseDataSource.md)
 - [DataLakeHTTPStore](docs/DataLakeHTTPStore.md)
 - [DataLakeHTTPStoreAllOf](docs/DataLakeHTTPStoreAllOf.md)
 - [DataLakeOnlineArchiveStore](docs/DataLakeOnlineArchiveStore.md)
 - [DataLakeOnlineArchiveStoreAllOf](docs/DataLakeOnlineArchiveStoreAllOf.md)
 - [DataLakeRegion](docs/DataLakeRegion.md)
 - [DataLakeS3Store](docs/DataLakeS3Store.md)
 - [DataLakeS3StoreAllOf](docs/DataLakeS3StoreAllOf.md)
 - [DataLakeStorage](docs/DataLakeStorage.md)
 - [DataLakeStore](docs/DataLakeStore.md)
 - [DataLakeTenant](docs/DataLakeTenant.md)
 - [DataLakeView](docs/DataLakeView.md)
 - [DataMetricAlertView](docs/DataMetricAlertView.md)
 - [DataMetricEventView](docs/DataMetricEventView.md)
 - [DataMetricUnits](docs/DataMetricUnits.md)
 - [DataMetricValueView](docs/DataMetricValueView.md)
 - [DatabaseUser](docs/DatabaseUser.md)
 - [DateCriteriaView](docs/DateCriteriaView.md)
 - [DateCriteriaViewAllOf](docs/DateCriteriaViewAllOf.md)
 - [DedicatedHardwareSpec](docs/DedicatedHardwareSpec.md)
 - [DefaultAlertViewForGroup](docs/DefaultAlertViewForGroup.md)
 - [DefaultAlertViewForGroupEventTypeName](docs/DefaultAlertViewForGroupEventTypeName.md)
 - [DefaultEventViewForGroup](docs/DefaultEventViewForGroup.md)
 - [DefaultEventViewForGroupEventTypeName](docs/DefaultEventViewForGroupEventTypeName.md)
 - [DefaultEventViewForOrg](docs/DefaultEventViewForOrg.md)
 - [DefaultEventViewForOrgEventTypeName](docs/DefaultEventViewForOrgEventTypeName.md)
 - [DefaultScheduleView](docs/DefaultScheduleView.md)
 - [Destination](docs/Destination.md)
 - [DiskBackupBaseRestoreMember](docs/DiskBackupBaseRestoreMember.md)
 - [DiskBackupCopySetting](docs/DiskBackupCopySetting.md)
 - [DiskBackupExportJob](docs/DiskBackupExportJob.md)
 - [DiskBackupExportJobRequest](docs/DiskBackupExportJobRequest.md)
 - [DiskBackupOnDemandSnapshotRequest](docs/DiskBackupOnDemandSnapshotRequest.md)
 - [DiskBackupReplicaSet](docs/DiskBackupReplicaSet.md)
 - [DiskBackupRestoreJob](docs/DiskBackupRestoreJob.md)
 - [DiskBackupShardedClusterSnapshot](docs/DiskBackupShardedClusterSnapshot.md)
 - [DiskBackupShardedClusterSnapshotMember](docs/DiskBackupShardedClusterSnapshotMember.md)
 - [DiskBackupSnapshot](docs/DiskBackupSnapshot.md)
 - [DiskBackupSnapshotAWSExportBucket](docs/DiskBackupSnapshotAWSExportBucket.md)
 - [DiskBackupSnapshotSchedule](docs/DiskBackupSnapshotSchedule.md)
 - [DiskGBAutoScaling](docs/DiskGBAutoScaling.md)
 - [EncryptionAtRest](docs/EncryptionAtRest.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointService](docs/EndpointService.md)
 - [EventTypeForGroup](docs/EventTypeForGroup.md)
 - [EventTypeForOrg](docs/EventTypeForOrg.md)
 - [EventViewForGroup](docs/EventViewForGroup.md)
 - [EventViewForOrg](docs/EventViewForOrg.md)
 - [ExampleResourceResponseView20210909](docs/ExampleResourceResponseView20210909.md)
 - [ExampleResourceResponseView20221018](docs/ExampleResourceResponseView20221018.md)
 - [FTSIndex](docs/FTSIndex.md)
 - [FTSIndexAuditType](docs/FTSIndexAuditType.md)
 - [FTSIndexAuditView](docs/FTSIndexAuditView.md)
 - [FTSSynonymMappingDefinition](docs/FTSSynonymMappingDefinition.md)
 - [FederatedUserView](docs/FederatedUserView.md)
 - [FieldTransformation](docs/FieldTransformation.md)
 - [FreeAutoScaling](docs/FreeAutoScaling.md)
 - [FreeProviderSettings](docs/FreeProviderSettings.md)
 - [FreeProviderSettingsAllOf](docs/FreeProviderSettingsAllOf.md)
 - [GCPAutoScaling](docs/GCPAutoScaling.md)
 - [GCPCloudProviderContainer](docs/GCPCloudProviderContainer.md)
 - [GCPCloudProviderContainerAllOf](docs/GCPCloudProviderContainerAllOf.md)
 - [GCPComputeAutoScaling](docs/GCPComputeAutoScaling.md)
 - [GCPConsumerForwardingRule](docs/GCPConsumerForwardingRule.md)
 - [GCPEndpointGroup](docs/GCPEndpointGroup.md)
 - [GCPEndpointService](docs/GCPEndpointService.md)
 - [GCPHardwareSpec](docs/GCPHardwareSpec.md)
 - [GCPPeerVpc](docs/GCPPeerVpc.md)
 - [GCPPeerVpcRequest](docs/GCPPeerVpcRequest.md)
 - [GCPProviderSettings](docs/GCPProviderSettings.md)
 - [GCPProviderSettingsAllOf](docs/GCPProviderSettingsAllOf.md)
 - [GCPRegionConfig](docs/GCPRegionConfig.md)
 - [GeoSharding](docs/GeoSharding.md)
 - [GeoShardingView](docs/GeoShardingView.md)
 - [GoogleCloudKMS](docs/GoogleCloudKMS.md)
 - [Group](docs/Group.md)
 - [GroupMaintenanceWindow](docs/GroupMaintenanceWindow.md)
 - [GroupPaginatedEventView](docs/GroupPaginatedEventView.md)
 - [GroupSettings](docs/GroupSettings.md)
 - [HardwareSpec](docs/HardwareSpec.md)
 - [HostAlertView](docs/HostAlertView.md)
 - [HostEventType](docs/HostEventType.md)
 - [HostEventTypeForGroup](docs/HostEventTypeForGroup.md)
 - [HostEventViewForGroup](docs/HostEventViewForGroup.md)
 - [HostMetricAlertView](docs/HostMetricAlertView.md)
 - [HostMetricEventType](docs/HostMetricEventType.md)
 - [HostMetricEventView](docs/HostMetricEventView.md)
 - [HostMetricValueView](docs/HostMetricValueView.md)
 - [IdentityProviderUpdate](docs/IdentityProviderUpdate.md)
 - [IdentityProviderView](docs/IdentityProviderView.md)
 - [IndexOptions](docs/IndexOptions.md)
 - [IngestionPipeline](docs/IngestionPipeline.md)
 - [IngestionPipelineRun](docs/IngestionPipelineRun.md)
 - [IngestionSink](docs/IngestionSink.md)
 - [IngestionSource](docs/IngestionSource.md)
 - [InheritedRole](docs/InheritedRole.md)
 - [InstanceSize](docs/InstanceSize.md)
 - [Label](docs/Label.md)
 - [LegacyClusterDescription](docs/LegacyClusterDescription.md)
 - [LegacyReplicationSpec](docs/LegacyReplicationSpec.md)
 - [Link](docs/Link.md)
 - [LinkAtlas](docs/LinkAtlas.md)
 - [LiveMigrationRequestView](docs/LiveMigrationRequestView.md)
 - [LiveMigrationResponseView](docs/LiveMigrationResponseView.md)
 - [MaintenanceWindowConfigAuditType](docs/MaintenanceWindowConfigAuditType.md)
 - [MaintenanceWindowConfigAuditView](docs/MaintenanceWindowConfigAuditView.md)
 - [ManagedNamespaceView](docs/ManagedNamespaceView.md)
 - [ManagedNamespaces](docs/ManagedNamespaces.md)
 - [MonthlyScheduleView](docs/MonthlyScheduleView.md)
 - [MonthlyScheduleViewAllOf](docs/MonthlyScheduleViewAllOf.md)
 - [NDSAuditTypeForGroup](docs/NDSAuditTypeForGroup.md)
 - [NDSAuditTypeForOrg](docs/NDSAuditTypeForOrg.md)
 - [NDSAuditViewForGroup](docs/NDSAuditViewForGroup.md)
 - [NDSAuditViewForOrg](docs/NDSAuditViewForOrg.md)
 - [NDSAutoScalingAuditTypeForGroup](docs/NDSAutoScalingAuditTypeForGroup.md)
 - [NDSAutoScalingAuditViewForGroup](docs/NDSAutoScalingAuditViewForGroup.md)
 - [NDSLDAP](docs/NDSLDAP.md)
 - [NDSLDAPVerifyConnectivityJobRequest](docs/NDSLDAPVerifyConnectivityJobRequest.md)
 - [NDSLDAPVerifyConnectivityJobRequestParams](docs/NDSLDAPVerifyConnectivityJobRequestParams.md)
 - [NDSLDAPVerifyConnectivityJobRequestValidation](docs/NDSLDAPVerifyConnectivityJobRequestValidation.md)
 - [NDSLabel](docs/NDSLabel.md)
 - [NDSServerlessInstanceAuditType](docs/NDSServerlessInstanceAuditType.md)
 - [NDSServerlessInstanceAuditView](docs/NDSServerlessInstanceAuditView.md)
 - [NDSTenantEndpointAuditType](docs/NDSTenantEndpointAuditType.md)
 - [NDSTenantEndpointAuditView](docs/NDSTenantEndpointAuditView.md)
 - [NDSUserToDNMapping](docs/NDSUserToDNMapping.md)
 - [NetworkPermissionEntry](docs/NetworkPermissionEntry.md)
 - [NetworkPermissionEntryStatus](docs/NetworkPermissionEntryStatus.md)
 - [NumberMetricAlertView](docs/NumberMetricAlertView.md)
 - [NumberMetricEventView](docs/NumberMetricEventView.md)
 - [NumberMetricUnits](docs/NumberMetricUnits.md)
 - [NumberMetricValueView](docs/NumberMetricValueView.md)
 - [OnDemandCpsSnapshotSource](docs/OnDemandCpsSnapshotSource.md)
 - [OnDemandCpsSnapshotSourceAllOf](docs/OnDemandCpsSnapshotSourceAllOf.md)
 - [OnlineArchive](docs/OnlineArchive.md)
 - [OnlineArchiveSchedule](docs/OnlineArchiveSchedule.md)
 - [OrgEventType](docs/OrgEventType.md)
 - [OrgEventView](docs/OrgEventView.md)
 - [OrgFederationSettingsView](docs/OrgFederationSettingsView.md)
 - [OrgGroupView](docs/OrgGroupView.md)
 - [OrgPaginatedEventView](docs/OrgPaginatedEventView.md)
 - [PaginatedAWSPeerVpcView](docs/PaginatedAWSPeerVpcView.md)
 - [PaginatedAlertConfigView](docs/PaginatedAlertConfigView.md)
 - [PaginatedAlertView](docs/PaginatedAlertView.md)
 - [PaginatedApiApiUserView](docs/PaginatedApiApiUserView.md)
 - [PaginatedApiAppUserView](docs/PaginatedApiAppUserView.md)
 - [PaginatedApiAtlasCheckpointView](docs/PaginatedApiAtlasCheckpointView.md)
 - [PaginatedApiAtlasDatabaseUserView](docs/PaginatedApiAtlasDatabaseUserView.md)
 - [PaginatedApiAtlasDiskBackupExportJobView](docs/PaginatedApiAtlasDiskBackupExportJobView.md)
 - [PaginatedApiAtlasProviderRegionsView](docs/PaginatedApiAtlasProviderRegionsView.md)
 - [PaginatedApiInvoiceView](docs/PaginatedApiInvoiceView.md)
 - [PaginatedApiUserAccessListView](docs/PaginatedApiUserAccessListView.md)
 - [PaginatedAppUserView](docs/PaginatedAppUserView.md)
 - [PaginatedAtlasGroupView](docs/PaginatedAtlasGroupView.md)
 - [PaginatedAzurePeerNetworkView](docs/PaginatedAzurePeerNetworkView.md)
 - [PaginatedBackupSnapshotExportBucketView](docs/PaginatedBackupSnapshotExportBucketView.md)
 - [PaginatedBackupSnapshotView](docs/PaginatedBackupSnapshotView.md)
 - [PaginatedCloudBackupReplicaSetView](docs/PaginatedCloudBackupReplicaSetView.md)
 - [PaginatedCloudBackupRestoreJobView](docs/PaginatedCloudBackupRestoreJobView.md)
 - [PaginatedCloudBackupShardedClusterSnapshotView](docs/PaginatedCloudBackupShardedClusterSnapshotView.md)
 - [PaginatedCloudProviderContainerView](docs/PaginatedCloudProviderContainerView.md)
 - [PaginatedClusterDescriptionV15View](docs/PaginatedClusterDescriptionV15View.md)
 - [PaginatedDatabaseView](docs/PaginatedDatabaseView.md)
 - [PaginatedDiskPartitionView](docs/PaginatedDiskPartitionView.md)
 - [PaginatedGCPPeerVpcView](docs/PaginatedGCPPeerVpcView.md)
 - [PaginatedHostViewAtlas](docs/PaginatedHostViewAtlas.md)
 - [PaginatedIntegrationView](docs/PaginatedIntegrationView.md)
 - [PaginatedLegacyClusterView](docs/PaginatedLegacyClusterView.md)
 - [PaginatedNetworkAccessView](docs/PaginatedNetworkAccessView.md)
 - [PaginatedOnlineArchiveView](docs/PaginatedOnlineArchiveView.md)
 - [PaginatedOrgGroupView](docs/PaginatedOrgGroupView.md)
 - [PaginatedOrganizationView](docs/PaginatedOrganizationView.md)
 - [PaginatedPipelineRunView](docs/PaginatedPipelineRunView.md)
 - [PaginatedPrivateLinkConnectionView](docs/PaginatedPrivateLinkConnectionView.md)
 - [PaginatedRestoreJobView](docs/PaginatedRestoreJobView.md)
 - [PaginatedServerlessBackupRestoreJobViewManual](docs/PaginatedServerlessBackupRestoreJobViewManual.md)
 - [PaginatedServerlessBackupSnapshotViewManual](docs/PaginatedServerlessBackupSnapshotViewManual.md)
 - [PaginatedServerlessInstanceDescriptionView](docs/PaginatedServerlessInstanceDescriptionView.md)
 - [PaginatedSnapshotView](docs/PaginatedSnapshotView.md)
 - [PaginatedTeamRoleView](docs/PaginatedTeamRoleView.md)
 - [PaginatedTeamView](docs/PaginatedTeamView.md)
 - [PaginatedTenantRestoreView](docs/PaginatedTenantRestoreView.md)
 - [PaginatedTenantSnapshotView](docs/PaginatedTenantSnapshotView.md)
 - [PaginatedUserCertView](docs/PaginatedUserCertView.md)
 - [PartitionField](docs/PartitionField.md)
 - [PartitionFieldView](docs/PartitionFieldView.md)
 - [PemFileInfo](docs/PemFileInfo.md)
 - [PemFileInfoView](docs/PemFileInfoView.md)
 - [PerformanceAdvisorOpStatsView](docs/PerformanceAdvisorOpStatsView.md)
 - [PerformanceAdvisorOperationView](docs/PerformanceAdvisorOperationView.md)
 - [PeriodicCpsSnapshotSource](docs/PeriodicCpsSnapshotSource.md)
 - [PeriodicCpsSnapshotSourceAllOf](docs/PeriodicCpsSnapshotSourceAllOf.md)
 - [PipelineRunStats](docs/PipelineRunStats.md)
 - [PrivateIPModeView](docs/PrivateIPModeView.md)
 - [PrivateNetworkEndpointIdEntry](docs/PrivateNetworkEndpointIdEntry.md)
 - [ProjectSettingItemView](docs/ProjectSettingItemView.md)
 - [ProviderRegions](docs/ProviderRegions.md)
 - [Raw](docs/Raw.md)
 - [RawMetricAlertView](docs/RawMetricAlertView.md)
 - [RawMetricEventView](docs/RawMetricEventView.md)
 - [RawMetricUnits](docs/RawMetricUnits.md)
 - [RawMetricValueView](docs/RawMetricValueView.md)
 - [RegionConfig](docs/RegionConfig.md)
 - [RegionSpec](docs/RegionSpec.md)
 - [ReplicaSetAlertView](docs/ReplicaSetAlertView.md)
 - [ReplicaSetEventType](docs/ReplicaSetEventType.md)
 - [ReplicaSetEventTypeForGroup](docs/ReplicaSetEventTypeForGroup.md)
 - [ReplicaSetEventViewForGroup](docs/ReplicaSetEventViewForGroup.md)
 - [ReplicationSpec](docs/ReplicationSpec.md)
 - [ReturnAllNetworkPeeringConnectionsInOneProject200Response](docs/ReturnAllNetworkPeeringConnectionsInOneProject200Response.md)
 - [ReturnOneNetworkPeeringConnectionInOneProject200Response](docs/ReturnOneNetworkPeeringConnectionInOneProject200Response.md)
 - [Role](docs/Role.md)
 - [RoleAssignment](docs/RoleAssignment.md)
 - [RoleMappingView](docs/RoleMappingView.md)
 - [SampleDatasetStatus](docs/SampleDatasetStatus.md)
 - [ServerlessAWSTenantEndpoint](docs/ServerlessAWSTenantEndpoint.md)
 - [ServerlessAWSTenantEndpointUpdate](docs/ServerlessAWSTenantEndpointUpdate.md)
 - [ServerlessAWSTenantEndpointUpdateAllOf](docs/ServerlessAWSTenantEndpointUpdateAllOf.md)
 - [ServerlessAzureTenantEndpoint](docs/ServerlessAzureTenantEndpoint.md)
 - [ServerlessAzureTenantEndpointUpdate](docs/ServerlessAzureTenantEndpointUpdate.md)
 - [ServerlessAzureTenantEndpointUpdateAllOf](docs/ServerlessAzureTenantEndpointUpdateAllOf.md)
 - [ServerlessBackupOptions](docs/ServerlessBackupOptions.md)
 - [ServerlessInstanceDescription](docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](docs/ServerlessInstanceDescriptionConnectionStrings.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint](docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint](docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ServerlessInstanceDescriptionCreate](docs/ServerlessInstanceDescriptionCreate.md)
 - [ServerlessInstanceDescriptionUpdate](docs/ServerlessInstanceDescriptionUpdate.md)
 - [ServerlessProviderSettings](docs/ServerlessProviderSettings.md)
 - [ServerlessTenantEndpoint](docs/ServerlessTenantEndpoint.md)
 - [ServerlessTenantEndpointCreate](docs/ServerlessTenantEndpointCreate.md)
 - [ServerlessTenantEndpointUpdate](docs/ServerlessTenantEndpointUpdate.md)
 - [SnapshotRetention](docs/SnapshotRetention.md)
 - [SnapshotSchedule](docs/SnapshotSchedule.md)
 - [Source](docs/Source.md)
 - [SynonymSource](docs/SynonymSource.md)
 - [TargetOrgRequestView](docs/TargetOrgRequestView.md)
 - [TargetOrgView](docs/TargetOrgView.md)
 - [TeamEventType](docs/TeamEventType.md)
 - [TeamEventTypeForGroup](docs/TeamEventTypeForGroup.md)
 - [TeamEventView](docs/TeamEventView.md)
 - [TeamEventViewForGroup](docs/TeamEventViewForGroup.md)
 - [TenantHardwareSpec](docs/TenantHardwareSpec.md)
 - [TenantRegionConfig](docs/TenantRegionConfig.md)
 - [TenantRestore](docs/TenantRestore.md)
 - [TenantSnapshot](docs/TenantSnapshot.md)
 - [TimeMetricAlertView](docs/TimeMetricAlertView.md)
 - [TimeMetricEventView](docs/TimeMetricEventView.md)
 - [TimeMetricUnits](docs/TimeMetricUnits.md)
 - [TimeMetricValueView](docs/TimeMetricValueView.md)
 - [TokenFilterasciiFolding](docs/TokenFilterasciiFolding.md)
 - [TokenFilterdaitchMokotoffSoundex](docs/TokenFilterdaitchMokotoffSoundex.md)
 - [TokenFilteredgeGram](docs/TokenFilteredgeGram.md)
 - [TokenFiltericuFolding](docs/TokenFiltericuFolding.md)
 - [TokenFiltericuNormalizer](docs/TokenFiltericuNormalizer.md)
 - [TokenFilterlength](docs/TokenFilterlength.md)
 - [TokenFilterlowercase](docs/TokenFilterlowercase.md)
 - [TokenFilternGram](docs/TokenFilternGram.md)
 - [TokenFilterregex](docs/TokenFilterregex.md)
 - [TokenFilterreverse](docs/TokenFilterreverse.md)
 - [TokenFiltershingle](docs/TokenFiltershingle.md)
 - [TokenFiltersnowballStemming](docs/TokenFiltersnowballStemming.md)
 - [TokenFilterstopword](docs/TokenFilterstopword.md)
 - [TokenFiltertrim](docs/TokenFiltertrim.md)
 - [TokenizeredgeGram](docs/TokenizeredgeGram.md)
 - [Tokenizerkeyword](docs/Tokenizerkeyword.md)
 - [TokenizernGram](docs/TokenizernGram.md)
 - [TokenizerregexCaptureGroup](docs/TokenizerregexCaptureGroup.md)
 - [TokenizerregexSplit](docs/TokenizerregexSplit.md)
 - [Tokenizerstandard](docs/Tokenizerstandard.md)
 - [TokenizeruaxUrlEmail](docs/TokenizeruaxUrlEmail.md)
 - [Tokenizerwhitespace](docs/Tokenizerwhitespace.md)
 - [TriggerIngestionRequest](docs/TriggerIngestionRequest.md)
 - [UpdateCustomDBRole](docs/UpdateCustomDBRole.md)
 - [UserCert](docs/UserCert.md)
 - [UserEventTypeForGroup](docs/UserEventTypeForGroup.md)
 - [UserEventTypeForOrg](docs/UserEventTypeForOrg.md)
 - [UserEventViewForGroup](docs/UserEventViewForGroup.md)
 - [UserEventViewForOrg](docs/UserEventViewForOrg.md)
 - [UserScope](docs/UserScope.md)
 - [UserSecurity](docs/UserSecurity.md)
 - [ValidationView](docs/ValidationView.md)
 - [WeeklyScheduleView](docs/WeeklyScheduleView.md)
 - [WeeklyScheduleViewAllOf](docs/WeeklyScheduleViewAllOf.md)
 - [X509Certificate](docs/X509Certificate.md)
 - [X509CertificateView](docs/X509CertificateView.md)




