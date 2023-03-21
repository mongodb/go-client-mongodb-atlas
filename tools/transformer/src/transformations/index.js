const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations, transformOneOf } = require("./oneOf");
const { applyDiscriminatorTransformations } = require("./discriminator");
const { applyGoNameTransformation } = require("./properties");
module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyGoNameTransformation,
  applyDiscriminatorTransformations,
};
