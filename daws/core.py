from pathlib import Path, PurePath
from platform import system
from typing import Optional

from .daw import Daw
from .consts import SUPPORTED_DAWS
from .errors import UnsupportedSystem


class Core:
    """
    Base class for the package.
    """

    def __init__(self):
        self.__ROOT: Path = Path(PurePath(Path.home().root))
        self.__OS: str = system()
        self.daws = [daw for daw in SUPPORTED_DAWS if self.__is_installed(daw.name)]

    @staticmethod
    def __resolve(path: str | Path | PurePath) -> Path:
        if isinstance(path, Path):
            try:
                path.resolve(strict=True)
            except FileNotFoundError as error:
                raise error
            else:
                return path.resolve(strict=True)
        # TODO:
        elif isinstance(path, PurePath):
            pass
        # TODO:
        elif isinstance(path, str):
            pass
        # TODO:
        else:
            pass

    def __get_default_path(self) -> Path:
        """
        Returns the default Path in which DAW apps will be installed, as well as the operating system.
        """
        system_root = Path(Path(__file__).root).resolve(strict=True)
        if self.__OS == "Darwin":
            return self.__resolve(system_root.joinpath("Applications"))
        elif self.__OS == "Linux":
            # TODO: Place NotWsl and is_wsl elsewhere?
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
            return self.__resolve(program_files)
        elif self.__OS == "Windows":
            return self.__resolve(system_root.joinpath("Program Files"))
        else:
            raise UnsupportedSystem(f"This app is not designed to run on {self.__OS}!")

    # FIXME:
    def __get_installed(self) -> Optional[list[Daw]]:
        """
        Gets all instances of `Daw`.
        """
        default_path = self.__get_default_path()
        try:
            default_path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        # else:
        #     if daw is not None:
        #         if isinstance(daw, list):
        #         # TODO:
        #         # for d in daw:
        #         else:
        #             # TODO: adapt to be DAW-agnostic
        #             installations: list[CubaseApp] = []
        #             app_paths = [file for file in self.__get_default_path().iterdir() if file.is_dir() and daw in file.name]
        #             for p in app_paths:
        #                 extracted_number: list = [
        #                     char for char in p.stem.split() if char.isdigit()
        #                 ]
        #                 version_number = int(extracted_number[0])
        #                 app = CubaseApp(p, version_number)
        #                 installations.append(app)
        #             return installations
        #     # TODO: get all installations of all DAWs
        #     else:
        #         pass

    # TODO: make some `@abstractmethod`s. 
    # These not only provide an opportunity to force the subclass to have to 
    # define the methods marked with this decorator, but also their presence 
    # effectively makes this class a read-only template from which to make Daw 
    # objects from, and prevents it from being instantiated itself.
    # This is a great way to deal with functionality that is essential to all 
    # DAWs, but the execution of that functionality is DAW-specific.
    # ...or perhaps go with Protocol instead? See: https://www.youtube.com/watch?v=xvb5hGLoK0A

