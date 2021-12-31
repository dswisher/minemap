// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System.Collections.Generic;

namespace MineMap.Lib.Util
{
    public class Sparse2DMatrix<TX, TY, TValue>
    {
        private readonly Dictionary<MatrixKey<TX, TY>, TValue> dict = new Dictionary<MatrixKey<TX, TY>, TValue>();


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
            }
        }


        private MatrixKey<TX, TY> MakeKey(TX x, TY y)
        {
            return new MatrixKey<TX, TY>(x, y);
        }
    }
}
