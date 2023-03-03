const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations, transformOneOf } = require("./oneOf");
const { applyDiscriminatorTransformations } = require("./discriminator");

module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyDiscriminatorTransformations,
};
