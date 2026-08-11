#!/usr/bin/env python3

import subprocess
import sys
from pathlib import Path


def resolve_binary():
    package_dir = Path(__file__).resolve().parent
    binary_name = "rqdata.exe" if sys.platform == "win32" else "rqdata"
    binary_path = package_dir / "bin" / binary_name

    if not binary_path.exists():
        raise FileNotFoundError(
            f"rqdata binary is missing: {binary_path}. Reinstall the wheel for your platform."
        )

    return binary_path


def main():
    binary_path = resolve_binary()
    result = subprocess.run([str(binary_path), *sys.argv[1:]])
    return result.returncode


if __name__ == "__main__":
    raise SystemExit(main())
