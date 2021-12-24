// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using CommandLine;

namespace MineMap.Cli.Options
{
    [Verb("dump-chunk", HelpText = "Parse a region file, extract the specified chunk, and write some info to stdout.")]
    public class DumpChunkOptions
    {
        [Option("world", HelpText = "The name of the world in which to find the region.")]
        public string WorldName { get; set; }

        [Option("world-dir", HelpText = "The directory containing the world in which to find the region.")]
        public string WorldDir { get; set; }

        [Option("cx", Default = 0, HelpText = "The X coordinate of the chunk.")]
        public int X { get; set; }

        [Option("cy", Default = 0, HelpText = "The Y coordinate of the slice of the chunk to display.")]
        public int Y { get; set; }

        [Option("cz", Default = 0, HelpText = "The Z coordinate of the chunk.")]
        public int Z { get; set; }
    }
}
