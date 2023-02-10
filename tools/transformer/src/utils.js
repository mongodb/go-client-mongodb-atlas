/**
 * Fetch all children JS objects and their paths that have a property which satisfies a set of conditions
 * @param {Object} apiObject - object to be searched
 * @param {String} key - the property that is searched
 * @param {(key, value) => boolean} [predicate] - validation function for the property
 * @returns {{ path: String, obj: Object }[]}
 */
function getAllObjectsWitProperty(
  apiObject,
  key,
  predicate = (_k, _v) => true
) {
  const objs = [];

  // Add root object at the top of the recursion stack
  const recursionStack = [{ path: "", obj: apiObject }];

  while (recursionStack.length > 0) {
    const { path, obj: currentObj } = recursionStack.pop();

    // Check validation function for the property on current object
    if (key in currentObj && predicate(key, currentObj[key])) {
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

function getObjectFromReference(obj, api) {
  if (obj.type || obj.properties) {
    // We dealing with object
    return obj;
  }

  if (obj && obj["$ref"]) {
    const referenceName = getObjectNameFromReference(obj);
    return api.components.schemas[referenceName];
  }
}

function getObjectNameFromReference(obj) {
  if (obj && obj["$ref"]) {
    return obj["$ref"].replace("#/components/schemas/", "");
  }
}

function detectDuplicates(objArray) {
  const allKeys = new Set();
  const duplicates = [];
  for (const obj of objArray) {
    if (obj) {
      for (const key of Object.keys(obj)) {
        if (allKeys.has(key)) {
          duplicates.push(key);
        } else {
          allKeys.add(key);
        }
      }
    }
  }
  return duplicates;
}

function removeParentFromAllOf(child, parentName) {
  child.allOf = child.allOf.filter((childAllOfItem) => {
    if (childAllOfItem["$ref"]) {
      if (
        childAllOfItem["$ref"].endsWith("#/components/schemas/" + parentName)
      ) {
        return false;
      }
    }
    return true;
  });
}

function getNameFromYamlPath(path) {
  return path.split(".").pop();
}

module.exports = {
  getNameFromYamlPath,
  getObjectFromReference,
  getAllObjectsWitProperty,
  removeParentFromAllOf,
  getObjectNameFromReference,
  detectDuplicates,
};
