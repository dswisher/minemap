// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using SixLabors.ImageSharp.PixelFormats;

namespace MineMap.Lib.Graphics
{
    public class ColorMap
    {
        private Transform1D transform;
        private long minVal = long.MaxValue;
        private long maxVal = long.MinValue;


        public void AddSample(long val)
        {
            if (val < minVal)
            {
                minVal = val;
            }

            if (val > maxVal)
            {
                maxVal = val;
            }

            transform = null;
        }


        public Rgba32 GetColor(long val)
        {
            if (transform == null)
            {
                transform = new Transform1D();

                transform.InputRange(minVal, maxVal);
                transform.OutputRange(128, 255);
            }

            var red = (byte)transform.Transform(val);

            return new Rgba32(red, 0x00, 0x00, 0xFF);
        }
    }
}
