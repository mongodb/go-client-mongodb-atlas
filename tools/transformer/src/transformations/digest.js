/**
 * Remove digest auth from paths and security schemes
 */
function applyDigestTransformations(api) {
  const hasSchemas = api && api.components && api.components.securitySchemes;
  if (hasSchemas) {
    delete api.components.securitySchemes.DigestAuth;
    Object.keys(api.paths).forEach((path) => {
      Object.keys(api.paths[path]).forEach((method) => {
        delete api.paths[path][method].security;
      });
    });
  }

  return api;
}

module.exports = {
  applyDigestTransformations,
};
