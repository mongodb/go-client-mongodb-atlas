/**
 * Fetch all children JS objects and their paths that have a property which satisfies a set of conditions
 *
 * @param {Object} apiObject - object to be searched
 * @param {String} key - the property that is searched
 * @param {(key, value) => boolean} [predicate] - validation function for the property
 * @returns {{ path: String, obj: Object }[]}
 */
function getAllObjectsWithProperty(
  apiObject,
  key,
  predicate = (_k, _v) => true
) {
  return getAllObjects(
    apiObject,
    (obj) => key in obj && predicate(key, obj[key])
  );
}

/**
 * Fetch all transformation objects
 *
 * @param {*} object
 * @param {*} filter
 * @returns
 */
function getAllObjects(object, filter = (_obj) => true) {
  const objs = [];

  // Add root object at the top of the recursion stack
  const recursionStack = [{ path: "", obj: object }];

  while (recursionStack.length > 0) {
    const { path, obj: currentObj } = recursionStack.pop();

    // Check validation function for the property on current object
    if (filter(currentObj)) {
      objs.push({
        path: path,
        obj: currentObj,
      });
    }

    // If current object is an array, add all its elements to the recursion stack
    if (Array.isArray(currentObj)) {
      for (let i = 0; i < currentObj.length; i++) {
        if (typeof currentObj[i] === "object" && currentObj[i]) {
          recursionStack.push({
            path: `${path}.${i}`,
            obj: currentObj[i],
          });
        }
      }
    }
    // Add all properties of the object to the recurssion stack
    else {
      for (let key of Object.keys(currentObj)) {
        if (typeof currentObj[key] === "object" && currentObj[key]) {
          recursionStack.push({
            path: `${path}.${key}`,
            obj: currentObj[key],
          });
        }
      }
    }
  }

  return objs;
}

function filterObjectProperties(object, filter = (_k, _v) => true) {
  const filteredObj = Object.keys(object)
    .filter((key) => filter(key, object[key]))
    .reduce((aggregationObj, key) => {
      aggregationObj[key] = object[key];
      return aggregationObj;
    }, {});

  return filteredObj;
}

function getObjectProperties(obj) {
  const exclusionList = ["oneOf", "discriminator"];

  return filterObjectProperties(obj, (key, _v) => {
    return !(key in exclusionList);
  });
}

function getObjectFromReference(obj, api) {
  if (obj.type || obj.properties) {
    // We dealing with object
    return obj;
  }

  if (obj && obj["$ref"]) {
    const referenceName = getObjectNameFromReference(obj);
    if (!api.components.schemas[referenceName]) {
      throw new Error(`Invalid or missing reference: ${referenceName}`);
    }
    return api.components.schemas[referenceName];
  }
}

function getObjectNameFromReference(obj) {
  if (obj && obj["$ref"]) {
    return obj["$ref"].replace("#/components/schemas/", "");
  }
}

function getObjectNameFromReferenceString(objString) {
  return objString.replace("#/components/schemas/", "");
}

function getNameFromYamlPath(path) {
  return path.split(".").pop();
}

function getObjectFromYamlPath(path, obj) {
  const propertiesStack = path.split(".").reverse();
  let currObj = obj;

  propertiesStack.pop();
  while (propertiesStack.length > 0) {
    const property = propertiesStack.pop();

    if (Array.isArray(currObj)) {
      const index = parseInt(property);
      if (index < 0 || index >= currObj.length) {
        throw new Error(`Invalid path: ${path}`);
      }

      currObj = currObj[parseInt(property)];
    } else if (property in currObj) {
      currObj = currObj[property];
    } else {
      throw new Error(`Invalid path: ${path}`);
    }
  }

  return currObj;
}

module.exports = {
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getObjectFromReference,
  getAllObjectsWithProperty,
  getAllObjects,
  getObjectNameFromReference,
  getObjectNameFromReferenceString,
};
