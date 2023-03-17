const {
    applyAllOfTransformations,
    applyOneOfTransformations,
    applyModelNameTransformations,
    applyDiscriminatorTransformations,
  } = require("./transformations");
  
const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
    openapi = applyDiscriminatorTransformations(openapi);
    openapi = applyOneOfTransformations(openapi);
    openapi = applyAllOfTransformations(openapi);
  
    openapi = applyModelNameTransformations(openapi, "Api", "", ignoredModelNames);
    openapi = applyModelNameTransformations(openapi, "Atlas", "", ignoredModelNames);
    openapi = applyModelNameTransformations(openapi, "", "Manual", ignoredModelNames);
    openapi = applyModelNameTransformations(openapi, "", "View", ignoredModelNames);
    
    return openapi
  }
