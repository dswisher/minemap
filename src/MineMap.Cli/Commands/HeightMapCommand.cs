// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Cli.Helpers;
using MineMap.Cli.Options;
using MineMap.Lib.Files;
using MineMap.Lib.Graphics;
using MineMap.Lib.Util;

using SixLabors.ImageSharp;
using SixLabors.ImageSharp.PixelFormats;

namespace MineMap.Cli.Commands
{
    public class HeightMapCommand
    {
        public void Run(HeightMapOptions options)
        {
            // Set up the world, which will be used to find the files in the world.
            var world = options.GetWorld();

            // Use the data to create the image
            using (var image = new Image<Rgba32>(options.Width, options.Height))
            {
                foreach (var regionPath in options.GetRegionsInRect())
                {
                    // TODO - read each chunk and render the height map
                }

                Console.WriteLine("Height map is not yet implemented!");

                image.SaveAsPng(options.OutputPath);
                Console.WriteLine("Wrote image to {0}.", options.OutputPath);
            }
        }
    }
}
