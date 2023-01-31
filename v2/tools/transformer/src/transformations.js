const {
  getObjectFromReference,
  removeParentFromAllOf,
  getNameFromYamlPath,
  getObjectNameFromReference,
  detectDuplicates,
} = require('./utils');

/**
 * Transforms provided API JSON file using allOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} allOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyAllOfTransformations(api, allOfTransformations) {
  for (let { path, obj } of allOfTransformations) {
    if (!obj || !path) {
      throw new Error('Invalid transformation object');
    }
    const name = getNameFromYamlPath(path);
    transformAllOf(name, obj, api);
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
function applyOneOfTransformations(api, oneOfTransformations) {
  for (let { obj } of oneOfTransformations) {
    if (!obj) {
      throw new Error('Invalid transformation object');
    }
    transformOneOf(obj, api);
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
      const schemaTransformedName = schemaKey.replace(prefix, '').replace(suffix, '');
      if (api.components.schemas[schemaTransformedName]) {
        throw new Error('components.schemas already contain key', schemaTransformedName);
      }
      api.components.schemas[schemaTransformedName] = api.components.schemas[schemaKey];
      delete api.components.schemas[schemaKey];
    }
  }

  let regexp = new RegExp(`"#/components/schemas/${prefix}(.*)${suffix}"`, 'g');
  let apiString = JSON.stringify(api, undefined, 2);
  let apiToReplace = apiString.replace(regexp, '"#/components/schemas/$1"');
  api = JSON.parse(apiToReplace);
  return api;
}

function transformAllOf(objectName, parentObject, api) {
  if (!(parentObject && parentObject.oneOf)) {
    const errorData = JSON.stringify(parentObject);
    throw new Error(`Invalid object for AllOf Transformation: ${errorData}`);
  }

  if (!parentObject.properties) {
    return console.error(`Transformation cannot be performed. ${objectName}.properties is empty`);
  }

  const parentProperties = parentObject.properties;
  delete parentObject.properties;

  for (let obj of parentObject.oneOf) {
    const child = getObjectFromReference(obj, api);
    if (!child.allOf) {
      const errorData = JSON.stringify(child);
      throw new Error(`Invalid child for AllOf Transformation: ${errorData}`);
    }

    removeParentFromAllOf(child, objectName);
    if (child.allOf.length !== 1) {
      const errorData = JSON.stringify(child);
      throw new Error(`Child contains more than single parent: ${errorData}`);
    }
    if (!child.properties) {
      child.properties = {};
    }
    const childName = getObjectNameFromReference(obj);
    if (child.allOf[0].properties) {
      console.debug(`${objectName}: moving parent properties into ${childName} properties`);
      var propertiesParent = JSON.parse(JSON.stringify(parentProperties));
      var childProperties = child.properties;
      var propertiesAllOfChild = child.allOf[0].properties;
      var duplicates = detectDuplicates([propertiesParent, childProperties, propertiesAllOfChild]);
      if (duplicates.length > 0) {
        console.warn(`${childName}: detected properties that would be overriden: ${duplicates}`);
      }
      child.properties = { ...propertiesParent, ...childProperties, ...propertiesAllOfChild };
    }
    delete child.allOf;
  }
}

function transformOneOf(parentObject, api) {
  if (!(parentObject && parentObject.oneOf)) {
    const errorObj = JSON.stringify(parentObject);
    throw new Error(`Invalid object for OneOf Transformation ${errorObj}`);
  }

  for (const objRef of parentObject.oneOf) {
    const oneOfObject = getObjectFromReference(objRef, api);
    if (!oneOfObject) {
      const errorData = JSON.stringify(oneOfObject);
      throw new Error(`Invalid or missing reference: ${errorData}`);
    }

    if (oneOfObject.properties && oneOfObject.enum) {
      const errorData = JSON.stringify(oneOfObject);
      throw new Error(`Object cannot contain both enum and properties: ${errorData}`);
    }

    // Handle properties
    if (oneOfObject.properties) {
      const propertiesOneOf = JSON.parse(JSON.stringify(oneOfObject.properties));
      console.debug(`${oneOfObject.title}: moving child properties into parent`);
      const duplicates = detectDuplicates([parentObject.properties, propertiesOneOf]);
      if (duplicates.length > 0) {
        const duplicatesSource = oneOfObject.title || '';
        console.warn(`## ${duplicatesSource} - Detected properties that would be overriden: ${duplicates}\n`);
      }
      parentObject.properties = { ...parentObject.properties, ...propertiesOneOf };
    }

    if (oneOfObject.enum) {
      console.debug(`${oneOfObject.title}: moving child enum values into parent`);
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

module.exports = {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations
};
