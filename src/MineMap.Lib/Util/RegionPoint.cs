// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Util
{
    public class RegionPoint
    {
        public RegionPoint()
        {
        }


        public RegionPoint(int x, int z)
        {
            X = x;
            Z = z;
        }


        public int X { get; set; }
        public int Z { get; set; }


        public override string ToString()
        {
            return $"({X}, {Z})";
        }


        public override int GetHashCode()
        {
            int hash = 17;

            hash = (hash * 23) + X.GetHashCode();
            hash = (hash * 23) + Z.GetHashCode();

            return hash;
        }


        public override bool Equals(object obj)
        {
            return Equals(obj as RegionPoint);
        }


        public bool Equals(RegionPoint obj)
        {
            return (obj != null) && (obj.X == X) && (obj.Z == Z);
        }
    }
}
