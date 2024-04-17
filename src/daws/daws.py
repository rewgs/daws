from pathlib import Path, PurePath
from types import NoneType
from platform import system

from .consts import DAWS
from .errors import UnsupportedDaws, WrongOsBuddy

class DawApp:
    """
    Class for a single instance of a DAW application, e.g. one for Cubase 12, another for 
    Cubase 13, another for Pro Tools 11, another for Pro Tools 12, etc...
    """
    def __init__(self, name: str, path: Path, version: int):
        self.__check(name)
        self.path: Path = path
        self.version: int = version

    def __check(self, name):
        match = [daw for daw in DAWS if daw.name == self.name]
        if len(match) != 1:
            if len(match) == 0:
                raise UnsupportedDaw(f"The DAW {name} is not supported by this library!")
            # TODO: not sure what to do here. Panic? This shouldn't ever happen.
            if len(match) > 1:
                pass
        else:
            daw = match[0]
            self.name = daw.name
            self.developer = daw.developer
            self.operating_systems = daw.operating_systems

    def is_open(self) -> bool: 
        return True if len([proc for proc in psutil.process_iter(["pid", "name", "username"]) if f"{self.name} {self.version}" in proc.name() and proc.is_running()]) > 0 else False


class Daw:
    """
    Abstract base class for all DAWs (Cubase, Reaper, etc).
    """
    def __init__(self, name: str | NoneType = None):
        self.name = name


    @staticmethod
    def __get_installed(daw: str | list[str] | NoneType = None) -> DawApp | list[DawApp] | NoneType:
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


class Daws:
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
        """
        Returns the default Path in which DAW apps will be installed, as well 
        as the operating system.
        """

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
