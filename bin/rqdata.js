#!/usr/bin/env node

const { spawnSync } = require('node:child_process');
const { findCurrentTarget } = require('../scripts/npm/platforms');

function resolveBinary() {
  const target = findCurrentTarget();

  if (!target) {
    console.error(`Unsupported platform: ${process.platform}-${process.arch}`);
    process.exit(1);
  }

  let binary;

  try {
    ({ binary } = require(target.packageName));
  } catch (error) {
    console.error(`Missing platform package: ${target.packageName}`);
    console.error('Reinstall the package on a supported platform or verify optional dependencies are enabled.');
    process.exit(1);
  }

  if (!binary) {
    console.error('rqdata binary is missing. Reinstall the package or rebuild release artifacts.');
    process.exit(1);
  }

  return binary;
}

const result = spawnSync(resolveBinary(), process.argv.slice(2), {
  stdio: 'inherit'
});

if (result.error) {
  console.error(result.error.message);
  process.exit(1);
}

process.exit(result.status === null ? 1 : result.status);
