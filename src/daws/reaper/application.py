from pathlib import Path, PurePath
from types import NoneType

from ../daws/daw import Daw

class Reaper(Daw):
    def __init__(self):
        self.name = "Reaper"
        # self.executable = 

        super.__init__(self.name)
