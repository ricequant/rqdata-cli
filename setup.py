#!/usr/bin/env python3

import json
import os
import shutil
import subprocess
from pathlib import Path

from setuptools import find_packages, setup
from setuptools.command.build_py import build_py as _build_py

try:
    from wheel.bdist_wheel import bdist_wheel as _bdist_wheel
except ImportError:  # pragma: no cover - build backend installs wheel for real builds.
    _bdist_wheel = None


ROOT = Path(__file__).resolve().parent
PACKAGE_JSON = ROOT / "package.json"
README = ROOT / "README.md"
PYTHON_SRC = ROOT / "python" / "src"
DIST_DIR = ROOT / "dist"
DEFAULT_GOCACHE_DIR = ROOT / ".cache" / "go-build"

PACKAGE_NAME = "rqdata-cli"
IMPORT_NAME = "rqdata_cli"

PLATFORM_TARGETS = {
    "linux-x64": {
        "goos": "linux",
        "goarch": "amd64",
        "binary_name": "rqdata",
        "dist_name": "rqdata-linux-amd64",
        "wheel_plat_name": "manylinux2014_x86_64",
    },
    "linux-arm64": {
        "goos": "linux",
        "goarch": "arm64",
        "binary_name": "rqdata",
        "dist_name": "rqdata-linux-arm64",
        "wheel_plat_name": "manylinux2014_aarch64",
    },
    "darwin-x64": {
        "goos": "darwin",
        "goarch": "amd64",
        "binary_name": "rqdata",
        "dist_name": "rqdata-macos-amd64",
        "wheel_plat_name": "macosx_10_15_x86_64",
    },
    "darwin-arm64": {
        "goos": "darwin",
        "goarch": "arm64",
        "binary_name": "rqdata",
        "dist_name": "rqdata-macos-arm64",
        "wheel_plat_name": "macosx_11_0_arm64",
    },
    "win32-x64": {
        "goos": "windows",
        "goarch": "amd64",
        "binary_name": "rqdata.exe",
        "dist_name": "rqdata-windows-amd64.exe",
        "wheel_plat_name": "win_amd64",
    },
}

CURRENT_TARGETS = {
    ("linux", "x86_64"): "linux-x64",
    ("linux", "amd64"): "linux-x64",
    ("linux", "aarch64"): "linux-arm64",
    ("linux", "arm64"): "linux-arm64",
    ("darwin", "x86_64"): "darwin-x64",
    ("darwin", "amd64"): "darwin-x64",
    ("darwin", "arm64"): "darwin-arm64",
    ("win32", "x86_64"): "win32-x64",
    ("win32", "amd64"): "win32-x64",
}


def load_package_json():
    return json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))


def detect_current_target():
    import platform
    import sys

    machine = platform.machine().lower()
    key = (sys.platform, machine)
    target = CURRENT_TARGETS.get(key)
    if target:
        return target
    raise RuntimeError(
        f"unsupported platform for wheel build: {sys.platform}-{machine}. "
        "Set RQDATA_PY_TARGET explicitly."
    )


def resolve_target():
    key = os.environ.get("RQDATA_PY_TARGET") or detect_current_target()
    target = PLATFORM_TARGETS.get(key)
    if not target:
        supported = ", ".join(sorted(PLATFORM_TARGETS))
        raise RuntimeError(
            f"unsupported RQDATA_PY_TARGET={key!r}. Supported values: {supported}"
        )
    return key, target


def ensure_binary(version, target_key, target):
    binary_path = DIST_DIR / target["dist_name"]
    if binary_path.exists():
        return binary_path

    DIST_DIR.mkdir(parents=True, exist_ok=True)
    gocache_dir = Path(os.environ.get("GOCACHE_DIR", DEFAULT_GOCACHE_DIR))
    gocache_dir.mkdir(parents=True, exist_ok=True)
    ldflags = f"-s -w -X github.com/ricequant/rqdata-cli/cmd.Version={version}"

    env = os.environ.copy()
    env.update(
        {
            "CGO_ENABLED": "0",
            "GOOS": target["goos"],
            "GOARCH": target["goarch"],
            "GOCACHE": str(gocache_dir),
        }
    )

    print(f"Building Go binary for {target_key} -> {binary_path.name}")
    subprocess.run(
        ["go", "build", "-trimpath", "-ldflags", ldflags, "-o", str(binary_path), "main.go"],
        cwd=ROOT,
        env=env,
        check=True,
    )
    if target["binary_name"] != "rqdata.exe":
        binary_path.chmod(0o755)
    return binary_path


class build_py(_build_py):
    def run(self):
        super().run()

        package_info = load_package_json()
        version = package_info["version"]
        target_key, target = resolve_target()
        binary_source = ensure_binary(version, target_key, target)

        destination_dir = Path(self.build_lib) / IMPORT_NAME / "bin"
        shutil.rmtree(destination_dir, ignore_errors=True)
        destination_dir.mkdir(parents=True, exist_ok=True)
        destination = destination_dir / target["binary_name"]
        shutil.copy2(binary_source, destination)
        if target["binary_name"] != "rqdata.exe":
            destination.chmod(0o755)


if _bdist_wheel is not None:

    class bdist_wheel(_bdist_wheel):
        def finalize_options(self):
            super().finalize_options()
            _, target = resolve_target()
            self.root_is_pure = False
            self.plat_name = target["wheel_plat_name"]
            self.plat_name_supplied = True

        def get_tag(self):
            _, target = resolve_target()
            return ("py3", "none", target["wheel_plat_name"])


else:  # pragma: no cover
    bdist_wheel = None


package_info = load_package_json()

cmdclass = {
    "build_py": build_py,
}
if bdist_wheel is not None:
    cmdclass["bdist_wheel"] = bdist_wheel

setup(
    name=PACKAGE_NAME,
    version=package_info["version"],
    description=package_info["description"],
    long_description=README.read_text(encoding="utf-8"),
    long_description_content_type="text/markdown",
    license=package_info["license"],
    url=package_info["homepage"],
    project_urls={
        "Source": package_info["repository"]["url"],
        "Issues": package_info["bugs"]["url"],
    },
    python_requires=">=3.8",
    packages=find_packages(where=str(PYTHON_SRC)),
    package_dir={"": "python/src"},
    include_package_data=True,
    package_data={IMPORT_NAME: ["bin/*"]},
    entry_points={
        "console_scripts": [
            "rqdata=rqdata_cli.__main__:main",
        ]
    },
    cmdclass=cmdclass,
    keywords=package_info["keywords"],
    classifiers=[
        "Development Status :: 4 - Beta",
        "Environment :: Console",
        "Intended Audience :: Developers",
        "Intended Audience :: Financial and Insurance Industry",
        "Operating System :: MacOS",
        "Operating System :: Microsoft :: Windows",
        "Operating System :: POSIX :: Linux",
        "Programming Language :: Go",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3 :: Only",
        "Topic :: Office/Business :: Financial",
        "Topic :: Software Development :: Libraries :: Python Modules",
    ],
)
