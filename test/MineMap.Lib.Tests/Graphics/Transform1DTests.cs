// Copyright (c) Doug Swisher. All Rights Reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

using FluentAssertions;
using MineMap.Lib.Graphics;
using Xunit;

namespace MineMap.Lib.Tests.Graphics
{
    public class Transform1DTests
    {
        private readonly Transform1D transform = new Transform1D();


        [Theory]
        [InlineData(0, 0)]
        [InlineData(100, 100)]
        public void CanDoIdentityTransform(int input, int expected)
        {
            // Arrange
            transform.InputRange(0, 100);
            transform.OutputRange(0, 100);

            // Act
            var result = transform.Transform(input);

            // Assert
            result.Should().Be(expected);
        }


        [Theory]
        [InlineData(-1, 0)]
        [InlineData(101, 100)]
        public void OutputIsClamped(int input, int expected)
        {
            // Arrange
            transform.InputRange(0, 100);
            transform.OutputRange(0, 100);

            // Act
            var result = transform.Transform(input);

            // Assert
            result.Should().Be(expected);
        }


        [Theory]
        [InlineData(-50, 0)]
        [InlineData(50, 100)]
        public void CanShiftNegative(int input, int expected)
        {
            // Arrange
            transform.InputRange(-50, 50);
            transform.OutputRange(0, 100);

            // Act
            var result = transform.Transform(input);

            // Assert
            result.Should().Be(expected);
        }


        [Theory]
        [InlineData(0, 128)]
        [InlineData(128, 255)]
        public void CanShiftColor(int input, int expected)
        {
            // Arrange
            transform.InputRange(0, 128);
            transform.OutputRange(128, 255);

            // Act
            var result = transform.Transform(input);

            // Assert
            result.Should().Be(expected);
        }
    }
}
