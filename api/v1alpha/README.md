# Go API client
## Documentation for API Endpoints

All URIs are relative to *https://cloud.mongodb.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AWSClustersDNSApi* | [**GetAWSCustomDNS**](docs/AWSClustersDNSApi.md#getawscustomdns) | **Get** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Return One Custom DNS Configuration for Atlas Clusters on AWS
*AWSClustersDNSApi* | [**ToggleAWSCustomDNS**](docs/AWSClustersDNSApi.md#toggleawscustomdns) | **Patch** /api/atlas/v2/groups/{groupId}/awsCustomDNS | Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS
*AccessTrackingApi* | [**ListAccessLogsByClusterName**](docs/AccessTrackingApi.md#listaccesslogsbyclustername) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/clusters/{clusterName} | Return Database Access History for One Cluster using Its Cluster Name
*AccessTrackingApi* | [**ListAccessLogsByHostname**](docs/AccessTrackingApi.md#listaccesslogsbyhostname) | **Get** /api/atlas/v2/groups/{groupId}/dbAccessHistory/processes/{hostname} | Return Database Access History for One Cluster using Its Hostname
*AlertConfigurationsApi* | [**CreateAlertConfiguration**](docs/AlertConfigurationsApi.md#createalertconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/alertConfigs | Create One Alert Configuration in One Project
*AlertConfigurationsApi* | [**DeleteAlertConfiguration**](docs/AlertConfigurationsApi.md#deletealertconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Remove One Alert Configuration from One Project
*AlertConfigurationsApi* | [**GetAlertConfiguration**](docs/AlertConfigurationsApi.md#getalertconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Return One Alert Configuration from One Project
*AlertConfigurationsApi* | [**ListAlertConfigurationMatchersFieldNames**](docs/AlertConfigurationsApi.md#listalertconfigurationmatchersfieldnames) | **Get** /api/atlas/v2/alertConfigs/matchers/fieldNames | Get All Alert Configuration Matchers Field Names
*AlertConfigurationsApi* | [**ListAlertConfigurations**](docs/AlertConfigurationsApi.md#listalertconfigurations) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs | Return All Alert Configurations for One Project
*AlertConfigurationsApi* | [**ListAlertConfigurationsByAlertId**](docs/AlertConfigurationsApi.md#listalertconfigurationsbyalertid) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId}/alertConfigs | Return All Alert Configurations Set for One Alert
*AlertConfigurationsApi* | [**ToggleAlertConfiguration**](docs/AlertConfigurationsApi.md#togglealertconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Toggle One State of One Alert Configuration in One Project
*AlertConfigurationsApi* | [**UpdateAlertConfiguration**](docs/AlertConfigurationsApi.md#updatealertconfiguration) | **Put** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId} | Update One Alert Configuration for One Project
*AlertsApi* | [**AcknowledgeAlert**](docs/AlertsApi.md#acknowledgealert) | **Patch** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Acknowledge One Alert from One Project
*AlertsApi* | [**GetAlert**](docs/AlertsApi.md#getalert) | **Get** /api/atlas/v2/groups/{groupId}/alerts/{alertId} | Return One Alert from One Project
*AlertsApi* | [**ListAlerts**](docs/AlertsApi.md#listalerts) | **Get** /api/atlas/v2/groups/{groupId}/alerts | Return All Alerts from One Project
*AlertsApi* | [**ListAlertsByAlertConfigurationId**](docs/AlertsApi.md#listalertsbyalertconfigurationid) | **Get** /api/atlas/v2/groups/{groupId}/alertConfigs/{alertConfigId}/alerts | Return All Open Alerts for Alert Configuration
*AtlasSearchApi* | [**CreateAtlasSearchIndex**](docs/AtlasSearchApi.md#createatlassearchindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes | Create One Atlas Search Index
*AtlasSearchApi* | [**DeleteAtlasSearchIndex**](docs/AtlasSearchApi.md#deleteatlassearchindex) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Remove One Atlas Search Index
*AtlasSearchApi* | [**GetAtlasSearchIndex**](docs/AtlasSearchApi.md#getatlassearchindex) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Return One Atlas Search Index
*AtlasSearchApi* | [**ListAtlasSearchIndexes**](docs/AtlasSearchApi.md#listatlassearchindexes) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName} | Return All Atlas Search Indexes for One Collection
*AtlasSearchApi* | [**UpdateAtlasSearchIndex**](docs/AtlasSearchApi.md#updateatlassearchindex) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId} | Update One Atlas Search Index
*AuditingApi* | [**GetAuditingConfiguration**](docs/AuditingApi.md#getauditingconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/auditLog | Return the Auditing Configuration for One Project
*AuditingApi* | [**UpdateAuditingConfiguration**](docs/AuditingApi.md#updateauditingconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/auditLog | Update Auditing Configuration for One Project
*CloudBackupsApi* | [**CancelBackupRestoreJob**](docs/CloudBackupsApi.md#cancelbackuprestorejob) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Cancel One Restore Job of One Cluster
*CloudBackupsApi* | [**CreateBackupExportJob**](docs/CloudBackupsApi.md#createbackupexportjob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Create One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [**CreateBackupRestoreJob**](docs/CloudBackupsApi.md#createbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Cluster
*CloudBackupsApi* | [**CreateExportBucket**](docs/CloudBackupsApi.md#createexportbucket) | **Post** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Grant Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**CreateServerlessBackupRestoreJob**](docs/CloudBackupsApi.md#createserverlessbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Restore One Snapshot of One Serverless Instance
*CloudBackupsApi* | [**DeleteAllBackupSchedules**](docs/CloudBackupsApi.md#deleteallbackupschedules) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Remove All Cloud Backup Schedules
*CloudBackupsApi* | [**DeleteExportBucket**](docs/CloudBackupsApi.md#deleteexportbucket) | **Delete** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Revoke Access to AWS S3 Bucket for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**DeleteReplicaSetBackup**](docs/CloudBackupsApi.md#deletereplicasetbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Remove One Replica Set Cloud Backup
*CloudBackupsApi* | [**DeleteShardedClusterBackup**](docs/CloudBackupsApi.md#deleteshardedclusterbackup) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Remove One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [**GetBackupExportJob**](docs/CloudBackupsApi.md#getbackupexportjob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports/{exportId} | Return One Cloud Backup Snapshot Export Job
*CloudBackupsApi* | [**GetBackupRestoreJob**](docs/CloudBackupsApi.md#getbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job of One Cluster
*CloudBackupsApi* | [**GetBackupSchedule**](docs/CloudBackupsApi.md#getbackupschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Return One Cloud Backup Schedule
*CloudBackupsApi* | [**GetDataProtectionSettings**](docs/CloudBackupsApi.md#getdataprotectionsettings) | **Get** /api/atlas/v2/groups/{groupId}/dataProtection | Return the Data Protection Policy settings
*CloudBackupsApi* | [**GetExportBucket**](docs/CloudBackupsApi.md#getexportbucket) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets/{exportBucketId} | Return One AWS S3 Bucket Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**GetReplicaSetBackup**](docs/CloudBackupsApi.md#getreplicasetbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Return One Replica Set Cloud Backup
*CloudBackupsApi* | [**GetServerlessBackup**](docs/CloudBackupsApi.md#getserverlessbackup) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots/{snapshotId} | Return One Snapshot of One Serverless Instance
*CloudBackupsApi* | [**GetServerlessBackupRestoreJob**](docs/CloudBackupsApi.md#getserverlessbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs/{restoreJobId} | Return One Restore Job for One Serverless Instance
*CloudBackupsApi* | [**GetShardedClusterBackup**](docs/CloudBackupsApi.md#getshardedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedCluster/{snapshotId} | Return One Sharded Cluster Cloud Backup
*CloudBackupsApi* | [**ListBackupExportJobs**](docs/CloudBackupsApi.md#listbackupexportjobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/exports | Return All Cloud Backup Snapshot Export Jobs
*CloudBackupsApi* | [**ListBackupRestoreJobs**](docs/CloudBackupsApi.md#listbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Cluster
*CloudBackupsApi* | [**ListExportBuckets**](docs/CloudBackupsApi.md#listexportbuckets) | **Get** /api/atlas/v2/groups/{groupId}/backup/exportBuckets | Return All AWS S3 Buckets Used for Cloud Backup Snapshot Exports
*CloudBackupsApi* | [**ListReplicaSetBackups**](docs/CloudBackupsApi.md#listreplicasetbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Return All Replica Set Cloud Backups
*CloudBackupsApi* | [**ListServerlessBackupRestoreJobs**](docs/CloudBackupsApi.md#listserverlessbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/restoreJobs | Return All Restore Jobs for One Serverless Instance
*CloudBackupsApi* | [**ListServerlessBackups**](docs/CloudBackupsApi.md#listserverlessbackups) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{clusterName}/backup/snapshots | Return All Snapshots of One Serverless Instance
*CloudBackupsApi* | [**ListShardedClusterBackups**](docs/CloudBackupsApi.md#listshardedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/shardedClusters | Return All Sharded Cluster Cloud Backups
*CloudBackupsApi* | [**TakeSnapshot**](docs/CloudBackupsApi.md#takesnapshot) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots | Take One On-Demand Snapshot
*CloudBackupsApi* | [**UpdateBackupSchedule**](docs/CloudBackupsApi.md#updatebackupschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/schedule | Update Cloud Backup Schedule for One Cluster
*CloudBackupsApi* | [**UpdateDataProtectionSettings**](docs/CloudBackupsApi.md#updatedataprotectionsettings) | **Put** /api/atlas/v2/groups/{groupId}/dataProtection | Update or enable the Data Protection Policy settings
*CloudBackupsApi* | [**UpdateSnapshotRetention**](docs/CloudBackupsApi.md#updatesnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/snapshots/{snapshotId} | Change Expiration Date for One Cloud Backup
*CloudMigrationServiceApi* | [**CreateLinkToken**](docs/CloudMigrationServiceApi.md#createlinktoken) | **Post** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Create One Link-Token
*CloudMigrationServiceApi* | [**CreatePushMigration**](docs/CloudMigrationServiceApi.md#createpushmigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations | Migrate One Local Managed Cluster to MongoDB Atlas
*CloudMigrationServiceApi* | [**CutoverMigration**](docs/CloudMigrationServiceApi.md#cutovermigration) | **Put** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover | Cut Over the Migrated Cluster
*CloudMigrationServiceApi* | [**DeleteLinkToken**](docs/CloudMigrationServiceApi.md#deletelinktoken) | **Delete** /api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens | Remove One Link-Token
*CloudMigrationServiceApi* | [**GetPushMigration**](docs/CloudMigrationServiceApi.md#getpushmigration) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId} | Return One Migration Job
*CloudMigrationServiceApi* | [**GetValidationStatus**](docs/CloudMigrationServiceApi.md#getvalidationstatus) | **Get** /api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId} | Return One Migration Validation Job
*CloudMigrationServiceApi* | [**ListSourceProjects**](docs/CloudMigrationServiceApi.md#listsourceprojects) | **Get** /api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects | Return All Projects Available for Migration
*CloudMigrationServiceApi* | [**ValidateMigration**](docs/CloudMigrationServiceApi.md#validatemigration) | **Post** /api/atlas/v2/groups/{groupId}/liveMigrations/validate | Validate One Migration Request
*CloudProviderAccessApi* | [**AuthorizeCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#authorizecloudprovideraccessrole) | **Patch** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Authorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [**CreateCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#createcloudprovideraccessrole) | **Post** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Create One Cloud Provider Access Role
*CloudProviderAccessApi* | [**DeauthorizeCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#deauthorizecloudprovideraccessrole) | **Delete** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId} | Deauthorize One Cloud Provider Access Role
*CloudProviderAccessApi* | [**GetCloudProviderAccessRole**](docs/CloudProviderAccessApi.md#getcloudprovideraccessrole) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId} | Return specified Cloud Provider Access Role
*CloudProviderAccessApi* | [**ListCloudProviderAccessRoles**](docs/CloudProviderAccessApi.md#listcloudprovideraccessroles) | **Get** /api/atlas/v2/groups/{groupId}/cloudProviderAccess | Return All Cloud Provider Access Roles
*ClusterOutageSimulationApi* | [**EndOutageSimulation**](docs/ClusterOutageSimulationApi.md#endoutagesimulation) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | End an Outage Simulation
*ClusterOutageSimulationApi* | [**GetOutageSimulation**](docs/ClusterOutageSimulationApi.md#getoutagesimulation) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Return One Outage Simulation
*ClusterOutageSimulationApi* | [**StartOutageSimulation**](docs/ClusterOutageSimulationApi.md#startoutagesimulation) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/outageSimulation | Start an Outage Simulation
*ClustersApi* | [**GetClusterAdvancedConfiguration**](docs/ClustersApi.md#getclusteradvancedconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Return One Advanced Configuration Options for One Cluster
*ClustersApi* | [**GetClusterStatus**](docs/ClustersApi.md#getclusterstatus) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status | Return Status of All Cluster Operations
*ClustersApi* | [**GetSampleDatasetLoadStatus**](docs/ClustersApi.md#getsampledatasetloadstatus) | **Get** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId} | Check Status of Cluster Sample Dataset Request
*ClustersApi* | [**ListCloudProviderRegions**](docs/ClustersApi.md#listcloudproviderregions) | **Get** /api/atlas/v2/groups/{groupId}/clusters/provider/regions | Return All Cloud Provider Regions
*ClustersApi* | [**ListClustersForAllProjects**](docs/ClustersApi.md#listclustersforallprojects) | **Get** /api/atlas/v2/clusters | Return All Authorized Clusters in All Projects
*ClustersApi* | [**LoadSampleDataset**](docs/ClustersApi.md#loadsampledataset) | **Post** /api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name} | Load Sample Dataset Request into Cluster
*ClustersApi* | [**UpdateClusterAdvancedConfiguration**](docs/ClustersApi.md#updateclusteradvancedconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs | Update Advanced Configuration Options for One Cluster
*ClustersApi* | [**UpgradeSharedCluster**](docs/ClustersApi.md#upgradesharedcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade | Upgrade One Shared-tier Cluster
*ClustersApi* | [**UpgradeSharedClusterToServerless**](docs/ClustersApi.md#upgradesharedclustertoserverless) | **Post** /api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless | Upgrades One Shared-Tier Cluster to the Serverless Instance
*CustomDatabaseRolesApi* | [**CreateCustomDatabaseRole**](docs/CustomDatabaseRolesApi.md#createcustomdatabaserole) | **Post** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Create One Custom Role
*CustomDatabaseRolesApi* | [**DeleteCustomDatabaseRole**](docs/CustomDatabaseRolesApi.md#deletecustomdatabaserole) | **Delete** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Remove One Custom Role from One Project
*CustomDatabaseRolesApi* | [**GetCustomDatabaseRole**](docs/CustomDatabaseRolesApi.md#getcustomdatabaserole) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Return One Custom Role in One Project
*CustomDatabaseRolesApi* | [**ListCustomDatabaseRoles**](docs/CustomDatabaseRolesApi.md#listcustomdatabaseroles) | **Get** /api/atlas/v2/groups/{groupId}/customDBRoles/roles | Return All Custom Roles in One Project
*CustomDatabaseRolesApi* | [**UpdateCustomDatabaseRole**](docs/CustomDatabaseRolesApi.md#updatecustomdatabaserole) | **Patch** /api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName} | Update One Custom Role in One Project
*DataFederationApi* | [**CreateDataFederationPrivateEndpoint**](docs/DataFederationApi.md#createdatafederationprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Create One Federated Database Instance and Online Archive Private Endpoint for One Project
*DataFederationApi* | [**CreateFederatedDatabase**](docs/DataFederationApi.md#createfederateddatabase) | **Post** /api/atlas/v2/groups/{groupId}/dataFederation | Create One Federated Database Instance in One Project
*DataFederationApi* | [**CreateOneDataFederationQueryLimit**](docs/DataFederationApi.md#createonedatafederationquerylimit) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Configure One Query Limit for One Federated Database Instance
*DataFederationApi* | [**DeleteDataFederationPrivateEndpoint**](docs/DataFederationApi.md#deletedatafederationprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Remove One Federated Database Instance and Online Archive Private Endpoint from One Project
*DataFederationApi* | [**DeleteFederatedDatabase**](docs/DataFederationApi.md#deletefederateddatabase) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Remove One Federated Database Instance from One Project
*DataFederationApi* | [**DeleteOneDataFederationInstanceQueryLimit**](docs/DataFederationApi.md#deleteonedatafederationinstancequerylimit) | **Delete** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Delete One Query Limit For One Federated Database Instance
*DataFederationApi* | [**DownloadFederatedDatabaseQueryLogs**](docs/DataFederationApi.md#downloadfederateddatabasequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/queryLogs.gz | Download Query Logs for One Federated Database Instance
*DataFederationApi* | [**GetDataFederationPrivateEndpoint**](docs/DataFederationApi.md#getdatafederationprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds/{endpointId} | Return One Federated Database Instance and Online Archive Private Endpoint in One Project
*DataFederationApi* | [**GetFederatedDatabase**](docs/DataFederationApi.md#getfederateddatabase) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Return One Federated Database Instance in One Project
*DataFederationApi* | [**ListDataFederationPrivateEndpoints**](docs/DataFederationApi.md#listdatafederationprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateNetworkSettings/endpointIds | Return All Federated Database Instance and Online Archive Private Endpoints in One Project
*DataFederationApi* | [**ListFederatedDatabases**](docs/DataFederationApi.md#listfederateddatabases) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation | Return All Federated Database Instances in One Project
*DataFederationApi* | [**ReturnFederatedDatabaseQueryLimit**](docs/DataFederationApi.md#returnfederateddatabasequerylimit) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits/{limitName} | Return One Federated Database Instance Query Limit for One Project
*DataFederationApi* | [**ReturnFederatedDatabaseQueryLimits**](docs/DataFederationApi.md#returnfederateddatabasequerylimits) | **Get** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName}/limits | Return All Query Limits for One Federated Database Instance
*DataFederationApi* | [**UpdateFederatedDatabase**](docs/DataFederationApi.md#updatefederateddatabase) | **Patch** /api/atlas/v2/groups/{groupId}/dataFederation/{tenantName} | Update One Federated Database Instance in One Project
*DataLakePipelinesApi* | [**CreatePipeline**](docs/DataLakePipelinesApi.md#createpipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines | Create One Data Lake Pipeline
*DataLakePipelinesApi* | [**DeletePipeline**](docs/DataLakePipelinesApi.md#deletepipeline) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Remove One Data Lake Pipeline
*DataLakePipelinesApi* | [**DeletePipelineRunDataset**](docs/DataLakePipelinesApi.md#deletepipelinerundataset) | **Delete** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Delete Pipeline Run Dataset
*DataLakePipelinesApi* | [**GetPipeline**](docs/DataLakePipelinesApi.md#getpipeline) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Return One Data Lake Pipeline
*DataLakePipelinesApi* | [**GetPipelineRun**](docs/DataLakePipelinesApi.md#getpipelinerun) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId} | Return One Data Lake Pipeline Run
*DataLakePipelinesApi* | [**ListPipelineRuns**](docs/DataLakePipelinesApi.md#listpipelineruns) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs | Return All Data Lake Pipeline Runs from One Project
*DataLakePipelinesApi* | [**ListPipelineSchedules**](docs/DataLakePipelinesApi.md#listpipelineschedules) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules | Return Available Ingestion Schedules for One Data Lake Pipeline
*DataLakePipelinesApi* | [**ListPipelineSnapshots**](docs/DataLakePipelinesApi.md#listpipelinesnapshots) | **Get** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots | Return Available Backup Snapshots for One Data Lake Pipeline
*DataLakePipelinesApi* | [**ListPipelines**](docs/DataLakePipelinesApi.md#listpipelines) | **Get** /api/atlas/v2/groups/{groupId}/pipelines | Return All Data Lake Pipelines from One Project
*DataLakePipelinesApi* | [**PausePipeline**](docs/DataLakePipelinesApi.md#pausepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause | Pause One Data Lake Pipeline
*DataLakePipelinesApi* | [**ResumePipeline**](docs/DataLakePipelinesApi.md#resumepipeline) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume | Resume One Data Lake Pipeline
*DataLakePipelinesApi* | [**TriggerSnapshotIngestion**](docs/DataLakePipelinesApi.md#triggersnapshotingestion) | **Post** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger | Trigger on demand snapshot ingestion
*DataLakePipelinesApi* | [**UpdatePipeline**](docs/DataLakePipelinesApi.md#updatepipeline) | **Patch** /api/atlas/v2/groups/{groupId}/pipelines/{pipelineName} | Update One Data Lake Pipeline
*DatabaseUsersApi* | [**CreateDatabaseUser**](docs/DatabaseUsersApi.md#createdatabaseuser) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers | Create One Database User in One Project
*DatabaseUsersApi* | [**DeleteDatabaseUser**](docs/DatabaseUsersApi.md#deletedatabaseuser) | **Delete** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Remove One Database User from One Project
*DatabaseUsersApi* | [**GetDatabaseUser**](docs/DatabaseUsersApi.md#getdatabaseuser) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Return One Database User from One Project
*DatabaseUsersApi* | [**ListDatabaseUsers**](docs/DatabaseUsersApi.md#listdatabaseusers) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers | Return All Database Users from One Project
*DatabaseUsersApi* | [**UpdateDatabaseUser**](docs/DatabaseUsersApi.md#updatedatabaseuser) | **Patch** /api/atlas/v2/groups/{groupId}/databaseUsers/{databaseName}/{username} | Update One Database User in One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [**ReturnOneConfigurationForEncryptionAtRestUsingCustomerManagedKeysForOneProject**](docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#returnoneconfigurationforencryptionatrestusingcustomermanagedkeysforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Return One Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EncryptionAtRestUsingCustomerKeyManagementApi* | [**UpdateEncryptionAtRest**](docs/EncryptionAtRestUsingCustomerKeyManagementApi.md#updateencryptionatrest) | **Patch** /api/atlas/v2/groups/{groupId}/encryptionAtRest | Update Configuration for Encryption at Rest using Customer-Managed Keys for One Project
*EventsApi* | [**GetOrganizationEvent**](docs/EventsApi.md#getorganizationevent) | **Get** /api/atlas/v2/orgs/{orgId}/events/{eventId} | Return One Event from One Organization
*EventsApi* | [**GetProjectEvent**](docs/EventsApi.md#getprojectevent) | **Get** /api/atlas/v2/groups/{groupId}/events/{eventId} | Return One Event from One Project
*EventsApi* | [**ListOrganizationEvents**](docs/EventsApi.md#listorganizationevents) | **Get** /api/atlas/v2/orgs/{orgId}/events | Return All Events from One Organization
*EventsApi* | [**ListProjectEvents**](docs/EventsApi.md#listprojectevents) | **Get** /api/atlas/v2/groups/{groupId}/events | Return All Events from One Project
*FederatedAuthenticationApi* | [**CreateRoleMapping**](docs/FederatedAuthenticationApi.md#createrolemapping) | **Post** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Add One Role Mapping to One Organization
*FederatedAuthenticationApi* | [**DeleteFederationApp**](docs/FederatedAuthenticationApi.md#deletefederationapp) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId} | Delete the federation settings instance.
*FederatedAuthenticationApi* | [**DeleteRoleMapping**](docs/FederatedAuthenticationApi.md#deleterolemapping) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Remove One Role Mapping from One Organization
*FederatedAuthenticationApi* | [**GetConnectedOrgConfig**](docs/FederatedAuthenticationApi.md#getconnectedorgconfig) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Return One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**GetFederationSettings**](docs/FederatedAuthenticationApi.md#getfederationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/federationSettings | Return Federation Settings for One Organization
*FederatedAuthenticationApi* | [**GetIdentityProvider**](docs/FederatedAuthenticationApi.md#getidentityprovider) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Return one identity provider from the specified federation.
*FederatedAuthenticationApi* | [**GetIdentityProviderMetadata**](docs/FederatedAuthenticationApi.md#getidentityprovidermetadata) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId}/metadata.xml | Return the metadata of one identity provider in the specified federation.
*FederatedAuthenticationApi* | [**GetRoleMapping**](docs/FederatedAuthenticationApi.md#getrolemapping) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Return One Role Mapping from One Organization
*FederatedAuthenticationApi* | [**ListConnectedOrgConfigs**](docs/FederatedAuthenticationApi.md#listconnectedorgconfigs) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs | Return All Connected Org Configs from the Federation
*FederatedAuthenticationApi* | [**ListIdentityProviders**](docs/FederatedAuthenticationApi.md#listidentityproviders) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders | Return all identity providers from the specified federation.
*FederatedAuthenticationApi* | [**ListRoleMappings**](docs/FederatedAuthenticationApi.md#listrolemappings) | **Get** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings | Return All Role Mappings from One Organization
*FederatedAuthenticationApi* | [**RemoveConnectedOrgConfig**](docs/FederatedAuthenticationApi.md#removeconnectedorgconfig) | **Delete** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Remove One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**UpdateConnectedOrgConfig**](docs/FederatedAuthenticationApi.md#updateconnectedorgconfig) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId} | Update One Org Config Connected to One Federation
*FederatedAuthenticationApi* | [**UpdateIdentityProvider**](docs/FederatedAuthenticationApi.md#updateidentityprovider) | **Patch** /api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders/{identityProviderId} | Update the identity provider.
*FederatedAuthenticationApi* | [**UpdateRoleMapping**](docs/FederatedAuthenticationApi.md#updaterolemapping) | **Put** /api/atlas/v2/federationSettings/{federationSettingsId}/connectedOrgConfigs/{orgId}/roleMappings/{id} | Update One Role Mapping in One Organization
*GlobalClustersApi* | [**CreateCustomZoneMapping**](docs/GlobalClustersApi.md#createcustomzonemapping) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Add One Entry to One Custom Zone Mapping
*GlobalClustersApi* | [**CreateManagedNamespace**](docs/GlobalClustersApi.md#createmanagednamespace) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Create One Managed Namespace in One Global Multi-Cloud Cluster
*GlobalClustersApi* | [**DeleteAllCustomZoneMappings**](docs/GlobalClustersApi.md#deleteallcustomzonemappings) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping | Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster
*GlobalClustersApi* | [**DeleteManagedNamespace**](docs/GlobalClustersApi.md#deletemanagednamespace) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces | Remove One Managed Namespace from One Global Multi-Cloud Cluster
*GlobalClustersApi* | [**GetManagedNamespace**](docs/GlobalClustersApi.md#getmanagednamespace) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites | Return One Managed Namespace in One Global Multi-Cloud Cluster
*InvoicesApi* | [**DownloadInvoiceCSV**](docs/InvoicesApi.md#downloadinvoicecsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Organization Invoice as CSV
*InvoicesApi* | [**GetInvoice**](docs/InvoicesApi.md#getinvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Organization Invoice
*InvoicesApi* | [**ListInvoices**](docs/InvoicesApi.md#listinvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization
*InvoicesApi* | [**ListPendingInvoices**](docs/InvoicesApi.md#listpendinginvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
*LDAPConfigurationApi* | [**DeleteLDAPConfiguration**](docs/LDAPConfigurationApi.md#deleteldapconfiguration) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping | Remove the Current LDAP User to DN Mapping
*LDAPConfigurationApi* | [**GetLDAPConfiguration**](docs/LDAPConfigurationApi.md#getldapconfiguration) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity | Return the Current LDAP or X.509 Configuration
*LDAPConfigurationApi* | [**GetLDAPConfigurationStatus**](docs/LDAPConfigurationApi.md#getldapconfigurationstatus) | **Get** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId} | Return the Status of One Verify LDAP Configuration Request
*LDAPConfigurationApi* | [**SaveLDAPConfiguration**](docs/LDAPConfigurationApi.md#saveldapconfiguration) | **Patch** /api/atlas/v2/groups/{groupId}/userSecurity | Edit the LDAP or X.509 Configuration
*LDAPConfigurationApi* | [**VerifyLDAPConfiguration**](docs/LDAPConfigurationApi.md#verifyldapconfiguration) | **Post** /api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify | Verify the LDAP Configuration in One Project
*LegacyBackupApi* | [**DeleteLegacySnapshot**](docs/LegacyBackupApi.md#deletelegacysnapshot) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Remove One Legacy Backup Snapshot
*LegacyBackupApi* | [**GetLegacyBackupCheckpoint**](docs/LegacyBackupApi.md#getlegacybackupcheckpoint) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId} | Return One Legacy Backup Checkpoint
*LegacyBackupApi* | [**GetLegacyBackupRestoreJob**](docs/LegacyBackupApi.md#getlegacybackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId} | Return One Legacy Backup Restore Job
*LegacyBackupApi* | [**GetLegacySnapshot**](docs/LegacyBackupApi.md#getlegacysnapshot) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Return One Legacy Backup Snapshot
*LegacyBackupApi* | [**GetLegacySnapshotSchedule**](docs/LegacyBackupApi.md#getlegacysnapshotschedule) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Return One Snapshot Schedule
*LegacyBackupApi* | [**ListLegacyBackupCheckpoints**](docs/LegacyBackupApi.md#listlegacybackupcheckpoints) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints | Return All Legacy Backup Checkpoints
*LegacyBackupApi* | [**ListLegacyBackupRestoreJobs**](docs/LegacyBackupApi.md#listlegacybackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Return All Legacy Backup Restore Jobs
*LegacyBackupApi* | [**ListLegacySnapshots**](docs/LegacyBackupApi.md#listlegacysnapshots) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots | Return All Legacy Backup Snapshots
*LegacyBackupApi* | [**UpdateLegacySnapshotRetention**](docs/LegacyBackupApi.md#updatelegacysnapshotretention) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId} | Change One Legacy Backup Snapshot Expiration
*LegacyBackupApi* | [**UpdateLegacySnapshotSchedule**](docs/LegacyBackupApi.md#updatelegacysnapshotschedule) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule | Update Snapshot Schedule for One Cluster
*LegacyBackupRestoreJobsApi* | [**CreateLegacyBackupRestoreJob**](docs/LegacyBackupRestoreJobsApi.md#createlegacybackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs | Create One Legacy Backup Restore Job
*MaintenanceWindowsApi* | [**DeferMaintenanceWindow**](docs/MaintenanceWindowsApi.md#defermaintenancewindow) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/defer | Defer One Maintenance Window for One Project
*MaintenanceWindowsApi* | [**GetMaintenanceWindow**](docs/MaintenanceWindowsApi.md#getmaintenancewindow) | **Get** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Return One Maintenance Window for One Project
*MaintenanceWindowsApi* | [**ResetMaintenanceWindow**](docs/MaintenanceWindowsApi.md#resetmaintenancewindow) | **Delete** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Reset One Maintenance Window for One Project
*MaintenanceWindowsApi* | [**ToggleMaintenanceAutoDefer**](docs/MaintenanceWindowsApi.md#togglemaintenanceautodefer) | **Post** /api/atlas/v2/groups/{groupId}/maintenanceWindow/autoDefer | Toggle Automatic Deferral of Maintenance for One Project
*MaintenanceWindowsApi* | [**UpdateMaintenanceWindow**](docs/MaintenanceWindowsApi.md#updatemaintenancewindow) | **Patch** /api/atlas/v2/groups/{groupId}/maintenanceWindow | Update Maintenance Window for One Project
*MongoDBCloudUsersApi* | [**CreateUser**](docs/MongoDBCloudUsersApi.md#createuser) | **Post** /api/atlas/v2/users | Create One MongoDB Cloud User
*MongoDBCloudUsersApi* | [**GetUser**](docs/MongoDBCloudUsersApi.md#getuser) | **Get** /api/atlas/v2/users/{userId} | Return One MongoDB Cloud User using Its ID
*MongoDBCloudUsersApi* | [**GetUserByUsername**](docs/MongoDBCloudUsersApi.md#getuserbyusername) | **Get** /api/atlas/v2/users/byName/{userName} | Return One MongoDB Cloud User using Their Username
*MonitoringAndLogsApi* | [**GetAtlasProcess**](docs/MonitoringAndLogsApi.md#getatlasprocess) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId} | Return One MongoDB Process by ID
*MonitoringAndLogsApi* | [**GetDatabase**](docs/MonitoringAndLogsApi.md#getdatabase) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName} | Return One Database for a MongoDB Process
*MonitoringAndLogsApi* | [**GetDatabaseMeasurements**](docs/MonitoringAndLogsApi.md#getdatabasemeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases/{databaseName}/measurements | Return Measurements of One Database for One MongoDB Process
*MonitoringAndLogsApi* | [**GetDiskMeasurements**](docs/MonitoringAndLogsApi.md#getdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName}/measurements | Return Measurements of One Disk for One MongoDB Process
*MonitoringAndLogsApi* | [**GetHostLogs**](docs/MonitoringAndLogsApi.md#gethostlogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{hostName}/logs/{logName}.gz | Download Logs for One Multi-Cloud Cluster Host in One Project
*MonitoringAndLogsApi* | [**GetHostMeasurements**](docs/MonitoringAndLogsApi.md#gethostmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/measurements | Return Measurements for One MongoDB Process
*MonitoringAndLogsApi* | [**GetIndexMetrics**](docs/MonitoringAndLogsApi.md#getindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/{indexName}/measurements | Return Atlas Search Metrics for One Index in One Specified Namespace
*MonitoringAndLogsApi* | [**GetMeasurements**](docs/MonitoringAndLogsApi.md#getmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/measurements | Return Atlas Search Hardware and Status Metrics
*MonitoringAndLogsApi* | [**ListAtlasProcesses**](docs/MonitoringAndLogsApi.md#listatlasprocesses) | **Get** /api/atlas/v2/groups/{groupId}/processes | Return All MongoDB Processes in One Project
*MonitoringAndLogsApi* | [**ListDatabases**](docs/MonitoringAndLogsApi.md#listdatabases) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/databases | Return Available Databases for One MongoDB Process
*MonitoringAndLogsApi* | [**ListDiskMeasurements**](docs/MonitoringAndLogsApi.md#listdiskmeasurements) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks/{partitionName} | Return Measurements of One Disk
*MonitoringAndLogsApi* | [**ListDiskPartitions**](docs/MonitoringAndLogsApi.md#listdiskpartitions) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/disks | Return Available Disks for One MongoDB Process
*MonitoringAndLogsApi* | [**ListIndexMetrics**](docs/MonitoringAndLogsApi.md#listindexmetrics) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics/indexes/{databaseName}/{collectionName}/measurements | Return All Atlas Search Index Metrics for One Namespace
*MonitoringAndLogsApi* | [**ListMetricTypes**](docs/MonitoringAndLogsApi.md#listmetrictypes) | **Get** /api/atlas/v2/groups/{groupId}/hosts/{processId}/fts/metrics | Return All Atlas Search Metric Types for One Process
*MultiCloudClustersApi* | [**CreateCluster**](docs/MultiCloudClustersApi.md#createcluster) | **Post** /api/atlas/v2/groups/{groupId}/clusters | Create One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**DeleteCluster**](docs/MultiCloudClustersApi.md#deletecluster) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Remove One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**GetCluster**](docs/MultiCloudClustersApi.md#getcluster) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Return One Multi-Cloud Cluster from One Project
*MultiCloudClustersApi* | [**ListClusters**](docs/MultiCloudClustersApi.md#listclusters) | **Get** /api/atlas/v2/groups/{groupId}/clusters | Return All Multi-Cloud Clusters from One Project
*MultiCloudClustersApi* | [**TestFailover**](docs/MultiCloudClustersApi.md#testfailover) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries | Test Failover for One Multi-Cloud Cluster
*MultiCloudClustersApi* | [**UpdateCluster**](docs/MultiCloudClustersApi.md#updatecluster) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName} | Modify One Multi-Cloud Cluster from One Project
*NetworkPeeringApi* | [**CreatePeeringConnection**](docs/NetworkPeeringApi.md#createpeeringconnection) | **Post** /api/atlas/v2/groups/{groupId}/peers | Create One New Network Peering Connection
*NetworkPeeringApi* | [**CreatePeeringContainer**](docs/NetworkPeeringApi.md#createpeeringcontainer) | **Post** /api/atlas/v2/groups/{groupId}/containers | Create One New Network Peering Container
*NetworkPeeringApi* | [**DeletePeeringConnection**](docs/NetworkPeeringApi.md#deletepeeringconnection) | **Delete** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Remove One Existing Network Peering Connection
*NetworkPeeringApi* | [**DeletePeeringContainer**](docs/NetworkPeeringApi.md#deletepeeringcontainer) | **Delete** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Remove One Network Peering Container
*NetworkPeeringApi* | [**DisablePeering**](docs/NetworkPeeringApi.md#disablepeering) | **Patch** /api/atlas/v2/groups/{groupId}/privateIpMode | Disable Connect via Peering Only Mode for One Project
*NetworkPeeringApi* | [**GetPeeringConnection**](docs/NetworkPeeringApi.md#getpeeringconnection) | **Get** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Return One Network Peering Connection in One Project
*NetworkPeeringApi* | [**GetPeeringContainer**](docs/NetworkPeeringApi.md#getpeeringcontainer) | **Get** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Return One Network Peering Container
*NetworkPeeringApi* | [**ListPeeringConnections**](docs/NetworkPeeringApi.md#listpeeringconnections) | **Get** /api/atlas/v2/groups/{groupId}/peers | Return All Network Peering Connections in One Project
*NetworkPeeringApi* | [**ListPeeringContainerByCloudProvider**](docs/NetworkPeeringApi.md#listpeeringcontainerbycloudprovider) | **Get** /api/atlas/v2/groups/{groupId}/containers | Return All Network Peering Containers in One Project for One Cloud Provider
*NetworkPeeringApi* | [**ListPeeringContainers**](docs/NetworkPeeringApi.md#listpeeringcontainers) | **Get** /api/atlas/v2/groups/{groupId}/containers/all | Return All Network Peering Containers in One Project
*NetworkPeeringApi* | [**UpdatePeeringConnection**](docs/NetworkPeeringApi.md#updatepeeringconnection) | **Patch** /api/atlas/v2/groups/{groupId}/peers/{peerId} | Update One New Network Peering Connection
*NetworkPeeringApi* | [**UpdatePeeringContainer**](docs/NetworkPeeringApi.md#updatepeeringcontainer) | **Patch** /api/atlas/v2/groups/{groupId}/containers/{containerId} | Update One Network Peering Container
*NetworkPeeringApi* | [**VerifyConnectViaPeeringOnlyModeForOneProject**](docs/NetworkPeeringApi.md#verifyconnectviapeeringonlymodeforoneproject) | **Get** /api/atlas/v2/groups/{groupId}/privateIpMode | Verify Connect via Peering Only Mode for One Project
*OnlineArchiveApi* | [**CreateOnlineArchive**](docs/OnlineArchiveApi.md#createonlinearchive) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Create One Online Archive
*OnlineArchiveApi* | [**DeleteOnlineArchive**](docs/OnlineArchiveApi.md#deleteonlinearchive) | **Delete** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Remove One Online Archive
*OnlineArchiveApi* | [**DownloadOnlineArchiveQueryLogs**](docs/OnlineArchiveApi.md#downloadonlinearchivequerylogs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz | Download Online Archive Query Logs
*OnlineArchiveApi* | [**GetOnlineArchive**](docs/OnlineArchiveApi.md#getonlinearchive) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Return One Online Archive
*OnlineArchiveApi* | [**ListOnlineArchives**](docs/OnlineArchiveApi.md#listonlinearchives) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives | Return All Online Archives for One Cluster
*OnlineArchiveApi* | [**UpdateOnlineArchive**](docs/OnlineArchiveApi.md#updateonlinearchive) | **Patch** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId} | Update One Online Archive
*OrganizationsApi* | [**CreateOrganization**](docs/OrganizationsApi.md#createorganization) | **Post** /api/atlas/v2/orgs | Create One Organization
*OrganizationsApi* | [**CreateOrganizationInvitation**](docs/OrganizationsApi.md#createorganizationinvitation) | **Post** /api/atlas/v2/orgs/{orgId}/invites | Invite One MongoDB Cloud User to Join One Atlas Organization
*OrganizationsApi* | [**DeleteOrganization**](docs/OrganizationsApi.md#deleteorganization) | **Delete** /api/atlas/v2/orgs/{orgId} | Remove One Organization
*OrganizationsApi* | [**DeleteOrganizationInvitation**](docs/OrganizationsApi.md#deleteorganizationinvitation) | **Delete** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Cancel One Organization Invitation
*OrganizationsApi* | [**GetOrganization**](docs/OrganizationsApi.md#getorganization) | **Get** /api/atlas/v2/orgs/{orgId} | Return One Organization
*OrganizationsApi* | [**GetOrganizationInvitation**](docs/OrganizationsApi.md#getorganizationinvitation) | **Get** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Return One Organization Invitation
*OrganizationsApi* | [**GetOrganizationSettings**](docs/OrganizationsApi.md#getorganizationsettings) | **Get** /api/atlas/v2/orgs/{orgId}/settings | Return Settings for One Organization
*OrganizationsApi* | [**ListOrganizationInvitations**](docs/OrganizationsApi.md#listorganizationinvitations) | **Get** /api/atlas/v2/orgs/{orgId}/invites | Return All Organization Invitations
*OrganizationsApi* | [**ListOrganizationProjects**](docs/OrganizationsApi.md#listorganizationprojects) | **Get** /api/atlas/v2/orgs/{orgId}/groups | Return One or More Projects in One Organization
*OrganizationsApi* | [**ListOrganizationUsers**](docs/OrganizationsApi.md#listorganizationusers) | **Get** /api/atlas/v2/orgs/{orgId}/users | Return All MongoDB Cloud Users in One Organization
*OrganizationsApi* | [**ListOrganizations**](docs/OrganizationsApi.md#listorganizations) | **Get** /api/atlas/v2/orgs | Return All Organizations
*OrganizationsApi* | [**RenameOrganization**](docs/OrganizationsApi.md#renameorganization) | **Patch** /api/atlas/v2/orgs/{orgId} | Rename One Organization
*OrganizationsApi* | [**UpdateOrganizationInvitation**](docs/OrganizationsApi.md#updateorganizationinvitation) | **Patch** /api/atlas/v2/orgs/{orgId}/invites | Update One Organization Invitation
*OrganizationsApi* | [**UpdateOrganizationInvitationById**](docs/OrganizationsApi.md#updateorganizationinvitationbyid) | **Patch** /api/atlas/v2/orgs/{orgId}/invites/{invitationId} | Update One Organization Invitation by Invitation ID
*OrganizationsApi* | [**UpdateOrganizationSettings**](docs/OrganizationsApi.md#updateorganizationsettings) | **Patch** /api/atlas/v2/orgs/{orgId}/settings | Update Settings for One Organization
*PerformanceAdvisorApi* | [**DisableSlowOperationThresholding**](docs/PerformanceAdvisorApi.md#disableslowoperationthresholding) | **Delete** /api/atlas/v2/groups/{groupId}/managedSlowMs/disable | Disable Managed Slow Operation Threshold
*PerformanceAdvisorApi* | [**EnableSlowOperationThresholding**](docs/PerformanceAdvisorApi.md#enableslowoperationthresholding) | **Post** /api/atlas/v2/groups/{groupId}/managedSlowMs/enable | Enable Managed Slow Operation Threshold
*PerformanceAdvisorApi* | [**ListSlowQueries**](docs/PerformanceAdvisorApi.md#listslowqueries) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs | Return Slow Queries
*PerformanceAdvisorApi* | [**ListSlowQueryNamespaces**](docs/PerformanceAdvisorApi.md#listslowquerynamespaces) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces | Return All Namespaces for One Host
*PerformanceAdvisorApi* | [**ListSuggestedIndexes**](docs/PerformanceAdvisorApi.md#listsuggestedindexes) | **Get** /api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes | Return Suggested Indexes
*PrivateEndpointServicesApi* | [**CreatePrivateEndpoint**](docs/PrivateEndpointServicesApi.md#createprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint | Create One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**CreatePrivateEndpointService**](docs/PrivateEndpointServicesApi.md#createprivateendpointservice) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/endpointService | Create One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**DeletePrivateEndpoint**](docs/PrivateEndpointServicesApi.md#deleteprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Remove One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**DeletePrivateEndpointService**](docs/PrivateEndpointServicesApi.md#deleteprivateendpointservice) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Remove One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**GetPrivateEndpoint**](docs/PrivateEndpointServicesApi.md#getprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId}/endpoint/{endpointId} | Return One Private Endpoint for One Provider
*PrivateEndpointServicesApi* | [**GetPrivateEndpointService**](docs/PrivateEndpointServicesApi.md#getprivateendpointservice) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService/{endpointServiceId} | Return One Private Endpoint Service for One Provider
*PrivateEndpointServicesApi* | [**GetRegionalizedPrivateEndpointSetting**](docs/PrivateEndpointServicesApi.md#getregionalizedprivateendpointsetting) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Return Regionalized Private Endpoint Status
*PrivateEndpointServicesApi* | [**ListPrivateEndpointServices**](docs/PrivateEndpointServicesApi.md#listprivateendpointservices) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/{cloudProvider}/endpointService | Return All Private Endpoint Services for One Provider
*PrivateEndpointServicesApi* | [**ToggleRegionalizedPrivateEndpointSetting**](docs/PrivateEndpointServicesApi.md#toggleregionalizedprivateendpointsetting) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/regionalMode | Toggle Regionalized Private Endpoint Status
*ProgrammaticAPIKeysApi* | [**AddProjectApiKey**](docs/ProgrammaticAPIKeysApi.md#addprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [**CreateApiKey**](docs/ProgrammaticAPIKeysApi.md#createapikey) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys | Create One Organization API Key
*ProgrammaticAPIKeysApi* | [**CreateApiKeyAccessList**](docs/ProgrammaticAPIKeysApi.md#createapikeyaccesslist) | **Post** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Create Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [**CreateProjectApiKey**](docs/ProgrammaticAPIKeysApi.md#createprojectapikey) | **Post** /api/atlas/v2/groups/{groupId}/apiKeys | Create and Assign One Organization API Key to One Project
*ProgrammaticAPIKeysApi* | [**DeleteApiKey**](docs/ProgrammaticAPIKeysApi.md#deleteapikey) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Remove One Organization API Key
*ProgrammaticAPIKeysApi* | [**DeleteApiKeyAcessList**](docs/ProgrammaticAPIKeysApi.md#deleteapikeyacesslist) | **Delete** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Remove One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [**GetApiKey**](docs/ProgrammaticAPIKeysApi.md#getapikey) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Return One Organization API Key
*ProgrammaticAPIKeysApi* | [**GetApiKeyAccessList**](docs/ProgrammaticAPIKeysApi.md#getapikeyaccesslist) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList/{ipAddress} | Return One Access List Entry for One Organization API Key
*ProgrammaticAPIKeysApi* | [**ListApiKeyAccessListsEntries**](docs/ProgrammaticAPIKeysApi.md#listapikeyaccesslistsentries) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId}/accessList | Return All Access List Entries for One Organization API Key
*ProgrammaticAPIKeysApi* | [**ListApiKeys**](docs/ProgrammaticAPIKeysApi.md#listapikeys) | **Get** /api/atlas/v2/orgs/{orgId}/apiKeys | Return All Organization API Keys
*ProgrammaticAPIKeysApi* | [**ListProjectApiKeys**](docs/ProgrammaticAPIKeysApi.md#listprojectapikeys) | **Get** /api/atlas/v2/groups/{groupId}/apiKeys | Return All Organization API Keys Assigned to One Project
*ProgrammaticAPIKeysApi* | [**RemoveProjectApiKey**](docs/ProgrammaticAPIKeysApi.md#removeprojectapikey) | **Delete** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Unassign One Organization API Key from One Project
*ProgrammaticAPIKeysApi* | [**UpdateApiKey**](docs/ProgrammaticAPIKeysApi.md#updateapikey) | **Patch** /api/atlas/v2/orgs/{orgId}/apiKeys/{apiUserId} | Update One Organization API Key
*ProgrammaticAPIKeysApi* | [**UpdateApiKeyRoles**](docs/ProgrammaticAPIKeysApi.md#updateapikeyroles) | **Patch** /api/atlas/v2/groups/{groupId}/apiKeys/{apiUserId} | Update Roles of One Organization API Key to One Project
*ProjectIPAccessListApi* | [**CreateProjectIpAccessList**](docs/ProjectIPAccessListApi.md#createprojectipaccesslist) | **Post** /api/atlas/v2/groups/{groupId}/accessList | Add Entries to Project IP Access List
*ProjectIPAccessListApi* | [**DeleteProjectIpAccessList**](docs/ProjectIPAccessListApi.md#deleteprojectipaccesslist) | **Delete** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Remove One Entry from One Project IP Access List
*ProjectIPAccessListApi* | [**GetProjectIpAccessListStatus**](docs/ProjectIPAccessListApi.md#getprojectipaccessliststatus) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue}/status | Return Status of One Project IP Access List Entry
*ProjectIPAccessListApi* | [**GetProjectIpList**](docs/ProjectIPAccessListApi.md#getprojectiplist) | **Get** /api/atlas/v2/groups/{groupId}/accessList/{entryValue} | Return One Project IP Access List Entry
*ProjectIPAccessListApi* | [**ListProjectIpAccessLists**](docs/ProjectIPAccessListApi.md#listprojectipaccesslists) | **Get** /api/atlas/v2/groups/{groupId}/accessList | Return Project IP Access List
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /api/atlas/v2/groups | Create One Project
*ProjectsApi* | [**CreateProjectInvitation**](docs/ProjectsApi.md#createprojectinvitation) | **Post** /api/atlas/v2/groups/{groupId}/invites | Invite One MongoDB Cloud User to Join One Project
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /api/atlas/v2/groups/{groupId} | Remove One Project
*ProjectsApi* | [**DeleteProjectInvitation**](docs/ProjectsApi.md#deleteprojectinvitation) | **Delete** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Cancel One Project Invitation
*ProjectsApi* | [**DeleteProjectLimit**](docs/ProjectsApi.md#deleteprojectlimit) | **Delete** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Remove One Project Limit
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /api/atlas/v2/groups/{groupId} | Return One Project
*ProjectsApi* | [**GetProjectByName**](docs/ProjectsApi.md#getprojectbyname) | **Get** /api/atlas/v2/groups/byName/{groupName} | Return One Project using Its Name
*ProjectsApi* | [**GetProjectInvitation**](docs/ProjectsApi.md#getprojectinvitation) | **Get** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Return One Project Invitation
*ProjectsApi* | [**GetProjectLimit**](docs/ProjectsApi.md#getprojectlimit) | **Get** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Return One Limit for One Project
*ProjectsApi* | [**GetProjectSettings**](docs/ProjectsApi.md#getprojectsettings) | **Get** /api/atlas/v2/groups/{groupId}/settings | Return One Project Settings
*ProjectsApi* | [**ListProjectInvitations**](docs/ProjectsApi.md#listprojectinvitations) | **Get** /api/atlas/v2/groups/{groupId}/invites | Return All Project Invitations
*ProjectsApi* | [**ListProjectLimits**](docs/ProjectsApi.md#listprojectlimits) | **Get** /api/atlas/v2/groups/{groupId}/limits | Return All Limits for One Project
*ProjectsApi* | [**ListProjectUsers**](docs/ProjectsApi.md#listprojectusers) | **Get** /api/atlas/v2/groups/{groupId}/users | Return All Users in One Project
*ProjectsApi* | [**ListProjects**](docs/ProjectsApi.md#listprojects) | **Get** /api/atlas/v2/groups | Return All Projects
*ProjectsApi* | [**RemoveProjectUser**](docs/ProjectsApi.md#removeprojectuser) | **Delete** /api/atlas/v2/groups/{groupId}/users/{userId} | Remove One User from One Project
*ProjectsApi* | [**SetProjectLimit**](docs/ProjectsApi.md#setprojectlimit) | **Patch** /api/atlas/v2/groups/{groupId}/limits/{limitName} | Set One Project Limit
*ProjectsApi* | [**UpdateProject**](docs/ProjectsApi.md#updateproject) | **Patch** /api/atlas/v2/groups/{groupId} | Update One Project Name
*ProjectsApi* | [**UpdateProjectInvitation**](docs/ProjectsApi.md#updateprojectinvitation) | **Patch** /api/atlas/v2/groups/{groupId}/invites | Update One Project Invitation
*ProjectsApi* | [**UpdateProjectInvitationById**](docs/ProjectsApi.md#updateprojectinvitationbyid) | **Patch** /api/atlas/v2/groups/{groupId}/invites/{invitationId} | Update One Project Invitation by Invitation ID
*ProjectsApi* | [**UpdateProjectSettings**](docs/ProjectsApi.md#updateprojectsettings) | **Patch** /api/atlas/v2/groups/{groupId}/settings | Update One Project Settings
*RollingIndexApi* | [**CreateRollingIndex**](docs/RollingIndexApi.md#createrollingindex) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/index | Create One Rolling Index
*RootApi* | [**GetSystemStatus**](docs/RootApi.md#getsystemstatus) | **Get** /api/atlas/v2 | Return the status of this MongoDB application
*ServerlessInstancesApi* | [**CreateServerlessInstance**](docs/ServerlessInstancesApi.md#createserverlessinstance) | **Post** /api/atlas/v2/groups/{groupId}/serverless | Create One Serverless Instance in One Project
*ServerlessInstancesApi* | [**DeleteServerlessInstance**](docs/ServerlessInstancesApi.md#deleteserverlessinstance) | **Delete** /api/atlas/v2/groups/{groupId}/serverless/{name} | Remove One Serverless Instance from One Project
*ServerlessInstancesApi* | [**GetServerlessInstance**](docs/ServerlessInstancesApi.md#getserverlessinstance) | **Get** /api/atlas/v2/groups/{groupId}/serverless/{name} | Return One Serverless Instance from One Project
*ServerlessInstancesApi* | [**ListServerlessInstances**](docs/ServerlessInstancesApi.md#listserverlessinstances) | **Get** /api/atlas/v2/groups/{groupId}/serverless | Return All Serverless Instances from One Project
*ServerlessInstancesApi* | [**UpdateServerlessInstance**](docs/ServerlessInstancesApi.md#updateserverlessinstance) | **Patch** /api/atlas/v2/groups/{groupId}/serverless/{name} | Update One Serverless Instance in One Project
*ServerlessPrivateEndpointsApi* | [**CreateServerlessPrivateEndpoint**](docs/ServerlessPrivateEndpointsApi.md#createserverlessprivateendpoint) | **Post** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Create One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**DeleteServerlessPrivateEndpoint**](docs/ServerlessPrivateEndpointsApi.md#deleteserverlessprivateendpoint) | **Delete** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Remove One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**GetServerlessPrivateEndpoint**](docs/ServerlessPrivateEndpointsApi.md#getserverlessprivateendpoint) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Return One Private Endpoint for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**ListServerlessPrivateEndpoints**](docs/ServerlessPrivateEndpointsApi.md#listserverlessprivateendpoints) | **Get** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint | Return All Private Endpoints for One Serverless Instance
*ServerlessPrivateEndpointsApi* | [**UpdateServerlessPrivateEndpoint**](docs/ServerlessPrivateEndpointsApi.md#updateserverlessprivateendpoint) | **Patch** /api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId} | Update One Private Endpoint for One Serverless Instance
*SharedTierRestoreJobsApi* | [**CreateSharedClusterBackupRestoreJob**](docs/SharedTierRestoreJobsApi.md#createsharedclusterbackuprestorejob) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore | Create One Restore Job from One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [**GetSharedClusterBackupRestoreJob**](docs/SharedTierRestoreJobsApi.md#getsharedclusterbackuprestorejob) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId} | Return One Restore Job for One M2 or M5 Cluster
*SharedTierRestoreJobsApi* | [**ListSharedClusterBackupRestoreJobs**](docs/SharedTierRestoreJobsApi.md#listsharedclusterbackuprestorejobs) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores | Return All Restore Jobs for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [**DownloadSharedClusterBackup**](docs/SharedTierSnapshotsApi.md#downloadsharedclusterbackup) | **Post** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download | Download One M2 or M5 Cluster Snapshot
*SharedTierSnapshotsApi* | [**GetSharedClusterBackup**](docs/SharedTierSnapshotsApi.md#getsharedclusterbackup) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId} | Return One Snapshot for One M2 or M5 Cluster
*SharedTierSnapshotsApi* | [**ListSharedClusterBackups**](docs/SharedTierSnapshotsApi.md#listsharedclusterbackups) | **Get** /api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots | Return All Snapshots for One M2 or M5 Cluster
*TeamsApi* | [**AddAllTeamsToProject**](docs/TeamsApi.md#addallteamstoproject) | **Post** /api/atlas/v2/groups/{groupId}/teams | Add One or More Teams to One Project
*TeamsApi* | [**AddTeamUser**](docs/TeamsApi.md#addteamuser) | **Post** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Assign MongoDB Cloud Users from One Organization to One Team
*TeamsApi* | [**CreateTeam**](docs/TeamsApi.md#createteam) | **Post** /api/atlas/v2/orgs/{orgId}/teams | Create One Team in One Organization
*TeamsApi* | [**DeleteTeam**](docs/TeamsApi.md#deleteteam) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Remove One Team from One Organization
*TeamsApi* | [**GetTeamById**](docs/TeamsApi.md#getteambyid) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Return One Team using its ID
*TeamsApi* | [**GetTeamByName**](docs/TeamsApi.md#getteambyname) | **Get** /api/atlas/v2/orgs/{orgId}/teams/byName/{teamName} | Return One Team using its Name
*TeamsApi* | [**ListOrganizationTeams**](docs/TeamsApi.md#listorganizationteams) | **Get** /api/atlas/v2/orgs/{orgId}/teams | Return All Teams in One Organization
*TeamsApi* | [**ListProjectTeams**](docs/TeamsApi.md#listprojectteams) | **Get** /api/atlas/v2/groups/{groupId}/teams | Return All Teams in One Project
*TeamsApi* | [**ListTeamUsers**](docs/TeamsApi.md#listteamusers) | **Get** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users | Return All MongoDB Cloud Users Assigned to One Team
*TeamsApi* | [**RemoveProjectTeam**](docs/TeamsApi.md#removeprojectteam) | **Delete** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Remove One Team from One Project
*TeamsApi* | [**RemoveTeamUser**](docs/TeamsApi.md#removeteamuser) | **Delete** /api/atlas/v2/orgs/{orgId}/teams/{teamId}/users/{userId} | Remove One MongoDB Cloud User from One Team
*TeamsApi* | [**RenameTeam**](docs/TeamsApi.md#renameteam) | **Patch** /api/atlas/v2/orgs/{orgId}/teams/{teamId} | Rename One Team
*TeamsApi* | [**UpdateTeamRoles**](docs/TeamsApi.md#updateteamroles) | **Patch** /api/atlas/v2/groups/{groupId}/teams/{teamId} | Update Team Roles in One Project
*TestApi* | [**VersionedExample**](docs/TestApi.md#versionedexample) | **Get** /api/atlas/v2/example/info | Example resource info for versioning of the Atlas API
*ThirdPartyIntegrationsApi* | [**CreateThirdPartyIntegration**](docs/ThirdPartyIntegrationsApi.md#createthirdpartyintegration) | **Post** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Configure One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**DeleteThirdPartyIntegration**](docs/ThirdPartyIntegrationsApi.md#deletethirdpartyintegration) | **Delete** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Remove One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**GetThirdPartyIntegration**](docs/ThirdPartyIntegrationsApi.md#getthirdpartyintegration) | **Get** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Return One Third-Party Service Integration
*ThirdPartyIntegrationsApi* | [**ListThirdPartyIntegrations**](docs/ThirdPartyIntegrationsApi.md#listthirdpartyintegrations) | **Get** /api/atlas/v2/groups/{groupId}/integrations | Return All Active Third-Party Service Integrations
*ThirdPartyIntegrationsApi* | [**UpdateThirdPartyIntegration**](docs/ThirdPartyIntegrationsApi.md#updatethirdpartyintegration) | **Put** /api/atlas/v2/groups/{groupId}/integrations/{integrationType} | Update One Third-Party Service Integration
*X509AuthenticationApi* | [**CreateDatabaseUserCertificate**](docs/X509AuthenticationApi.md#createdatabaseusercertificate) | **Post** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Create One X.509 Certificate for One MongoDB User
*X509AuthenticationApi* | [**DisableCustomerManagedX509**](docs/X509AuthenticationApi.md#disablecustomermanagedx509) | **Delete** /api/atlas/v2/groups/{groupId}/userSecurity/customerX509 | Disable Customer-Managed X.509
*X509AuthenticationApi* | [**ListDatabaseUserCertificates**](docs/X509AuthenticationApi.md#listdatabaseusercertificates) | **Get** /api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs | Return All X.509 Certificates Assigned to One MongoDB User


## Documentation For Models

 - [AWSAutoScaling](docs/AWSAutoScaling.md)
 - [AWSCloudProviderContainer](docs/AWSCloudProviderContainer.md)
 - [AWSComputeAutoScaling](docs/AWSComputeAutoScaling.md)
 - [AWSCustomDNSEnabled](docs/AWSCustomDNSEnabled.md)
 - [AWSHardwareSpec](docs/AWSHardwareSpec.md)
 - [AWSInterfaceEndpoint](docs/AWSInterfaceEndpoint.md)
 - [AWSKMS](docs/AWSKMS.md)
 - [AWSPeerVpc](docs/AWSPeerVpc.md)
 - [AWSPeerVpcRequest](docs/AWSPeerVpcRequest.md)
 - [AWSPrivateLinkConnection](docs/AWSPrivateLinkConnection.md)
 - [AWSProviderSettings](docs/AWSProviderSettings.md)
 - [AWSRegionConfig](docs/AWSRegionConfig.md)
 - [AccessListItem](docs/AccessListItem.md)
 - [AddUserToTeam](docs/AddUserToTeam.md)
 - [Alert](docs/Alert.md)
 - [AlertAudit](docs/AlertAudit.md)
 - [AlertAuditType](docs/AlertAuditType.md)
 - [AlertConfig](docs/AlertConfig.md)
 - [AlertConfigAudit](docs/AlertConfigAudit.md)
 - [AlertConfigAuditType](docs/AlertConfigAuditType.md)
 - [AlertConfigViewForNdsGroup](docs/AlertConfigViewForNdsGroup.md)
 - [AlertViewForNdsGroup](docs/AlertViewForNdsGroup.md)
 - [ApiInstanceSizeView](docs/ApiInstanceSizeView.md)
 - [ApiRoleAssignmentView](docs/ApiRoleAssignmentView.md)
 - [ApiUserEventTypeViewForNdsGroup](docs/ApiUserEventTypeViewForNdsGroup.md)
 - [ApiUserEventTypeViewForOrg](docs/ApiUserEventTypeViewForOrg.md)
 - [ApiUserEventViewForNdsGroup](docs/ApiUserEventViewForNdsGroup.md)
 - [ApiUserEventViewForOrg](docs/ApiUserEventViewForOrg.md)
 - [AppServiceAlert](docs/AppServiceAlert.md)
 - [AppServiceAlertConfigViewForNdsGroup](docs/AppServiceAlertConfigViewForNdsGroup.md)
 - [AppServiceEvent](docs/AppServiceEvent.md)
 - [AppServiceEventType](docs/AppServiceEventType.md)
 - [AppServiceEventTypeViewAlertable](docs/AppServiceEventTypeViewAlertable.md)
 - [AppServiceEventTypeViewAlertableNoThreshold](docs/AppServiceEventTypeViewAlertableNoThreshold.md)
 - [AppServiceEventTypeViewAlertableWithThreshold](docs/AppServiceEventTypeViewAlertableWithThreshold.md)
 - [AppServiceMetricAlertConfigViewForNdsGroup](docs/AppServiceMetricAlertConfigViewForNdsGroup.md)
 - [AppServiceMetricMatcher](docs/AppServiceMetricMatcher.md)
 - [AppServiceMetricMatcherField](docs/AppServiceMetricMatcherField.md)
 - [AppServiceMetricThreshold](docs/AppServiceMetricThreshold.md)
 - [AppUser](docs/AppUser.md)
 - [AtlasContainerPeerViewRequest](docs/AtlasContainerPeerViewRequest.md)
 - [AtlasFTSAnalyzersViewManual](docs/AtlasFTSAnalyzersViewManual.md)
 - [AtlasFTSAnalyzersViewManualCharFiltersInner](docs/AtlasFTSAnalyzersViewManualCharFiltersInner.md)
 - [AtlasFTSAnalyzersViewManualTokenFiltersInner](docs/AtlasFTSAnalyzersViewManualTokenFiltersInner.md)
 - [AtlasFTSAnalyzersViewManualTokenizer](docs/AtlasFTSAnalyzersViewManualTokenizer.md)
 - [AtlasFTSMappingsViewManual](docs/AtlasFTSMappingsViewManual.md)
 - [AtlasNetPeerRequestBase](docs/AtlasNetPeerRequestBase.md)
 - [AuditLog](docs/AuditLog.md)
 - [AutoExportPolicy](docs/AutoExportPolicy.md)
 - [AutoScaling](docs/AutoScaling.md)
 - [AutoScalingV15](docs/AutoScalingV15.md)
 - [AutomationConfigEvent](docs/AutomationConfigEvent.md)
 - [AutomationConfigEventType](docs/AutomationConfigEventType.md)
 - [AvailableDeployment](docs/AvailableDeployment.md)
 - [AvailableProject](docs/AvailableProject.md)
 - [AvailableRegion](docs/AvailableRegion.md)
 - [AzureAutoScaling](docs/AzureAutoScaling.md)
 - [AzureCloudProviderContainer](docs/AzureCloudProviderContainer.md)
 - [AzureComputeAutoScaling](docs/AzureComputeAutoScaling.md)
 - [AzureHardwareSpec](docs/AzureHardwareSpec.md)
 - [AzureKeyVault](docs/AzureKeyVault.md)
 - [AzurePeerNetwork](docs/AzurePeerNetwork.md)
 - [AzurePeerNetworkRequest](docs/AzurePeerNetworkRequest.md)
 - [AzurePrivateEndpoint](docs/AzurePrivateEndpoint.md)
 - [AzurePrivateLinkConnection](docs/AzurePrivateLinkConnection.md)
 - [AzureProviderSettings](docs/AzureProviderSettings.md)
 - [AzureRegionConfig](docs/AzureRegionConfig.md)
 - [BSONTimestamp](docs/BSONTimestamp.md)
 - [BiConnector](docs/BiConnector.md)
 - [BillingEventTypeViewAlertableWithThreshold](docs/BillingEventTypeViewAlertableWithThreshold.md)
 - [BillingEventTypeViewForNdsGroup](docs/BillingEventTypeViewForNdsGroup.md)
 - [BillingEventTypeViewForOrg](docs/BillingEventTypeViewForOrg.md)
 - [BillingEventViewForNdsGroup](docs/BillingEventViewForNdsGroup.md)
 - [BillingEventViewForOrg](docs/BillingEventViewForOrg.md)
 - [BillingThresholdAlertConfigViewForNdsGroup](docs/BillingThresholdAlertConfigViewForNdsGroup.md)
 - [CharFilterhtmlStrip](docs/CharFilterhtmlStrip.md)
 - [CharFiltericuNormalize](docs/CharFiltericuNormalize.md)
 - [CharFiltermapping](docs/CharFiltermapping.md)
 - [CharFiltermappingMappings](docs/CharFiltermappingMappings.md)
 - [CharFilterpersian](docs/CharFilterpersian.md)
 - [Checkpoint](docs/Checkpoint.md)
 - [CheckpointPart](docs/CheckpointPart.md)
 - [CloudProviderAccess](docs/CloudProviderAccess.md)
 - [CloudProviderAccessAWSIAMRole](docs/CloudProviderAccessAWSIAMRole.md)
 - [CloudProviderAccessDataLakeFeatureUsage](docs/CloudProviderAccessDataLakeFeatureUsage.md)
 - [CloudProviderAccessEncryptionAtRestFeatureUsage](docs/CloudProviderAccessEncryptionAtRestFeatureUsage.md)
 - [CloudProviderAccessExportSnapshotFeatureUsage](docs/CloudProviderAccessExportSnapshotFeatureUsage.md)
 - [CloudProviderAccessFeatureUsage](docs/CloudProviderAccessFeatureUsage.md)
 - [CloudProviderAccessFeatureUsageDataLakeFeatureId](docs/CloudProviderAccessFeatureUsageDataLakeFeatureId.md)
 - [CloudProviderAccessFeatureUsageExportSnapshotFeatureId](docs/CloudProviderAccessFeatureUsageExportSnapshotFeatureId.md)
 - [CloudProviderAccessRole](docs/CloudProviderAccessRole.md)
 - [CloudProviderContainer](docs/CloudProviderContainer.md)
 - [Cluster](docs/Cluster.md)
 - [ClusterAlert](docs/ClusterAlert.md)
 - [ClusterAlertConfigViewForNdsGroup](docs/ClusterAlertConfigViewForNdsGroup.md)
 - [ClusterDescriptionConnectionStrings](docs/ClusterDescriptionConnectionStrings.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpoint](docs/ClusterDescriptionConnectionStringsPrivateEndpoint.md)
 - [ClusterDescriptionConnectionStringsPrivateEndpointEndpoint](docs/ClusterDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ClusterDescriptionProcessArgs](docs/ClusterDescriptionProcessArgs.md)
 - [ClusterDescriptionV15](docs/ClusterDescriptionV15.md)
 - [ClusterEventTypeViewAlertable](docs/ClusterEventTypeViewAlertable.md)
 - [ClusterEventTypeViewForNdsGroup](docs/ClusterEventTypeViewForNdsGroup.md)
 - [ClusterEventViewForNdsGroup](docs/ClusterEventViewForNdsGroup.md)
 - [ClusterMatcher](docs/ClusterMatcher.md)
 - [ClusterMatcherField](docs/ClusterMatcherField.md)
 - [ClusterOutageSimulation](docs/ClusterOutageSimulation.md)
 - [ClusterOutageSimulationOutageFilter](docs/ClusterOutageSimulationOutageFilter.md)
 - [ClusterProviderSettings](docs/ClusterProviderSettings.md)
 - [ClusterStatus](docs/ClusterStatus.md)
 - [Collation](docs/Collation.md)
 - [ComputeAutoScaling](docs/ComputeAutoScaling.md)
 - [ComputeAutoScalingV15](docs/ComputeAutoScalingV15.md)
 - [ConnectedOrgConfig](docs/ConnectedOrgConfig.md)
 - [ContainerPeer](docs/ContainerPeer.md)
 - [CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold](docs/CpsBackupEventTypeViewForNdsGroupAlertableWithThreshold.md)
 - [CpsBackupThresholdAlertConfigViewForNdsGroup](docs/CpsBackupThresholdAlertConfigViewForNdsGroup.md)
 - [CreateAWSEndpointRequest](docs/CreateAWSEndpointRequest.md)
 - [CreateAWSEndpointRequestAllOf](docs/CreateAWSEndpointRequestAllOf.md)
 - [CreateApiKey](docs/CreateApiKey.md)
 - [CreateAzureEndpointRequest](docs/CreateAzureEndpointRequest.md)
 - [CreateAzureEndpointRequestAllOf](docs/CreateAzureEndpointRequestAllOf.md)
 - [CreateEndpointServiceRequest](docs/CreateEndpointServiceRequest.md)
 - [CreateGCPEndpointGroupRequest](docs/CreateGCPEndpointGroupRequest.md)
 - [CreateGCPEndpointGroupRequestAllOf](docs/CreateGCPEndpointGroupRequestAllOf.md)
 - [CreateGCPForwardingRuleRequest](docs/CreateGCPForwardingRuleRequest.md)
 - [CreateOrganizationRequest](docs/CreateOrganizationRequest.md)
 - [CreateOrganizationResponse](docs/CreateOrganizationResponse.md)
 - [CreatePeeringConnection200Response](docs/CreatePeeringConnection200Response.md)
 - [CreatePrivateEndpointRequest](docs/CreatePrivateEndpointRequest.md)
 - [Criteria](docs/Criteria.md)
 - [CustomCriteria](docs/CustomCriteria.md)
 - [CustomDBRole](docs/CustomDBRole.md)
 - [CustomZoneMappings](docs/CustomZoneMappings.md)
 - [CustomerX509](docs/CustomerX509.md)
 - [DBAction](docs/DBAction.md)
 - [DBResource](docs/DBResource.md)
 - [DLSIngestionSink](docs/DLSIngestionSink.md)
 - [DailySchedule](docs/DailySchedule.md)
 - [DataExplorerAccessedEvent](docs/DataExplorerAccessedEvent.md)
 - [DataExplorerAccessedEventType](docs/DataExplorerAccessedEventType.md)
 - [DataFederationQueryLimit](docs/DataFederationQueryLimit.md)
 - [DataFederationTenantQueryLimit](docs/DataFederationTenantQueryLimit.md)
 - [DataLake](docs/DataLake.md)
 - [DataLakeAWSCloudProviderConfig](docs/DataLakeAWSCloudProviderConfig.md)
 - [DataLakeAtlasStore](docs/DataLakeAtlasStore.md)
 - [DataLakeAtlasStoreReadPreference](docs/DataLakeAtlasStoreReadPreference.md)
 - [DataLakeAtlasStoreReadPreferenceTag](docs/DataLakeAtlasStoreReadPreferenceTag.md)
 - [DataLakeCloudProviderConfig](docs/DataLakeCloudProviderConfig.md)
 - [DataLakeDataProcessRegion](docs/DataLakeDataProcessRegion.md)
 - [DataLakeDatabase](docs/DataLakeDatabase.md)
 - [DataLakeDatabaseCollection](docs/DataLakeDatabaseCollection.md)
 - [DataLakeDatabaseDataSource](docs/DataLakeDatabaseDataSource.md)
 - [DataLakeHTTPStore](docs/DataLakeHTTPStore.md)
 - [DataLakeOnlineArchiveStore](docs/DataLakeOnlineArchiveStore.md)
 - [DataLakeRegion](docs/DataLakeRegion.md)
 - [DataLakeS3Store](docs/DataLakeS3Store.md)
 - [DataLakeStorage](docs/DataLakeStorage.md)
 - [DataLakeStore](docs/DataLakeStore.md)
 - [DataLakeTenant](docs/DataLakeTenant.md)
 - [DataMetricAlert](docs/DataMetricAlert.md)
 - [DataMetricEvent](docs/DataMetricEvent.md)
 - [DataMetricThreshold](docs/DataMetricThreshold.md)
 - [DataMetricUnits](docs/DataMetricUnits.md)
 - [DataMetricValue](docs/DataMetricValue.md)
 - [DataProtectionSettings](docs/DataProtectionSettings.md)
 - [Database](docs/Database.md)
 - [DatabaseUser](docs/DatabaseUser.md)
 - [Datadog](docs/Datadog.md)
 - [DatadogNotification](docs/DatadogNotification.md)
 - [DateCriteria](docs/DateCriteria.md)
 - [DedicatedHardwareSpec](docs/DedicatedHardwareSpec.md)
 - [DefaultAlertConfigViewForNdsGroup](docs/DefaultAlertConfigViewForNdsGroup.md)
 - [DefaultAlertViewForNdsGroup](docs/DefaultAlertViewForNdsGroup.md)
 - [DefaultEventViewForNdsGroup](docs/DefaultEventViewForNdsGroup.md)
 - [DefaultEventViewForOrg](docs/DefaultEventViewForOrg.md)
 - [DefaultSchedule](docs/DefaultSchedule.md)
 - [DeleteCopiedBackups](docs/DeleteCopiedBackups.md)
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
 - [DiskPartition](docs/DiskPartition.md)
 - [EmailNotification](docs/EmailNotification.md)
 - [EncryptionAtRest](docs/EncryptionAtRest.md)
 - [EncryptionKeyAlertConfigViewForNdsGroup](docs/EncryptionKeyAlertConfigViewForNdsGroup.md)
 - [EncryptionKeyEventTypeViewAlertable](docs/EncryptionKeyEventTypeViewAlertable.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointService](docs/EndpointService.md)
 - [Error](docs/Error.md)
 - [EventTypeForNdsGroup](docs/EventTypeForNdsGroup.md)
 - [EventTypeForOrg](docs/EventTypeForOrg.md)
 - [EventViewForNdsGroup](docs/EventViewForNdsGroup.md)
 - [EventViewForOrg](docs/EventViewForOrg.md)
 - [ExampleResourceResponseView20230101](docs/ExampleResourceResponseView20230101.md)
 - [ExampleResourceResponseView20230201](docs/ExampleResourceResponseView20230201.md)
 - [ExportStatus](docs/ExportStatus.md)
 - [FTSIndex](docs/FTSIndex.md)
 - [FTSIndexAudit](docs/FTSIndexAudit.md)
 - [FTSIndexAuditType](docs/FTSIndexAuditType.md)
 - [FTSMetric](docs/FTSMetric.md)
 - [FTSMetrics](docs/FTSMetrics.md)
 - [FTSSynonymMappingDefinition](docs/FTSSynonymMappingDefinition.md)
 - [FederatedUser](docs/FederatedUser.md)
 - [FieldTransformation](docs/FieldTransformation.md)
 - [ForNdsGroup](docs/ForNdsGroup.md)
 - [FreeAutoScaling](docs/FreeAutoScaling.md)
 - [FreeProviderSettings](docs/FreeProviderSettings.md)
 - [GCPAutoScaling](docs/GCPAutoScaling.md)
 - [GCPCloudProviderContainer](docs/GCPCloudProviderContainer.md)
 - [GCPComputeAutoScaling](docs/GCPComputeAutoScaling.md)
 - [GCPConsumerForwardingRule](docs/GCPConsumerForwardingRule.md)
 - [GCPEndpointGroup](docs/GCPEndpointGroup.md)
 - [GCPEndpointService](docs/GCPEndpointService.md)
 - [GCPHardwareSpec](docs/GCPHardwareSpec.md)
 - [GCPPeerVpc](docs/GCPPeerVpc.md)
 - [GCPPeerVpcRequest](docs/GCPPeerVpcRequest.md)
 - [GCPProviderSettings](docs/GCPProviderSettings.md)
 - [GCPRegionConfig](docs/GCPRegionConfig.md)
 - [GeoSharding](docs/GeoSharding.md)
 - [GetPeeringConnection200Response](docs/GetPeeringConnection200Response.md)
 - [GoogleCloudKMS](docs/GoogleCloudKMS.md)
 - [GreaterThanDaysThreshold](docs/GreaterThanDaysThreshold.md)
 - [GreaterThanRawThreshold](docs/GreaterThanRawThreshold.md)
 - [GreaterThanRawThresholdAlertConfigViewForNdsGroup](docs/GreaterThanRawThresholdAlertConfigViewForNdsGroup.md)
 - [GreaterThanTimeThreshold](docs/GreaterThanTimeThreshold.md)
 - [Group](docs/Group.md)
 - [GroupInvitation](docs/GroupInvitation.md)
 - [GroupInvitationRequest](docs/GroupInvitationRequest.md)
 - [GroupInvitationUpdateRequest](docs/GroupInvitationUpdateRequest.md)
 - [GroupMaintenanceWindow](docs/GroupMaintenanceWindow.md)
 - [GroupName](docs/GroupName.md)
 - [GroupNotification](docs/GroupNotification.md)
 - [GroupPaginatedEvent](docs/GroupPaginatedEvent.md)
 - [GroupPaginatedIntegration](docs/GroupPaginatedIntegration.md)
 - [GroupSettings](docs/GroupSettings.md)
 - [HardwareSpec](docs/HardwareSpec.md)
 - [HipChatNotification](docs/HipChatNotification.md)
 - [HostAlertConfigViewForNdsGroup](docs/HostAlertConfigViewForNdsGroup.md)
 - [HostAlertViewForNdsGroup](docs/HostAlertViewForNdsGroup.md)
 - [HostEventTypeViewForNdsGroup](docs/HostEventTypeViewForNdsGroup.md)
 - [HostEventTypeViewForNdsGroupAlertable](docs/HostEventTypeViewForNdsGroupAlertable.md)
 - [HostEventViewForNdsGroup](docs/HostEventViewForNdsGroup.md)
 - [HostMatcher](docs/HostMatcher.md)
 - [HostMatcherField](docs/HostMatcherField.md)
 - [HostMetricAlert](docs/HostMetricAlert.md)
 - [HostMetricAlertConfigViewForNdsGroup](docs/HostMetricAlertConfigViewForNdsGroup.md)
 - [HostMetricEvent](docs/HostMetricEvent.md)
 - [HostMetricEventType](docs/HostMetricEventType.md)
 - [HostMetricEventTypeViewAlertable](docs/HostMetricEventTypeViewAlertable.md)
 - [HostMetricThreshold](docs/HostMetricThreshold.md)
 - [HostMetricValue](docs/HostMetricValue.md)
 - [HostViewAtlas](docs/HostViewAtlas.md)
 - [IdentityProvider](docs/IdentityProvider.md)
 - [IdentityProviderUpdate](docs/IdentityProviderUpdate.md)
 - [IndexOptions](docs/IndexOptions.md)
 - [IndexRequest](docs/IndexRequest.md)
 - [IngestionPipeline](docs/IngestionPipeline.md)
 - [IngestionPipelineRun](docs/IngestionPipelineRun.md)
 - [IngestionSink](docs/IngestionSink.md)
 - [IngestionSource](docs/IngestionSource.md)
 - [InheritedRole](docs/InheritedRole.md)
 - [InstanceSize](docs/InstanceSize.md)
 - [Integration](docs/Integration.md)
 - [IntegrationViewForNdsGroup](docs/IntegrationViewForNdsGroup.md)
 - [Invoice](docs/Invoice.md)
 - [Key](docs/Key.md)
 - [Label](docs/Label.md)
 - [LegacyClusterDescription](docs/LegacyClusterDescription.md)
 - [LegacyReplicationSpec](docs/LegacyReplicationSpec.md)
 - [LessThanDaysThreshold](docs/LessThanDaysThreshold.md)
 - [LessThanTimeThreshold](docs/LessThanTimeThreshold.md)
 - [LessThanTimeThresholdAlertConfigViewForNdsGroup](docs/LessThanTimeThresholdAlertConfigViewForNdsGroup.md)
 - [Limit](docs/Limit.md)
 - [LineItem](docs/LineItem.md)
 - [Link](docs/Link.md)
 - [LinkAtlas](docs/LinkAtlas.md)
 - [ListPeeringConnections200Response](docs/ListPeeringConnections200Response.md)
 - [LiveMigrationRequest](docs/LiveMigrationRequest.md)
 - [LiveMigrationResponse](docs/LiveMigrationResponse.md)
 - [ManagedNamespace](docs/ManagedNamespace.md)
 - [ManagedNamespaces](docs/ManagedNamespaces.md)
 - [Matcher](docs/Matcher.md)
 - [MatcherField](docs/MatcherField.md)
 - [MatcherHostType](docs/MatcherHostType.md)
 - [Measurement](docs/Measurement.md)
 - [MeasurementViewAtlas](docs/MeasurementViewAtlas.md)
 - [MeasurementsGeneralViewAtlas](docs/MeasurementsGeneralViewAtlas.md)
 - [MeasurementsIndexes](docs/MeasurementsIndexes.md)
 - [MeasurementsNonIndex](docs/MeasurementsNonIndex.md)
 - [MetricDataPoint](docs/MetricDataPoint.md)
 - [MetricDataPointViewAtlas](docs/MetricDataPointViewAtlas.md)
 - [MicrosoftTeams](docs/MicrosoftTeams.md)
 - [MicrosoftTeamsNotification](docs/MicrosoftTeamsNotification.md)
 - [MongoDBAccessLogs](docs/MongoDBAccessLogs.md)
 - [MongoDBAccessLogsList](docs/MongoDBAccessLogsList.md)
 - [MonthlySchedule](docs/MonthlySchedule.md)
 - [NDSAuditTypeViewForNdsGroup](docs/NDSAuditTypeViewForNdsGroup.md)
 - [NDSAuditTypeViewForOrg](docs/NDSAuditTypeViewForOrg.md)
 - [NDSAuditViewForNdsGroup](docs/NDSAuditViewForNdsGroup.md)
 - [NDSAuditViewForOrg](docs/NDSAuditViewForOrg.md)
 - [NDSAutoScalingAuditTypeViewForNdsGroup](docs/NDSAutoScalingAuditTypeViewForNdsGroup.md)
 - [NDSAutoScalingAuditViewForNdsGroup](docs/NDSAutoScalingAuditViewForNdsGroup.md)
 - [NDSLDAP](docs/NDSLDAP.md)
 - [NDSLDAPVerifyConnectivityJobRequest](docs/NDSLDAPVerifyConnectivityJobRequest.md)
 - [NDSLDAPVerifyConnectivityJobRequestParams](docs/NDSLDAPVerifyConnectivityJobRequestParams.md)
 - [NDSLDAPVerifyConnectivityJobRequestValidation](docs/NDSLDAPVerifyConnectivityJobRequestValidation.md)
 - [NDSLabel](docs/NDSLabel.md)
 - [NDSNotification](docs/NDSNotification.md)
 - [NDSServerlessInstanceAudit](docs/NDSServerlessInstanceAudit.md)
 - [NDSServerlessInstanceAuditType](docs/NDSServerlessInstanceAuditType.md)
 - [NDSTenantEndpointAudit](docs/NDSTenantEndpointAudit.md)
 - [NDSTenantEndpointAuditType](docs/NDSTenantEndpointAuditType.md)
 - [NDSUserToDNMapping](docs/NDSUserToDNMapping.md)
 - [NDSX509UserAuthenticationAlertConfigViewForNdsGroup](docs/NDSX509UserAuthenticationAlertConfigViewForNdsGroup.md)
 - [NDSX509UserAuthenticationEventTypeViewAlertable](docs/NDSX509UserAuthenticationEventTypeViewAlertable.md)
 - [NamespaceObj](docs/NamespaceObj.md)
 - [Namespaces](docs/Namespaces.md)
 - [NetworkPermissionEntry](docs/NetworkPermissionEntry.md)
 - [NetworkPermissionEntryStatus](docs/NetworkPermissionEntryStatus.md)
 - [NewRelic](docs/NewRelic.md)
 - [Notification](docs/Notification.md)
 - [NotificationViewForNdsGroup](docs/NotificationViewForNdsGroup.md)
 - [NumberMetricAlert](docs/NumberMetricAlert.md)
 - [NumberMetricEvent](docs/NumberMetricEvent.md)
 - [NumberMetricThreshold](docs/NumberMetricThreshold.md)
 - [NumberMetricUnits](docs/NumberMetricUnits.md)
 - [NumberMetricValue](docs/NumberMetricValue.md)
 - [OnDemandCpsSnapshotSource](docs/OnDemandCpsSnapshotSource.md)
 - [OnlineArchive](docs/OnlineArchive.md)
 - [OnlineArchiveSchedule](docs/OnlineArchiveSchedule.md)
 - [Operator](docs/Operator.md)
 - [OpsGenie](docs/OpsGenie.md)
 - [OpsGenieNotification](docs/OpsGenieNotification.md)
 - [OrgEventTypeViewForOrg](docs/OrgEventTypeViewForOrg.md)
 - [OrgEventViewForOrg](docs/OrgEventViewForOrg.md)
 - [OrgFederationSettings](docs/OrgFederationSettings.md)
 - [OrgGroup](docs/OrgGroup.md)
 - [OrgNotification](docs/OrgNotification.md)
 - [OrgPaginatedEvent](docs/OrgPaginatedEvent.md)
 - [Organization](docs/Organization.md)
 - [OrganizationInvitation](docs/OrganizationInvitation.md)
 - [OrganizationInvitationRequest](docs/OrganizationInvitationRequest.md)
 - [OrganizationInvitationUpdateRequest](docs/OrganizationInvitationUpdateRequest.md)
 - [OrganizationSettings](docs/OrganizationSettings.md)
 - [PagerDuty](docs/PagerDuty.md)
 - [PagerDutyNotification](docs/PagerDutyNotification.md)
 - [PaginatedAWSPeerVpc](docs/PaginatedAWSPeerVpc.md)
 - [PaginatedAlert](docs/PaginatedAlert.md)
 - [PaginatedAlertConfig](docs/PaginatedAlertConfig.md)
 - [PaginatedApiApiUser](docs/PaginatedApiApiUser.md)
 - [PaginatedApiAppUser](docs/PaginatedApiAppUser.md)
 - [PaginatedApiAtlasCheckpoint](docs/PaginatedApiAtlasCheckpoint.md)
 - [PaginatedApiAtlasDatabaseUser](docs/PaginatedApiAtlasDatabaseUser.md)
 - [PaginatedApiAtlasDiskBackupExportJob](docs/PaginatedApiAtlasDiskBackupExportJob.md)
 - [PaginatedApiAtlasProviderRegions](docs/PaginatedApiAtlasProviderRegions.md)
 - [PaginatedApiAtlasServerlessBackupRestoreJob](docs/PaginatedApiAtlasServerlessBackupRestoreJob.md)
 - [PaginatedApiAtlasServerlessBackupSnapshot](docs/PaginatedApiAtlasServerlessBackupSnapshot.md)
 - [PaginatedApiInvoice](docs/PaginatedApiInvoice.md)
 - [PaginatedApiUserAccessList](docs/PaginatedApiUserAccessList.md)
 - [PaginatedAppUser](docs/PaginatedAppUser.md)
 - [PaginatedAtlasGroup](docs/PaginatedAtlasGroup.md)
 - [PaginatedAzurePeerNetwork](docs/PaginatedAzurePeerNetwork.md)
 - [PaginatedBackupSnapshot](docs/PaginatedBackupSnapshot.md)
 - [PaginatedBackupSnapshotExportBucket](docs/PaginatedBackupSnapshotExportBucket.md)
 - [PaginatedCloudBackupReplicaSet](docs/PaginatedCloudBackupReplicaSet.md)
 - [PaginatedCloudBackupRestoreJob](docs/PaginatedCloudBackupRestoreJob.md)
 - [PaginatedCloudBackupShardedClusterSnapshot](docs/PaginatedCloudBackupShardedClusterSnapshot.md)
 - [PaginatedCloudProviderContainer](docs/PaginatedCloudProviderContainer.md)
 - [PaginatedClusterDescriptionV15](docs/PaginatedClusterDescriptionV15.md)
 - [PaginatedDatabase](docs/PaginatedDatabase.md)
 - [PaginatedDiskPartition](docs/PaginatedDiskPartition.md)
 - [PaginatedGCPPeerVpc](docs/PaginatedGCPPeerVpc.md)
 - [PaginatedHostViewAtlas](docs/PaginatedHostViewAtlas.md)
 - [PaginatedLegacyCluster](docs/PaginatedLegacyCluster.md)
 - [PaginatedNetworkAccess](docs/PaginatedNetworkAccess.md)
 - [PaginatedOnlineArchive](docs/PaginatedOnlineArchive.md)
 - [PaginatedOrgGroup](docs/PaginatedOrgGroup.md)
 - [PaginatedOrganization](docs/PaginatedOrganization.md)
 - [PaginatedPipelineRun](docs/PaginatedPipelineRun.md)
 - [PaginatedPrivateLinkConnection](docs/PaginatedPrivateLinkConnection.md)
 - [PaginatedRestoreJob](docs/PaginatedRestoreJob.md)
 - [PaginatedServerlessInstanceDescription](docs/PaginatedServerlessInstanceDescription.md)
 - [PaginatedSnapshot](docs/PaginatedSnapshot.md)
 - [PaginatedTeam](docs/PaginatedTeam.md)
 - [PaginatedTeamRole](docs/PaginatedTeamRole.md)
 - [PaginatedTenantRestore](docs/PaginatedTenantRestore.md)
 - [PaginatedTenantSnapshot](docs/PaginatedTenantSnapshot.md)
 - [PaginatedUserCert](docs/PaginatedUserCert.md)
 - [PartitionField](docs/PartitionField.md)
 - [PartitionFieldView](docs/PartitionFieldView.md)
 - [Payment](docs/Payment.md)
 - [PemFileInfo](docs/PemFileInfo.md)
 - [PemFileInfoView](docs/PemFileInfoView.md)
 - [PerformanceAdvisorIndex](docs/PerformanceAdvisorIndex.md)
 - [PerformanceAdvisorOpStats](docs/PerformanceAdvisorOpStats.md)
 - [PerformanceAdvisorOperation](docs/PerformanceAdvisorOperation.md)
 - [PerformanceAdvisorResponse](docs/PerformanceAdvisorResponse.md)
 - [PerformanceAdvisorShape](docs/PerformanceAdvisorShape.md)
 - [PerformanceAdvisorSlowQuery](docs/PerformanceAdvisorSlowQuery.md)
 - [PerformanceAdvisorSlowQueryList](docs/PerformanceAdvisorSlowQueryList.md)
 - [PeriodicCpsSnapshotSource](docs/PeriodicCpsSnapshotSource.md)
 - [PipelineRunStats](docs/PipelineRunStats.md)
 - [Policy](docs/Policy.md)
 - [PolicyItem](docs/PolicyItem.md)
 - [PrivateIPMode](docs/PrivateIPMode.md)
 - [PrivateNetworkEndpointIdEntry](docs/PrivateNetworkEndpointIdEntry.md)
 - [ProjectSettingItem](docs/ProjectSettingItem.md)
 - [Prometheus](docs/Prometheus.md)
 - [ProviderRegions](docs/ProviderRegions.md)
 - [RPUMetricThreshold](docs/RPUMetricThreshold.md)
 - [Raw](docs/Raw.md)
 - [RawMetricAlert](docs/RawMetricAlert.md)
 - [RawMetricEvent](docs/RawMetricEvent.md)
 - [RawMetricThreshold](docs/RawMetricThreshold.md)
 - [RawMetricUnits](docs/RawMetricUnits.md)
 - [RawMetricValue](docs/RawMetricValue.md)
 - [Refund](docs/Refund.md)
 - [RegionConfig](docs/RegionConfig.md)
 - [RegionSpec](docs/RegionSpec.md)
 - [ReplicaSetAlertConfigViewForNdsGroup](docs/ReplicaSetAlertConfigViewForNdsGroup.md)
 - [ReplicaSetAlertViewForNdsGroup](docs/ReplicaSetAlertViewForNdsGroup.md)
 - [ReplicaSetEventTypeViewForNdsGroup](docs/ReplicaSetEventTypeViewForNdsGroup.md)
 - [ReplicaSetEventTypeViewForNdsGroupAlertable](docs/ReplicaSetEventTypeViewForNdsGroupAlertable.md)
 - [ReplicaSetEventTypeViewForNdsGroupAlertableNoThreshold](docs/ReplicaSetEventTypeViewForNdsGroupAlertableNoThreshold.md)
 - [ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold](docs/ReplicaSetEventTypeViewForNdsGroupAlertableWithThreshold.md)
 - [ReplicaSetMatcher](docs/ReplicaSetMatcher.md)
 - [ReplicaSetMatcherField](docs/ReplicaSetMatcherField.md)
 - [ReplicaSetThresholdAlertConfigViewForNdsGroup](docs/ReplicaSetThresholdAlertConfigViewForNdsGroup.md)
 - [ReplicationSpec](docs/ReplicationSpec.md)
 - [RestoreJob](docs/RestoreJob.md)
 - [RestoreJobDelivery](docs/RestoreJobDelivery.md)
 - [RestoreJobFileHash](docs/RestoreJobFileHash.md)
 - [Role](docs/Role.md)
 - [RoleAssignment](docs/RoleAssignment.md)
 - [RoleMapping](docs/RoleMapping.md)
 - [SMSNotification](docs/SMSNotification.md)
 - [SampleDatasetStatus](docs/SampleDatasetStatus.md)
 - [ServerlessAWSTenantEndpoint](docs/ServerlessAWSTenantEndpoint.md)
 - [ServerlessAWSTenantEndpointUpdate](docs/ServerlessAWSTenantEndpointUpdate.md)
 - [ServerlessAzureTenantEndpoint](docs/ServerlessAzureTenantEndpoint.md)
 - [ServerlessAzureTenantEndpointUpdate](docs/ServerlessAzureTenantEndpointUpdate.md)
 - [ServerlessBackupOptions](docs/ServerlessBackupOptions.md)
 - [ServerlessBackupRestoreJob](docs/ServerlessBackupRestoreJob.md)
 - [ServerlessBackupSnapshot](docs/ServerlessBackupSnapshot.md)
 - [ServerlessEventTypeViewAlertable](docs/ServerlessEventTypeViewAlertable.md)
 - [ServerlessInstanceDescription](docs/ServerlessInstanceDescription.md)
 - [ServerlessInstanceDescriptionConnectionStrings](docs/ServerlessInstanceDescriptionConnectionStrings.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint](docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint.md)
 - [ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint](docs/ServerlessInstanceDescriptionConnectionStringsPrivateEndpointEndpoint.md)
 - [ServerlessInstanceDescriptionCreate](docs/ServerlessInstanceDescriptionCreate.md)
 - [ServerlessInstanceDescriptionUpdate](docs/ServerlessInstanceDescriptionUpdate.md)
 - [ServerlessMetricAlertConfigViewForNdsGroup](docs/ServerlessMetricAlertConfigViewForNdsGroup.md)
 - [ServerlessMetricThreshold](docs/ServerlessMetricThreshold.md)
 - [ServerlessMetricUnits](docs/ServerlessMetricUnits.md)
 - [ServerlessProviderSettings](docs/ServerlessProviderSettings.md)
 - [ServerlessTenantEndpoint](docs/ServerlessTenantEndpoint.md)
 - [ServerlessTenantEndpointCreate](docs/ServerlessTenantEndpointCreate.md)
 - [ServerlessTenantEndpointUpdate](docs/ServerlessTenantEndpointUpdate.md)
 - [Slack](docs/Slack.md)
 - [SlackNotification](docs/SlackNotification.md)
 - [Snapshot](docs/Snapshot.md)
 - [SnapshotPart](docs/SnapshotPart.md)
 - [SnapshotRetention](docs/SnapshotRetention.md)
 - [SnapshotSchedule](docs/SnapshotSchedule.md)
 - [Source](docs/Source.md)
 - [SummaryNotification](docs/SummaryNotification.md)
 - [SynonymSource](docs/SynonymSource.md)
 - [SystemStatus](docs/SystemStatus.md)
 - [TargetOrg](docs/TargetOrg.md)
 - [TargetOrgRequest](docs/TargetOrgRequest.md)
 - [Team](docs/Team.md)
 - [TeamEvent](docs/TeamEvent.md)
 - [TeamEventType](docs/TeamEventType.md)
 - [TeamEventTypeViewForNdsGroup](docs/TeamEventTypeViewForNdsGroup.md)
 - [TeamEventViewForNdsGroup](docs/TeamEventViewForNdsGroup.md)
 - [TeamNotification](docs/TeamNotification.md)
 - [TeamResponse](docs/TeamResponse.md)
 - [TeamRole](docs/TeamRole.md)
 - [TenantHardwareSpec](docs/TenantHardwareSpec.md)
 - [TenantRegionConfig](docs/TenantRegionConfig.md)
 - [TenantRestore](docs/TenantRestore.md)
 - [TenantSnapshot](docs/TenantSnapshot.md)
 - [ThresholdViewInteger](docs/ThresholdViewInteger.md)
 - [TimeMetricAlert](docs/TimeMetricAlert.md)
 - [TimeMetricEvent](docs/TimeMetricEvent.md)
 - [TimeMetricThreshold](docs/TimeMetricThreshold.md)
 - [TimeMetricUnits](docs/TimeMetricUnits.md)
 - [TimeMetricValue](docs/TimeMetricValue.md)
 - [Toggle](docs/Toggle.md)
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
 - [User](docs/User.md)
 - [UserAccessList](docs/UserAccessList.md)
 - [UserCert](docs/UserCert.md)
 - [UserEventTypeViewForNdsGroup](docs/UserEventTypeViewForNdsGroup.md)
 - [UserEventTypeViewForOrg](docs/UserEventTypeViewForOrg.md)
 - [UserEventViewForNdsGroup](docs/UserEventViewForNdsGroup.md)
 - [UserEventViewForOrg](docs/UserEventViewForOrg.md)
 - [UserNotification](docs/UserNotification.md)
 - [UserRoleAssignment](docs/UserRoleAssignment.md)
 - [UserScope](docs/UserScope.md)
 - [UserSecurity](docs/UserSecurity.md)
 - [Validation](docs/Validation.md)
 - [VictorOps](docs/VictorOps.md)
 - [VictorOpsNotification](docs/VictorOpsNotification.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookNotification](docs/WebhookNotification.md)
 - [WeeklySchedule](docs/WeeklySchedule.md)
 - [X509Certificate](docs/X509Certificate.md)
 - [X509CertificateView](docs/X509CertificateView.md)
 - [ZoneMapping](docs/ZoneMapping.md)




