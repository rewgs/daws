from abc import ABC, abstractmethod
from pathlib import Path, PurePath
from types import NoneType
from platform import system

from .consts import DAWS
from .errors import UnsupportedDaws, WrongOsBuddy
import daw


class Core(ABC):
    """
    Base class for the package.
    """

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

    # FIXME:
    def __get_installed(self) -> DawApp | list[DawApp] | NoneType:
        """
        Gets all instances of `daw`.
        """
        default_path = self.__get_default_path()
        try:
            default_path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        else:
            if daw is not None:
                if isinstance(daw, list):
                    # TODO:
                    # for d in daw:
                else:
                    # TODO: adapt to be DAW-agnostic
                    installations: list[CubaseApp] = []
                    app_paths = [ file for file in get_default_path().iterdir() if file.is_dir() and daw in file.name ]
                    for p in app_paths:
                        extracted_number: list = [
                            char for char in p.stem.split() if char.isdigit()
                        ]
                        version_number = int(extracted_number[0])
                        app = CubaseApp(p, version_number)
                        installations.append(app)
                    return installations
            # TODO: get all installations of all DAWs
            else:
                pass

    # TODO: make some `@abstractmethod`s. 
    # These not only provide an opportunity to force the subclass to have to 
    # define the methods marked with this decorator, but also their presence 
    # effectively makes this class a read-only template from which to make Daw 
    # objects from, and prevents it from being instantiated itself.
    # This is a great way to deal with functionality that is essential to all 
    # DAWs, but the execution of that functionality is DAW-specific.
    # ...or perhaps go with Protocol instead? See: https://www.youtube.com/watch?v=xvb5hGLoK0A
