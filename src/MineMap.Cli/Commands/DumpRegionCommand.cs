// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;

using MineMap.Cli.Helpers;
using MineMap.Cli.Options;
using MineMap.Lib.Files;
using MineMap.Lib.Util;

namespace MineMap.Cli.Commands
{
    public class DumpRegionCommand
    {
        public void Run(DumpRegionOptions options)
        {
            // Set up the world, which will be used to find the region file.
            var world = options.GetWorld();

            // Get the path to the region file
            var pt = new Coordinate2D(options.X, options.Z, CoordinateType2D.Region);

            var regionPath = world.GetRegionPath(pt);

            // Does the region actually exist?
            if (!File.Exists(regionPath))
            {
                Console.WriteLine("Region file '{0}' does not exist.", regionPath);
                return;
            }

            // Open up the region file, extract the desired info, and write it out.
            using (var region = new Region(regionPath))
            {
                if (options.ShowChunkMap)
                {
                    WriteChunkMap(region);
                }

                // TODO - other options
            }
        }


        private void WriteChunkMap(Region region)
        {
            Console.WriteLine("** Chunk Map **");
            Console.WriteLine("      X         1 1 1 1 1 1 2 2 2 2 3");
            Console.WriteLine("      0 2 4 6 8 0 2 4 6 8 0 2 4 6 8 0");
            for (var z = 0; z < 32; z++)
            {
                Console.Write("{0,2} {1,-2} ", z == 0 ? "Z" : string.Empty, z);

                for (var x = 0; x < 32; x++)
                {
                    // TODO - xyzzy - Type should be RegionChunk or ChunkRegion, not just Chunk
                    var pt = new Coordinate2D(x, z, CoordinateType2D.Chunk);

                    if (region.HasChunk(pt))
                    {
                        Console.Write("X");
                    }
                    else
                    {
                        Console.Write(".");
                    }
                }

                Console.WriteLine();
            }
        }
    }
}
