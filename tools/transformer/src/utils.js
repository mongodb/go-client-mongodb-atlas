/**
 * Fetch all children JS objects and their paths that have a property which satisfies a set of conditions
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

function getObjectFromReference(obj, api) {
  if (obj.type || obj.properties) {
    // We dealing with object
    return obj;
  }

  if (obj && obj["$ref"]) {
    const referenceName = getObjectNameFromReference(obj);
    if(!api.components.schemas[referenceName]) {
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

function expandReference(obj, apiObject) {
  if (!obj || !obj["$ref"]) {
    return obj;
  }

  const dereferencedObj = getObjectFromReference(obj, apiObject);
  delete obj["$ref"];

  Object.keys(dereferencedObj).forEach((key) => {
    obj[key] = dereferencedObj[key];
  });

  return obj;
}

function flattenAllOfObject(obj, apiObject) {
  if (!obj.allOf) {
    return obj;
  }

  if (!obj.properties) {
    obj.properties = {};
  }

  if (!obj.required) {
    obj.required = new Set();
  } else {
    obj.required = new Set(obj.required);
  }

  for (let parent of obj.allOf) {
    parent = expandReference(parent, apiObject);
    obj.properties = mergeObjects(obj.properties, parent.properties);

    if (parent.required) {
      parent.required.forEach((item) => obj.required.add(item));
    }
  }
  obj.required = [...obj.required];
  delete obj.allOf;

  return obj;
}

function mergeObjects(...objs) {
  let mergedObj = {};

  for (let obj of objs) {
    if (obj) {
      mergedObj = { ...mergedObj, ...obj };
    }
  }

  return mergedObj;
}

function removeParentFromAllOf(child, parent, api) {
  if (!child.allOf) {
    return false;
  }

  const initialLength = child.allOf.length;
  child.allOf = child.allOf.filter((childAllOfItem) => {
    const obj = getObjectFromReference(childAllOfItem, api);
    return obj !== parent;
  });

  if (initialLength === child.allOf.length) {
    return false;
  }

  return true;
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
  removeParentFromAllOf,
  getObjectNameFromReference,
  detectDuplicates,
  mergeObjects,
  filterObjectProperties,
  flattenAllOfObject,
};
