// https://github.com/ReaTeam/Doc/blob/master/REAPER-CLI.md
//
// Usage:
// reaper.exe [options] [projectfile.rpp | mediafile.wav | scriptfile.lua [...]] | fxchainpreset.RfxChain | vstbank.fxb | vstpatch.fxp | vstpatch.vstpreset
// Multiple media files and/or scripts may be specified, and will be added or run in order. -nonewinst can be used to add media files and/or run scripts in an already-running instance of REAPER.
// FX chain, VST plugin bank and preset will be applied to a new automatically added track. The bank/preset will trigger instantiation of the compatible plugin on the newly added track. The feature is supported since build 6.43.
// Passing both project/template and a media file isn't supported as of Oct 10, 2021 (see Limitations below).
// Since build 6.80 passing either project/template or a media file AND a script file is supported, e.g.
// reaper.exe projectfile.rpp scriptfile.lua
// reaper.exe -nonewinst media.wav scriptfile.lua

package reaper
