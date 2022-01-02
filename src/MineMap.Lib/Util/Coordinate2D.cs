// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using MineMap.Lib.Exceptions;

namespace MineMap.Lib.Util
{
    public enum CoordinateType2D
    {
        /// <summary>
        /// The coordinates of a block within the world
        /// </summary>
        Block,

        /// <summary>
        /// The coordinates of a chunk within the world
        /// </summary>
        Chunk,

        /// <summary>
        /// The coordinates of a region
        /// </summary>
        Region
    }


    /// <summary>
    /// A coordinate (point) within a specified coordinate system.
    /// </summary>
    public class Coordinate2D
    {
        public Coordinate2D()
            : this(CoordinateType2D.Block)
        {
        }


        public Coordinate2D(CoordinateType2D type)
            : this(0, 0, type)
        {
        }


        public Coordinate2D(int x, int z, CoordinateType2D type)
        {
            X = x;
            Z = z;
            Type = type;
        }


        public CoordinateType2D Type { get; set; }
        public int X { get; set; }
        public int Z { get; set; }


        public override string ToString()
        {
            // TODO - come up with abbreviations for the various types
            return $"{Type}({X},{Z})";
        }


        public Coordinate2D ToChunk()
        {
            if (Type == CoordinateType2D.Chunk)
            {
                return this;
            }
            else if (Type == CoordinateType2D.Block)
            {
                return new Coordinate2D(X >> 4, Z >> 4, CoordinateType2D.Chunk);
            }
            else
            {
                throw new CoordinateConversionException(Type, CoordinateType2D.Chunk);
            }
        }


        public Coordinate2D ToRegion()
        {
            if (Type == CoordinateType2D.Region)
            {
                return this;
            }
            else if (Type == CoordinateType2D.Chunk)
            {
                return new Coordinate2D(X >> 5, Z >> 5, CoordinateType2D.Region);
            }
            else if (Type == CoordinateType2D.Block)
            {
                return new Coordinate2D(X >> 9, Z >> 9, CoordinateType2D.Region);
            }
            else
            {
                throw new CoordinateConversionException(Type, CoordinateType2D.Region);
            }
        }
    }
}
