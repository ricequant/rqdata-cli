#!/usr/bin/env python3

import argparse
import importlib.util
import os
import shutil
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
if str(ROOT) not in sys.path:
    sys.path.insert(0, str(ROOT))

from scripts.python.platforms import PLATFORM_TARGETS, ROOT, WHEELHOUSE_DIR, detect_current_target


def run(command, env=None):
    subprocess.run(command, cwd=ROOT, env=env, check=True)


def has_module(name):
    return importlib.util.find_spec(name) is not None


def build_wheel(target_key, out_dir):
    env = os.environ.copy()
    env["RQDATA_PY_TARGET"] = target_key
    if has_module("build.__main__"):
        run(
            [
                sys.executable,
                "-m",
                "build.__main__",
                "--wheel",
                "--outdir",
                str(out_dir),
            ],
            env=env,
        )
        return

    if has_module("build"):
        raise SystemExit(
            "The installed 'build' package does not expose a runnable module. "
            "Upgrade it or install 'wheel', for example: python3 -m pip install -U build"
        )

    if has_module("wheel"):
        run(
            [
                sys.executable,
                "setup.py",
                "bdist_wheel",
                "--dist-dir",
                str(out_dir),
            ],
            env=env,
        )
        return

    raise SystemExit(
        "Building wheels requires either the 'build' package or the 'wheel' package. "
        "Install one of them first, for example: python3 -m pip install build"
    )


def main():
    parser = argparse.ArgumentParser(description="Build platform-specific Python wheels for rqdata-cli.")
    parser.add_argument(
        "--target",
        choices=["all", *sorted(PLATFORM_TARGETS)],
        default=detect_current_target(),
        help="Target wheel to build. Use 'all' to build every supported platform wheel.",
    )
    parser.add_argument(
        "--out-dir",
        default=str(WHEELHOUSE_DIR),
        help="Directory where built wheels are written.",
    )
    parser.add_argument(
        "--clean",
        action="store_true",
        help="Delete the output directory before building.",
    )
    args = parser.parse_args()

    out_dir = Path(args.out_dir).resolve()
    if args.clean and out_dir.exists():
        shutil.rmtree(out_dir)
    out_dir.mkdir(parents=True, exist_ok=True)

    targets = sorted(PLATFORM_TARGETS) if args.target == "all" else [args.target]
    for target in targets:
        print(f"Building wheel for {target}")
        build_wheel(target, out_dir)

    print(f"Wheel build complete: {out_dir}")


if __name__ == "__main__":
    main()
