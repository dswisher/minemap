// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using CommandLine;

namespace MineMap.Cli.Options
{
    [Verb("height-map", HelpText = "Create a map that shows the height of the world surface.")]
    public class HeightMapOptions : IWorldOptions, IBlockRectOptions
    {
        [Option("world", HelpText = "The name of the world to use.")]
        public string WorldName { get; set; }

        [Option("world-dir", HelpText = "The directory containing the world to use.")]
        public string WorldDir { get; set; }

        [Option("output", Default = "height-map.png", HelpText = "The path of the image file to create.")]
        public string OutputPath { get; set; }

        [Option("width", Default = 1024, HelpText = "The width of the output image.")]
        public int Width { get; set; }

        [Option("height", Default = 768, HelpText = "The height of the output image.")]
        public int Height { get; set; }

        [Option("cx", Default = 0, HelpText = "The X coordinate of the block in the center of the image.")]
        public int CenterX { get; set; }

        [Option("cz", Default = 0, HelpText = "The Z coordinate of the block in the center of the image.")]
        public int CenterZ { get; set; }
    }
}
