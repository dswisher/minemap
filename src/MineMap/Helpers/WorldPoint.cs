// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Helpers
{
    public class WorldPoint
    {
        public WorldPoint(int x, int y, int z)
        {
            X = x;
            Y = y;
            Z = z;
        }


        public int X { get; set; }
        public int Y { get; set; }
        public int Z { get; set; }


        public ChunkPoint ToChunk()
        {
            return new ChunkPoint(X / 16, Y / 16, Z / 16);
        }


        public override string ToString()
        {
            return $"({X}, {Y}, {Z})";
        }
    }
}
