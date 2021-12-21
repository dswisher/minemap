// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Util
{
    public enum CoordinateType
    {
        /// <summary>
        /// The coordinates of a chunk within the world
        /// </summary>
        ChunkGlobal,

        /// <summary>
        /// The coordinates of a chunk within a region
        /// </summary>
        ChunkRegion,

        /// <summary>
        /// The coordinates of a region
        /// </summary>
        Region,

        /// <summary>
        /// The coordinates of a block within a chunk
        /// </summary>
        BlockChunk,

        /// <summary>
        /// The coordinates of a block within the world
        /// </summary>
        BlockGlobal
    }


    /// <summary>
    /// A coordinate (point) within a specified coordinate system.
    /// </summary>
    public class Coordinate
    {
        public Coordinate()
            : this(CoordinateType.BlockGlobal)
        {
        }


        public Coordinate(CoordinateType type)
            : this(0, 0, 0, type)
        {
        }


        public Coordinate(int x, int y, int z, CoordinateType type)
        {
            X = x;
            Y = y;
            Z = z;
            Type = type;
        }


        public CoordinateType Type { get; set; }
        public int X { get; set; }
        public int Y { get; set; }
        public int Z { get; set; }


        public override string ToString()
        {
            return $"{Type}({X},{Y},{Z})";
        }
    }
}
