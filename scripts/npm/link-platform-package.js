#!/usr/bin/env node

const { mkdirSync, rmSync, symlinkSync, existsSync } = require('node:fs');
const path = require('node:path');
const { findCurrentTarget } = require('./platforms');

const ROOT = path.resolve(__dirname, '..', '..');
const NODE_MODULES = path.join(ROOT, 'node_modules');

const target = findCurrentTarget();

if (!target) {
  console.warn(`Unsupported platform for local link: ${process.platform}-${process.arch}`);
  process.exit(0);
}

const destination = path.join(NODE_MODULES, target.packageName);
const destinationDir = path.dirname(destination);

mkdirSync(NODE_MODULES, { recursive: true });
mkdirSync(destinationDir, { recursive: true });
rmSync(destination, { recursive: true, force: true });
rmSync(path.join(NODE_MODULES, target.packageName.split('/').pop()), { recursive: true, force: true });

if (!existsSync(target.packageDir)) {
  console.error(`Platform package has not been built: ${target.packageDir}`);
  console.error('Run `npm run build:npm` first.');
  process.exit(1);
}

symlinkSync(target.packageDir, destination, 'dir');
console.log(`Linked ${target.packageName} -> ${target.packageDir}`);
