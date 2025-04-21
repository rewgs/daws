from abc import (
    ABC,
    # abstractmethod
)
from pathlib import (
    Path,
    # PurePath
)

import psutil

# from types import NoneType
# from platform import system
#
# from .consts import DAWS
# from .errors import UnsupportedDaws, WrongOsBuddy


class Daw(ABC):
    """
    Abstract Base Class for a single instance of a DAW application, e.g. one
    for Cubase 12, another for Cubase 13, another for Pro Tools 11, another for
    Pro Tools 12, etc...
    """

    def __init__(self, name: str, path: Path, version: int):
        self.name = name
        self.path = path
        self.version = version

    def is_open(self) -> bool:
        """ """
        return (
            True
            if len(
                [
                    proc
                    for proc in psutil.process_iter(["pid", "name", "username"])
                    if f"{self.name} {self.version}" in proc.name()
                    and proc.is_running()
                ]
            )
            > 0
            else False
        )

    # TODO:
    def __is_installed(self, daw: str) -> bool:
        """ """
        pass
