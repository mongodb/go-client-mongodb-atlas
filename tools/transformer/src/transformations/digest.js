function applyDigestTransformations(api) {
  const hasSchemas = api && api.components && api.components.securitySchemes;
  if (hasSchemas) {
    delete api.components.securitySchemes.DigestAuth;
  }
  return api;
}

module.exports = {
  applyDigestTransformations,
};
