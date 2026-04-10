#!/usr/bin/env node

const path = require('node:path');
const { spawnSync } = require('node:child_process');
const { PLATFORM_TARGETS } = require('./platforms');

const ROOT = path.resolve(__dirname, '..', '..');
const NPM_TAG = process.env.NPM_TAG || 'latest';
const NPM_DRY_RUN = process.env.NPM_DRY_RUN === '1';

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

function publishArgs() {
  const args = ['publish', '--access', 'public', '--tag', NPM_TAG];

  if (NPM_DRY_RUN) {
    args.push('--dry-run');
  }

  return args;
}

run('node', ['scripts/npm/build-packages.js']);

for (const target of PLATFORM_TARGETS) {
  run('npm', publishArgs(), target.packageDir);
}

run('npm', publishArgs());
