from pathlib import Path, PurePath
from types import NoneType
import os

from daws import Daw


class SessionFile:
    """
    A single session file within a Session folder, e.g. a .rpp file, .ptx file, .cpr, etc.
    """

    def __init__(self, path: Path):
        self.path = path
        # self.version = 
        # self.last_modified = 
        # self.birth_time = 


class Session:
    """
    A single folder when creating a cue. Contains DAW files (.rpp, .ptx, etc), 
    probably an "Audio Files" folder, etc.
    """

    def __determine_daw(self):
        """
        Examines contents of a Session folder and determines which DAW it's from.
        """

    def __get_session_files(self):
        pass

    def __init__(self, daw: Daw | NoneType = None):
        self.daw = self.__determine_daw() if daw is None else daw
        self.session_files = self.__get_session_files()


class Collection:
    """
    A folder within a Project that contains one or more Sessions.
    """

    def __init__(self, path: Path):
        self.path = path

    def __get_sessions(self) -> list[Session]:
        """
        """
        pass

    def get_session_files(self) -> list[SessionFile]:
        """
        """
        pass


class Project:
    """
    The root of a scoring project, within which *all* assets would presumably be located -- not 
    just DAW sessions, but also picture, mixes, documents...anything.
    """

    # TODO: replace me with better function (forget where it is, but it also takes PurePaths and strings)
    @staticmethod
    def validate_selected_dir(dir: Path) -> Path:
        try:
            dir.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        finally:
            return dir.resolve(strict=True)

    @staticmethod
    def convert_str_to_path(path: str) -> Path:
        # .parts() returns a tuple; must be converted to a list since I'm editing it
        parts = list(PurePath(path).parts)

        # Path() can't take a string with `\` in it, so they need to be removed:
        for i, part in enumerate(parts):
            if "\\" in part:
                parts[i] = part.replace("\\", "")

        # Joins the parts back together before feeding into Path()...
        joined_parts = os.path.join(*parts)
        project_path = Path(joined_parts)
        return project_path

    def __prompt_for_path(self) -> Path:
        project_dir: str = input("Enter the full path of the project folder's root:\n").strip()
        project_path: Path = convert_str_to_path(project_dir)
    
        try:
            res_project_path = project_path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        else:
            return res_project_path


    def __get_session_collections(self) -> list[Collection]:
        """
        """
        pass

    def __init__(self, path: Path | NoneType = None, daw: Daw | NoneType = None):
        # FIXME: validate_selected_path() not working, so not using it for now
        # self.path = self.__prompt_for_path() if path is None else validate_selected_dir(path)
        self.path = self.__prompt_for_path() if path is None else path
        self.daw = self.__prompt_for_daw() if daw is None else daw
        self.session_collections: list[Collection] = self.__get_session_collections()

    def set_as_session_collection(self, path: Path) -> Collection:
        """
        Returns a Collection for a given Path.
        """
        collection_path = self.convert_str_to_path(path)
        collection = Collection(collection_path)
        # FIXME: CODESMELL
        if self.session_collections:
            if collection not in self.session_collections:
                self.session_collections.append(collection)
        else:
            self.session_collections = []
            self.session_collections.append(collection)
        return collection

    def list_session_collections(self):
        for collection in self.session_collections:
            print(collection.path)

    def get_session_files(self) -> list[SessionFile]:
        """
        """
        pass

    def get_sessions(self) -> list[Session]:
        """
        """
        pass



