// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using MineMap.Cli.Exceptions;
using MineMap.Cli.Options;
using MineMap.Lib.Files;

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
    }
}
