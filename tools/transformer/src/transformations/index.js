const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations, transformOneOf } = require("./oneOf");
const { applyDiscriminatorTransformations } = require("./discriminator");
const { applyDigestTransformations } = require("./digest");

module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyDiscriminatorTransformations,
  applyDigestTransformations,
};
