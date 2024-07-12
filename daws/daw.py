from abc import ABC, abstractmethod
from pathlib import Path, PurePath
from types import NoneType
from platform import system

from .consts import DAWS
from .errors import UnsupportedDaws, WrongOsBuddy


class Daw(ABC):
    """
    Abstract Base Class for a single instance of a DAW application, e.g. one 
    for Cubase 12, another for Cubase 13, another for Pro Tools 11, another for 
    Pro Tools 12, etc...
    """
    def __init__(self, name: str, path: Path, version: int):
        self.__check(name)
        self.path: Path = path
        self.version: int = version

    # This might not be needed, or could better be put in `Daw()`.
    # def __check(self, name):
    #     match = [daw for daw in DAWS if daw.name == self.name]
    #     if len(match) != 1:
    #         if len(match) == 0:
    #             raise UnsupportedDaw(f"The DAW {name} is not supported by this library!")
    #         # TODO: not sure what to do here. Panic? This shouldn't ever happen.
    #         if len(match) > 1:
    #             pass
    #     else:
    #         daw = match[0]
    #         self.name = daw.name
    #         self.developer = daw.developer
    #         self.operating_systems = daw.operating_systems

    def is_open(self) -> bool: 
        return True if len([proc for proc in psutil.process_iter(["pid", "name", "username"]) if f"{self.name} {self.version}" in proc.name() and proc.is_running()]) > 0 else False
