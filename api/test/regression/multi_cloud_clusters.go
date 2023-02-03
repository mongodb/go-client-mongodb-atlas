package regression

import (
	"github.com/openlyinc/pointy"
	mongodbatlas "go.mongodb.org/atlas/api/v1"
)

func TestMultiCloudClustersModelRegression() {
	_true := true
	_string := pointy.String("")
	_int := pointy.Int32(0)
	labelView := mongodbatlas.NDSLabel{
		Key:   _string,
		Value: _string,
	}

	instanceSize := mongodbatlas.INSTANCESIZE_M10

	autoScale := &mongodbatlas.AutoScalingV15{
		Compute: &mongodbatlas.ComputeAutoScalingV15{
			Enabled:          &_true,
			MaxInstanceSize:  &instanceSize,
			MinInstanceSize:  &instanceSize,
			ScaleDownEnabled: &_true,
		},
		DiskGB: &mongodbatlas.DiskGBAutoScaling{
			Enabled: &_true,
		},
	}

	hardwareSpec := &mongodbatlas.DedicatedHardwareSpec{
		NodeCount:     new(int32),
		DiskIOPS:      new(int32),
		EbsVolumeType: new(string),
		InstanceSize:  _string,
	}

	regionAWS := mongodbatlas.RegionConfig{
		AWSRegionConfig: &mongodbatlas.AWSRegionConfig{
			ElectableSpecs: &mongodbatlas.HardwareSpec{
				DiskIOPS:      _int,
				EbsVolumeType: _string,
				InstanceSize:  _string,
				NodeCount:     _int,
			},
			Priority:             new(int32),
			ProviderName:         new(string),
			RegionName:           new(string),
			AnalyticsAutoScaling: autoScale,
			AnalyticsSpecs: &mongodbatlas.DedicatedHardwareSpec{
				NodeCount:     new(int32),
				DiskIOPS:      new(int32),
				EbsVolumeType: new(string),
				InstanceSize:  _string,
			},
			AutoScaling:   autoScale,
			ReadOnlySpecs: hardwareSpec,
		},
	}

	regionGcp := mongodbatlas.RegionConfig{
		GCPRegionConfig: &mongodbatlas.GCPRegionConfig{
			ElectableSpecs: &mongodbatlas.HardwareSpec{
				DiskIOPS:      _int,
				EbsVolumeType: _string,
				InstanceSize:  _string,
				NodeCount:     _int,
			},
			Priority:     new(int32),
			ProviderName: new(string),
			RegionName:   new(string),
		},
	}

	regionAzure := mongodbatlas.RegionConfig{
		AzureRegionConfig: &mongodbatlas.AzureRegionConfig{
			ElectableSpecs: &mongodbatlas.HardwareSpec{
				DiskIOPS:      _int,
				EbsVolumeType: _string,
				InstanceSize:  _string,
				NodeCount:     _int,
			},
			Priority:     new(int32),
			ProviderName: new(string),
			RegionName:   new(string),
		},
	}

	regionDedicated := mongodbatlas.RegionConfig{
		TenantRegionConfig: &mongodbatlas.TenantRegionConfig{
			ElectableSpecs: &mongodbatlas.HardwareSpec{
				DiskIOPS:      _int,
				EbsVolumeType: _string,
				InstanceSize:  _string,
				NodeCount:     _int,
			},
			Priority:            new(int32),
			ProviderName:        new(string),
			RegionName:          new(string),
			BackingProviderName: new(string),
		},
	}

	replicationSpec := mongodbatlas.ReplicationSpec{
		Id:            _string,
		NumShards:     _int,
		RegionConfigs: []mongodbatlas.RegionConfig{regionAWS, regionGcp, regionAzure, regionDedicated},
		ZoneName:      _string,
	}

	replicationSpec2 := mongodbatlas.ReplicationSpec{
		Id:            _string,
		NumShards:     _int,
		RegionConfigs: []mongodbatlas.RegionConfig{regionGcp},
		ZoneName:      _string,
	}

	_ = mongodbatlas.ClusterDescriptionV15{
		BackupEnabled: &_true,
		BiConnector:   &mongodbatlas.BiConnector{Enabled: &_true},
		ClusterType:   nil,
		ConnectionStrings: &mongodbatlas.ClusterDescriptionConnectionStrings{
			AwsPrivateLink:    &map[string]string{},
			AwsPrivateLinkSrv: &map[string]string{},
			Private:           _string,
			PrivateEndpoint:   []mongodbatlas.ClusterDescriptionConnectionStringsPrivateEndpoint{},
			PrivateSrv:        _string,
			Standard:          _string,
			StandardSrv:       _string,
		},
		// TODO this should be removed as they are part of the response body.
		// See: https://jira.mongodb.org/browse/CLOUDP-151153
		CreateDate: nil,
		Paused:     nil,
		Links:      nil,
		Id:         nil,
		StateName:  _string,

		DiskSizeGB:                   nil,
		EncryptionAtRestProvider:     nil,
		Labels:                       []mongodbatlas.NDSLabel{labelView},
		MongoDBMajorVersion:          nil,
		MongoDBVersion:               nil,
		Name:                         _string,
		ReplicationSpecs:             []mongodbatlas.ReplicationSpec{replicationSpec, replicationSpec2},
		RootCertType:                 nil,
		TerminationProtectionEnabled: &_true,
		VersionReleaseSystem:         nil,
	}
}
