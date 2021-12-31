// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Util
{
    public class ChunkPoint
    {
        public ChunkPoint(int x, int z)
        {
            X = x;
            Z = z;
        }


        public int X { get; set; }
        public int Z { get; set; }


        public RegionPoint ToRegion()
        {
            return new RegionPoint
            {
                X = X >> 5,
                Z = Z >> 5
            };
        }


        public override string ToString()
        {
            return $"({X}, {Z})";
        }
    }
}
