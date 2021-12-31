// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

using MineMap.Lib.Util;

namespace MineMap.Lib.Files
{
    public class World
    {
        private readonly DirectoryInfo worldRoot;

        public World(string path)
        {
            // TODO - parse info out of level.dat?
            // TODO - make sure this is really a world directory?
            worldRoot = new DirectoryInfo(path);
        }


        public static World FromDirectory(string path)
        {
            return new World(path);
        }


        public static World FromName(string name)
        {
            // TODO - use the proper "root" for each platform
            var root = Path.Join(Environment.GetFolderPath(Environment.SpecialFolder.Personal), "Library/Application Support/minecraft/saves");

            return new World(Path.Join(root, name));
        }


        public string GetRegionPath(RegionPoint pt)
        {
            return Path.Join(worldRoot.FullName, "region", $"r.{pt.X}.{pt.Z}.mca");
        }


        public IEnumerable<string> ListRegionPaths()
        {
            foreach (var regionPath in Directory.GetFiles(Path.Join(worldRoot.FullName, "region")).OrderBy(x => x))
            {
                yield return regionPath;
            }
        }
    }
}
