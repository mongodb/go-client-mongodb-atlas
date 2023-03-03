const { getObjectFromReference } = require("../engine/readers");

function filterObjectProperties(object, filter = (_k, _v) => true) {
  const filteredObj = Object.keys(object)
    .filter((key) => filter(key, object[key]))
    .reduce((aggregationObj, key) => {
      aggregationObj[key] = object[key];
      return aggregationObj;
    }, {});

  return filteredObj;
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

module.exports = {
  removeParentFromAllOf,
  detectDuplicates,
  mergeObjects,
  filterObjectProperties,
  flattenAllOfObject,
};
