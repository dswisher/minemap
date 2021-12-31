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
    public class DumpChunkCommand
    {
        public void Run(DumpChunkOptions options)
        {
            // Set up the world, which will be used to find the region file.
            var world = options.GetWorld();

            // Determine the region containing the chunk, and get the path to the region
            var chunkPoint = new ChunkPoint(options.X, options.Z);
            var regionPoint = chunkPoint.ToRegion();
            var regionPath = world.GetRegionPath(regionPoint);

            // Does the region actually exist?
            if (!File.Exists(regionPath))
            {
                Console.WriteLine("Region file '{0}' does not exist.", regionPath);
                return;
            }

            // Open up the region file, extract the desired info, and write it out.
            using (var region = new Region(regionPath))
            {
                if (!region.HasChunk(chunkPoint))
                {
                    Console.WriteLine("Chunk {0} not found in region {1}.", chunkPoint, regionPoint);
                    return;
                }

                var chunk = region.GetChunk(chunkPoint);

                Console.WriteLine("Chunk ({0},{1},{2}):", chunk.X, chunk.Y, chunk.Z);
                Console.WriteLine("   Inhabited Time: {0}", chunk.InhabitedTime);

                // TODO - dump more info
            }
        }
    }
}
