// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

namespace MineMap.Lib.Util
{
    internal class MatrixKey<TX, TY> : IEquatable<MatrixKey<TX, TY>>
    {
        public MatrixKey(TX x, TY y)
        {
            X = x;
            Y = y;
        }


        public TX X { get; private set; }
        public TY Y { get; private set; }


        public bool Equals(MatrixKey<TX, TY> other)
        {
            if (other == null)
            {
                return false;
            }

            if (X.Equals(other.X) && Y.Equals(other.Y))
            {
                return true;
            }

            return false;
        }


        public override bool Equals(object obj)
        {
            if (obj == null)
            {
                return false;
            }

            var other = obj as MatrixKey<TX, TY>;
            if (other == null)
            {
                return false;
            }

            return Equals(other);
        }


        public override int GetHashCode()
        {
            int hash = 17;

            hash = (hash * 23) + X.GetHashCode();
            hash = (hash * 23) + Y.GetHashCode();

            return hash;
        }
    }
}
