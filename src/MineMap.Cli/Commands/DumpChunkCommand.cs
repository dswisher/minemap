// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.IO;

using MineMap.Cli.Options;
using MineMap.Lib.Chunks;
using MineMap.Lib.Files;
using MineMap.Lib.Nbt;
using MineMap.Lib.Util;

namespace MineMap.Cli.Commands
{
    public class DumpChunkCommand
    {
        public void Run(DumpChunkOptions options)
        {
            // Set up the world, which will be used to find the region file.
            // TODO - extract this out to a common place, and use here and in DumpRegionCommand
            World world = null;
            if (!string.IsNullOrEmpty(options.WorldDir))
            {
                world = World.FromDirectory(options.WorldDir);
            }
            else if (!string.IsNullOrEmpty(options.WorldName))
            {
                world = World.FromName(options.WorldName);
            }
            else
            {
                // TODO - throw a custom exception and catch it in Program, to properly set the exit status.
                Console.WriteLine("You must specify either a world directory or a world name.");
                return;
            }

            // Determine the region containing the chunk, and get the path to the region
            var chunkPoint = new ChunkPoint(options.X, options.Y, options.Z);
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

                using (var chunkStream = region.GetChunkStream(chunkPoint))
                using (var wrapper = new StreamWrapper(chunkStream))
                using (var reader = new NbtReader(wrapper))
                {
                    var rootTag = reader.ReadTag().AsCompound();

                    var chunk = Chunk.LoadFrom(rootTag);

                    // TODO - dump some info
                }
            }

            // TODO
            Console.WriteLine("Dump chunk is not yet implemented!");
        }
    }
}
