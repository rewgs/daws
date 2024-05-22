from dataclasses import dataclass
from pathlib import Path, PurePath
from platform import system
from types import NoneType
import psutil

from .application import CubaseApp, Cubase


@dataclass
class Pref:
    """
    A single instance of Cubase preferences, either default or custom.
    """
    name: str
    version: int
    description: str
    default: bool  # False for custom
    main_prefs_path: Path
    user_presets_path: Path
    associated_installation: CubaseApp


class Preferences(Cubase):
    """
    Deals with locating, storing, and creating Cubase preferences, both default and custom.
    """

    @staticmethod
    def __get_default_main_path() -> Path:
        default_main_path: Path = Path(PurePath(Path.home()))

        if system() == "Darwin":
            default_main_path = default_main_path.joinpath("Library", "Preferences")
        # TODO:
        elif system() == "Windows":
            pass
        try:
            default_main_path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        else:
            return default_main_path.resolve(strict=True)

    @staticmethod
    def __get_default_user_path() -> Path:
        default_user_path: Path = Path(PurePath(Path.home()).joinpath("Documents"))

        if system() == "Darwin":
            default_user_path = default_user_path.joinpath("Steinberg", "Cubase", "User Presets")
        # TODO:
        elif system() == "Windows":
            pass

        try:
            default_user_path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        else:
            return default_user_path.resolve(strict=True)

    def __get_default_preferences(self, installed_apps: list[CubaseApp]) -> list[Pref] | None:
        """
        Returns a list of default preferences for all verisons of Cubase.
        """

        default_prefs: list[Pref] = []

        for file in self.default_main_path.iterdir():
            if file.is_dir():
                for i in installed_apps:
                    if file.name == i.path.stem:
                        try:
                            file.resolve(strict=True)
                        except FileNotFoundError as error:
                            raise error
                        else:
                            pref = Pref(file.name, i.version, "", True, file.resolve(strict=True),
                                        self.default_user_path, i, )
                            default_prefs.append(pref)
        return default_prefs


    def __init__(self, custom_preferences_store=None):
        super().__init__()
        self.custom_preferences_store: Path | NoneType = custom_preferences_store
        self.default_main_path: Path = self.__get_default_main_path()
        self.default_user_path: Path = self.__get_default_user_path()
        self.default_preferences: list[Pref] = self.__get_default_preferences(self.apps)
        self.custom_preferences: list[Pref] = []
        # self.current_preferences: Pref = # TODO:

    def add_custom(self, name: str, path: Path, version: int | NoneType = None,
                   description: str | NoneType = None) -> Pref:
        """Adds a new Pref in a custom location. Default behavior is to target the most recent version of Cubase, but
        an older version can be overridden.
        @param name: The name of the new Pref.
        @param path: The path to the new Pref.
        @param version: The version of Cubase the new Pref is to target.
        @param description: Optional long description of the new Pref.
        @return: The new Pref.
        """

        try:
            path.resolve(strict=True)
        except FileNotFoundError as error:
            raise error
        else:
            if version is None:
                version == self.newest

            custom_pref_path = path.resolve(strict=True)

            # remaining attr declarations
            custom_pref_associated_installation = self.get_by_version(version)

            custom_pref_user_preset_path: Path = self.default_user_path

            custom_pref = Pref(name, version, description, False, custom_pref_path, custom_pref_user_preset_path,
                               self.get_by_version(version))

            self.custom_preferences.append(custom_pref)
            return custom_pref

    # reference: https://realpython.com/python-json/
    def write(self):
        """Writes the CustomPref to disk."""
        pass

    def read(self):
        """Reads the CustomPref from disk."""
        pass
