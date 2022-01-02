// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using CommandLine;

namespace MineMap.Cli.Options
{
    [Verb("time-map", HelpText = "Create a map that shows inhabited time for each chunk using a gradient.")]
    public class TimeMapOptions : IWorldOptions
    {
        [Option("world", HelpText = "The name of the world to use.")]
        public string WorldName { get; set; }

        [Option("world-dir", HelpText = "The directory containing the world to use.")]
        public string WorldDir { get; set; }

        [Option("output", Default = "time-map.png", HelpText = "The path of the image file to create.")]
        public string OutputPath { get; set; }

        [Option("chunk-size", Default = 2, HelpText = "The width and height of a chunk, in pixels.")]
        public int ChunkSize { get; set; }
    }
}
