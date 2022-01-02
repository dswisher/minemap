// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using System;

using MineMap.Lib.Util;

namespace MineMap.Lib.Exceptions
{
    public class CoordinateConversionException : Exception
    {
        public CoordinateConversionException(string message)
            : base(message)
        {
        }


        public CoordinateConversionException(CoordinateType2D from, CoordinateType2D to)
            : this($"Cannot convert {from} coordinate to {to} coordinate.")
        {
        }
    }
}
