# Minemap

Generate map images from Minecraft world files. Currently only supports the Java edition.

# Possibly Useful Links

* Formats
    * NBT - [Fandom](https://minecraft.fandom.com/wiki/NBT_format)
    * [NBT Spec](http://web.archive.org/web/20110723210920/http://www.minecraft.net/docs/NBT.txt) (web archive)
    * Region - [Fandom](https://minecraft.fandom.com/wiki/Region_file_format)
    * Chunks - [Fandom](https://minecraft.fandom.com/wiki/Chunk_format)
    * [Level](https://minecraft.fandom.com/wiki/Java_Edition_level_format) and [Player](https://minecraft.fandom.com/wiki/Player.dat_format)
* Data Values: [Java Edition](https://minecraft.gamepedia.com/Java_Edition_data_values)
* Articles: [Clockwork Codex](http://clockworkcodex.blogspot.com/2011/06/minecraft-mapping-reading-minecraft.html)
* Tools: [Dinnerbone Coords](https://dinnerbone.com/minecraft/tools/coordinates/)
* Colors: [Kenneth Moreland](https://www.kennethmoreland.com/color-advice/)


## Dependencies

* [ImageSharp](https://github.com/SixLabors/ImageSharp) - used for cross-platform image creation
    * Docs - [Getting Started](https://docs.sixlabors.com/articles/imagesharp/gettingstarted.html)
* [SharpZipLib](https://github.com/icsharpcode/SharpZipLib) - used to decompress chunks
* [CommandLine](https://github.com/commandlineparser/commandline) - used to parse command line args in the CLI


## Similar Tools

* [jaquadro/NBTExplorer](https://github.com/jaquadro/NBTExplorer) (C#) - A graphical NBT editor for all Minecraft NBT data sources - [Substrate](https://github.com/minecraft-dotnet/Substrate) library
* [AMIDST](https://github.com/crbednarz/AMIDST) (Java) - Advanced Minecraft Interface and Data/Structure Tracking
* Mapcrafter (C++) - [Home](https://docs.mapcrafter.org/builds/stable/) - [github](https://github.com/mapcrafter/mapcrafter) - High Performance Minecraft Map Renderer.
* Minecraft X-Ray (Java) - [Home](https://apocalyptech.com/minecraft-xray/) - [Github](https://github.com/apocalyptech/minecraftxray)
* [Minecraft-Overviewer](https://github.com/overviewer/Minecraft-Overviewer) (Python) - Render high-resolution maps of a Minecraft world with a Leaflet powered interface
    * [Design Doc](https://docs.overviewer.org/en/latest/design/designdoc/#)
    * [Leaflet JS](https://leafletjs.com/)
* [Hugobros3/Enklume](https://github.com/Hugobros3/Enklume) (Java) - Java library for parsing Minecraft save files

