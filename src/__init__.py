from abc import ABC, abstractmethod
from pathlib import Path, PurePath
from types import NoneType
from platform import system

from .consts import DAWS
from .errors import UnsupportedDaws, WrongOsBuddy
import daw


class Daws:
    """Base class for the package."""

    def __init__(self):
        self.__ROOT: Path = Path(PurePath(Path.home().root))
        self.__OS: str = system()
        self.daws = [ daw for daw in DAWS if self.__check_if_daw_installed(daw.name) ]

    def __check_if_daw_installed(self, daw: str) -> bool:
        pass

    @staticmethod
    def __get_default_path() -> tuple[Path, str]:
        """Returns the default Path in which DAW apps will be installed, as well as the operating system."""

        # NOTE: moved these to Daws class
        # system_root: Path = Path(PurePath(Path.home().root))
        # OS = system()

        if OS == "Darwin":
            return system_root.joinpath("Applications"), OS
        elif OS == "Linux":
            class NotWsl(Exception):
                """Raises a basic exception if not running on WSL."""

            def is_wsl(path: Path) -> bool:
                try:
                    path.resolve(strict=True)
                except FileNotFoundError:
                    return False
                else:
                    return True

            program_files = system_root.joinpath("mnt", "c", "Program Files")
            if not is_wsl(program_files):
                raise NotWsl(f"Expected to find {program_files.as_posix()}, but didn't. Are you running this on WSL?")
            return program_files

        # TODO:
        elif OS == "Windows":
            pass
        else:
            raise WrongOsBuddy(f"This app is not designed to run on {OS}!")
