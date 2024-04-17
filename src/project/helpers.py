from pathlib import Path, PurePath


# TODO: replace me with better function (forget where it is, but it also takes PurePaths and strings)
def validate_selected_dir(dir: Path) -> Path:
    try:
        dir.resolve(strict=True)
    except FileNotFoundError as error:
        raise error
    finally:
        return dir.resolve(strict=True)
