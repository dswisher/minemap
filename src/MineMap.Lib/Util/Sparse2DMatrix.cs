// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;
using System.Collections.Generic;

namespace MineMap.Lib.Util
{
    public class Sparse2DMatrix<TX, TY, TValue>
        where TX : IComparable<TX>
        where TY : IComparable<TY>
    {
        private readonly Dictionary<MatrixKey<TX, TY>, TValue> dict = new Dictionary<MatrixKey<TX, TY>, TValue>();


        public TX MinX { get; private set; }
        public TX MaxX { get; private set; }
        public TY MinY { get; private set; }
        public TY MaxY { get; private set; }


        public TValue this[TX x, TY y]
        {
            get
            {
                TValue val;

                if (!dict.TryGetValue(MakeKey(x, y), out val))
                {
                    return default(TValue);
                }

                return val;
            }

            set
            {
                dict[MakeKey(x, y)] = value;

                if (dict.Count == 1)
                {
                    MinX = x;
                    MaxX = x;

                    MinY = y;
                    MaxY = y;
                }
                else
                {
                    if (x.CompareTo(MinX) < 0)
                    {
                        MinX = x;
                    }

                    if (x.CompareTo(MaxX) > 0)
                    {
                        MaxX = x;
                    }

                    if (y.CompareTo(MinY) < 0)
                    {
                        MinY = y;
                    }

                    if (y.CompareTo(MaxY) > 0)
                    {
                        MaxY = y;
                    }
                }
            }
        }


        private MatrixKey<TX, TY> MakeKey(TX x, TY y)
        {
            return new MatrixKey<TX, TY>(x, y);
        }
    }
}
