// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

namespace MineMap.Lib.Graphics
{
    public class Transform1D
    {
        private long minInput;
        private long maxInput;
        private long minOutput;
        private long maxOutput;

        private double scale;
        private long delta;


        public void InputRange(long min, long max)
        {
            minInput = min;
            maxInput = max;

            Recalc();
        }


        public void OutputRange(long min, long max)
        {
            minOutput = min;
            maxOutput = max;

            Recalc();
        }


        public long Transform(long val)
        {
            var result = (long)(val * scale) + delta;

            if (result < minOutput)
            {
                result = minOutput;
            }
            else if (result > maxOutput)
            {
                result = maxOutput;
            }

            return result;
        }


        private void Recalc()
        {
            var deltaIn = maxInput - minInput;
            var deltaOut = maxOutput - minOutput;

            scale = (double)deltaOut / deltaIn;
            delta = minOutput - minInput;
        }
    }
}
