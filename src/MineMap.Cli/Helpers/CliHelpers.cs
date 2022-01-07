// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System.Collections.Generic;

using MineMap.Cli.Exceptions;
using MineMap.Cli.Options;
using MineMap.Lib.Files;
using MineMap.Lib.Util;

namespace MineMap.Cli.Helpers
{
    public static class CliHelpers
    {
        public static World GetWorld(this IWorldOptions options)
        {
            // Create the world object
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
                throw new CliException("You must specify either a world directory or a world name.");
            }

            return world;
        }


        public static IEnumerable<Coordinate2D> GetRegionsInRect(this IBlockRectOptions options)
        {
            // Calculate the bounds of the plot area, in blocks
            var (block1, block2) = GetBounds(options);

            // Convert both to region coords
            var region1 = block1.ToRegion();
            var region2 = block2.ToRegion();

            // Enumerate and return all the regions
            for (var dx = region1.X; dx <= region2.X; dx++)
            {
                for (var dz = region1.Z; dz <= region2.Z; dz++)
                {
                    yield return new Coordinate2D(dx, dz, CoordinateType2D.Region);
                }
            }
        }


        public static IEnumerable<Coordinate2D> GetChunksInRect(this IBlockRectOptions options, Coordinate2D regionPt)
        {
            // Calculate the bounds of the plot area, in blocks
            var (block1, block2) = GetBounds(options);

            // TODO
            yield return new Coordinate2D();    // HACK!
        }


        private static (Coordinate2D, Coordinate2D) GetBounds(IBlockRectOptions options)
        {
            // Calculate the bounds of the plot area, in blocks
            var x1 = options.CenterX - (options.Width / 2);
            var z1 = options.CenterZ - (options.Height / 2);

            var x2 = x1 + options.Width;
            var z2 = z1 + options.Height;

            var block1 = new Coordinate2D(x1, z1, CoordinateType2D.Block);
            var block2 = new Coordinate2D(x2, z2, CoordinateType2D.Block);

            // Return what we've got
            return (block1, block2);
        }
    }
}
