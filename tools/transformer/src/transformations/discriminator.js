const {
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getAllObjects,
  getObjectNameFromReferenceString,
} = require("../engine/readers");

const isModelIgnored = require("../engine/ignoreModel");
const ignoreModels = require("../discriminator.ignore.json").ignoreModels;
const { filterObjectProperties } = require("../engine/transformers");

/**
 * Transforms provided API JSON file using discriminatorOneOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @returns OpenAPI JSON File
 */
function applyDiscriminatorTransformations(api) {
  const transformations = getAllObjects(api, isTransformable);

  console.info(
    "# Discriminator transformations: ",
    transformations.map((e) => e.path)
  );

  for (let { path } of transformations) {
    transformDiscriminatorOneOf(path, api);
  }
  return api;
}

// Adds oneOf when discriminator is present
function transformDiscriminatorOneOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!isTransformable(parentObject)) {
    throw new Error(`Invalid object for AllOf Transformation: ${parentName}`);
  }

  if (!parentObject.oneOf && parentObject.discriminator) {
    if (!parentObject.discriminator.mapping) {
      throw new Error(`OpenAPI object schema contains discriminator but missing oneOf or discriminator mapping.
      Please consider adding oneOf or discriminator mapping section to object: ${parentName}`);
    }
    console.info(
      `Setting oneOf based on discriminator for allOf transformation in ${parentName}`
    );
    // Ignore objects that point to themselves
    oneOfReferences = Object.values(parentObject.discriminator.mapping);
    // Remove duplicates in referneces
    oneOfReferences = [...new Set(oneOfReferences)];

    for (referenceObj of oneOfReferences) {
      if (getObjectNameFromReferenceString(referenceObj) == parentName) {
        if (isModelIgnored(parentName, ignoreModels)) {
          console.warn("Ignored discriminator reference for: " + parentName);
          return;
        }
        throw new Error(`${parentName}.discriminator.mapping contains $ref to itself.
        Please consider replacing reference with new type.`);
      }
    }

    parentObject.oneOf = oneOfReferences.map((obj) => {
      return { $ref: obj };
    });
  }
}

function isTransformable(obj) {
  return !!obj.discriminator;
}

function getObjectProperties(obj) {
  const exclusionList = ["oneOf", "discriminator"];

  return filterObjectProperties(obj, (key, _v) => {
    return !(key in exclusionList);
  });
}

module.exports = {
  applyDiscriminatorTransformations,
};
