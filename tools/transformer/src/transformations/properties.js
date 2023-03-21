/**
 * Rename object based on provided property names
 * @returns
 */
function applyGoNameTransformation(api, namesToMatch) {
  if (!api.components) {
    throw new Error("Invalid schema for goStandardProperties");
  }
  var components = api.components;
  // Loop through each schema in the components
  for (let schemaName in components.schemas) {
    let schema = components.schemas[schemaName];
    // Loop through each property in the schema
    for (let propName in schema.properties) {
      let uppercasePropName = propName.toUpperCase();
      // Check if the property name matches any of the names to match
      if (namesToMatch.includes(uppercasePropName)) {
        schema.properties[propName]['x-go-name'] = uppercasePropName;
      }
    }
  }
  return api;
}

module.exports = {
  applyGoNameTransformation,
};
