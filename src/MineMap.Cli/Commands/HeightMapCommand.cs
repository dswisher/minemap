// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;

using MineMap.Cli.Helpers;
using MineMap.Cli.Options;
using MineMap.Lib.Files;

using SixLabors.ImageSharp;
using SixLabors.ImageSharp.Drawing.Processing;
using SixLabors.ImageSharp.PixelFormats;
using SixLabors.ImageSharp.Processing;

namespace MineMap.Cli.Commands
{
    public class HeightMapCommand
    {
        public void Run(HeightMapOptions options)
        {
            // Set up the world, which will be used to find the files in the world.
            var world = options.GetWorld();

            // Create the image
            using (var image = new Image<Rgba32>(options.Width, options.Height))
            {
                // Set the image to black
                image.Mutate(x => x.Fill(Color.Black));

                // Go through all the regions that are needed to generate the image. Note that
                // only a handful of chunks may be needed from the region.
                foreach (var regionPt in options.GetRegionsInRect())
                {
                    var regionPath = world.GetRegionPath(regionPt);

                    if (!File.Exists(regionPath))
                    {
                        continue;
                    }

                    var region = new Region(regionPath);

                    Console.WriteLine("-> {0}", regionPath);

                    // Go through all the chunks within this region that are needed. Note that not all
                    // blocks within the chunk may lie within the output rectangle.
                    foreach (var chunkPt in options.GetChunksInRect(regionPt))
                    {
                        if (!region.HasChunk(chunkPt))
                        {
                            continue;
                        }

                        // Go through all the needed blocks within this chunk
                        // TODO

                        // TODO - read each chunk and render the height map
                    }
                }

                Console.WriteLine("Height map is not yet implemented!");

                // Save the image
                image.SaveAsPng(options.OutputPath);
                Console.WriteLine("Wrote image to {0}.", options.OutputPath);
            }
        }
    }
}
