#!/usr/bin/env node

const { readFileSync } = require('node:fs');
const path = require('node:path');
const { spawnSync } = require('node:child_process');
const { PLATFORM_TARGETS } = require('./platforms');

const ROOT = path.resolve(__dirname, '..', '..');
const pkg = JSON.parse(readFileSync(path.join(ROOT, 'package.json'), 'utf8'));

function run(command, args, cwd = ROOT) {
  const result = spawnSync(command, args, {
    cwd,
    stdio: 'inherit',
    env: process.env
  });

  if (result.status !== 0) {
    process.exit(result.status === null ? 1 : result.status);
  }
}

run('node', ['scripts/npm/build-packages.js']);

for (const target of PLATFORM_TARGETS) {
  run('npm', ['pack'], target.packageDir);
}

run('npm', ['pack']);

console.log(`Packed ${pkg.name}@${pkg.version} and ${PLATFORM_TARGETS.length} platform packages.`);
