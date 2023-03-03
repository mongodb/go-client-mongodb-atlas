const {
  getObjectFromReference,
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getAllObjects,
} = require("../engine/readers");
const { detectDuplicates } = require("../engine/transformers");

// Name of the key (OpenAPI extension) to trigger transformation
const extensionKey = "x-xgen-go-transform";
const extensionOneOfValue = "merge-oneOf";

/**
 * Transforms provided API JSON file using oneOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} oneOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyOneOfTransformations(api) {
  const oneOfTransformations = getAllObjects(api, (obj) => {
    return canApplyOneOfTransformation(obj, api);
  });

  console.info(
    "# OneOf transformations: ",
    oneOfTransformations.map((e) => e.path)
  );

  for (let { path } of oneOfTransformations) {
    transformOneOf(path, api);
  }
  return api;
}

// Moves all the properties/enum values of the children into the parent
function transformOneOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!parentObject.oneOf) {
    throw new Error(`Invalid object for OneOf Transformation: ${parentName}`);
  }

  // Expand references
  const childObjects = parentObject.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );
  const isEnum = childObjects.reduce(
    (isEnum, childObject) => isEnum && childObject.enum,
    true
  );

  if (isEnum) {
    transformOneOfEnum(parentObject, api);
  } else {
    transformOneOfProperties(parentObject, api);
  }
}

// Moves all the enum values of the children into the parent
function transformOneOfEnum(parentObject, api) {
  const childObjects = parentObject.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );

  if (!parentObject.enum) {
    parentObject.enum = [];
  }
  parentObject.enum = new Set(parentObject.enum);

  for (let childObject of childObjects) {
    console.debug(`${childObject.title}: moving child enum values into parent`);
    childObject.enum.forEach((enumValue) => parentObject.enum.add(enumValue));
    // Requires all enums to be same type
    parentObject.type = childObject.type;
  }

  parentObject.enum = [...parentObject.enum];

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.oneOf;
}

// Moves all the propertis of the children into the parent
function transformOneOfProperties(parentObject, api) {
  const childObjects = parentObject.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );

  for (let childObject of childObjects) {
    const childProperties = JSON.parse(JSON.stringify(childObject.properties));
    console.debug(`${childObject.title}: moving child properties into parent`);
    const duplicates = detectDuplicates([
      parentObject.properties,
      childProperties,
    ]);
    if (duplicates.length > 0) {
      const duplicatesSource = childObject.title || "";
      console.info(
        `## ${duplicatesSource} - Detected properties that would be overriden: ${duplicates}\n`
      );
    }
    parentObject.properties = {
      ...parentObject.properties,
      ...childProperties,
    };
  }

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.oneOf;
}

function canApplyOneOfTransformation(obj, api) {
  if (!obj.oneOf) {
    return false;
  }

  if (obj[extensionKey] === extensionOneOfValue) {
    return true;
  }

  const children = obj.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );
  const isEnum = children.reduce((isEnum, child) => isEnum && child.enum, true);

  return isEnum;
}

module.exports = {
  applyOneOfTransformations,
  transformOneOf,
};
