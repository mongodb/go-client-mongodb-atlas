// Results of the transformations
module.exports = {
  ChildAllOf: `
{
  "properties": {
    "analyticsAutoScaling": {
      "$ref": "#/components/schemas/ApiAtlasAutoScalingV15View",
    },
    "analyticsSpecs": {
      "$ref": "#/components/schemas/ApiAtlasDedicatedHardwareSpecView",
    },
    "autoScaling": {
      "$ref": "#/components/schemas/ApiAtlasAutoScalingV15View",
    },
    "electableSpecs": {
      "$ref": "#/components/schemas/ApiAtlasHardwareSpecView",
    },
    "priority": {
      "format": "int32",
      "maximum": 7,
      "minimum": 0,
      "type": "integer",
    },
    "providerName": {
      "enum": [
        "AWS",
        "AZURE",
        "GCP",
        "TENANT",
      ],
      "type": "string",
    },
    "readOnlySpecs": {
      "$ref": "#/components/schemas/ApiAtlasDedicatedHardwareSpecView",
    },
    "regionName": {
      "oneOf": [
        {
          "enum": [
            "US_GOV_WEST_1",
            "US_GOV_EAST_1",
            "US_EAST_1",
            "US_EAST_2",
            "US_WEST_1",
            "US_WEST_2",
          ],
          "title": "AWS Regions",
          "type": "string",
        },
        {
          "enum": [
            "US_CENTRAL",
            "US_EAST",
            "US_EAST_2",
          ],
          "title": "Azure Regions",
          "type": "string",
        },
        {
          "enum": [
            "EASTERN_US",
            "US_EAST_4",
            "US_WEST_2",
            "US_WEST_3",
            "US_WEST_4",
          ],
          "title": "GCP Regions",
          "type": "string",
        },
      ],
      "type": "object",
    },
  },
  "required": [],
  "title": "AWS Regional Replication Specifications",
  "type": "object",
}
`,
  ParentAllOf: `
{
  "discriminator": {
    "mapping": {
      "AWS": "#/components/schemas/ApiAtlasAWSRegionConfigView",
      "AZURE": "#/components/schemas/ApiAtlasAzureRegionConfigView",
      "GCP": "#/components/schemas/ApiAtlasGCPRegionConfigView",
      "TENANT": "#/components/schemas/ApiAtlasTenantRegionConfigView",
    },
    "propertyName": "providerName",
  },
  "oneOf": [
    {
      "$ref": "#/components/schemas/ApiAtlasAWSRegionConfigView",
    },
    {
      "$ref": "#/components/schemas/ApiAtlasAzureRegionConfigView",
    },
    {
      "$ref": "#/components/schemas/ApiAtlasGCPRegionConfigView",
    },
    {
      "$ref": "#/components/schemas/ApiAtlasTenantRegionConfigView",
    },
  ],
  "title": "Cloud Service Provider Settings for Multi-Cloud Clusters",
  "type": "object",
}
`,
  EnumOneOf: `
{
  "enum": [
    "US_GOV_WEST_1",
    "US_GOV_EAST_1",
    "US_EAST_1",
    "US_EAST_2",
    "US_WEST_1",
    "US_WEST_2",
    "US_CENTRAL",
    "US_EAST",
    "EASTERN_US",
    "US_EAST_4",
    "US_WEST_3",
    "US_WEST_4",
  ],
  "type": "string",
}
`,
  PropertiesOneOf: `
{
  "diskIOPS": {
    "format": "int32",
    "type": "integer",
  },
  "ebsVolumeType": {
    "default": "STANDARD",
    "enum": [
      "STANDARD",
      "PROVISIONED",
    ],
    "type": "string",
  },
  "instanceSize": {
    "enum": [
      "M0",
      "M2",
      "M5",
    ],
    "title": "Tenant Instance Sizes",
    "type": "string",
  },
  "nodeCount": {
    "format": "int32",
    "type": "integer",
  },
}
`,
};
