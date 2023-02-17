const {
  getObjectFromReference,
  removeParentFromAllOf,
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getObjectNameFromReference,
  detectDuplicates,
  getAllObjectsWithProperty,
  getAllObjects,
  filterObjectProperties,
  flattenAllOfObject,
} = require('./utils');

// Name of the key (OpenAPI extension) to trigger transformation
const extensionKey = 'x-xgen-go-transform';
const extensionAllOfValue = 'merge-allOf';
const extensionOneOfValue = 'merge-oneOf';

/**
 * Transforms provided API JSON file using allOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} allOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyAllOfTransformations(api) {
  const allOfTransformations = getAllObjects(api, isAllOfTransformable);

  console.info(
    '# AllOf transformations: ',
    allOfTransformations.map((e) => e.path)
  );

  for (let { path, obj } of allOfTransformations) {
    if (!obj || !path) {
      throw new Error("Invalid transformation object");
    }
    transformAllOf(path, api);
  }
  return api;
}

/**
 * Transforms provided API JSON file using oneOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} oneOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyOneOfTransformations(api) {
  const oneOfTransformations = getAllObjectsWithProperty(api, extensionKey, (_k, v) => v === extensionOneOfValue);

  console.info(
    '# OneOf transformations: ',
    oneOfTransformations.map((e) => e.path)
  );

  for (let { path, obj } of oneOfTransformations) {
    if (!obj) {
      throw new Error("Invalid transformation object");
    }
    transformOneOf(path, api);
  }
  return api;
}

/**
 *  Transforms provided API JSON file by removing prefix and suffix from the
 *  API models
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} prefix model prefix or empty string if only suffix is needed.
 * @param {*} suffix model suffix or empty string if only prefix is needed.
 * @returns OpenAPI JSON File
 */
function applyModelNameTransformations(api, prefix, suffix) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error();
  }

  for (const schemaKey of Object.keys(api.components.schemas)) {
    if (schemaKey.startsWith(prefix) && schemaKey.endsWith(suffix)) {
      const schemaTransformedName = schemaKey
        .replace(prefix, "")
        .replace(suffix, "");
      if (api.components.schemas[schemaTransformedName]) {
        throw new Error(
          "components.schemas already contain key",
          schemaTransformedName
        );
      }
      api.components.schemas[schemaTransformedName] =
        api.components.schemas[schemaKey];
      delete api.components.schemas[schemaKey];
    }
  }

  let regexp = new RegExp(`"#/components/schemas/${prefix}(.*)${suffix}"`, "g");
  let apiString = JSON.stringify(api, undefined, 2);
  let apiToReplace = apiString.replace(regexp, '"#/components/schemas/$1"');
  api = JSON.parse(apiToReplace);
  return api;
}

function transformAllOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!(parentObject && parentObject.oneOf)) {
    throw new Error(`Invalid object for AllOf Transformation: ${parentName}`);
  }

  if (!parentObject.properties) {
    return console.error(
      `AllOf: Transformation cannot be performed. ${parentName}.properties is empty`
    );
  }

  const expandedParent = getObjectProperties(parentObject);

  // Remove invalid fields
  delete parentObject.properties;
  delete parentObject.required;

  for (let obj of parentObject.oneOf) {
    const childObject = getObjectFromReference(obj, api);
    const childName = getObjectNameFromReference(obj);

    if (removeParentFromAllOf(childObject, parentObject, api)) {
      console.debug(`AllOf: Moving ${parentName} (parent) properties into ${childName} (child) properties`);
      if (!childObject.allOf) {
        childObject.allOf = {};
      }
      // Expand parent in child allOf
      childObject.allOf.push(expandedParent);
      // Flatten child allOf
      flattenAllOfObject(childObject);
    }
  }
}

function transformOneOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);

  if (!(parentObject && parentObject.oneOf)) {
    throw new Error(`Invalid object for OneOf Transformation: ${objectPath}`);
  }

  for (const objRef of parentObject.oneOf) {
    const oneOfObject = getObjectFromReference(objRef, api);
    if (!oneOfObject) {
      const errorData = JSON.stringify(oneOfObject);
      throw new Error(`Invalid or missing reference: ${errorData}`);
    }

    if (oneOfObject.properties && oneOfObject.enum) {
      const errorData = JSON.stringify(oneOfObject);
      throw new Error(
        `Object cannot contain both enum and properties: ${errorData}`
      );
    }

    // Handle properties
    if (oneOfObject.properties) {
      const propertiesOneOf = JSON.parse(
        JSON.stringify(oneOfObject.properties)
      );
      console.debug(
        `${oneOfObject.title}: moving child properties into parent`
      );
      const duplicates = detectDuplicates([
        parentObject.properties,
        propertiesOneOf,
      ]);
      if (duplicates.length > 0) {
        const duplicatesSource = oneOfObject.title || "";
        console.warn(
          `## ${duplicatesSource} - Detected properties that would be overriden: ${duplicates}\n`
        );
      }
      parentObject.properties = {
        ...parentObject.properties,
        ...propertiesOneOf,
      };
    }

    if (oneOfObject.enum) {
      console.debug(
        `${oneOfObject.title}: moving child enum values into parent`
      );
      if (!parentObject.enum) {
        parentObject.enum = [];
      }
      parentObject.enum = parentObject.enum.concat(oneOfObject.enum);
      parentObject.enum = [...new Set(parentObject.enum)];
      // Requires all enums to be same type
      parentObject.type = oneOfObject.type;
    }
  }

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.oneOf;
}

function isAllOfTransformable(obj) {
  return obj.properties && obj.oneOf;
}

function getObjectProperties(obj) {
  const exclusionList = ['oneOf', 'discriminator'];

  return filterObjectProperties(obj, (key, _v) => {
    return !(key in exclusionList);
  });
}

module.exports = {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  transformOneOf,
  transformAllOf
};
