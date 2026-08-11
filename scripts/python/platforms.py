from pathlib import Path
import platform
import sys


ROOT = Path(__file__).resolve().parents[2]
DIST_DIR = ROOT / "dist"
WHEELHOUSE_DIR = ROOT / "wheelhouse"

PLATFORM_TARGETS = {
    "linux-x64": {
        "goos": "linux",
        "goarch": "amd64",
        "wheel_plat_name": "manylinux2014_x86_64",
    },
    "darwin-x64": {
        "goos": "darwin",
        "goarch": "amd64",
        "wheel_plat_name": "macosx_10_15_x86_64",
    },
    "darwin-arm64": {
        "goos": "darwin",
        "goarch": "arm64",
        "wheel_plat_name": "macosx_11_0_arm64",
    },
    "win32-x64": {
        "goos": "windows",
        "goarch": "amd64",
        "wheel_plat_name": "win_amd64",
    },
}


def detect_current_target():
    machine = platform.machine().lower()
    key = (sys.platform, machine)
    mapping = {
        ("linux", "x86_64"): "linux-x64",
        ("linux", "amd64"): "linux-x64",
        ("darwin", "x86_64"): "darwin-x64",
        ("darwin", "amd64"): "darwin-x64",
        ("darwin", "arm64"): "darwin-arm64",
        ("win32", "x86_64"): "win32-x64",
        ("win32", "amd64"): "win32-x64",
    }
    target = mapping.get(key)
    if not target:
        supported = ", ".join(sorted(PLATFORM_TARGETS))
        raise RuntimeError(
            f"unsupported host platform: {sys.platform}-{machine}. Supported targets: {supported}"
        )
    return target
