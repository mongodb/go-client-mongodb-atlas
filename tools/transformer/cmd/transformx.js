#!/usr/bin/env node
const yargs = require('yargs')

yargs.scriptName("openapi-transformx")
  .usage('$0 <cmd> [args]')
  .command('transform [openapi]', 'transform openapi spec', (yargs) => {
    return yargs
      .positional('openapi', {
        describe: 'openapi file location to transform',
      })
  }, (argv) => {
    if(!argv.openapi){
      throw new Error("Missing openapi file")
    }
    // TODO replace with method
    process.env.OPENAPI_FILE=argv.openapi
    require("../src/transform.js")
  }).help()
  .argv
