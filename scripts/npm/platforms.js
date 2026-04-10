const path = require('node:path');

const ROOT = path.resolve(__dirname, '..', '..');
const DIST_DIR = path.join(ROOT, 'dist');
const PACKAGES_DIR = path.join(ROOT, 'npm');
const NPM_SCOPE = '@ricequant2026';

const PLATFORM_TARGETS = [
  {
    key: 'linux-x64',
    goos: 'linux',
    goarch: 'amd64',
    os: 'linux',
    cpu: 'x64',
    packageName: `${NPM_SCOPE}/rqdata-cli-linux-x64`,
    packageDir: path.join(PACKAGES_DIR, 'rqdata-cli-linux-x64'),
    output: 'rqdata-linux-amd64',
    binName: 'rqdata'
  },
  {
    key: 'darwin-x64',
    goos: 'darwin',
    goarch: 'amd64',
    os: 'darwin',
    cpu: 'x64',
    packageName: `${NPM_SCOPE}/rqdata-cli-darwin-x64`,
    packageDir: path.join(PACKAGES_DIR, 'rqdata-cli-darwin-x64'),
    output: 'rqdata-macos-amd64',
    binName: 'rqdata'
  },
  {
    key: 'darwin-arm64',
    goos: 'darwin',
    goarch: 'arm64',
    os: 'darwin',
    cpu: 'arm64',
    packageName: `${NPM_SCOPE}/rqdata-cli-darwin-arm64`,
    packageDir: path.join(PACKAGES_DIR, 'rqdata-cli-darwin-arm64'),
    output: 'rqdata-macos-arm64',
    binName: 'rqdata'
  },
  {
    key: 'win32-x64',
    goos: 'windows',
    goarch: 'amd64',
    os: 'win32',
    cpu: 'x64',
    packageName: `${NPM_SCOPE}/rqdata-cli-win32-x64`,
    packageDir: path.join(PACKAGES_DIR, 'rqdata-cli-win32-x64'),
    output: 'rqdata-windows-amd64.exe',
    binName: 'rqdata.exe'
  }
];

function currentPlatformKey() {
  return `${process.platform}-${process.arch}`;
}

function findCurrentTarget() {
  return PLATFORM_TARGETS.find((target) => target.key === currentPlatformKey()) || null;
}

function findTargetByKey(key) {
  return PLATFORM_TARGETS.find((target) => target.key === key) || null;
}

module.exports = {
  DIST_DIR,
  NPM_SCOPE,
  PACKAGES_DIR,
  PLATFORM_TARGETS,
  findTargetByKey,
  findCurrentTarget
};
