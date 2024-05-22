A library for working with all major DAWs and their files.

Contains two submodules: `daw` and `project`:
- `daw`: deals with the DAW applications -- paths, various versions, etc.
- `project`: deals with session files, project files, etc.

- daws
    - daw: Concerned with the abstract notion of the DAW in question -- the various application versions installed, whether any of said applications are open, etc. For example, Cubase *IS* a DAW, so it inherits from `daw`.
        - dawApp: A single application installation (i.e. one for Cubase 12, another for Cubase 13, etc).
        - pref: A daw or dawApp's preferences files.
        - name
        - developer
        - operatingSystems
    - project: The root of a scoring project, within which *all* assets would presumably be located -- not just DAW sessions, but also picture, mixes, documents...anything.
        - sessionCollection: A folder within a Project that contains one or more Sessions.
            - sessionFolder: A single folder when creating a cue. Contains DAW files (.rpp, .ptx, etc), probably an "Audio Files" folder, etc.
                - sessionFile: A single session file within a Session folder, e.g. a .rpp file, .ptx file, .cpr, etc.

example daw:
- Cubase(daw)
    - Cubase 12(dawApp)
    - Cubase 13(dawApp)
- Live
- Logic
- ProTools
- Reaper

example project:
- LBP(project)
    - cues(sessionCollection)
        - 1m01(sessionFolder)
            - 1m01_v01.cpr(sessionFile)
