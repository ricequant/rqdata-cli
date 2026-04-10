#!/usr/bin/env node

const { mkdirSync, rmSync, readFileSync, writeFileSync, copyFileSync, chmodSync } = require('node:fs');
const path = require('node:path');
const { spawnSync } = require('node:child_process');
const { DIST_DIR, PLATFORM_TARGETS, PACKAGES_DIR } = require('./platforms');

const ROOT = path.resolve(__dirname, '..', '..');
const pkg = JSON.parse(readFileSync(path.join(ROOT, 'package.json'), 'utf8'));
const VERSION = process.env.VERSION || pkg.version;
const GOCACHE_DIR = process.env.GOCACHE_DIR || path.join(ROOT, '.cache', 'go-build');

function run(command, args, options = {}) {
  const result = spawnSync(command, args, {
    cwd: options.cwd || ROOT,
    stdio: 'inherit',
    env: {
      ...process.env,
      ...options.env
    }
  });

  if (result.status !== 0) {
    process.exit(result.status === null ? 1 : result.status);
  }
}

function buildTargets() {
  rmSync(DIST_DIR, { recursive: true, force: true });
  mkdirSync(DIST_DIR, { recursive: true });
  mkdirSync(GOCACHE_DIR, { recursive: true });

  for (const target of PLATFORM_TARGETS) {
    const outputPath = path.join(DIST_DIR, target.output);
    const ldflags = `-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=${VERSION}`;

    console.log(`Building ${target.key} -> ${target.output}`);
    run('go', ['build', '-trimpath', '-ldflags', ldflags, '-o', outputPath, 'main.go'], {
      env: {
        CGO_ENABLED: '0',
        GOCACHE: GOCACHE_DIR,
        GOOS: target.goos,
        GOARCH: target.goarch
      }
    });
  }
}

function writePlatformPackage(target) {
  const packageDir = target.packageDir;
  const binarySource = path.join(DIST_DIR, target.output);
  const binaryTarget = path.join(packageDir, 'bin', target.binName);

  rmSync(packageDir, { recursive: true, force: true });
  mkdirSync(path.join(packageDir, 'bin'), { recursive: true });

  copyFileSync(binarySource, binaryTarget);
  if (target.os !== 'win32') {
    chmodSync(binaryTarget, 0o755);
  }

  writeFileSync(
    path.join(packageDir, 'index.js'),
    [
      "const path = require('node:path');",
      '',
      `module.exports.binary = path.join(__dirname, 'bin', '${target.binName}');`,
      ''
    ].join('\n')
  );

  writeFileSync(
    path.join(packageDir, 'README.md'),
    `# ${target.packageName}\n\nPlatform binary package for @ricequant2026/rqdata-cli.\n`
  );

  writeFileSync(
    path.join(packageDir, 'package.json'),
    JSON.stringify(
      {
        name: target.packageName,
        version: VERSION,
        description: `Platform binary for @ricequant2026/rqdata-cli (${target.key})`,
        license: pkg.license,
        os: [target.os],
        cpu: [target.cpu],
        files: ['bin/', 'index.js', 'README.md', 'LICENSE'],
        main: 'index.js'
      },
      null,
      2
    ) + '\n'
  );

  copyFileSync(path.join(ROOT, 'LICENSE'), path.join(packageDir, 'LICENSE'));
}

function materializePackages() {
  rmSync(PACKAGES_DIR, { recursive: true, force: true });
  mkdirSync(PACKAGES_DIR, { recursive: true });

  for (const target of PLATFORM_TARGETS) {
    writePlatformPackage(target);
  }
}

buildTargets();
materializePackages();
