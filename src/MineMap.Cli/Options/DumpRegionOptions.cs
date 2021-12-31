// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using CommandLine;

namespace MineMap.Cli.Options
{
    [Verb("dump-region", HelpText = "Parse a region file and write some info to stdout.")]
    public class DumpRegionOptions : IWorldOptions
    {
        [Option("world", HelpText = "The name of the world in which to find the region.")]
        public string WorldName { get; set; }

        [Option("world-dir", HelpText = "The directory containing the world in which to find the region.")]
        public string WorldDir { get; set; }

        [Option("rx", Default = 0, HelpText = "The X coordinate of the region.")]
        public int X { get; set; }

        [Option("rz", Default = 0, HelpText = "The Z coordinate of the region.")]
        public int Z { get; set; }

        [Option("chunk-map", HelpText = "Show a small map of which chunks are present in the region.")]
        public bool ShowChunkMap { get; set; }
    }
}
