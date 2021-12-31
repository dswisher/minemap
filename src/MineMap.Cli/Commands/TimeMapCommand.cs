// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Cli.Helpers;
using MineMap.Cli.Options;
using MineMap.Lib.Files;
using MineMap.Lib.Util;

namespace MineMap.Cli.Commands
{
    public class TimeMapCommand
    {
        public void Run(TimeMapOptions options)
        {
            // Set up the world, which will be used to find the files in the world.
            var world = options.GetWorld();

            // To properly scale the histogram, all chunks must be scanned. A 2D sparse matrix is used
            // to store the inhabited times for each chunk. While that is being populated, the histogram
            // is also created.
            var matrix = new Sparse2DMatrix<int, int, long>();

            foreach (var regionPath in world.ListRegionPaths())
            {
                var region = new Region(regionPath);

                Console.WriteLine("Region {0}, X={1}, Z={2}:", regionPath, region.X, region.Z);

                var found = 0;
                for (var x = 0; x < 32; x++)
                {
                    for (var z = 0; z < 32; z++)
                    {
                        var pt = new ChunkPoint((region.X * 32) + x, (region.Z * 32) + z);

                        if (region.HasChunk(pt))
                        {
                            found += 1;

                            var chunk = region.GetChunk(pt);
                            var time = chunk.InhabitedTime;

                            matrix[pt.X, pt.Z] = time;

                            // TODO - add to the gradient
                        }
                    }
                }

                Console.WriteLine("   {0} chunks", found);
            }

            // Use the data to create the image
            // TODO - create the image
            Console.WriteLine("TimeMap is not yet implemented!");
        }
    }
}
